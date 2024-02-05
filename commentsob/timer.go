package commentsob

import (
	"time"

	"github.com/WangHongshuo/acfun_comments_observer_backend/internal/util"
)

func (c *CommentsOb) startObserveNextArticleTimer() {
	c.timer.SendOnce(util.GetRandomDuration(c.config.PerArticleMinDelay, c.config.PerArticleMaxDelay, int64(c.instId)),
		c.pid, &observeNextArticle{})
}

func (c *CommentsOb) startObserveNextCommentsPageTimer(msg *observeNextCommentsPage) {
	c.timer.SendOnce(util.GetRandomDuration(c.config.PerCommentsPageMinDelay, c.config.PerCommentsPageMaxDelay,
		int64(c.instId)), c.pid, msg)
}

func (c *CommentsOb) startRetryTimer(msg *observeNextCommentsPage) {
	log.Errorf("%v start retry: %v", c.pid.Id, c.retryCount)
	c.timer.SendOnce(time.Duration(c.config.RetryInterval)*time.Second, c.pid, msg)
}
