// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// fleet HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package server

import (
	"context"
	"io"
	"net/http"

	fleet "github.com/proepkes/speeddate/src/spawnsvc/gen/fleet"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// EncodeAddResponse returns an encoder for responses returned by the fleet add
// endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// EncodeClearResponse returns an encoder for responses returned by the fleet
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

// EncodeConfigurationResponse returns an encoder for responses returned by the
// fleet configuration endpoint.
func EncodeConfigurationResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*fleet.GameserverTemplate)
		enc := encoder(ctx, w)
		body := NewConfigurationResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeConfigureResponse returns an encoder for responses returned by the
// fleet configure endpoint.
func EncodeConfigureResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeConfigureRequest returns a decoder for requests sent to the fleet
// configure endpoint.
func DecodeConfigureRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ConfigureRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = body.Validate()
		if err != nil {
			return nil, err
		}
		payload := NewConfigureGameserverTemplate(&body)

		return payload, nil
	}
}
