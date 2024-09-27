// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineUpdateCallAnnotationRequest engine update call annotation request
//
// swagger:model engineUpdateCallAnnotationRequest
type EngineUpdateCallAnnotationRequest struct {

	// call id
	CallID string `json:"call_id,omitempty"`

	// end sec
	EndSec int32 `json:"end_sec,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// note
	Note string `json:"note,omitempty"`

	// start sec
	StartSec int32 `json:"start_sec,omitempty"`
}

// Validate validates this engine update call annotation request
func (m *EngineUpdateCallAnnotationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine update call annotation request based on context it is used
func (m *EngineUpdateCallAnnotationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineUpdateCallAnnotationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineUpdateCallAnnotationRequest) UnmarshalBinary(b []byte) error {
	var res EngineUpdateCallAnnotationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
