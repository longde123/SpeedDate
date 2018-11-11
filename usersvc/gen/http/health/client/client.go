// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// health client HTTP transport
//
// Command:
// $ goa gen github.com/proepkes/speeddate/usersvc/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the health service endpoint HTTP clients.
type Client struct {
	// CheckHealth Doer is the HTTP client used to make requests to the checkHealth
	// endpoint.
	CheckHealthDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the health service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CheckHealthDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// CheckHealth returns an endpoint that makes HTTP requests to the health
// service checkHealth server.
func (c *Client) CheckHealth() goa.Endpoint {
	var (
		decodeResponse = DecodeCheckHealthResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCheckHealthRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CheckHealthDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("health", "checkHealth", err)
		}
		return decodeResponse(resp)
	}
}
