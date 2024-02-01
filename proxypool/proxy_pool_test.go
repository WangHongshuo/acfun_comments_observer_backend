package proxypool

import (
	"log"
	"testing"

	_ "github.com/WangHongshuo/acfuncommentsspider-go/cfg"
	"github.com/stretchr/testify/assert"
)

func Test_ProxyPool(t *testing.T) {
	Init()

	proxyAddr, err := GlobalProxyPool.GetHttpsProxy()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", proxyAddr)
	log.Println("proxyAddr:", proxyAddr)
}
