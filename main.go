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

var system = actor.NewActorSystem()
var log = logger.RootLogger

func main() {
	log.Info("spawn actors")

	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &spiderctrl.SpiderController{} })
	context.Spawn(props)

	log.Info("spawn succ")
	console.ReadLine()
}
