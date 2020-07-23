package cfg

import (
	"net/http"
	"time"
)

var httpClient *http.Client

// default package trigger
func initHTTP() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	httpClient = &http.Client{Transport: transport}
	return httpClient
}
