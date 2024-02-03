package commentsob

import (
	"github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/WangHongshuo/acfun_comments_observer_backend/dao"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/util"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/scheduler"
)

func (c *CommentsOb) init(ctx actor.Context) error {
	log.Infof("CommentsOb init")

	c.pid = ctx.Self()
	c.parent = ctx.Parent()
	c.instId, _ = util.GetInstIdFromPid(c.pid)
	c.ctx = ctx
	c.timer = scheduler.NewTimerScheduler(ctx)
	c.config = cfg.GlobalConfig.Observers["comments"]

	return nil
}

func (c *CommentsOb) initResource(ctx actor.Context) {
	c.db = dao.GlobalPgDb

	ctx.RequestWithCustomSender(ctx.Sender(), &msg.CommentsObReadyMsg{}, c.pid)
}
