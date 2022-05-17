package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	buff := bytes.NewBuffer(commentsGetter())

	comments := &CommentsJsonResponse{}
	decoder := json.NewDecoder(buff)
	err := decoder.Decode(comments)
	if err != nil {
		fmt.Println(err)
		return
	}
	// CommentIDs是保序的的
	for _, v := range comments.CommentIDS {
		if v, ok := comments.CommentsMap[fmt.Sprintf("c%v", v)]; ok {
			fmt.Printf("Floor: %v, Cid: %v, UserName: %v, Uid: %v, Time: %v, Comment:%v\n", v.Floor, v.Cid, v.UserName, v.UserID, v.Timestamp, v.Content)
		}
	}

}
