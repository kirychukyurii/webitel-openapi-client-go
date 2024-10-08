// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIDeviceID DeviceId lookup value
//
// swagger:model api.DeviceId
type APIDeviceID struct {

	// Object ID
	ID string `json:"id,omitempty"`

	// Display Name
	Name string `json:"name,omitempty"`
}

// Validate validates this api device Id
func (m *APIDeviceID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api device Id based on context it is used
func (m *APIDeviceID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIDeviceID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIDeviceID) UnmarshalBinary(b []byte) error {
	var res APIDeviceID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
