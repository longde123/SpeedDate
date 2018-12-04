// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// fleet client
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package fleet

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "fleet" service client.
type Client struct {
	AddEndpoint           goa.Endpoint
	ClearEndpoint         goa.Endpoint
	ConfigurationEndpoint goa.Endpoint
	ConfigureEndpoint     goa.Endpoint
}

// NewClient initializes a "fleet" service client given the endpoints.
func NewClient(add, clear, configuration, configure goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:           add,
		ClearEndpoint:         clear,
		ConfigurationEndpoint: configuration,
		ConfigureEndpoint:     configure,
	}
}

// Add calls the "add" endpoint of the "fleet" service.
func (c *Client) Add(ctx context.Context) (res string, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Clear calls the "clear" endpoint of the "fleet" service.
func (c *Client) Clear(ctx context.Context) (res string, err error) {
	var ires interface{}
	ires, err = c.ClearEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Configuration calls the "configuration" endpoint of the "fleet" service.
func (c *Client) Configuration(ctx context.Context) (res *GameserverTemplate, err error) {
	var ires interface{}
	ires, err = c.ConfigurationEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*GameserverTemplate), nil
}

// Configure calls the "configure" endpoint of the "fleet" service.
func (c *Client) Configure(ctx context.Context, p *GameserverTemplate) (res string, err error) {
	var ires interface{}
	ires, err = c.ConfigureEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
