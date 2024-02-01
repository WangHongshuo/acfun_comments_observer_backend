package getter

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	_ "github.com/WangHongshuo/acfuncommentsspider-go/cfg"
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

	ret, err := ArticlesListGetter("http://159.138.57.97:4003", cfg.GlobalConfig.ArticleUrl[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Logf("\n%+v", ret)
	assert.Equal(t, true, len(ret) > 0)
}
