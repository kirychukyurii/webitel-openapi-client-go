// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListSkillAgentListSkillAgg list skill agent list skill agg
//
// swagger:model ListSkillAgentListSkillAgg
type ListSkillAgentListSkillAgg struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this list skill agent list skill agg
func (m *ListSkillAgentListSkillAgg) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list skill agent list skill agg based on context it is used
func (m *ListSkillAgentListSkillAgg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListSkillAgentListSkillAgg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListSkillAgentListSkillAgg) UnmarshalBinary(b []byte) error {
	var res ListSkillAgentListSkillAgg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
