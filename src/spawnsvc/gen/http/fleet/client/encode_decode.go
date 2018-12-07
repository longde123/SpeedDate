// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// fleet HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	fleet "github.com/proepkes/speeddate/src/spawnsvc/gen/fleet"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "fleet" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddFleetPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("fleet", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeAddResponse returns a decoder for responses returned by the fleet add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				return nil, goahttp.ErrDecodingError("fleet", "add", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("fleet", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "fleet" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateFleetPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("fleet", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the fleet create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*fleet.Fleet)
		if !ok {
			return goahttp.ErrInvalidType("fleet", "create", "*fleet.Fleet", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("fleet", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the fleet
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				return nil, goahttp.ErrDecodingError("fleet", "create", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("fleet", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "fleet" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListFleetPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("fleet", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the fleet list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*fleet.NamespacePayload)
		if !ok {
			return goahttp.ErrInvalidType("fleet", "list", "*fleet.NamespacePayload", v)
		}
		values := req.URL.Query()
		if p.Namespace != nil {
			values.Add("namespace", *p.Namespace)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the fleet
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("fleet", "list", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := e.Validate(); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("fleet", "list", err)
			}
			res := NewListStoredFleetOK(body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("fleet", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildClearRequest instantiates a HTTP request object with method and path
// set to call the "fleet" service "clear" endpoint
func (c *Client) BuildClearRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ClearFleetPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("fleet", "clear", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeClearResponse returns a decoder for responses returned by the fleet
// clear endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeClearResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("fleet", "clear", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("fleet", "clear", resp.StatusCode, string(body))
		}
	}
}

// BuildConfigurationRequest instantiates a HTTP request object with method and
// path set to call the "fleet" service "configuration" endpoint
func (c *Client) BuildConfigurationRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ConfigurationFleetPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("fleet", "configuration", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeConfigurationResponse returns a decoder for responses returned by the
// fleet configuration endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeConfigurationResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusOK:
			var (
				body ConfigurationResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("fleet", "configuration", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("fleet", "configuration", err)
			}
			res := NewConfigurationGameserverTemplateOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("fleet", "configuration", resp.StatusCode, string(body))
		}
	}
}

// BuildConfigureRequest instantiates a HTTP request object with method and
// path set to call the "fleet" service "configure" endpoint
func (c *Client) BuildConfigureRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ConfigureFleetPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("fleet", "configure", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeConfigureRequest returns an encoder for requests sent to the fleet
// configure server.
func EncodeConfigureRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*fleet.GameserverTemplate)
		if !ok {
			return goahttp.ErrInvalidType("fleet", "configure", "*fleet.GameserverTemplate", v)
		}
		body := NewConfigureRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("fleet", "configure", err)
		}
		return nil
	}
}

// DecodeConfigureResponse returns a decoder for responses returned by the
// fleet configure endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeConfigureResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("fleet", "configure", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("fleet", "configure", resp.StatusCode, string(body))
		}
	}
}

// marshalObjectMetaToObjectMetaRequestBody builds a value of type
// *ObjectMetaRequestBody from a value of type *fleet.ObjectMeta.
func marshalObjectMetaToObjectMetaRequestBody(v *fleet.ObjectMeta) *ObjectMetaRequestBody {
	if v == nil {
		return nil
	}
	res := &ObjectMetaRequestBody{
		GenerateName: v.GenerateName,
		Namespace:    v.Namespace,
	}

	return res
}

// marshalFleetSpecToFleetSpecRequestBody builds a value of type
// *FleetSpecRequestBody from a value of type *fleet.FleetSpec.
func marshalFleetSpecToFleetSpecRequestBody(v *fleet.FleetSpec) *FleetSpecRequestBody {
	res := &FleetSpecRequestBody{
		Replicas: v.Replicas,
	}
	if v.Template != nil {
		res.Template = marshalGameserverTemplateToGameserverTemplateRequestBody(v.Template)
	}

	return res
}

// marshalGameserverTemplateToGameserverTemplateRequestBody builds a value of
// type *GameserverTemplateRequestBody from a value of type
// *fleet.GameserverTemplate.
func marshalGameserverTemplateToGameserverTemplateRequestBody(v *fleet.GameserverTemplate) *GameserverTemplateRequestBody {
	res := &GameserverTemplateRequestBody{}
	if v.ObjectMeta != nil {
		res.ObjectMeta = marshalObjectMetaToObjectMetaRequestBody(v.ObjectMeta)
	}
	if v.GameServerSpec != nil {
		res.GameServerSpec = marshalGameServerSpecToGameServerSpecRequestBody(v.GameServerSpec)
	}

	return res
}

// marshalGameServerSpecToGameServerSpecRequestBody builds a value of type
// *GameServerSpecRequestBody from a value of type *fleet.GameServerSpec.
func marshalGameServerSpecToGameServerSpecRequestBody(v *fleet.GameServerSpec) *GameServerSpecRequestBody {
	res := &GameServerSpecRequestBody{
		PortPolicy:     v.PortPolicy,
		ContainerName:  v.ContainerName,
		ContainerImage: v.ContainerImage,
		ContainerPort:  v.ContainerPort,
	}

	return res
}

// marshalObjectMetaRequestBodyToObjectMeta builds a value of type
// *fleet.ObjectMeta from a value of type *ObjectMetaRequestBody.
func marshalObjectMetaRequestBodyToObjectMeta(v *ObjectMetaRequestBody) *fleet.ObjectMeta {
	if v == nil {
		return nil
	}
	res := &fleet.ObjectMeta{
		GenerateName: v.GenerateName,
		Namespace:    v.Namespace,
	}

	return res
}

// marshalFleetSpecRequestBodyToFleetSpec builds a value of type
// *fleet.FleetSpec from a value of type *FleetSpecRequestBody.
func marshalFleetSpecRequestBodyToFleetSpec(v *FleetSpecRequestBody) *fleet.FleetSpec {
	res := &fleet.FleetSpec{
		Replicas: v.Replicas,
	}
	if v.Template != nil {
		res.Template = marshalGameserverTemplateRequestBodyToGameserverTemplate(v.Template)
	}

	return res
}

// marshalGameserverTemplateRequestBodyToGameserverTemplate builds a value of
// type *fleet.GameserverTemplate from a value of type
// *GameserverTemplateRequestBody.
func marshalGameserverTemplateRequestBodyToGameserverTemplate(v *GameserverTemplateRequestBody) *fleet.GameserverTemplate {
	res := &fleet.GameserverTemplate{}
	if v.ObjectMeta != nil {
		res.ObjectMeta = marshalObjectMetaRequestBodyToObjectMeta(v.ObjectMeta)
	}
	if v.GameServerSpec != nil {
		res.GameServerSpec = marshalGameServerSpecRequestBodyToGameServerSpec(v.GameServerSpec)
	}

	return res
}

// marshalGameServerSpecRequestBodyToGameServerSpec builds a value of type
// *fleet.GameServerSpec from a value of type *GameServerSpecRequestBody.
func marshalGameServerSpecRequestBodyToGameServerSpec(v *GameServerSpecRequestBody) *fleet.GameServerSpec {
	res := &fleet.GameServerSpec{
		PortPolicy:     v.PortPolicy,
		ContainerName:  v.ContainerName,
		ContainerImage: v.ContainerImage,
		ContainerPort:  v.ContainerPort,
	}

	return res
}

// unmarshalObjectMetaResponseBodyToObjectMeta builds a value of type
// *fleet.ObjectMeta from a value of type *ObjectMetaResponseBody.
func unmarshalObjectMetaResponseBodyToObjectMeta(v *ObjectMetaResponseBody) *fleet.ObjectMeta {
	res := &fleet.ObjectMeta{
		GenerateName: *v.GenerateName,
		Namespace:    *v.Namespace,
	}

	return res
}

// unmarshalFleetSpecResponseBodyToFleetSpec builds a value of type
// *fleet.FleetSpec from a value of type *FleetSpecResponseBody.
func unmarshalFleetSpecResponseBodyToFleetSpec(v *FleetSpecResponseBody) *fleet.FleetSpec {
	res := &fleet.FleetSpec{
		Replicas: *v.Replicas,
	}
	res.Template = unmarshalGameserverTemplateResponseBodyToGameserverTemplate(v.Template)

	return res
}

// unmarshalGameserverTemplateResponseBodyToGameserverTemplate builds a value
// of type *fleet.GameserverTemplate from a value of type
// *GameserverTemplateResponseBody.
func unmarshalGameserverTemplateResponseBodyToGameserverTemplate(v *GameserverTemplateResponseBody) *fleet.GameserverTemplate {
	res := &fleet.GameserverTemplate{}
	if v.ObjectMeta != nil {
		res.ObjectMeta = unmarshalObjectMetaResponseBodyToObjectMeta(v.ObjectMeta)
	}
	res.GameServerSpec = unmarshalGameServerSpecResponseBodyToGameServerSpec(v.GameServerSpec)

	return res
}

// unmarshalGameServerSpecResponseBodyToGameServerSpec builds a value of type
// *fleet.GameServerSpec from a value of type *GameServerSpecResponseBody.
func unmarshalGameServerSpecResponseBodyToGameServerSpec(v *GameServerSpecResponseBody) *fleet.GameServerSpec {
	res := &fleet.GameServerSpec{
		PortPolicy:     *v.PortPolicy,
		ContainerName:  *v.ContainerName,
		ContainerImage: *v.ContainerImage,
		ContainerPort:  *v.ContainerPort,
	}

	return res
}

// unmarshalFleetStatusResponseBodyToFleetStatus builds a value of type
// *fleet.FleetStatus from a value of type *FleetStatusResponseBody.
func unmarshalFleetStatusResponseBodyToFleetStatus(v *FleetStatusResponseBody) *fleet.FleetStatus {
	if v == nil {
		return nil
	}
	res := &fleet.FleetStatus{
		Replicas:          *v.Replicas,
		ReadyReplicas:     *v.ReadyReplicas,
		AllocatedReplicas: *v.AllocatedReplicas,
	}

	return res
}
