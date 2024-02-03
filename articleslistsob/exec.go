package articleslistsob

import (
	"fmt"

	"github.com/WangHongshuo/acfun_comments_observer_backend/articleslistsob/getter"
	"github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/logger"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/proxypool"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/scheduler"
)

const actorName = "ArticlesListOb"

var log = logger.NewLogger(actorName)

type ArticlesListOb struct {
	pid    *actor.PID
	parent *actor.PID
	instId int
	ctx    actor.Context
	timer  *scheduler.TimerScheduler
	config cfg.ObserverConfig

	children       []*actor.PID
	notReadyMap    map[string]struct{}
	notFinishedMap map[string]struct{}

	observeConfig cfg.ArticleUrlConfig
}

func (a *ArticlesListOb) Receive(ctx actor.Context) {
	log.Infof("%v recv msg: %T", ctx.Self().Id, ctx.Message())

	switch ctxMsg := ctx.Message().(type) {
	case *actor.Started:
		a.init(ctx)
	case *msg.ResourceReadyMsg:
		a.initResource(ctx)
	case *msg.ObserveArticlesListTaskMsg:
		a.procArticlesListTaskMsg(ctxMsg)
	case *msg.CommentsObReadyMsg:
		a.procCommentsObReadyMsg(ctx)
	case *msg.ObserveCommentsTaskFinishedMsg:
		a.procObserveCommentsTaskFinishedMsg(ctx)
	default:
		log.Infof("%v recv unknow msg: %T", ctx.Self().Id, ctxMsg)
	}
}

func (a *ArticlesListOb) procCommentsObReadyMsg(ctx actor.Context) {
	delete(a.notReadyMap, ctx.Sender().Id)
	if len(a.notReadyMap) == 0 {
		log.Errorf("%v all comments exec ready", ctx.Self().Id)
		ctx.RequestWithCustomSender(a.parent, &msg.ArticlesListObReadyMsg{}, a.pid)
	}
}

func (a *ArticlesListOb) procArticlesListTaskMsg(ctxMsg *msg.ObserveArticlesListTaskMsg) {
	if ctxMsg == nil {
		log.Errorf("%v recv empty msg: %T", a.pid.Id, ctxMsg)
		return
	}
	a.observeConfig = ctxMsg.Target
	if err := a.observeArticlesListAndDispatchToChildren(a.observeConfig); err != nil {
		log.Errorf("%v observe articles list error: %v", a.pid.Id, err)
	}
}

func (a *ArticlesListOb) observeArticlesListAndDispatchToChildren(config cfg.ArticleUrlConfig) error {
	proxyAddr, err := proxypool.GlobalProxyPool.GetHttpsProxy()
	if err != nil {
		return fmt.Errorf("%v get proxy error: %v", a.pid.Id, err)
	}

	articlesList, err := getter.ArticlesListGetter(proxyAddr, config)
	if err != nil {
		return fmt.Errorf("%v get articles list by %v error: %v", a.pid.Id, proxyAddr, err)
	}

	selfCommentsExecutorNum := len(a.children)
	aidList := make([][]int64, selfCommentsExecutorNum)

	i := 0
	for _, article := range articlesList {
		aidList[i%selfCommentsExecutorNum] = append(aidList[i%selfCommentsExecutorNum], article.ArticleID)
		i++
	}
	for i, pid := range a.children {
		a.ctx.RequestWithCustomSender(pid, &msg.ObserveCommentsTaskMsg{Aids: aidList[i]}, a.pid)
		a.notFinishedMap[pid.Id] = struct{}{}
	}
	return nil
}

func (a *ArticlesListOb) procObserveCommentsTaskFinishedMsg(ctx actor.Context) {
	delete(a.notFinishedMap, ctx.Sender().Id)
	if len(a.notFinishedMap) == 0 {
		a.startIdleTimer()
	}
}
