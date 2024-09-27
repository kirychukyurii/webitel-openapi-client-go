// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineRedialCallRequest engine redial call request
//
// swagger:model engineRedialCallRequest
type EngineRedialCallRequest struct {

	// call id
	CallID string `json:"call_id,omitempty"`
}

// Validate validates this engine redial call request
func (m *EngineRedialCallRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine redial call request based on context it is used
func (m *EngineRedialCallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineRedialCallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineRedialCallRequest) UnmarshalBinary(b []byte) error {
	var res EngineRedialCallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
