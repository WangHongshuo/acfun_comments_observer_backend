package articleslistsob

import (
	"github.com/WangHongshuo/acfun_comments_observer_backend/articleslistsob/getter"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/logger"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/proxypool"
	"github.com/asynkron/protoactor-go/actor"
)

const actorName = "ArticlesListOb"

var log = logger.NewLogger(actorName)

type ArticlesListOb struct {
	pid    *actor.PID
	parent *actor.PID
	instId int
	ctx    actor.Context

	children    []*actor.PID
	notReadyMap map[string]struct{}
}

func (a *ArticlesListOb) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T\n", ctx.Self().Id, ctx.Message())

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		a.init(ctx)
	case *msg.ResourceReadyMsg:
		a.initResource(ctx)
	case *msg.ArticlesListTaskMsg:
		a.procArticlesListTaskMsg(ctxMsg, ctx)
	case *msg.CommentsObReadyMsg:
		a.procCommentsObReadyMsg(ctx)
	default:
		log.Infof("%v recv unknow msg: %T\n", ctx.Self().Id, ctxMsg)
	}
}

func (a *ArticlesListOb) procCommentsObReadyMsg(ctx actor.Context) {
	delete(a.notReadyMap, ctx.Sender().Id)
	if len(a.notReadyMap) == 0 {
		log.Errorf("%v all comments exec ready\n", ctx.Self().Id)
		ctx.RequestWithCustomSender(a.parent, &msg.ArticlesListObReadyMsg{}, a.pid)
	}
}

func (a *ArticlesListOb) procArticlesListTaskMsg(ctxMsg *msg.ArticlesListTaskMsg, ctx actor.Context) {
	if ctxMsg == nil {
		log.Errorf("%v recv empty msg: %T\n", ctx.Self().Id, ctxMsg)
		return
	}

	proxyAddr, err := proxypool.GlobalProxyPool.GetHttpsProxy()
	if err != nil {
		log.Errorf("%v get proxy error: %v\n", ctx.Self().Id, err)
		return
	}

	articlesList, err := getter.ArticlesListGetter(proxyAddr, ctxMsg.Target)
	if err != nil {
		log.Errorf("%v get articles list by %v error: %v\n", ctx.Self().Id, proxyAddr, err)
		return
	}

	selfCommentsExecutorNum := len(a.children)
	aidList := make([][]int64, selfCommentsExecutorNum)

	i := 0
	for _, article := range articlesList {
		aidList[i%selfCommentsExecutorNum] = append(aidList[i%selfCommentsExecutorNum], article.ArticleID)
		i++
	}
	for i, pid := range a.children {
		ctx.RequestWithCustomSender(pid, &msg.CommentsTaskMsg{Aids: aidList[i]}, a.pid)
	}
}
