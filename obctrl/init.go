package obctrl

import (
	"fmt"

	"github.com/WangHongshuo/acfun_comments_observer_backend/articleslistsob"
	"github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/dao"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/util"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/asynkron/protoactor-go/actor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (s *ObController) init(ctx actor.Context) error {
	log.Infof("ObController init")
	s.pid = ctx.Self()

	s.spawnArticlesListExecutors(ctx)
	s.initGlobalPgDb(ctx)

	return nil
}

func (s *ObController) spawnArticlesListExecutors(ctx actor.Context) error {

	props := actor.PropsFromProducer(func() actor.Actor { return &articleslistsob.ArticlesListOb{} })
	config := cfg.GlobalConfig.Observers["articles"]
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

func (s *ObController) initGlobalPgDb(ctx actor.Context) error {
	dbConfig := cfg.GlobalConfig.Database
	actorsSpec := cfg.GlobalConfig.Observers["comments"].Spec + cfg.GlobalConfig.Observers["articles"].Spec

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
		ctx.RequestWithCustomSender(pid, &msg.ResourceReadyMsg{}, s.pid)
	}

	return nil
}
