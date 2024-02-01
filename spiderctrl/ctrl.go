package spiderctrl

import (
	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/WangHongshuo/acfuncommentsspider-go/msg"
	"github.com/asynkron/protoactor-go/actor"
)

const actorName = "SpiderCtrl"

var log = logger.NewLogger(actorName)

type SpiderController struct {
	pid      *actor.PID
	children []*actor.PID
}

func (s *SpiderController) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T\n", ctx.Self().Id, ctx.Message())

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		s.init(ctx)
	case *msg.ArticlesListExecReadyMsg:
		s.procArticlesListExecReadyMsg(ctx)
	default:
		log.Infof("%v recv unknow msg: %T\n", ctx.Self().Id, ctxMsg)
	}
}

func (s *SpiderController) procArticlesListExecReadyMsg(ctx actor.Context) {
	config := cfg.GlobalConfig.ArticleUrl[0].Clone()
	ctx.Send(ctx.Sender(), &msg.ArticlesListTaskMsg{Target: config})
}
