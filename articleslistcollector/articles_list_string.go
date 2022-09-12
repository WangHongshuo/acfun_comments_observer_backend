package articleslistcollector

import (
	"bytes"
	"fmt"
)

// For Test
func (a ArticlesList) String() string {
	var buff bytes.Buffer
	for i := range a.Articles {
		buff.WriteString("----------\n")
		buff.WriteString(a.Articles[i].String())
	}
	return buff.String()
}

// For Test
func (a Article) String() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("ArticleID: %v\n", a.ArticleID))
	buff.WriteString(fmt.Sprintf("Title: %v\n", a.Title))
	buff.WriteString(fmt.Sprintf("CommentCount: %v\n", a.CommentCount))
	buff.WriteString(fmt.Sprintf("UserID: %v\n", a.UserID))
	buff.WriteString(fmt.Sprintf("UserName: %v\n", a.UserName))
	buff.WriteString(fmt.Sprintf("RealmID: %v\n", a.RealmID))
	return buff.String()
}