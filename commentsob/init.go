package commentsob

import (
	"github.com/WangHongshuo/acfun_comments_observer_backend/dao"
	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/util"
	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
	"github.com/asynkron/protoactor-go/actor"
)

func (c *CommentsOb) init(ctx actor.Context) error {
	log.Infof("CommentsOb init")

	c.pid = ctx.Self()
	c.instId, _ = util.GetInstIdFromPid(c.pid)

	return nil
}

func (c *CommentsOb) initResource(ctx actor.Context) {
	c.db = dao.GlobalPgDb

	ctx.RequestWithCustomSender(ctx.Sender(), &msg.CommentsObReadyMsg{}, c.pid)
}
