package obctrl

import (
	"github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/logger"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/asynkron/protoactor-go/actor"
)

const actorName = "ObCtrl"

var log = logger.NewLogger(actorName)

type ObController struct {
	pid      *actor.PID
	children []*actor.PID
}

func (s *ObController) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T\n", ctx.Self().Id, ctx.Message())

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		s.init(ctx)
	case *msg.ArticlesListObReadyMsg:
		s.procArticlesListObReadyMsg(ctx)
	default:
		log.Infof("%v recv unknow msg: %T\n", ctx.Self().Id, ctxMsg)
	}
}

func (s *ObController) procArticlesListObReadyMsg(ctx actor.Context) {
	config := cfg.GlobalConfig.ArticleUrl[0].Clone()
	ctx.Send(ctx.Sender(), &msg.ArticlesListTaskMsg{Target: config})
}
