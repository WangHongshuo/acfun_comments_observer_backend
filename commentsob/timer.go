package commentsob

import "github.com/WangHongshuo/acfun_comments_observer_backend/internal/util"

func (c *CommentsOb) startObserveNextArticleTimer() {
	c.timer.SendOnce(util.GetRandomDuration(c.perArticleMinDelay, c.perArticleMaxDelay, int64(c.instId)),
		c.pid, &observeNextArticle{})
}

func (c *CommentsOb) startObserveNextCommentsPageTimer(msg *observeNextCommentsPage) {
	c.timer.SendOnce(util.GetRandomDuration(c.perCommentsPageMinDelay, c.perCommentsPageMinDelay,
		int64(c.instId)), c.pid, msg)
}
