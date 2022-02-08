// +build !go1.7

package oss

import "net/http"

// Client OSS client
type Client struct {
	Config     *Config      // OSS client configuration
	Conn       *Conn        // Send HTTP request
	HTTPClient *http.Client //http.Client to use - if nil will make its own
}
