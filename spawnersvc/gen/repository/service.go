// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// repository service
//
// Command:
// $ goa gen github.com/proepkes/speeddate/gamehostsvc/design

package repository

import (
	"context"

	repositoryviews "github.com/proepkes/speeddate/gamehostsvc/gen/repository/views"
	"goa.design/goa"
	"goa.design/goa/security"
)

// The service makes it possible to insert, delete or get users.
type Service interface {
	// Add new user and return its ID.
	Insert(context.Context, *User) (res string, err error)
	// Remove user from storage
	Delete(context.Context, *DeletePayload) (err error)
	// Get implements get.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Get(context.Context, *GetPayload) (res *StoredUser, view string, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "repository"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"insert", "delete", "get"}

// User is the payload type of the repository service insert method.
type User struct {
	// The username
	Name string
}

// DeletePayload is the payload type of the repository service delete method.
type DeletePayload struct {
	// ID of user to remove
	ID string
}

// GetPayload is the payload type of the repository service get method.
type GetPayload struct {
	// JWT used for authentication
	Token *string
	// Get user by ID
	ID string
	// View to render
	View *string
}

// StoredUser is the result type of the repository service get method.
type StoredUser struct {
	// UUID is the unique id of the user.
	ID string `gorm:"TYPE:uuid; COLUMN:id; PRIMARY_KEY; DEFAULT: gen_random_uuid()" json:"id"`
	// The username
	Name string
	// Indicates whether the user is currently online.
	Online *bool
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewStoredUser initializes result type StoredUser from viewed result type
// StoredUser.
func NewStoredUser(vres *repositoryviews.StoredUser) *StoredUser {
	var res *StoredUser
	switch vres.View {
	case "default", "":
		res = newStoredUser(vres.Projected)
	case "tiny":
		res = newStoredUserTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredUser initializes viewed result type StoredUser from result
// type StoredUser using the given view.
func NewViewedStoredUser(res *StoredUser, view string) *repositoryviews.StoredUser {
	var vres *repositoryviews.StoredUser
	switch view {
	case "default", "":
		p := newStoredUserView(res)
		vres = &repositoryviews.StoredUser{p, "default"}
	case "tiny":
		p := newStoredUserViewTiny(res)
		vres = &repositoryviews.StoredUser{p, "tiny"}
	}
	return vres
}

// newStoredUser converts projected type StoredUser to service type StoredUser.
func newStoredUser(vres *repositoryviews.StoredUserView) *StoredUser {
	res := &StoredUser{
		Online: vres.Online,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newStoredUserTiny converts projected type StoredUser to service type
// StoredUser.
func newStoredUserTiny(vres *repositoryviews.StoredUserView) *StoredUser {
	res := &StoredUser{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newStoredUserView projects result type StoredUser into projected type
// StoredUserView using the "default" view.
func newStoredUserView(res *StoredUser) *repositoryviews.StoredUserView {
	vres := &repositoryviews.StoredUserView{
		ID:     &res.ID,
		Name:   &res.Name,
		Online: res.Online,
	}
	return vres
}

// newStoredUserViewTiny projects result type StoredUser into projected type
// StoredUserView using the "tiny" view.
func newStoredUserViewTiny(res *StoredUser) *repositoryviews.StoredUserView {
	vres := &repositoryviews.StoredUserView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	return vres
}
