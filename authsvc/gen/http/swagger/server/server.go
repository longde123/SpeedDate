// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// swagger HTTP server
//
// Command:
// $ goa gen github.com/proepkes/speeddate/authsvc/design

package server

import (
	"context"
	"net/http"

	swagger "github.com/proepkes/speeddate/authsvc/gen/swagger"
	goahttp "goa.design/goa/http"
)

// Server lists the swagger service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the swagger service endpoints.
func New(
	e *swagger.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"../../gen/http/openapi.json", "GET", "/swagger/swagger.json"},
		},
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "swagger" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
}

// Mount configures the mux to serve the swagger endpoints.
func Mount(mux goahttp.Muxer) {
	MountGenHTTPOpenapiJSON(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../gen/http/openapi.json")
	}))
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/swagger/swagger.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/swagger/swagger.json", h.ServeHTTP)
}
