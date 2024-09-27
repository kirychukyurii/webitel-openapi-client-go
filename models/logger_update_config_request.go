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

// LoggerUpdateConfigRequest logger update config request
//
// swagger:model loggerUpdateConfigRequest
type LoggerUpdateConfigRequest struct {

	// config id
	ConfigID int32 `json:"config_id,omitempty"`

	// days to store
	DaysToStore int32 `json:"days_to_store,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// period
	Period int32 `json:"period,omitempty"`

	// storage
	Storage *LoggerLookup `json:"storage,omitempty"`
}

// Validate validates this logger update config request
func (m *LoggerUpdateConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoggerUpdateConfigRequest) validateStorage(formats strfmt.Registry) error {
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

// ContextValidate validate this logger update config request based on the context it is used
func (m *LoggerUpdateConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoggerUpdateConfigRequest) contextValidateStorage(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LoggerUpdateConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoggerUpdateConfigRequest) UnmarshalBinary(b []byte) error {
	var res LoggerUpdateConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
