package main

import (
	_ "github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/logger"
	"github.com/WangHongshuo/acfun_comments_observer_backend/obctrl"
	"github.com/WangHongshuo/acfun_comments_observer_backend/proxypool"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

var log = logger.NewLogger("Root")

func main() {
	log.Info("main start, spawn actors")

	proxypool.Init()
	system := actor.NewActorSystem()
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &obctrl.ObController{} })
	if _, err := context.SpawnNamed(props, "Root"); err != nil {
		panic(err)
	}

	log.Info("spawn succ")
	console.ReadLine()
}
