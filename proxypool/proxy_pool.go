package proxypool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/WangHongshuo/acfuncommentsspider-go/cfg"
)

var GlobalProxyPool *proxyPool

type proxyPool struct {
	client           *http.Client
	getHttpsPorxyUrl string
}

type proxyPoolResult struct {
	Proxy string `json:"proxy"`
}

func (p *proxyPool) GetHttpsProxy() (string, error) {
	if cfg.GlobalConfig.ProxyServer.CustomProxy != "" {
		return fmt.Sprintf("http://%v", cfg.GlobalConfig.ProxyServer.CustomProxy), nil
	}
	if p == nil {
		return "", fmt.Errorf("*proxyPool is nil")
	}
	if p.client == nil {
		return "", fmt.Errorf("proxyPool.client is nil")
	}

	req, err := http.NewRequest(http.MethodGet, p.getHttpsPorxyUrl, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	resp, err := p.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", fmt.Errorf("error making http request: %v", err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	buff := bytes.NewBuffer(bodyBytes)
	ret := &proxyPoolResult{}

	decoder := json.NewDecoder(buff)
	err = decoder.Decode(ret)
	if err != nil {
		return "", fmt.Errorf("error decoding json response: %v", err)
	}

	if ret.Proxy == "" {
		return ret.Proxy, fmt.Errorf("no proxy resource")
	}

	return fmt.Sprintf("http://%v", ret.Proxy), nil
}
