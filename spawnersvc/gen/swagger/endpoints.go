// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// swagger endpoints
//
// Command:
// $ goa gen github.com/proepkes/speeddate/gamehostsvc/design

package swagger

import (
	goa "goa.design/goa"
)

// Endpoints wraps the "swagger" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "swagger" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "swagger" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}
