package commentscollector

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"time"
)

var commentsHeaderTemplate = map[string]string{
	"Host":            "www.acfun.cn",
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39",
	"Accept":          "application/json, text/plain, */*",
	"Accept-Encoding": "gzip, deflate, br",
	"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	"Connection":      "keep-alive",
	"DNT":             "1",
	"Referer":         "https://www.acfun.cn/a/ac34857021",
	"TE":              "Trailers",
}

func addCommentsPageHeader(req *http.Request) {
	for k, v := range commentsHeaderTemplate {
		req.Header.Add(k, v)
	}
}

func commentsGetter() []byte {
	url := fmt.Sprintf("https://www.acfun.cn/rest/pc-direct/comment/listByFloor?sourceId=%v&sourceType=3&page=%v&pivotCommentId=0&newPivotCommentId=0&t=%v&supportZtEmot=true",
		"34857021", 1, time.Now().UnixMilli())
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	addCommentsPageHeader(req)
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("err: %v, code: %v", err, resp.StatusCode)
		return nil
	}
	// gzip解压
	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		if gr != nil {
			gr.Close()
		}
		return nil
	}
	defer gr.Close()
	json, err := io.ReadAll(gr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return json
}
