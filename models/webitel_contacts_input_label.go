// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebitelContactsInputLabel A Contact's associated Tag.
// Output purpose only.
//
// swagger:model webitel.contacts.InputLabel
type WebitelContactsInputLabel struct {

	// Unique ID of the latest version of an existing resorce.
	Etag string `json:"etag,omitempty"`

	// REQUIRED. Hashtag value;
	// NOTE: Keep in mind, hashtags are not case-sensitive,
	// but adding capital letters does make them easier to read:
	// #MakeAWish vs. #makeawish.
	Label string `json:"label,omitempty"`
}

// Validate validates this webitel contacts input label
func (m *WebitelContactsInputLabel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webitel contacts input label based on context it is used
func (m *WebitelContactsInputLabel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebitelContactsInputLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsInputLabel) UnmarshalBinary(b []byte) error {
	var res WebitelContactsInputLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
