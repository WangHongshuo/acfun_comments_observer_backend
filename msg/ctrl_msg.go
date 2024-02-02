package msg

import "github.com/WangHongshuo/acfun_comments_observer_backend/cfg"

type ArticlesListTaskMsg struct {
	Target cfg.ArticleUrlConfig
}

type CommentsTaskMsg struct {
	Aids []int64
}
