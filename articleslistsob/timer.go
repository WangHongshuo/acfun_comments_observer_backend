package articleslistsob

import (
	"time"

	"github.com/WangHongshuo/acfun_comments_observer_backend/msg"
)

func (a *ArticlesListOb) startIdleTimer() {
	idleTime := time.Duration(a.config.IdleTime) * time.Minute
	log.Infof("%v all children finished task, start idle: %v", a.pid.Id, idleTime)
	a.timer.SendOnce(idleTime, a.pid, &msg.ObserveArticlesListTaskMsg{Target: a.observeConfig})
}
