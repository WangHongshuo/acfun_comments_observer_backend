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
	db     *gorm.DB
}

func (c *CommentsExecutor) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T\n", ctx.Self().Id, ctx.Message())

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		c.init(ctx)
	case *msg.ResourceReadyMsg:
		c.initResource(ctx)
	case *msg.CommentsTaskMsg:
		log.Infof("%v recv: %v\n", ctx.Self().Id, ctxMsg)
	default:
		log.Infof("%v recv unknow msg: %T\n", ctx.Self().Id, ctxMsg)
	}
}
