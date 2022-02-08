// +build !go1.7

package oss

// Bucket implements the operations of object.
type Bucket struct {
	Client     Client
	BucketName string
}


// Private
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

	resp, err := bucket.Client.Conn.Do(method, bucket.BucketName, objectName,
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

	resp, err := client.Conn.Do(method, bucketName, "", params, headers, data, 0, nil)

	// get response header
	respHeader, _ := FindOption(options, responseHeader, nil)
	if respHeader != nil {
		pRespHeader := respHeader.(*http.Header)
		*pRespHeader = resp.Headers
	}

	return resp, err
}
