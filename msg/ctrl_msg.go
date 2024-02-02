package msg

import "github.com/WangHongshuo/acfuncommentsspider-go/cfg"

type ArticlesListTaskMsg struct {
	Target cfg.ArticleUrlConfig
}

type CommentsTaskMsg struct {
	Aids []int64
}
