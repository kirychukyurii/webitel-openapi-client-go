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

// WebitelContactsTranscriptLookup webitel contacts transcript lookup
//
// swagger:model webitel.contacts.TranscriptLookup
type WebitelContactsTranscriptLookup struct {

	// file
	File *WebitelContactsLookup `json:"file,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// locale
	Locale string `json:"locale,omitempty"`
}

// Validate validates this webitel contacts transcript lookup
func (m *WebitelContactsTranscriptLookup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsTranscriptLookup) validateFile(formats strfmt.Registry) error {
	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {
		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this webitel contacts transcript lookup based on the context it is used
func (m *WebitelContactsTranscriptLookup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsTranscriptLookup) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

	if m.File != nil {

		if swag.IsZero(m.File) { // not required
			return nil
		}

		if err := m.File.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebitelContactsTranscriptLookup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsTranscriptLookup) UnmarshalBinary(b []byte) error {
	var res WebitelContactsTranscriptLookup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
