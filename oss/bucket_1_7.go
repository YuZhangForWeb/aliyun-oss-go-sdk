// +build go1.7

package oss

import (
	"context"
	"io"
	"net/http"
)

// Bucket implements the operations of object.
type Bucket struct {
	Client     Client
	BucketName string
	ctx   context.Context
}

// WithContext support go1.7
func (bucket Bucket)WithContext(ctx context.Context) Bucket  {
	bucket.ctx = ctx
	return bucket
}


// Private with ctx support
func (bucket Bucket) do(method, objectName string, params map[string]interface{}, options []Option,
	data io.Reader, listener ProgressListener) (*Response, error) {
	headers := make(map[string]string)
	err := handleOptions(headers, options)
	if err != nil {
		return nil, err
	}

	err = CheckBucketName(bucket.BucketName)
	if len(bucket.BucketName) > 0 && err != nil {
		return nil, err
	}

	conn := bucket.Client.Conn.WithContext(bucket.ctx)
	resp, err := conn.Do(method, bucket.BucketName, objectName,
		params, headers, data, 0, listener)

	// get response header
	respHeader, _ := FindOption(options, responseHeader, nil)
	if respHeader != nil && resp != nil {
		pRespHeader := respHeader.(*http.Header)
		*pRespHeader = resp.Headers
	}

	return resp, err
}

// Private
func (client Client) do(method, bucketName string, params map[string]interface{},
	headers map[string]string, data io.Reader, options ...Option) (*Response, error) {
	err := CheckBucketName(bucketName)
	if len(bucketName) > 0 && err != nil {
		return nil, err
	}

	// option headers
	addHeaders := make(map[string]string)
	err = handleOptions(addHeaders, options)
	if err != nil {
		return nil, err
	}

	// merge header
	if headers == nil {
		headers = make(map[string]string)
	}

	for k, v := range addHeaders {
		if _, ok := headers[k]; !ok {
			headers[k] = v
		}
	}

	conn := client.Conn.WithContext(client.ctx)
	resp, err := conn.Do(method, bucketName, "", params, headers, data, 0, nil)

	// get response header
	respHeader, _ := FindOption(options, responseHeader, nil)
	if respHeader != nil {
		pRespHeader := respHeader.(*http.Header)
		*pRespHeader = resp.Headers
	}

	return resp, err
}
