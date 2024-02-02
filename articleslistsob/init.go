package articleslistsob

import (
	"fmt"

	"github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/commentsob"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/util"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/asynkron/protoactor-go/actor"
)

func (a *ArticlesListOb) init(ctx actor.Context) error {
	log.Infof("ArticlesListOb init")
	a.pid = ctx.Self()
	a.parent = ctx.Parent()
	a.instId, _ = util.GetInstIdFromPid(a.pid)
	a.ctx = ctx

	a.spawnCommentsObs(ctx)

	return nil
}

func (a *ArticlesListOb) spawnCommentsObs(ctx actor.Context) error {
	props := actor.PropsFromProducer(func() actor.Actor { return &commentsob.CommentsOb{} })
	commentsObConfig := cfg.GlobalConfig.Observers["comments"]
	commentsObSpec := commentsObConfig.Spec
	articlesListObSpec := cfg.GlobalConfig.Observers["articles"].Spec
	prefix := commentsObConfig.Prefix + util.ActorNameSuffixFmt
	a.notReadyMap = make(map[string]struct{}, 0)

	start, end := util.CalculateChildrenIdRangeFromInstSpec(articlesListObSpec, commentsObSpec, a.instId)

	for i := start; i <= end; i++ {
		name := fmt.Sprintf(prefix, i)
		pid, err := ctx.SpawnNamed(props, name)
		if err != nil {
			log.Errorf("SpawnNamed %v failed, error: %s", name, err.Error())
			continue
		}
		a.children = append(a.children, pid)
		a.notReadyMap[pid.Id] = struct{}{}
	}

	return nil
}

func (a *ArticlesListOb) initResource(ctx actor.Context) {
	for _, pid := range a.children {
		ctx.RequestWithCustomSender(pid, &msg.ResourceReadyMsg{}, a.pid)
	}
}
