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
	c.instId, _ = util.GetInstIdFromPid(c.pid)
	c.ctx = ctx
	c.timer = scheduler.NewTimerScheduler(ctx)

	return nil
}

func (c *CommentsOb) initResource(ctx actor.Context) {
	c.db = dao.GlobalPgDb
	obConfig := cfg.GlobalConfig.Observers["comments"]
	c.perArticleMinDelay = obConfig.PerArticleMinDelay
	c.perArticleMaxDelay = obConfig.PerArticleMaxDelay
	c.perCommentsPageMinDelay = obConfig.PerCommentsPageMinDelay
	c.perCommentsPageMaxDelay = obConfig.PerCommentsPageMaxDelay

	ctx.RequestWithCustomSender(ctx.Sender(), &msg.CommentsObReadyMsg{}, c.pid)
}
