// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineCreateTriggerJobRequest engine create trigger job request
//
// swagger:model engineCreateTriggerJobRequest
type EngineCreateTriggerJobRequest struct {

	// trigger id
	TriggerID int32 `json:"trigger_id,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`
}

// Validate validates this engine create trigger job request
func (m *EngineCreateTriggerJobRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine create trigger job request based on context it is used
func (m *EngineCreateTriggerJobRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineCreateTriggerJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineCreateTriggerJobRequest) UnmarshalBinary(b []byte) error {
	var res EngineCreateTriggerJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
