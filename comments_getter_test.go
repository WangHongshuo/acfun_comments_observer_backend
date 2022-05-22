package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Comments_Getter(t *testing.T) {
	buff := bytes.NewBuffer(commentsGetter())

	comments := &CommentsJsonResponse{}
	decoder := json.NewDecoder(buff)
	err := decoder.Decode(comments)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Logf("\n%v", comments)
}
