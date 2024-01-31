package spiderctrl

import (
	"fmt"

	"github.com/WangHongshuo/acfuncommentsspider-go/articleslistspider"
	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/WangHongshuo/acfuncommentsspider-go/dao"
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/util"
	"github.com/WangHongshuo/acfuncommentsspider-go/msg"
	"github.com/asynkron/protoactor-go/actor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (s *SpiderController) init(ctx actor.Context) error {
	log.Infof("SpiderController init")
	s.pid = ctx.Self()

	s.spawnArticlesListExecutors(ctx)
	s.initGlobalPgDb(ctx)

	return nil
}

func (s *SpiderController) spawnArticlesListExecutors(ctx actor.Context) error {

	props := actor.PropsFromProducer(func() actor.Actor { return &articleslistspider.ArticlesListExecutor{} })
	config := cfg.GlobalConfig.Spiders["articles"]
	prefix := config.Prefix + util.ActorNameSuffixFmt

	for i := 1; i <= config.Spec; i++ {
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

func (s *SpiderController) initGlobalPgDb(ctx actor.Context) error {
	dbConfig := cfg.GlobalConfig.Database
	actorsSpec := cfg.GlobalConfig.Spiders["comments"].Spec + cfg.GlobalConfig.Spiders["articles"].Spec

	dsn := fmt.Sprintf("user=%v password=%v port=%v dbname=%v sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.UserName, dbConfig.Password, dbConfig.Port, dbConfig.DbName)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	postgresDb, err := db.DB()
	if err != nil {
		return err
	}
	postgresDb.SetMaxOpenConns(actorsSpec + dbConfig.ReservedConn)
	postgresDb.SetMaxIdleConns(actorsSpec)
	dao.GlobalPgDb = db

	for _, pid := range s.children {
		ctx.Send(pid, &msg.ResourceReadyMsg{})
	}

	return nil
}
