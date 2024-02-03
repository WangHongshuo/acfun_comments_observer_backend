package commentsob

import "github.com/WangHongshuo/acfun_comments_observer_backend/dao/model"

func (c *CommentsOb) commitAll() {
	c.commitCommentsDataToDb(c.commentsCache)
	c.commitArticleDataToDb([]model.Article{c.articleCache})
}

func (c *CommentsOb) commitCommentsDataToDb(data []model.Comment) {
	c.db.Save(data)
}

func (c *CommentsOb) commitArticleDataToDb(data []model.Article) {
	c.db.Save(data)
}
