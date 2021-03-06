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

package workflow

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceHeaderSchemaWorkflow defines the interface for SimpleServiceHeaderSchema workflow
type SimpleServiceHeaderSchemaWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBazBaz.SimpleService_HeaderSchema_Args,
	) (*endpointsBazBaz.HeaderSchema, zanzibar.Header, error)
}

// NewSimpleServiceHeaderSchemaWorkflow creates a workflow
func NewSimpleServiceHeaderSchemaWorkflow(deps *module.Dependencies) SimpleServiceHeaderSchemaWorkflow {
	return &simpleServiceHeaderSchemaWorkflow{
		Clients: deps.Client,
		Logger:  deps.Default.Logger,
	}
}

// simpleServiceHeaderSchemaWorkflow calls thrift client Baz.HeaderSchema
type simpleServiceHeaderSchemaWorkflow struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
}

// Handle calls thrift client.
func (w simpleServiceHeaderSchemaWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBazBaz.SimpleService_HeaderSchema_Args,
) (*endpointsBazBaz.HeaderSchema, zanzibar.Header, error) {
	clientRequest := convertToHeaderSchemaClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("Auth")
	if ok {
		clientHeaders["Auth"] = h
	}
	h, ok = reqHeaders.Get("Content-Type")
	if ok {
		clientHeaders["Content-Type"] = h
	}
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	h, ok = reqHeaders.Get("X-Token")
	if ok {
		clientHeaders["X-Token"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}
	h, ok = reqHeaders.Get("X-Zanzibar-Use-Staging")
	if ok {
		clientHeaders["X-Zanzibar-Use-Staging"] = h
	}

	clientRespBody, _, err := w.Clients.Baz.HeaderSchema(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertHeaderSchemaAuthErr(
				errValue,
			)

			return nil, nil, serverErr

		case *clientsBazBaz.OtherAuthErr:
			serverErr := convertHeaderSchemaOtherAuthErr(
				errValue,
			)

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceHeaderSchemaClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToHeaderSchemaClientRequest(in *endpointsBazBaz.SimpleService_HeaderSchema_Args) *clientsBazBaz.SimpleService_HeaderSchema_Args {
	out := &clientsBazBaz.SimpleService_HeaderSchema_Args{}

	if in.Req != nil {
		out.Req = &clientsBazBaz.HeaderSchema{}
	} else {
		out.Req = nil
	}

	return out
}

func convertHeaderSchemaAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
func convertHeaderSchemaOtherAuthErr(
	clientError *clientsBazBaz.OtherAuthErr,
) *endpointsBazBaz.OtherAuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.OtherAuthErr{}
	return serverError
}

func convertSimpleServiceHeaderSchemaClientResponse(in *clientsBazBaz.HeaderSchema) *endpointsBazBaz.HeaderSchema {
	out := &endpointsBazBaz.HeaderSchema{}

	return out
}
