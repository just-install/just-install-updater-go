package util

import (
	"net"
	"net/http"
	"time"
)

// HTTPClient is the HTTP client used by all jiup-go requests. By default, it
// is the same as http.DefaultClient and http.DefaultTransport, but with limits
// more suitable for jiup-go's use.
var HTTPClient *http.Client = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second, // lower, as it isn't really reasonable to wait 30s
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          20, // lower than the limit from the default client to prevent random failures on bad CI nodes
		MaxConnsPerHost:       http.DefaultMaxIdleConnsPerHost,
		IdleConnTimeout:       15 * time.Second, // lower, since we won't really re-use connections after that, especially for Versioners
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: time.Second,
	},
}

// TODO: HTTP helpers (raw, JSON, HTML) (check headers, status)
