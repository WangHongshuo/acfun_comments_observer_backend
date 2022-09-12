package commentscollector

import (
	"bytes"
	"fmt"
	"time"
)

// For Test
func (c CommentsJsonResponse) String() string {
	var buff bytes.Buffer
	// CommentIDs是保序的的
	buff.WriteString(fmt.Sprintf("CurPage: %v, TotalPage: %v, PageSize:%v, TotalCount: %v\n",
		c.CurPage, c.TotalPage, c.PageSize, c.TotalCount))
	for _, v := range c.CommentIDS {
		if v, ok := c.CommentsMap[fmt.Sprintf("c%v", v)]; ok {
			buff.WriteString("----------\n")
			buff.WriteString(v.String())
		}
	}
	return buff.String()
}

// For Test
func (c Comment) String() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("Cid: %v, Floor: %v, UserName: %v, IsUp: %v, Timestamp: %v, SourceID: %v, IsDelete: %v, IsUpDelete: %v\n",
		c.Cid, c.Floor, c.UserName, c.IsUp, time.UnixMilli(c.Timestamp).Format("2006-01-02 15:04:05"), c.SourceID, c.IsDelete, c.IsDelete))
	buff.WriteString(fmt.Sprintf("Content: %v\n", c.Content))
	return buff.String()
}
