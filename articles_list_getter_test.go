package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Articles_Getter(t *testing.T) {
	buff := bytes.NewBuffer(articlesListGetter())
	articlesList := &ArticlesList{}

	decoder := json.NewDecoder(buff)
	err := decoder.Decode(articlesList)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Logf("%+v", articlesList)
}
