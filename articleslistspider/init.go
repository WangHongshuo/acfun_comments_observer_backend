package articleslistspider

import (
	"fmt"

	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/WangHongshuo/acfuncommentsspider-go/commentsspider"
	"github.com/asynkron/protoactor-go/actor"
)

func (a *ArticlesListExecutor) init(ctx actor.Context) error {
	log.Infof("SpiderController init")
	a.pid = ctx.Self()

	a.spawnCommentsExecutors(ctx)

	log.Infof("%+v", a)
	return nil
}

func (a *ArticlesListExecutor) spawnCommentsExecutors(ctx actor.Context) error {

	props := actor.PropsFromProducer(func() actor.Actor { return &commentsspider.CommentsExecutor{} })
	config := cfg.GlobalConfig.Spiders["comments"]
	prefix := config.Prefix + "%v"

	for i := 0; i < config.Spec; i++ {
		name := fmt.Sprintf(prefix, i)
		pid, err := ctx.SpawnNamed(props, name)
		if err != nil {
			log.Errorf("SpawnNamed %v failed, error: %s", name, err.Error())
			continue
		}
		a.children = append(a.children, pid)
	}

	return nil
}
