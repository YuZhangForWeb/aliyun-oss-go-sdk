// +build go1.7

package oss

import (
	"context"
	"net/http"
)

// Client OSS client
type Client struct {
	Config     *Config      // OSS client configuration
	Conn       *Conn        // Send HTTP request
	HTTPClient *http.Client //http.Client to use - if nil will make its own
	ctx context.Context
}

// WithContext support go1.7 context
func (client Client)WithContext(ctx context.Context) Client  {
	client.ctx = ctx
	return client
}
