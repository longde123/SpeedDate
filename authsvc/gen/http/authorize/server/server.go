// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// authorize HTTP server
//
// Command:
// $ goa gen github.com/proepkes/speeddate/authsvc/design

package server

import (
	"context"
	"net/http"

	authorize "github.com/proepkes/speeddate/authsvc/gen/authorize"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
	"goa.design/plugins/cors"
)

// Server lists the authorize service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Login  http.Handler
	CORS   http.Handler
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

// New instantiates HTTP handlers for all the authorize service endpoints.
func New(
	e *authorize.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Login", "POST", "/auth/login"},
			{"CORS", "OPTIONS", "/auth/login"},
		},
		Login: NewLoginHandler(e.Login, mux, dec, enc, eh),
		CORS:  NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "authorize" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Login = m(s.Login)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the authorize endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountLoginHandler(mux, h.Login)
	MountCORSHandler(mux, h.CORS)
}

// MountLoginHandler configures the mux to serve the "authorize" service
// "login" endpoint.
func MountLoginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleAuthorizeOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/auth/login", f)
}

// NewLoginHandler creates a HTTP handler which loads the HTTP request and
// calls the "authorize" service "login" endpoint.
func NewLoginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeLoginRequest(mux, dec)
		encodeResponse = EncodeLoginResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login")
		ctx = context.WithValue(ctx, goa.ServiceKey, "authorize")
		payload, err := decodeRequest(r)
		if err != nil {
			eh(ctx, w, err)
			return
		}

		res, err := endpoint(ctx, payload)

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

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service authorize.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleAuthorizeOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/auth/login", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleAuthorizeOrigin applies the CORS response headers corresponding to the
// origin for the service authorize.
func handleAuthorizeOrigin(h http.Handler) http.Handler {
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Expose-Headers", "Access-token")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
