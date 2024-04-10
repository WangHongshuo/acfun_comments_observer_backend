package main

import (
	// import config first
	_ "github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/logger"
	"github.com/WangHongshuo/acfun_comments_observer_backend/obctrl"
	"github.com/WangHongshuo/acfun_comments_observer_backend/proxypool"
	"github.com/asynkron/protoactor-go/actor"
)

var log = logger.NewLogger("Main")

func main() {
	defer logger.WatiLogger()
	log.Info("main start, spawn actors")

	proxypool.Init()
	system := actor.NewActorSystem()
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &obctrl.ObController{} })
	if _, err := context.SpawnNamed(props, "ObCtrl"); err != nil {
		panic(err)
	}

	log.Info("spawn succ")
	select {}
}
