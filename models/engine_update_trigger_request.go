// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineUpdateTriggerRequest engine update trigger request
//
// swagger:model engineUpdateTriggerRequest
type EngineUpdateTriggerRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// expression
	Expression string `json:"expression,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// schema
	Schema *EngineLookup `json:"schema,omitempty"`

	// timeout
	Timeout int32 `json:"timeout,omitempty"`

	// timezone
	Timezone *EngineLookup `json:"timezone,omitempty"`

	// type
	Type *EngineTriggerType `json:"type,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`
}

// Validate validates this engine update trigger request
func (m *EngineUpdateTriggerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineUpdateTriggerRequest) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *EngineUpdateTriggerRequest) validateTimezone(formats strfmt.Registry) error {
	if swag.IsZero(m.Timezone) { // not required
		return nil
	}

	if m.Timezone != nil {
		if err := m.Timezone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timezone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timezone")
			}
			return err
		}
	}

	return nil
}

func (m *EngineUpdateTriggerRequest) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine update trigger request based on the context it is used
func (m *EngineUpdateTriggerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimezone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineUpdateTriggerRequest) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {

		if swag.IsZero(m.Schema) { // not required
			return nil
		}

		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *EngineUpdateTriggerRequest) contextValidateTimezone(ctx context.Context, formats strfmt.Registry) error {

	if m.Timezone != nil {

		if swag.IsZero(m.Timezone) { // not required
			return nil
		}

		if err := m.Timezone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timezone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timezone")
			}
			return err
		}
	}

	return nil
}

func (m *EngineUpdateTriggerRequest) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineUpdateTriggerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineUpdateTriggerRequest) UnmarshalBinary(b []byte) error {
	var res EngineUpdateTriggerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
