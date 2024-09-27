// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartFileTranscriptResponseTranscriptJob start file transcript response transcript job
//
// swagger:model StartFileTranscriptResponseTranscriptJob
type StartFileTranscriptResponseTranscriptJob struct {

	// action
	Action string `json:"action,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// file id
	FileID string `json:"file_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this start file transcript response transcript job
func (m *StartFileTranscriptResponseTranscriptJob) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start file transcript response transcript job based on context it is used
func (m *StartFileTranscriptResponseTranscriptJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StartFileTranscriptResponseTranscriptJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartFileTranscriptResponseTranscriptJob) UnmarshalBinary(b []byte) error {
	var res StartFileTranscriptResponseTranscriptJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
