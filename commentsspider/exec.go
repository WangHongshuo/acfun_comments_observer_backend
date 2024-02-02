package commentsspider

import (
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/WangHongshuo/acfuncommentsspider-go/msg"
	"github.com/asynkron/protoactor-go/actor"
	"gorm.io/gorm"
)

const actorName = "CommentsExec"

var log = logger.NewLogger(actorName)

type CommentsExecutor struct {
	pid    *actor.PID
	instId int

	db      *gorm.DB
	aidList []int64
}

func (c *CommentsExecutor) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T\n", ctx.Self().Id, ctx.Message())

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		c.init(ctx)
	case *msg.ResourceReadyMsg:
		c.initResource(ctx)
	case *msg.CommentsTaskMsg:
		c.procCommentsTaskMsg(ctxMsg)
	default:
		log.Infof("%v recv unknow msg: %T\n", ctx.Self().Id, ctxMsg)
	}
}

func (c *CommentsExecutor) procCommentsTaskMsg(ctxMsg *msg.CommentsTaskMsg) {
	log.Infof("%v recv: %v\n", c.pid.Id, ctxMsg)
	if ctxMsg == nil {
		return
	}

	c.aidList = ctxMsg.Aids
}
