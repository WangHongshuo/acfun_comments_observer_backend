package getter

import (
	"bytes"
	"fmt"
	"time"
)

// For Test
func (a ArticlesList) String() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("----------Total: %v Cursor: %v----------\n", len(a.Articles), a.Cursor))
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
	buff.WriteString(fmt.Sprintf("CreateTime: %v\n", time.UnixMilli(a.CreateTime).Format(time.RFC3339)))
	buff.WriteString(fmt.Sprintf("CommentCount: %v\n", a.CommentCount))
	buff.WriteString(fmt.Sprintf("RealmID: %v\n", a.RealmID))
	return buff.String()
}
