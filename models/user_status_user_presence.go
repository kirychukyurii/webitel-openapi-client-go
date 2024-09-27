// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserStatusUserPresence user status user presence
//
// swagger:model UserStatusUserPresence
type UserStatusUserPresence struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this user status user presence
func (m *UserStatusUserPresence) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user status user presence based on context it is used
func (m *UserStatusUserPresence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserStatusUserPresence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserStatusUserPresence) UnmarshalBinary(b []byte) error {
	var res UserStatusUserPresence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
