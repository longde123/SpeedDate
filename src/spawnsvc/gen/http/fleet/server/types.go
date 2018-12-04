// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// fleet HTTP server types
//
// Command:
// $ goa gen github.com/proepkes/speeddate/src/spawnsvc/design

package server

import (
	"unicode/utf8"

	fleet "github.com/proepkes/speeddate/src/spawnsvc/gen/fleet"
	goa "goa.design/goa"
)

// ConfigureRequestBody is the type of the "fleet" service "configure" endpoint
// HTTP request body.
type ConfigureRequestBody struct {
	// Namespace where the gameserver will run in
	Namespace *string `form:"Namespace,omitempty" json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Prefix for the generated pod-name
	NamePrefix *string `form:"NamePrefix,omitempty" json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// Portpolicy either dynamic or static
	PortPolicy *string `form:"PortPolicy,omitempty" json:"PortPolicy,omitempty" xml:"PortPolicy,omitempty"`
	// Name of the gameserver-container
	ContainerName *string `form:"ContainerName,omitempty" json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// Image of the gameserver
	ContainerImage *string `form:"ContainerImage,omitempty" json:"ContainerImage,omitempty" xml:"ContainerImage,omitempty"`
	// Exposed port of the gameserver
	ContainerPort *string `form:"ContainerPort,omitempty" json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
}

// ConfigurationResponseBody is the type of the "fleet" service "configuration"
// endpoint HTTP response body.
type ConfigurationResponseBody struct {
	// Namespace where the gameserver will run in
	Namespace string `form:"Namespace" json:"Namespace" xml:"Namespace"`
	// Prefix for the generated pod-name
	NamePrefix string `form:"NamePrefix" json:"NamePrefix" xml:"NamePrefix"`
	// Portpolicy either dynamic or static
	PortPolicy string `form:"PortPolicy" json:"PortPolicy" xml:"PortPolicy"`
	// Name of the gameserver-container
	ContainerName string `form:"ContainerName" json:"ContainerName" xml:"ContainerName"`
	// Image of the gameserver
	ContainerImage string `form:"ContainerImage" json:"ContainerImage" xml:"ContainerImage"`
	// Exposed port of the gameserver
	ContainerPort string `form:"ContainerPort" json:"ContainerPort" xml:"ContainerPort"`
}

// NewConfigurationResponseBody builds the HTTP response body from the result
// of the "configuration" endpoint of the "fleet" service.
func NewConfigurationResponseBody(res *fleet.GameserverTemplate) *ConfigurationResponseBody {
	body := &ConfigurationResponseBody{
		Namespace:      res.Namespace,
		NamePrefix:     res.NamePrefix,
		PortPolicy:     res.PortPolicy,
		ContainerName:  res.ContainerName,
		ContainerImage: res.ContainerImage,
		ContainerPort:  res.ContainerPort,
	}
	return body
}

// NewConfigureGameserverTemplate builds a fleet service configure endpoint
// payload.
func NewConfigureGameserverTemplate(body *ConfigureRequestBody) *fleet.GameserverTemplate {
	v := &fleet.GameserverTemplate{
		Namespace:      *body.Namespace,
		NamePrefix:     *body.NamePrefix,
		PortPolicy:     *body.PortPolicy,
		ContainerName:  *body.ContainerName,
		ContainerImage: *body.ContainerImage,
		ContainerPort:  *body.ContainerPort,
	}
	return v
}

// Validate runs the validations defined on ConfigureRequestBody
func (body *ConfigureRequestBody) Validate() (err error) {
	if body.Namespace == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Namespace", "body"))
	}
	if body.NamePrefix == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("NamePrefix", "body"))
	}
	if body.PortPolicy == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("PortPolicy", "body"))
	}
	if body.ContainerName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ContainerName", "body"))
	}
	if body.ContainerImage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ContainerImage", "body"))
	}
	if body.ContainerPort == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ContainerPort", "body"))
	}
	if body.Namespace != nil {
		if utf8.RuneCountInString(*body.Namespace) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.Namespace", *body.Namespace, utf8.RuneCountInString(*body.Namespace), 100, false))
		}
	}
	if body.NamePrefix != nil {
		if utf8.RuneCountInString(*body.NamePrefix) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.NamePrefix", *body.NamePrefix, utf8.RuneCountInString(*body.NamePrefix), 100, false))
		}
	}
	return
}
