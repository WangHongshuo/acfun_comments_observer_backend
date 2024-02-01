package main

import (
	_ "github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/WangHongshuo/acfuncommentsspider-go/proxypool"
	"github.com/WangHongshuo/acfuncommentsspider-go/spiderctrl"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

var log = logger.NewLogger("Root")

func main() {
	log.Info("main start, spawn actors")

	proxypool.Init()
	system := actor.NewActorSystem()
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &spiderctrl.SpiderController{} })
	if _, err := context.SpawnNamed(props, "Root"); err != nil {
		panic(err)
	}

	log.Info("spawn succ")
	console.ReadLine()
}
