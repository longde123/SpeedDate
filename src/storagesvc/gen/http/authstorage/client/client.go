// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// authstorage client HTTP transport
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/storagesvc/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the authstorage service endpoint HTTP clients.
type Client struct {
	// Insert Doer is the HTTP client used to make requests to the insert endpoint.
	InsertDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// Get Doer is the HTTP client used to make requests to the get endpoint.
	GetDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the authstorage service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		InsertDoer:          doer,
		DeleteDoer:          doer,
		GetDoer:             doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Insert returns an endpoint that makes HTTP requests to the authstorage
// service insert server.
func (c *Client) Insert() goa.Endpoint {
	var (
		encodeRequest  = EncodeInsertRequest(c.encoder)
		decodeResponse = DecodeInsertResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildInsertRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.InsertDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("authstorage", "insert", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the authstorage
// service delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("authstorage", "delete", err)
		}
		return decodeResponse(resp)
	}
}

// Get returns an endpoint that makes HTTP requests to the authstorage service
// get server.
func (c *Client) Get() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetRequest(c.encoder)
		decodeResponse = DecodeGetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("authstorage", "get", err)
		}
		return decodeResponse(resp)
	}
}