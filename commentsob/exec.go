package commentsob

import (
	"fmt"

	"github.com/WangHongshuo/acfun_comments_observer_backend/commentsob/getter"
	"github.com/WangHongshuo/acfun_comments_observer_backend/dao/model"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/logger"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/util"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/proxypool"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/scheduler"
	"gorm.io/gorm"
)

const actorName = "CommentsOb"

var log = logger.NewLogger(actorName)

type CommentsOb struct {
	pid    *actor.PID
	instId int
	ctx    actor.Context
	timer  *scheduler.TimerScheduler

	db            *gorm.DB
	aidList       []int64
	commentsCache []model.Comment

	perArticleMinDelay      int
	perArticleMaxDelay      int
	perCommentsPageMinDelay int
	perCommentsPageMaxDelay int
}

func (c *CommentsOb) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T\n", ctx.Self().Id, ctx.Message())
	var err error

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		c.init(ctx)
	case *msg.ResourceReadyMsg:
		c.initResource(ctx)
	case *msg.CommentsTaskMsg:
		c.procCommentsTaskMsg(ctxMsg)
	case *observeNextCommentsPage:
		err = c.procObserveNextCommentsPageMsg(ctxMsg)
	case *observeNextArticle:
		c.startObserveNextArticle()
	default:
		log.Infof("%v recv unknow msg: %T\n", ctx.Self().Id, ctxMsg)
	}

	if err != nil {
		log.Errorf("%v proc msg err: %v\n", ctx.Self().Id, err)
	}
}

func (c *CommentsOb) procCommentsTaskMsg(ctxMsg *msg.CommentsTaskMsg) {
	log.Infof("%v recv: %v\n", c.pid.GetId(), ctxMsg)
	if ctxMsg == nil || len(ctxMsg.Aids) == 0 {
		return
	}

	c.aidList = ctxMsg.Aids
	c.timer.SendOnce(util.GetRandomDuration(c.perArticleMinDelay, c.perArticleMaxDelay, int64(c.instId)),
		c.pid, &observeNextArticle{})
}

func (c *CommentsOb) startObserveNextArticle() {
	if len(c.aidList) == 0 {
		log.Infof("%v all task finished\n", c.pid.GetId())
		return
	}

	n := len(c.aidList)
	aid := c.aidList[n-1]
	c.aidList = c.aidList[:n-1]
	log.Infof("%v start observe aid: %v\n", c.pid.GetId(), aid)
	if err := c.observeComments(aid); err != nil {
		log.Errorf("%v observe comments error: %v\n", c.pid.GetId(), err)
	}
}

func (c *CommentsOb) observeComments(aid int64) error {
	proxyAddr, err := proxypool.GlobalProxyPool.GetHttpsProxy()
	if err != nil {
		return fmt.Errorf("get proxy error: %v", err)
	}
	articleData := c.getArticleData(aid)
	oldFloor := articleData.LastFloorNumber
	comments, totalPage, err := getter.CommentsGetter(proxyAddr, aid, 1)
	if err != nil {
		return fmt.Errorf("get comments error: %v", err)
	}
	if len(comments) == 0 || comments[0].Floor <= int64(oldFloor) {
		return nil
	}

	newFloor := comments[0].Floor
	c.commentsCache = make([]model.Comment, 0)
	for i := range comments {
		if comments[i].Floor <= int64(oldFloor) {
			c.commitCommentsDataToDb(c.commentsCache)
			c.commitArticleDataToDb(model.Article{
				Aid:             aid,
				LastFloorNumber: int32(newFloor),
				IsCompleted:     true,
			})
			log.Infof("%v observe comments completed, aid: %v, last floor number: %v, new comments: %v", c.pid.GetId(), aid, newFloor, len(c.commentsCache))
			c.commentsCache = make([]model.Comment, 0)
			c.timer.SendOnce(util.GetRandomDuration(c.perArticleMinDelay, c.perArticleMaxDelay, int64(c.instId)), c.pid, &observeNextArticle{})
			return nil
		}
		// avoid duplicate comments when observe next page
		n := len(c.commentsCache)
		if n > 0 && c.commentsCache[n-1].FloorNumber <= int32(comments[i].Floor) {
			continue
		}
		c.commentsCache = append(c.commentsCache, model.Comment{
			Cid:         comments[i].Cid,
			Aid:         aid,
			FloorNumber: int32(comments[i].Floor),
			Comment:     comments[i].Content,
		})
	}

	if 1 == int(totalPage) {
		c.commitCommentsDataToDb(c.commentsCache)
		c.commitArticleDataToDb(model.Article{
			Aid:             aid,
			LastFloorNumber: int32(newFloor),
			IsCompleted:     true,
		})
		log.Infof("%v observe comments completed, aid: %v, last floor number: %v, new comments: %v", c.pid.GetId(), aid, newFloor, len(c.commentsCache))
		c.commentsCache = make([]model.Comment, 0)
		c.timer.SendOnce(util.GetRandomDuration(c.perArticleMinDelay, c.perArticleMaxDelay, int64(c.instId)), c.pid, &observeNextArticle{})
		return nil
	}

	c.timer.SendOnce(util.GetRandomDuration(c.perCommentsPageMinDelay, c.perCommentsPageMinDelay, int64(c.instId)), c.pid,
		&observeNextCommentsPage{
			aid:       aid,
			oldFloor:  int64(oldFloor),
			newFloor:  newFloor,
			nextPage:  2,
			totalPage: int(totalPage),
			proxyAddr: proxyAddr,
		})
	return nil
}

func (c *CommentsOb) procObserveNextCommentsPageMsg(ctxMsg *observeNextCommentsPage) error {
	log.Infof("ob next page for aid: %v, curr: %v, total: %v", ctxMsg.aid, ctxMsg.nextPage, ctxMsg.totalPage)
	comments, totalPage, err := getter.CommentsGetter(ctxMsg.proxyAddr, ctxMsg.aid, ctxMsg.nextPage)
	if err != nil {
		return fmt.Errorf("get comments error: %v", err)
	}

	for i := range comments {
		if comments[i].Floor <= int64(ctxMsg.oldFloor) {
			c.commitCommentsDataToDb(c.commentsCache)
			c.commitArticleDataToDb(model.Article{
				Aid:             ctxMsg.aid,
				LastFloorNumber: int32(ctxMsg.newFloor),
				IsCompleted:     true,
			})
			log.Infof("%v observe comments completed, aid: %v, last floor number: %v, new comments: %v", c.pid.GetId(), ctxMsg.aid, ctxMsg.newFloor, len(c.commentsCache))
			c.commentsCache = make([]model.Comment, 0)
			c.timer.SendOnce(util.GetRandomDuration(c.perArticleMinDelay, c.perArticleMaxDelay, int64(c.instId)), c.pid, &observeNextArticle{})
			return nil
		}
		// avoid duplicate comments when observe next page
		n := len(c.commentsCache)
		if n > 0 && c.commentsCache[n-1].FloorNumber <= int32(comments[i].Floor) {
			continue
		}
		c.commentsCache = append(c.commentsCache, model.Comment{
			Cid:         comments[i].Cid,
			Aid:         ctxMsg.aid,
			FloorNumber: int32(comments[i].Floor),
			Comment:     comments[i].Content,
		})
	}

	if ctxMsg.nextPage == int(totalPage) {
		c.commitCommentsDataToDb(c.commentsCache)
		c.commitArticleDataToDb(model.Article{
			Aid:             ctxMsg.aid,
			LastFloorNumber: int32(ctxMsg.newFloor),
			IsCompleted:     true,
		})
		log.Infof("%v observe comments completed, aid: %v, last floor number: %v, new comments: %v", c.pid.GetId(), ctxMsg.aid, ctxMsg.newFloor, len(c.commentsCache))
		c.commentsCache = make([]model.Comment, 0)
		c.timer.SendOnce(util.GetRandomDuration(c.perArticleMinDelay, c.perArticleMaxDelay, int64(c.instId)), c.pid, &observeNextArticle{})
		return nil
	}

	c.timer.SendOnce(util.GetRandomDuration(c.perCommentsPageMinDelay, c.perCommentsPageMinDelay, int64(c.instId)), c.pid,
		&observeNextCommentsPage{
			aid:       ctxMsg.aid,
			oldFloor:  int64(ctxMsg.oldFloor),
			newFloor:  ctxMsg.newFloor,
			nextPage:  ctxMsg.nextPage + 1,
			totalPage: int(totalPage),
			proxyAddr: ctxMsg.proxyAddr,
		})

	log.Infof("end ob next page for aid: %v", ctxMsg.aid)
	return nil
}

func (c *CommentsOb) commitCommentsDataToDb(data []model.Comment) {
	c.db.Save(data)
}

func (c *CommentsOb) commitArticleDataToDb(data model.Article) {
	c.db.Save([]model.Article{data})
}

func (c *CommentsOb) getArticleData(aid int64) model.Article {
	var result []model.Article
	c.db.Where("aid = ?", aid).Find(&result)
	if len(result) == 0 {
		return model.Article{Aid: aid}
	}
	return result[0]
}
