// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageTranscriptPhrase storage transcript phrase
//
// swagger:model storageTranscriptPhrase
type StorageTranscriptPhrase struct {

	// channel
	Channel int64 `json:"channel,omitempty"`

	// end sec
	EndSec float32 `json:"end_sec,omitempty"`

	// phrase
	Phrase string `json:"phrase,omitempty"`

	// start sec
	StartSec float32 `json:"start_sec,omitempty"`
}

// Validate validates this storage transcript phrase
func (m *StorageTranscriptPhrase) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage transcript phrase based on context it is used
func (m *StorageTranscriptPhrase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageTranscriptPhrase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageTranscriptPhrase) UnmarshalBinary(b []byte) error {
	var res StorageTranscriptPhrase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
