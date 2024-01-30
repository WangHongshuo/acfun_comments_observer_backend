package commentsspider

import (
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/asynkron/protoactor-go/actor"
)

const actorName = "CommentsExec"

var log = logger.NewLogger(actorName)

type CommentsExecutor struct{}

func (c *CommentsExecutor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	default:
		log.Infof("%v recv msg: %T\n", ctx.Self().Id, msg)
	}
}
