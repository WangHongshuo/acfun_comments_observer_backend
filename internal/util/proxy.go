package util

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

func NewHttpClient(proxyAddr string) (*http.Client, error) {
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		return nil, err
	}

	netTransport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
		Dial: func(network, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(network, addr, time.Second*time.Duration(10))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}, nil
}
