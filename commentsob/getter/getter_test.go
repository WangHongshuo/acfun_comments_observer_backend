package getter

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	_ "github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalJsonBody(t *testing.T) {
	f, err := os.Open("./comments.json")
	assert.Equal(t, nil, err)
	defer f.Close()

	b, err := io.ReadAll(f)
	assert.Equal(t, nil, err)

	comments := &CommentsJsonResponse{}
	err = json.Unmarshal(b, comments)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, len(comments.CommentsMap) > 0)
	t.Logf("\n%v", comments)
}

func Test_CommentsGetter(t *testing.T) {
	ret, totalPage, err := CommentsGetter("http://159.138.57.97:4003", 39658906, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Logf("\nTotalPage: %v", totalPage)
	t.Logf("\n%v", ret)
	assert.Equal(t, true, len(ret) > 0)
	assert.Equal(t, true, totalPage > 0)
}
