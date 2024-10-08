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

// WebitelContactsLanguage A Contact's locale preference.
// Output purpose only.
//
// swagger:model webitel.contacts.Language
type WebitelContactsLanguage struct {

	// uk              ; golang.org/x/text/language.MustParse(s).Base()
	Code string `json:"code,omitempty"`

	// The user who created this Field.
	CreatedAt string `json:"created_at,omitempty"`

	// Timestamp(milli) of the Field creation.
	CreatedBy *WebitelContactsLookup `json:"created_by,omitempty"`

	// Unique ID of the latest version of the update.
	// This ID changes after any update to the underlying value(s).
	Etag string `json:"etag,omitempty"`

	// The unique ID of the association. Never changes.
	ID string `json:"id,omitempty"`

	// Ukrainian       ; golang.org/x/text/language/display.English.Languages().Name(tag)
	Lang string `json:"lang,omitempty"`

	// українська      ; golang.org/x/text/language/display.Self().Name(tag)
	Name string `json:"name,omitempty"`

	// Indicates whether this association is the default
	// among others of the same type.
	Primary bool `json:"primary,omitempty"`

	// UA              ; golang.org/x/text/language.MustParse(s).Region()
	Region string `json:"region,omitempty"`

	// Cyrl; (Cyrillic); golang.org/x/text/language.MustParse(s).Script()
	Script string `json:"script,omitempty"`

	// The well-formed IETF BCP 47 language tag representing the locale.
	// https://www.rfc-editor.org/info/bcp47
	Tag string `json:"tag,omitempty"`

	// Timestamp(milli) of the last Field update.
	// Take part in Etag generation.
	UpdatedAt string `json:"updated_at,omitempty"`

	// The user who performed last Update.
	UpdatedBy *WebitelContactsLookup `json:"updated_by,omitempty"`

	// Version of the latest update. Numeric sequence.
	Ver int32 `json:"ver,omitempty"`
}

// Validate validates this webitel contacts language
func (m *WebitelContactsLanguage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
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

func (m *WebitelContactsLanguage) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *WebitelContactsLanguage) validateUpdatedBy(formats strfmt.Registry) error {
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

// ContextValidate validate this webitel contacts language based on the context it is used
func (m *WebitelContactsLanguage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
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

func (m *WebitelContactsLanguage) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WebitelContactsLanguage) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

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
func (m *WebitelContactsLanguage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsLanguage) UnmarshalBinary(b []byte) error {
	var res WebitelContactsLanguage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
