// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// repository views
//
// Command:
// $ goa gen github.com/proepkes/speeddate/usersvc/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa"
)

// StoredUser is the viewed result type that is projected based on a view.
type StoredUser struct {
	// Type to project
	Projected *StoredUserView
	// View to render
	View string
}

// StoredUserView is a type that runs validations on a projected type.
type StoredUserView struct {
	// UUID is the unique id of the user.
	ID *string `gorm:"TYPE:uuid; COLUMN:id; PRIMARY_KEY; DEFAULT: gen_random_uuid()" json:"id"`
	// The username
	Name *string
	// Indicates whether the user is currently online.
	Online *bool
}

// Validate runs the validations defined on the viewed result type StoredUser.
func (result *StoredUser) Validate() (err error) {
	switch result.View {
	case "default", "":
		err = result.Projected.Validate()
	case "tiny":
		err = result.Projected.ValidateTiny()
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// Validate runs the validations defined on StoredUserView using the "default"
// view.
func (result *StoredUserView) Validate() (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 50, false))
		}
	}
	return
}

// ValidateTiny runs the validations defined on StoredUserView using the "tiny"
// view.
func (result *StoredUserView) ValidateTiny() (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 50, false))
		}
	}
	return
}