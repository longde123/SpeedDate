// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// armada HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/http"
)

// EncodeAddResponse returns an encoder for responses returned by the armada
// add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// EncodeClearResponse returns an encoder for responses returned by the armada
// clear endpoint.
func EncodeClearResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}
