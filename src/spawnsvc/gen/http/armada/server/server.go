// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// armada HTTP server
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package server

import (
	"context"
	"net/http"

	armada "github.com/proepkes/speeddate/src/spawnsvc/gen/armada"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Server lists the armada service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Add    http.Handler
	Clear  http.Handler
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

// New instantiates HTTP handlers for all the armada service endpoints.
func New(
	e *armada.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Add", "POST", "/armada/add"},
			{"Clear", "POST", "/armada/clear"},
		},
		Add:   NewAddHandler(e.Add, mux, dec, enc, eh),
		Clear: NewClearHandler(e.Clear, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "armada" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Add = m(s.Add)
	s.Clear = m(s.Clear)
}

// Mount configures the mux to serve the armada endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAddHandler(mux, h.Add)
	MountClearHandler(mux, h.Clear)
}

// MountAddHandler configures the mux to serve the "armada" service "add"
// endpoint.
func MountAddHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/armada/add", f)
}

// NewAddHandler creates a HTTP handler which loads the HTTP request and calls
// the "armada" service "add" endpoint.
func NewAddHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeAddResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add")
		ctx = context.WithValue(ctx, goa.ServiceKey, "armada")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountClearHandler configures the mux to serve the "armada" service "clear"
// endpoint.
func MountClearHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/armada/clear", f)
}

// NewClearHandler creates a HTTP handler which loads the HTTP request and
// calls the "armada" service "clear" endpoint.
func NewClearHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeClearResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "clear")
		ctx = context.WithValue(ctx, goa.ServiceKey, "armada")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
