// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// fleet client HTTP transport
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the fleet service endpoint HTTP clients.
type Client struct {
	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Clear Doer is the HTTP client used to make requests to the clear endpoint.
	ClearDoer goahttp.Doer

	// Configuration Doer is the HTTP client used to make requests to the
	// configuration endpoint.
	ConfigurationDoer goahttp.Doer

	// Configure Doer is the HTTP client used to make requests to the configure
	// endpoint.
	ConfigureDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the fleet service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AddDoer:             doer,
		CreateDoer:          doer,
		ListDoer:            doer,
		ClearDoer:           doer,
		ConfigurationDoer:   doer,
		ConfigureDoer:       doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Add returns an endpoint that makes HTTP requests to the fleet service add
// server.
func (c *Client) Add() goa.Endpoint {
	var (
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("fleet", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Create returns an endpoint that makes HTTP requests to the fleet service
// create server.
func (c *Client) Create() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateRequest(c.encoder)
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("fleet", "create", err)
		}
		return decodeResponse(resp)
	}
}

// List returns an endpoint that makes HTTP requests to the fleet service list
// server.
func (c *Client) List() goa.Endpoint {
	var (
		encodeRequest  = EncodeListRequest(c.encoder)
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("fleet", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Clear returns an endpoint that makes HTTP requests to the fleet service
// clear server.
func (c *Client) Clear() goa.Endpoint {
	var (
		decodeResponse = DecodeClearResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildClearRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ClearDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("fleet", "clear", err)
		}
		return decodeResponse(resp)
	}
}

// Configuration returns an endpoint that makes HTTP requests to the fleet
// service configuration server.
func (c *Client) Configuration() goa.Endpoint {
	var (
		decodeResponse = DecodeConfigurationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConfigurationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConfigurationDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("fleet", "configuration", err)
		}
		return decodeResponse(resp)
	}
}

// Configure returns an endpoint that makes HTTP requests to the fleet service
// configure server.
func (c *Client) Configure() goa.Endpoint {
	var (
		encodeRequest  = EncodeConfigureRequest(c.encoder)
		decodeResponse = DecodeConfigureResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConfigureRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConfigureDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("fleet", "configure", err)
		}
		return decodeResponse(resp)
	}
}
