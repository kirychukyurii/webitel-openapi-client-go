// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WfmDeleteForecastCalculationResponse wfm delete forecast calculation response
//
// swagger:model wfmDeleteForecastCalculationResponse
type WfmDeleteForecastCalculationResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this wfm delete forecast calculation response
func (m *WfmDeleteForecastCalculationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this wfm delete forecast calculation response based on context it is used
func (m *WfmDeleteForecastCalculationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WfmDeleteForecastCalculationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmDeleteForecastCalculationResponse) UnmarshalBinary(b []byte) error {
	var res WfmDeleteForecastCalculationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
