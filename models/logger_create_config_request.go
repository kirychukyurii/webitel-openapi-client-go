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

// LoggerCreateConfigRequest logger create config request
//
// swagger:model loggerCreateConfigRequest
type LoggerCreateConfigRequest struct {

	// days to store
	DaysToStore int32 `json:"days_to_store,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// object
	Object *LoggerLookup `json:"object,omitempty"`

	// period
	Period int32 `json:"period,omitempty"`

	// storage
	Storage *LoggerLookup `json:"storage,omitempty"`
}

// Validate validates this logger create config request
func (m *LoggerCreateConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoggerCreateConfigRequest) validateObject(formats strfmt.Registry) error {
	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *LoggerCreateConfigRequest) validateStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.Storage) { // not required
		return nil
	}

	if m.Storage != nil {
		if err := m.Storage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this logger create config request based on the context it is used
func (m *LoggerCreateConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoggerCreateConfigRequest) contextValidateObject(ctx context.Context, formats strfmt.Registry) error {

	if m.Object != nil {

		if swag.IsZero(m.Object) { // not required
			return nil
		}

		if err := m.Object.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *LoggerCreateConfigRequest) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.Storage != nil {

		if swag.IsZero(m.Storage) { // not required
			return nil
		}

		if err := m.Storage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoggerCreateConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoggerCreateConfigRequest) UnmarshalBinary(b []byte) error {
	var res LoggerCreateConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
