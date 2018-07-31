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

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	clientsFooBaseBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/base/base"
	clientsFooFoo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/foo"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
	endpointsFooFoo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/foo/foo"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
	"go.uber.org/zap"
)

// BarTooManyArgsWorkflow defines the interface for BarTooManyArgs workflow
type BarTooManyArgsWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBarBar.Bar_TooManyArgs_Args,
	) (*endpointsBarBar.BarResponse, zanzibar.Header, error)
}

// NewBarTooManyArgsWorkflow creates a workflow
func NewBarTooManyArgsWorkflow(deps *module.Dependencies) BarTooManyArgsWorkflow {
	return &barTooManyArgsWorkflow{
		Clients: deps.Client,
		Logger:  deps.Default.Logger,
	}
}

// barTooManyArgsWorkflow calls thrift client Bar.TooManyArgs
type barTooManyArgsWorkflow struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
}

// Handle calls thrift client.
func (w barTooManyArgsWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_TooManyArgs_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToTooManyArgsClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
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

	clientRespBody, cliRespHeaders, err := w.Clients.Bar.TooManyArgs(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBarBar.BarException:
			serverErr := convertTooManyArgsBarException(
				errValue,
			)

			return nil, nil, serverErr

		case *clientsFooFoo.FooException:
			serverErr := convertTooManyArgsFooException(
				errValue,
			)

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "Bar"),
			)

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}
	resHeaders.Set("X-Token", cliRespHeaders["X-Token"])
	resHeaders.Set("X-Uuid", cliRespHeaders["X-Uuid"])

	response := convertBarTooManyArgsClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToTooManyArgsClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args) *clientsBarBar.Bar_TooManyArgs_Args {
	out := &clientsBarBar.Bar_TooManyArgs_Args{}

	if in.Request != nil {
		out.Request = &clientsBarBar.BarRequest{}
		out.Request.StringField = string(in.Request.StringField)
		out.Request.BoolField = bool(in.Request.BoolField)
		out.Request.BinaryField = []byte(in.Request.BinaryField)
		out.Request.Timestamp = clientsBarBar.Timestamp(in.Request.Timestamp)
		out.Request.EnumField = clientsBarBar.Fruit(in.Request.EnumField)
		out.Request.LongField = clientsBarBar.Long(in.Request.LongField)
	} else {
		out.Request = nil
	}
	if in.Foo != nil {
		out.Foo = &clientsFooFoo.FooStruct{}
		out.Foo.FooString = string(in.Foo.FooString)
		out.Foo.FooI32 = (*int32)(in.Foo.FooI32)
		out.Foo.FooI16 = (*int16)(in.Foo.FooI16)
		out.Foo.FooDouble = (*float64)(in.Foo.FooDouble)
		out.Foo.FooBool = (*bool)(in.Foo.FooBool)
		out.Foo.FooMap = make(map[string]string, len(in.Foo.FooMap))
		for key1, value2 := range in.Foo.FooMap {
			out.Foo.FooMap[key1] = string(value2)
		}
		if in.Foo.Message != nil {
			out.Foo.Message = &clientsFooBaseBase.Message{}
			out.Foo.Message.Body = string(in.Foo.Message.Body)
		} else {
			out.Foo.Message = nil
		}
	} else {
		out.Foo = nil
	}

	return out
}

func convertTooManyArgsBarException(
	clientError *clientsBarBar.BarException,
) *endpointsBarBar.BarException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBarBar.BarException{}
	return serverError
}
func convertTooManyArgsFooException(
	clientError *clientsFooFoo.FooException,
) *endpointsFooFoo.FooException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsFooFoo.FooException{}
	return serverError
}

func convertBarTooManyArgsClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
	for key1, value2 := range in.MapIntWithRange {
		out.MapIntWithRange[endpointsBarBar.UUID(key1)] = int32(value2)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key3, value4 := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key3] = int32(value4)
	}
	out.BinaryField = []byte(in.BinaryField)
	var convertBarResponseHelper5 func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse)
	convertBarResponseHelper5 = func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse) {
		if in != nil {
			out = &endpointsBarBar.BarResponse{}
			out.StringField = string(in.StringField)
			out.IntWithRange = int32(in.IntWithRange)
			out.IntWithoutRange = int32(in.IntWithoutRange)
			out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
			for key6, value7 := range in.MapIntWithRange {
				out.MapIntWithRange[endpointsBarBar.UUID(key6)] = int32(value7)
			}
			out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
			for key8, value9 := range in.MapIntWithoutRange {
				out.MapIntWithoutRange[key8] = int32(value9)
			}
			out.BinaryField = []byte(in.BinaryField)
			out.NextResponse = convertBarResponseHelper5(in.NextResponse)
		} else {
			out = nil
		}
		return
	}
	out.NextResponse = convertBarResponseHelper5(in.NextResponse)

	return out
}
