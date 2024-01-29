package spiderctrl

import (
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/asynkron/protoactor-go/actor"
)

var log = logger.SpiderControllerLogger

type SpiderController struct {
}

func (s *SpiderController) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	default:
		log.Infof("SpiderController recv msg: %T\n", msg)
	}
}
