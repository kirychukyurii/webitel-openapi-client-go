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

// EngineWebHook engine web hook
//
// swagger:model engineWebHook
type EngineWebHook struct {

	// authorization
	Authorization string `json:"authorization,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *EngineLookup `json:"created_by,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// origin
	Origin []string `json:"origin"`

	// schema
	Schema *EngineLookup `json:"schema,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy *EngineLookup `json:"updated_by,omitempty"`
}

// Validate validates this engine web hook
func (m *EngineWebHook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineWebHook) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *EngineWebHook) validateSchema(formats strfmt.Registry) error {
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

func (m *EngineWebHook) validateUpdatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine web hook based on the context it is used
func (m *EngineWebHook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineWebHook) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *EngineWebHook) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineWebHook) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedBy != nil {

		if swag.IsZero(m.UpdatedBy) { // not required
			return nil
		}

		if err := m.UpdatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineWebHook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineWebHook) UnmarshalBinary(b []byte) error {
	var res EngineWebHook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
