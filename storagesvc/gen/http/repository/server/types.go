// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// repository HTTP server types
//
// Command:
// $ goa gen github.com/proepkes/speeddate/storagesvc/design

package server

import (
	"unicode/utf8"

	repository "github.com/proepkes/speeddate/storagesvc/gen/repository"
	repositoryviews "github.com/proepkes/speeddate/storagesvc/gen/repository/views"
	goa "goa.design/goa"
)

// InsertRequestBody is the type of the "repository" service "insert" endpoint
// HTTP request body.
type InsertRequestBody struct {
	// The username
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// GetResponseBody is the type of the "repository" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// UUID is the unique id of the user.
	ID string `gorm:"TYPE:uuid; COLUMN:id; PRIMARY_KEY; DEFAULT: gen_random_uuid()" json:"id"`
	// The username
	Name string `form:"name" json:"name" xml:"name"`
	// Indicates whether the user is currently online.
	Online *bool `form:"online,omitempty" json:"online,omitempty" xml:"online,omitempty"`
}

// GetResponseBodyTiny is the type of the "repository" service "get" endpoint
// HTTP response body.
type GetResponseBodyTiny struct {
	// UUID is the unique id of the user.
	ID string `gorm:"TYPE:uuid; COLUMN:id; PRIMARY_KEY; DEFAULT: gen_random_uuid()" json:"id"`
	// The username
	Name string `form:"name" json:"name" xml:"name"`
}

// GetNotFoundResponseBody is the type of the "repository" service "get"
// endpoint HTTP response body for the "not_found" error.
type GetNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetUnauthorizedResponseBody is the type of the "repository" service "get"
// endpoint HTTP response body for the "unauthorized" error.
type GetUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "repository" service.
func NewGetResponseBody(res *repositoryviews.StoredUserView) *GetResponseBody {
	body := &GetResponseBody{
		ID:     *res.ID,
		Name:   *res.Name,
		Online: res.Online,
	}
	return body
}

// NewGetResponseBodyTiny builds the HTTP response body from the result of the
// "get" endpoint of the "repository" service.
func NewGetResponseBodyTiny(res *repositoryviews.StoredUserView) *GetResponseBodyTiny {
	body := &GetResponseBodyTiny{
		ID:   *res.ID,
		Name: *res.Name,
	}
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "repository" service.
func NewGetNotFoundResponseBody(res *goa.ServiceError) *GetNotFoundResponseBody {
	body := &GetNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUnauthorizedResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "repository" service.
func NewGetUnauthorizedResponseBody(res *goa.ServiceError) *GetUnauthorizedResponseBody {
	body := &GetUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewInsertUser builds a repository service insert endpoint payload.
func NewInsertUser(body *InsertRequestBody) *repository.User {
	v := &repository.User{
		Name: *body.Name,
	}
	return v
}

// NewDeletePayload builds a repository service delete endpoint payload.
func NewDeletePayload(id string) *repository.DeletePayload {
	return &repository.DeletePayload{
		ID: id,
	}
}

// NewGetPayload builds a repository service get endpoint payload.
func NewGetPayload(id string, view *string, token *string) *repository.GetPayload {
	return &repository.GetPayload{
		ID:    id,
		View:  view,
		Token: token,
	}
}

// Validate runs the validations defined on InsertRequestBody
func (body *InsertRequestBody) Validate() (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 50, false))
		}
	}
	return
}