package articleslistspider

import (
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/WangHongshuo/acfuncommentsspider-go/msg"
	"github.com/asynkron/protoactor-go/actor"
	"gorm.io/gorm"
)

const actorName = "ArticlesListExec"

var log = logger.NewLogger(actorName)

type ArticlesListExecutor struct {
	pid      *actor.PID
	instId   int
	children []*actor.PID
	db       *gorm.DB
}

func (a *ArticlesListExecutor) Receive(ctx actor.Context) {
	ctxMsg := ctx.Message()
	log.Infof("%v recv msg: %T\n", ctx.Self().Id, ctxMsg)

	switch ctxMsg.(type) {
	case *actor.Started:
		a.init(ctx)
	case *msg.ResourceReadyMsg:
		a.initResource(ctx)
	default:
		log.Infof("%v recv unknow msg: %T\n", ctx.Self().Id, ctxMsg)
	}
}
