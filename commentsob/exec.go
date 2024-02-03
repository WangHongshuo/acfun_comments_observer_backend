package commentsob

import (
	"fmt"

	"github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/commentsob/getter"
	"github.com/WangHongshuo/acfun_comments_observer_backend/dao/model"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/logger"
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
	parent *actor.PID
	ctx    actor.Context
	timer  *scheduler.TimerScheduler
	config cfg.ObserverConfig

	db            *gorm.DB
	aidList       []int64
	commentsCache []model.Comment
	articleCache  model.Article
}

func (c *CommentsOb) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T", ctx.Self().Id, ctx.Message())
	var err error

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		c.init(ctx)
	case *msg.ResourceReadyMsg:
		c.initResource(ctx)
	case *msg.ObserveCommentsTaskMsg:
		c.procObserveCommentsTaskMsg(ctxMsg)
	case *observeNextCommentsPage:
		err = c.procObserveNextCommentsPageMsg(ctxMsg)
	case *observeNextArticle:
		c.procObserveNextArticleMsg()
	default:
		log.Infof("%v recv unknow msg: %T", ctx.Self().Id, ctxMsg)
	}

	if err != nil {
		log.Errorf("%v proc msg err: %v", ctx.Self().Id, err)
	}
}

func (c *CommentsOb) procObserveCommentsTaskMsg(ctxMsg *msg.ObserveCommentsTaskMsg) {
	log.Infof("%v recv task: %v", c.pid.GetId(), ctxMsg)
	if ctxMsg == nil || len(ctxMsg.Aids) == 0 {
		return
	}

	c.aidList = ctxMsg.Aids
	c.startObserveNextArticleTimer()
}

func (c *CommentsOb) procObserveNextArticleMsg() {
	if len(c.aidList) == 0 {
		log.Infof("%v all task finished", c.pid.GetId())
		c.commentsCache = make([]model.Comment, 0)
		c.ctx.RequestWithCustomSender(c.parent, &msg.ObserveCommentsTaskFinishedMsg{}, c.pid)
		return
	}

	n := len(c.aidList)
	aid := c.aidList[n-1]
	c.aidList = c.aidList[:n-1]
	log.Infof("%v start observe aid: %v", c.pid.GetId(), aid)

	// init cache
	c.commentsCache = c.commentsCache[:0]
	c.articleCache = c.getArticleData(aid)
	c.startObserveNextCommentsPageTimer(&observeNextCommentsPage{
		isNewAid: true,
		aid:      aid,
	})
}

func (c *CommentsOb) procObserveNextCommentsPageMsg(ctxMsg *observeNextCommentsPage) error {
	if ctxMsg == nil {
		return fmt.Errorf("observe next comments page msg is nil")
	}

	if ctxMsg.isNewAid {
		proxyAddr, err := proxypool.GlobalProxyPool.GetHttpsProxy()
		if err != nil {
			c.startObserveNextArticleTimer()
			return fmt.Errorf("get proxy error: %v", err)
		}
		ctxMsg.proxyAddr = proxyAddr
		ctxMsg.nextPage = 1
		ctxMsg.oldFloor = int64(c.articleCache.LastFloorNumber)
	}

	comments, totalPage, err := getter.CommentsGetter(ctxMsg.proxyAddr, ctxMsg.aid, ctxMsg.nextPage)
	if err != nil {
		c.startObserveNextArticleTimer()
		return fmt.Errorf("get comments error: %v", err)
	}
	log.Infof("ob next page for aid: %v, curr: %v, total: %v", ctxMsg.aid, ctxMsg.nextPage, totalPage)

	isFinished := false
	for i := range comments {
		if comments[i].Floor <= int64(ctxMsg.oldFloor) {
			isFinished = true
			break
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

	if len(c.commentsCache) > 0 {
		c.articleCache.LastFloorNumber = int32(c.commentsCache[0].FloorNumber)
	}

	if isFinished || ctxMsg.nextPage == int(totalPage) {
		c.articleCache.IsCompleted = true
		oldCommentsCount := c.articleCache.CommentsCount
		c.articleCache.CommentsCount += int32(len(c.commentsCache))
		c.commitAll()
		log.Infof("%v ob aid: %v completed, new last floor number: %v, new comments: %v, from: %v, to: %v",
			c.pid.GetId(), ctxMsg.aid, c.articleCache.LastFloorNumber, len(c.commentsCache),
			oldCommentsCount, c.articleCache.CommentsCount)
		c.startObserveNextArticleTimer()
		return nil
	}

	c.startObserveNextCommentsPageTimer(&observeNextCommentsPage{
		aid:       ctxMsg.aid,
		oldFloor:  int64(ctxMsg.oldFloor),
		nextPage:  ctxMsg.nextPage + 1,
		proxyAddr: ctxMsg.proxyAddr,
	})

	log.Infof("end ob next page for aid: %v, curr: %v, total: %v", ctxMsg.aid, ctxMsg.nextPage, totalPage)
	return nil
}

func (c *CommentsOb) getArticleData(aid int64) model.Article {
	var result []model.Article
	c.db.Where("aid = ?", aid).Find(&result)
	if len(result) == 0 {
		return model.Article{Aid: aid}
	}
	return result[0]
}
