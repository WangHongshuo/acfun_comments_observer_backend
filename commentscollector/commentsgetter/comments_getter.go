package commentsgetter

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/WangHongshuo/acfuncommentsspider-go/internal/util"
)

var commentsHeaderTemplate = map[string]string{
	"Host":            "www.acfun.cn",
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39",
	"Accept":          "application/json, text/plain, */*",
	"Accept-Encoding": "gzip, deflate, br",
	"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	"Connection":      "keep-alive",
	"DNT":             "1",
	"Referer":         "",
	"TE":              "Trailers",
}

func addCommentsPageHeader(req *http.Request) {
	for k, v := range commentsHeaderTemplate {
		req.Header.Add(k, v)
	}
}

func commentsGetter(porxyAddr, aid string) ([]byte, error) {
	url := fmt.Sprintf("https://www.acfun.cn/rest/pc-direct/comment/listByFloor?sourceId=%v&sourceType=3&page=%v&pivotCommentId=0&newPivotCommentId=0&t=%v&supportZtEmot=true",
		aid, 1, time.Now().UnixMilli())

	client, err := util.NewHttpClient(porxyAddr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	addCommentsPageHeader(req)
	req.Header.Set("Referer", fmt.Sprintf("https://www.acfun.cn/a/ac%v", aid))
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("http resp is nil")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %v", resp.StatusCode)
	}

	// gzip解压
	gr, err := gzip.NewReader(resp.Body)

	if err != nil {
		if gr != nil {
			gr.Close()
		}
		return nil, err
	}

	defer gr.Close()
	json, err := io.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return json, nil
}
