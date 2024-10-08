// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIClassesUpdateClassBodyClass api classes update class body class
//
// swagger:model apiClassesUpdateClassBodyClass
type APIClassesUpdateClassBodyClass struct {

	// [a]ttribute-[b]ased [a]ccess [c]ontrol
	Abac bool `json:"abac,omitempty"`

	// (class::object).name
	Class string `json:"class,omitempty"`

	// [o]peration-[b]ased [a]ccess [c]ontrol (from::collection.dal)
	Obac bool `json:"obac,omitempty"`

	// [r]ecord-[b]ased [a]ccess [c]ontrol (from::resource.acl)
	Rbac bool `json:"rbac,omitempty"`
}

// Validate validates this api classes update class body class
func (m *APIClassesUpdateClassBodyClass) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api classes update class body class based on context it is used
func (m *APIClassesUpdateClassBodyClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIClassesUpdateClassBodyClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIClassesUpdateClassBodyClass) UnmarshalBinary(b []byte) error {
	var res APIClassesUpdateClassBodyClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
