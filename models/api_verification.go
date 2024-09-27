// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIVerification api verification
//
// swagger:model api.Verification
type APIVerification struct {

	// invalid
	Errors []string `json:"errors"`

	// non-fatal warnings, e.g.:
	Notify []string `json:"notify"`
}

// Validate validates this api verification
func (m *APIVerification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api verification based on context it is used
func (m *APIVerification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIVerification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIVerification) UnmarshalBinary(b []byte) error {
	var res APIVerification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
