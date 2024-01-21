package articleslistgetter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalJsonBody(t *testing.T) {
	f, err := os.Open("./articles_list.json")
	assert.Equal(t, nil, err)
	defer f.Close()

	b, err := io.ReadAll(f)
	assert.Equal(t, nil, err)

	articlesList := &ArticlesList{}
	err = json.Unmarshal(b, articlesList)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, len(articlesList.Articles) > 0)
	t.Logf("\n%+v", articlesList)
}

func Test_ArticlesGetter(t *testing.T) {
	ret, err := articlesListGetter("http://159.138.57.97:4003")
	if err != nil {
		fmt.Println(err)
		return
	}

	buff := bytes.NewBuffer(ret)
	articlesList := &ArticlesList{}

	decoder := json.NewDecoder(buff)
	err = decoder.Decode(articlesList)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Logf("\n%+v", articlesList)
	assert.Equal(t, true, len(articlesList.Articles) > 0)
}
