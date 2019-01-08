// Copyright (c) 2019 Uber Technologies, Inc.
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

package zanzibar

import (
	"context"
	"testing"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestWithEndpointField(t *testing.T) {
	expected := "someEndpoint"
	ctx := WithEndpointField(context.TODO(), expected)

	ek := ctx.Value(endpointKey)
	endpoint, ok := ek.(string)

	assert.True(t, ok)
	assert.Equal(t, endpoint, expected)
}

func TestGetRequestEndpointFromCtx(t *testing.T) {
	expected := "someEndpoint"
	ctx := WithEndpointField(context.TODO(), expected)
	endpoint := GetRequestEndpointFromCtx(ctx)
	assert.Equal(t, expected, endpoint)

	expected = ""
	ctx = context.TODO()
	endpoint = GetRequestEndpointFromCtx(ctx)
	assert.Equal(t, expected, endpoint)
}

func TestWithEndpointRequestHeadersField(t *testing.T) {
	expected := map[string]string{"region": "san_francisco", "dc": "sjc1"}
	ctx := WithEndpointRequestHeadersField(context.TODO(), expected)
	rh := ctx.Value(endpointRequestHeader)
	requestHeaders, ok := rh.(map[string]string)

	assert.True(t, ok)
	assert.Equal(t, requestHeaders, expected)
}

func TestGetEndpointRequestHeadersFromCtx(t *testing.T) {
	expected := map[string]string{"region": "san_francisco", "dc": "sjc1"}
	headers := map[string]string{"region": "san_francisco", "dc": "sjc1"}
	ctx := WithEndpointRequestHeadersField(context.TODO(), headers)
	requestHeaders := GetEndpointRequestHeadersFromCtx(ctx)
	assert.Equal(t, expected, requestHeaders)

	expected = map[string]string{}
	ctx = context.TODO()
	requestHeaders = GetEndpointRequestHeadersFromCtx(ctx)
	assert.Equal(t, expected, requestHeaders)
}

func TestWithScopeTags(t *testing.T) {
	expected := map[string]string{"endpoint": "tincup", "handler": "exchange"}
	ctx := WithScopeTags(context.TODO(), expected)
	rs := ctx.Value(scopeTags)
	scopes, ok := rs.(map[string]string)

	assert.True(t, ok)
	assert.Equal(t, expected, scopes)
}

func TestGetScopeTagsFromCtx(t *testing.T) {
	expected := map[string]string{"endpoint": "tincup", "handler": "exchange"}
	scope := map[string]string{"endpoint": "tincup", "handler": "exchange"}
	ctx := WithScopeTags(context.TODO(), scope)
	scopes := GetScopeTagsFromCtx(ctx)
	assert.Equal(t, expected, scopes)

	expected = map[string]string{}
	ctx = context.TODO()
	scopes = GetScopeTagsFromCtx(ctx)
	assert.Equal(t, expected, scopes)
}

func TestWithRequestFields(t *testing.T) {
	uid := uuid.NewUUID()
	ctx := withRequestUUID(context.TODO(), uid)

	u := ctx.Value(requestUUIDKey)
	u1, ok := u.(uuid.UUID)

	assert.NotNil(t, ctx)
	assert.Equal(t, uid, u1)
	assert.True(t, ok)
}

func TestGetRequestUUIDFromCtx(t *testing.T) {
	uid := uuid.NewUUID()
	ctx := withRequestUUID(context.TODO(), uid)

	requestUUID := GetRequestUUIDFromCtx(ctx)

	assert.NotNil(t, ctx)
	assert.Equal(t, uid, requestUUID)

	// Test Default Scenario where no uuid exists in the context
	requestUUID = GetRequestUUIDFromCtx(context.TODO())
	assert.Nil(t, requestUUID)
}

func TestWithRoutingDelegate(t *testing.T) {
	expected := "somewhere"
	ctx := WithRoutingDelegate(context.TODO(), expected)
	rd := ctx.Value(routingDelegateKey)
	routingDelegate, ok := rd.(string)

	assert.True(t, ok)
	assert.Equal(t, routingDelegate, expected)
}

func TestGetRoutingDelegateFromCtx(t *testing.T) {
	expected := "somewhere"
	ctx := WithRoutingDelegate(context.TODO(), expected)
	rd := GetRoutingDelegateFromCtx(ctx)

	assert.Equal(t, expected, rd)
}

func TestContextLogger(t *testing.T) {
	zapLoggerCore, logs := observer.New(zap.DebugLevel)
	zapLogger := zap.New(zapLoggerCore)
	contextLogger := NewContextLogger(zapLogger)
	ctx := context.Background()
	ctxWithField := WithLogFields(ctx, zap.String("ctxField", "ctxValue"))

	var logMessages []observer.LoggedEntry

	contextLogger.Debug(ctxWithField, "msg", zap.String("argField", "argValue"))
	logMessages = logs.TakeAll()
	assert.Len(t, logMessages, 1)
	assert.Equal(t, zap.DebugLevel, logMessages[0].Level)
	assert.Equal(t, logMessages[0].Context[0].Key, "ctxField")
	assert.Equal(t, logMessages[0].Context[0].String, "ctxValue")
	assert.Equal(t, logMessages[0].Context[1].Key, "argField")
	assert.Equal(t, logMessages[0].Context[1].String, "argValue")

	contextLogger.Info(ctxWithField, "msg", zap.String("argField", "argValue"))
	logMessages = logs.TakeAll()
	assert.Len(t, logMessages, 1)
	assert.Equal(t, zap.InfoLevel, logMessages[0].Level)

	contextLogger.Warn(ctxWithField, "msg", zap.String("argField", "argValue"))
	logMessages = logs.TakeAll()
	assert.Len(t, logMessages, 1)
	assert.Equal(t, zap.WarnLevel, logMessages[0].Level)

	contextLogger.Error(ctxWithField, "msg", zap.String("argField", "argValue"))
	logMessages = logs.TakeAll()
	assert.Len(t, logMessages, 1)
	assert.Equal(t, zap.ErrorLevel, logMessages[0].Level)
}

func TestContextLoggerPanic(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	zapNop := zap.NewNop()

	contextLogger := NewContextLogger(zapNop)
	ctx := context.Background()

	contextLogger.Panic(ctx, "msg", zap.String("argField", "argValue"))
}

func TestExtractScopeTag(t *testing.T) {
	headers := map[string]string{"x-uber-region-id": "san_francisco"}
	ctx := WithEndpointRequestHeadersField(context.TODO(), headers)
	contextScopeExtractors := []ContextScopeTagsExtractor{func(ctx context.Context) map[string]string {
		headers := GetEndpointRequestHeadersFromCtx(ctx)
		return map[string]string{"region-id": headers["x-uber-region-id"]}
	}}

	expected := map[string]string{"region-id": "san_francisco"}
	contextExtractors := &ContextExtractors{}
	for _, scopeExtractor := range contextScopeExtractors {
		contextExtractors.AddContextScopeTagsExtractor(scopeExtractor)
	}

	contextExtractor := contextExtractors.MakeContextExtractor()
	tags := contextExtractor.ExtractScopeTags(ctx)
	assert.Equal(t, tags, expected)
}

func TestExtractLogField(t *testing.T) {
	headers := map[string]string{"x-uber-region-id": "san_francisco"}
	ctx := WithEndpointRequestHeadersField(context.TODO(), headers)
	contextLogFieldsExtractor := []ContextLogFieldsExtractor{func(ctx context.Context) []zap.Field {
		var fields []zap.Field
		headers := GetEndpointRequestHeadersFromCtx(ctx)
		fields = append(fields, zap.String("region-id", headers["x-uber-region-id"]))
		return fields
	}}

	var expected []zap.Field
	expected = append(expected, zap.String("region-id", "san_francisco"))
	contextExtractors := &ContextExtractors{}
	for _, logFieldExtractor := range contextLogFieldsExtractor {
		contextExtractors.AddContextLogFieldsExtractor(logFieldExtractor)
	}

	contextExtractor := contextExtractors.MakeContextExtractor()
	fields := contextExtractor.ExtractLogFields(ctx)
	assert.Equal(t, expected, fields)
}

func TestLoggerWithFields(t *testing.T) {
	zapLoggerCore, logs := observer.New(zap.DebugLevel)
	zapLogger := zap.New(zapLoggerCore)
	fields := []zap.Field{
		zap.String("key", "value"),
	}
	logger := newLoggerWithFields(zapLogger, fields)

	var logMessages []observer.LoggedEntry

	logger.Debug("msg", zap.String("argField", "argValue"))
	logMessages = logs.TakeAll()
	assert.Len(t, logMessages, 1)
	assert.Equal(t, zap.DebugLevel, logMessages[0].Level)
	assert.Equal(t, logMessages[0].Context[0].Key, "key")
	assert.Equal(t, logMessages[0].Context[0].String, "value")
	assert.Equal(t, logMessages[0].Context[1].Key, "argField")
	assert.Equal(t, logMessages[0].Context[1].String, "argValue")

	logger.Info("msg", zap.String("argField", "argValue"))
	logMessages = logs.TakeAll()
	assert.Len(t, logMessages, 1)
	assert.Equal(t, zap.InfoLevel, logMessages[0].Level)
	assert.Equal(t, logMessages[0].Context[0].Key, "key")
	assert.Equal(t, logMessages[0].Context[0].String, "value")
	assert.Equal(t, logMessages[0].Context[1].Key, "argField")
	assert.Equal(t, logMessages[0].Context[1].String, "argValue")

	logger.Warn("msg", zap.String("argField", "argValue"))
	logMessages = logs.TakeAll()
	assert.Len(t, logMessages, 1)
	assert.Equal(t, zap.WarnLevel, logMessages[0].Level)
	assert.Equal(t, logMessages[0].Context[0].Key, "key")
	assert.Equal(t, logMessages[0].Context[0].String, "value")
	assert.Equal(t, logMessages[0].Context[1].Key, "argField")
	assert.Equal(t, logMessages[0].Context[1].String, "argValue")

	logger.Error("msg", zap.String("argField", "argValue"))
	logMessages = logs.TakeAll()
	assert.Len(t, logMessages, 1)
	assert.Equal(t, zap.ErrorLevel, logMessages[0].Level)
	assert.Equal(t, logMessages[0].Context[0].Key, "key")
	assert.Equal(t, logMessages[0].Context[0].String, "value")
	assert.Equal(t, logMessages[0].Context[1].Key, "argField")
	assert.Equal(t, logMessages[0].Context[1].String, "argValue")

	assert.Panics(t, func() {
		logger.Panic("msg", zap.String("argField", "argValue"))
		logMessages = logs.TakeAll()
		assert.Len(t, logMessages, 1)
		assert.Equal(t, zap.ErrorLevel, logMessages[0].Level)
		assert.Equal(t, logMessages[0].Context[0].Key, "key")
		assert.Equal(t, logMessages[0].Context[0].String, "value")
		assert.Equal(t, logMessages[0].Context[1].Key, "argField")
		assert.Equal(t, logMessages[0].Context[1].String, "argValue")
	})

	ce := logger.Check(zap.DebugLevel, "msg")
	assert.NotNil(t, ce)
}
