// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
)

// HandleArgNotStructRequest handles "/bar/arg-not-struct-path".
func HandleArgNotStructRequest(
	ctx context.Context,
	req *zanzibar.IncomingHTTPRequest,
	res *zanzibar.OutgoingHTTPResponse,
	g *zanzibar.Gateway,
	clients *clients.Clients,
) {
	// Handle request headers.
	h := http.Header{}

	// Handle request body.
	rawBody, ok := req.ReadAll()
	if !ok {
		return
	}
	var body ArgNotStructHTTPRequest
	if ok := req.UnmarshalBody(&body, rawBody); !ok {
		return
	}
	clientRequest := convertToArgNotStructClientRequest(&body)
	clientResp, err := clients.Bar.ArgNotStruct(ctx, clientRequest, h)
	if err != nil {
		g.Logger.Error("Could not make client request",
			zap.String("error", err.Error()),
		)
		res.SendError(500, errors.Wrap(err, "could not make client request:"))
		return
	}

	defer func() {
		if cerr := clientResp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Handle client respnse.
	if !res.IsOKResponse(clientResp.StatusCode, []int{200}) {
		g.Logger.Warn("Unknown response status code",
			zap.Int("status code", clientResp.StatusCode),
		)
	}
	res.WriteJSONBytes(clientResp.StatusCode, nil)
}

func convertToArgNotStructClientRequest(body *ArgNotStructHTTPRequest) *barClient.ArgNotStructHTTPRequest {
	clientRequest := &barClient.ArgNotStructHTTPRequest{}

	clientRequest.Request = string(body.Request)

	return clientRequest
}