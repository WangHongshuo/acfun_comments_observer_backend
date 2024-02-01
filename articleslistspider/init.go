package articleslistspider

import (
	"fmt"

	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/WangHongshuo/acfuncommentsspider-go/commentsspider"
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/util"
	"github.com/WangHongshuo/acfuncommentsspider-go/msg"
	"github.com/asynkron/protoactor-go/actor"
)

func (a *ArticlesListExecutor) init(ctx actor.Context) error {
	log.Infof("ArticlesListExecutor init")
	a.pid = ctx.Self()
	a.parent = ctx.Parent()
	a.instId, _ = util.GetInstIdFromPid(a.pid)

	a.spawnCommentsExecutors(ctx)

	return nil
}

func (a *ArticlesListExecutor) spawnCommentsExecutors(ctx actor.Context) error {
	props := actor.PropsFromProducer(func() actor.Actor { return &commentsspider.CommentsExecutor{} })
	commentsConfig := cfg.GlobalConfig.Spiders["comments"]
	commentsExecSpec := commentsConfig.Spec
	articlesListExecSpec := cfg.GlobalConfig.Spiders["articles"].Spec
	prefix := commentsConfig.Prefix + util.ActorNameSuffixFmt
	a.notReadyMap = make(map[string]struct{}, 0)

	start, end := util.CalculateChildrenIdRangeFromInstSpec(articlesListExecSpec, commentsExecSpec, a.instId)

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

func (a *ArticlesListExecutor) initResource(ctx actor.Context) {
	for _, pid := range a.children {
		ctx.RequestWithCustomSender(pid, &msg.ResourceReadyMsg{}, a.pid)
	}
}
