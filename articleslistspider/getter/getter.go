package getter

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/util"
)

var articlesListHeaderTemplate = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39",
	"accept":          "application/json, text/plain, */*",
	"accept-encoding": "gzip, deflate, br",
	"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	"origin":          "https://www.acfun.cn",
	"referer":         "",
	"content-type":    "application/x-www-form-urlencoded",
}

func addArticlesListPageHeader(req *http.Request, referer string) {
	for k, v := range articlesListHeaderTemplate {
		req.Header.Add(k, v)
	}
	req.Header.Set("referer", referer)
}

func getUrlPayload(realmIds []string) url.Values {
	v := url.Values{}
	v.Add("cursor", "first_page")
	v.Add("onlyOriginal", "false")
	v.Add("limit", "10")
	v.Add("sortType", "lastCommentTime")
	v.Add("timeRange", "all")

	for _, id := range realmIds {
		v.Add("realmId", id)
	}
	return v
}

func articlesListGetter(porxyAddr string, articleUrl cfg.ArticleUrlConfig) ([]byte, error) {
	client, err := util.NewHttpClient(porxyAddr)
	if err != nil {
		return nil, err
	}

	payload := getUrlPayload(articleUrl.RealmId).Encode()
	req, err := http.NewRequest("POST", cfg.GlobalConfig.ArticlesRequestUrl, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	addArticlesListPageHeader(req, articleUrl.Referer)

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
