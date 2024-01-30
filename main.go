package main

import (
	// 必须先初始化cfg
	_ "github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	// logger依赖cfg
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/WangHongshuo/acfuncommentsspider-go/spiderctrl"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

var log = logger.NewLogger("Root")

func main() {
	log.Info("main start, spawn actors")

	system := actor.NewActorSystem()
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &spiderctrl.SpiderController{} })
	if _, err := context.SpawnNamed(props, "Root"); err != nil {
		panic(err)
	}

	log.Info("spawn succ")
	console.ReadLine()
}
