package commentsob

import (
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/logger"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/util"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/asynkron/protoactor-go/actor"
	"gorm.io/gorm"
)

const actorName = "CommentsOb"

var log = logger.NewLogger(actorName)

type CommentsOb struct {
	pid    *actor.PID
	instId int

	db       *gorm.DB
	aidList  []int64
	minDelay int
	maxDelay int
}

func (c *CommentsOb) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T\n", ctx.Self().Id, ctx.Message())

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		c.init(ctx)
	case *msg.ResourceReadyMsg:
		c.initResource(ctx)
	case *msg.CommentsTaskMsg:
		c.procCommentsTaskMsg(ctxMsg, ctx)
	case *actor.ReceiveTimeout:
		c.procObserveCommentsTask()
	default:
		log.Infof("%v recv unknow msg: %T\n", ctx.Self().Id, ctxMsg)
	}
}

func (c *CommentsOb) procCommentsTaskMsg(ctxMsg *msg.CommentsTaskMsg, ctx actor.Context) {
	log.Infof("%v recv: %v\n", c.pid.Id, ctxMsg)
	if ctxMsg == nil || len(ctxMsg.Aids) == 0 {
		return
	}

	c.aidList = ctxMsg.Aids

	dur := util.GetRandomDuration(c.minDelay, c.maxDelay, int64(c.instId))
	log.Infof("%v timeout: %v", c.pid.Id, dur)
	ctx.SetReceiveTimeout(dur)
}

func (c *CommentsOb) procObserveCommentsTask() {
	if len(c.aidList) == 0 {
		log.Infof("%v all task finished\n", c.pid.Id)
		return
	}

	n := len(c.aidList)
	aid := c.aidList[n-1]
	c.aidList = c.aidList[:n-1]

	log.Infof("%v observe aid: %v\n", c.pid.Id, aid)
}
