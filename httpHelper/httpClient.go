package httpHelper

import (
	"crypto/tls"
	"net/http"
	"sync"
)

var httpClient *http.Client
var httpClientOnce sync.Once

func HttpClientInstance() *http.Client {
	httpClientOnce.Do(func() {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: tr}
	})
	return httpClient
}
