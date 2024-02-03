package msg

import "github.com/WangHongshuo/acfun_comments_observer_backend/cfg"

type ObserveArticlesListTaskMsg struct {
	Target cfg.ArticleUrlConfig
}

type ObserveCommentsTaskMsg struct {
	Aids []int64
}

type ObserveCommentsTaskFinishedMsg struct{}
