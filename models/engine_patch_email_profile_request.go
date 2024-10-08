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

// EnginePatchEmailProfileRequest engine patch email profile request
//
// swagger:model enginePatchEmailProfileRequest
type EnginePatchEmailProfileRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// fetch interval
	FetchInterval int32 `json:"fetch_interval,omitempty"`

	// fields
	Fields []string `json:"fields"`

	// id
	ID string `json:"id,omitempty"`

	// imap host
	ImapHost string `json:"imap_host,omitempty"`

	// imap port
	ImapPort int32 `json:"imap_port,omitempty"`

	// listen
	Listen bool `json:"listen,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// mailbox
	Mailbox string `json:"mailbox,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// schema
	Schema *EngineLookup `json:"schema,omitempty"`

	// smtp host
	SMTPHost string `json:"smtp_host,omitempty"`

	// smtp port
	SMTPPort int32 `json:"smtp_port,omitempty"`
}

// Validate validates this engine patch email profile request
func (m *EnginePatchEmailProfileRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnginePatchEmailProfileRequest) validateSchema(formats strfmt.Registry) error {
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

// ContextValidate validate this engine patch email profile request based on the context it is used
func (m *EnginePatchEmailProfileRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnginePatchEmailProfileRequest) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *EnginePatchEmailProfileRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnginePatchEmailProfileRequest) UnmarshalBinary(b []byte) error {
	var res EnginePatchEmailProfileRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
