package proxypool

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
)

func Init() {
	config := cfg.GlobalConfig.ProxyServer
	GlobalProxyPool = &proxyPool{
		client:           &http.Client{},
		getHttpsPorxyUrl: fmt.Sprintf("http://%v:%v/get/?type=https", config.Host, config.Port),
	}
	log.Printf("init proxy pool succ")
}
