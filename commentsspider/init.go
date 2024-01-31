package commentsspider

import (
	"github.com/WangHongshuo/acfuncommentsspider-go/dao"
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/util"
	"github.com/asynkron/protoactor-go/actor"
)

func (c *CommentsExecutor) init(ctx actor.Context) error {
	log.Infof("CommentsExecutor init")

	c.pid = ctx.Self()
	c.instId, _ = util.GetInstIdFromPid(c.pid)

	return nil
}

func (c *CommentsExecutor) initResource(ctx actor.Context) {
	c.db = dao.GlobalPgDb
}
