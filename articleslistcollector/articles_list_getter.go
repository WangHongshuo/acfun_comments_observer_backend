package articleslistcollector

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var articlesListHeaderTemplate = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39",
	"accept":          "application/json, text/plain, */*",
	"accept-encoding": "gzip, deflate, br",
	"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	"origin":          "https://www.acfun.cn",
	"referer":         "https://www.acfun.cn/v/list110/index.htm",
	"content-type":    "application/x-www-form-urlencoded",
}

func addArticlesListPageHeader(req *http.Request) {
	for k, v := range articlesListHeaderTemplate {
		req.Header.Add(k, v)
	}
}

func getUrlPayload() url.Values {
	v := url.Values{}
	v.Add("cursor", "first_page")
	v.Add("onlyOriginal", "false")
	v.Add("limit", "10")
	v.Add("sortType", "hotScore")
	v.Add("timeRange", "oneDay")
	v.Add("realmId", "5")  // 杂谈
	v.Add("realmId", "22") // 体育
	v.Add("realmId", "28") // 新闻资讯
	v.Add("realmId", "3")  // 影视
	return v
}

func articlesListGetter() []byte {
	url_path := "https://www.acfun.cn/rest/pc-direct/article/feed"
	client := &http.Client{}
	payload := getUrlPayload().Encode()
	req, err := http.NewRequest("POST", url_path, strings.NewReader(payload))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	addArticlesListPageHeader(req)

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
	json, err := ioutil.ReadAll(gr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return json
}
