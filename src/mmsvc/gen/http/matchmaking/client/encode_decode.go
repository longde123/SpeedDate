// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// matchmaking HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/mmsvc/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/http"
)

// BuildInsertRequest instantiates a HTTP request object with method and path
// set to call the "matchmaking" service "insert" endpoint
func (c *Client) BuildInsertRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: InsertMatchmakingPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("matchmaking", "insert", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeInsertResponse returns a decoder for responses returned by the
// matchmaking insert endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeInsertResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("matchmaking", "insert", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("matchmaking", "insert", resp.StatusCode, string(body))
		}
	}
}