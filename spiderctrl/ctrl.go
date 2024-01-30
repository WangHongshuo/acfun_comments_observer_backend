package spiderctrl

import (
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/asynkron/protoactor-go/actor"
)

const actorName = "SpiderCtrl"

var log = logger.NewLogger(actorName)

type SpiderController struct {
	pid      *actor.PID
	children []*actor.PID
}

func (s *SpiderController) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		s.init(ctx)
	default:
		log.Infof("%v recv msg: %T\n", ctx.Self().Id, msg)
	}
}
