package wooliesx

import (
	"net"
	"net/http"
	"time"
)

//SDK is a container to interact with the WooliesX API
type SDK struct {
	client *http.Client
	url    string
}

// New creates and returns a connection to the WookiesX API
func New(baseURL string) *SDK {
	return &SDK{
		client: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   time.Second * 10,
					KeepAlive: time.Second * 10,
				}).Dial,
				TLSHandshakeTimeout:   time.Second * 10,
				ResponseHeaderTimeout: time.Second * 10,
				ExpectContinueTimeout: time.Second * 10,
				MaxIdleConns:          10,
				MaxIdleConnsPerHost:   10,
				IdleConnTimeout:       time.Second * 10,
			},
		},
		url: baseURL,
	}
}
