package spiderctrl

import (
	"fmt"

	"github.com/WangHongshuo/acfuncommentsspider-go/articleslistspider"
	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/asynkron/protoactor-go/actor"
)

func (s *SpiderController) init(ctx actor.Context) error {
	log.Infof("SpiderController init")
	s.pid = ctx.Self()

	s.spawnArticlesListExecutors(ctx)

	log.Infof("%+v", s)
	return nil
}

func (s *SpiderController) spawnArticlesListExecutors(ctx actor.Context) error {

	props := actor.PropsFromProducer(func() actor.Actor { return &articleslistspider.ArticlesListExecutor{} })
	config := cfg.GlobalConfig.Spiders["articles"]
	prefix := config.Prefix + "%v"

	for i := 0; i < config.Spec; i++ {
		name := fmt.Sprintf(prefix, i)
		pid, err := ctx.SpawnNamed(props, name)
		if err != nil {
			log.Errorf("SpawnNamed %v failed, error: %s", name, err.Error())
			continue
		}
		s.children = append(s.children, pid)
	}

	return nil
}
