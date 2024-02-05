package articleslistsob

import (
	"time"

	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
)

func (a *ArticlesListOb) startIdleTimer() {
	idleTime := time.Duration(a.config.IdleTime) * time.Minute
	log.Infof("%v start idle: %v", a.pid.Id, idleTime)
	a.timer.SendOnce(idleTime, a.pid, &msg.ObserveArticlesListTaskMsg{Target: a.observeConfig})
}

func (a *ArticlesListOb) startRetryTimer() {
	log.Errorf("%v start retry: %v", a.pid.Id, a.retryCount)
	a.timer.SendOnce(time.Duration(a.config.RetryInterval)*time.Second, a.pid, &msg.ObserveArticlesListTaskMsg{Target: a.observeConfig})
}
