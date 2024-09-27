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

// EngineQueueBucket engine queue bucket
//
// swagger:model engineQueueBucket
type EngineQueueBucket struct {

	// bucket
	Bucket *EngineLookup `json:"bucket,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// ratio
	Ratio int32 `json:"ratio,omitempty"`
}

// Validate validates this engine queue bucket
func (m *EngineQueueBucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineQueueBucket) validateBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if m.Bucket != nil {
		if err := m.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine queue bucket based on the context it is used
func (m *EngineQueueBucket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBucket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineQueueBucket) contextValidateBucket(ctx context.Context, formats strfmt.Registry) error {

	if m.Bucket != nil {

		if swag.IsZero(m.Bucket) { // not required
			return nil
		}

		if err := m.Bucket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineQueueBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineQueueBucket) UnmarshalBinary(b []byte) error {
	var res EngineQueueBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
