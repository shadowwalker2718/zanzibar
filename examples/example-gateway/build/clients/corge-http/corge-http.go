// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package corgehttpclient

import (
	"context"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"

	zanzibar "github.com/uber/zanzibar/runtime"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/corge-http/module"
	clientsCorgeCorge "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/corge/corge"
)

// Client defines corge-http client interface.
type Client interface {
	HTTPClient() *zanzibar.HTTPClient

	EchoString(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsCorgeCorge.Corge_EchoString_Args,
	) (string, map[string]string, error)
}

// corgeHTTPClient is the http client.
type corgeHTTPClient struct {
	clientID   string
	httpClient *zanzibar.HTTPClient

	calleeHeader string
	callerHeader string
	callerName   string
	calleeName   string
}

// NewClient returns a new http client.
func NewClient(deps *module.Dependencies) Client {
	ip := deps.Default.Config.MustGetString("sidecarRouter.default.http.ip")
	port := deps.Default.Config.MustGetInt("sidecarRouter.default.http.port")
	callerHeader := deps.Default.Config.MustGetString("sidecarRouter.default.http.callerHeader")
	calleeHeader := deps.Default.Config.MustGetString("sidecarRouter.default.http.calleeHeader")
	callerName := deps.Default.Config.MustGetString("serviceName")
	calleeName := deps.Default.Config.MustGetString("clients.corge-http.serviceName")
	baseURL := fmt.Sprintf("http://%s:%d", ip, port)
	timeout := time.Duration(deps.Default.Config.MustGetInt("clients.corge-http.timeout")) * time.Millisecond
	defaultHeaders := make(map[string]string)
	if deps.Default.Config.ContainsKey("clients.corge-http.defaultHeaders") {
		deps.Default.Config.MustGetStruct("clients.corge-http.defaultHeaders", &defaultHeaders)
	}

	maxConcurrentRequests := deps.Default.Config.MustGetInt("clients.corge-http.maxConcurrentRequests")
	errorPercentThreshold := deps.Default.Config.MustGetInt("clients.corge-http.errorPercentThreshold")
	hystrix.ConfigureCommand("corge-http", hystrix.CommandConfig{
		MaxConcurrentRequests: int(maxConcurrentRequests),
		ErrorPercentThreshold: int(errorPercentThreshold),
	})

	return &corgeHTTPClient{
		clientID:     "corge-http",
		callerHeader: callerHeader,
		calleeHeader: calleeHeader,
		callerName:   callerName,
		calleeName:   calleeName,
		httpClient: zanzibar.NewHTTPClientContext(
			deps.Default.Logger, deps.Default.ContextMetrics,
			"corge-http",
			[]string{
				"EchoString",
			},
			baseURL,
			defaultHeaders,
			timeout,
		),
	}
}

// HTTPClient returns the underlying HTTP client, should only be
// used for internal testing.
func (c *corgeHTTPClient) HTTPClient() *zanzibar.HTTPClient {
	return c.httpClient
}

// EchoString calls "/echo/string" endpoint.
func (c *corgeHTTPClient) EchoString(
	ctx context.Context,
	headers map[string]string,
	r *clientsCorgeCorge.Corge_EchoString_Args,
) (string, map[string]string, error) {
	var defaultRes string
	req := zanzibar.NewClientHTTPRequest(ctx, c.clientID, "EchoString", c.httpClient)

	headers[c.callerHeader] = c.callerName
	headers[c.calleeHeader] = c.calleeName

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/echo" + "/string"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return defaultRes, nil, err
	}

	var res *zanzibar.ClientHTTPResponse
	err = hystrix.Do("corge-http", func() error {
		res, err = req.Do()
		return err
	}, nil)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody string
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}

		return responseBody, respHeaders, nil
	default:
		_, err = res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
	}

	return defaultRes, respHeaders, &zanzibar.UnexpectedHTTPError{
		StatusCode: res.StatusCode,
		RawBody:    res.GetRawBody(),
	}
}
