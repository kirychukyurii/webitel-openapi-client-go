// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebitelChatMessageFile Media File.
//
// swagger:model webitel.chat.MessageFile
type WebitelChatMessageFile struct {

	// File location
	ID string `json:"id,omitempty"`

	// Filename
	Name string `json:"name,omitempty"`

	// Size in bytes
	Size string `json:"size,omitempty"`

	// MIME media type
	Type string `json:"type,omitempty"`
}

// Validate validates this webitel chat message file
func (m *WebitelChatMessageFile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webitel chat message file based on context it is used
func (m *WebitelChatMessageFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebitelChatMessageFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelChatMessageFile) UnmarshalBinary(b []byte) error {
	var res WebitelChatMessageFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
