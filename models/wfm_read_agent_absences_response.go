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

// WfmReadAgentAbsencesResponse wfm read agent absences response
//
// swagger:model wfmReadAgentAbsencesResponse
type WfmReadAgentAbsencesResponse struct {

	// item
	Item *WfmAgentAbsences `json:"item,omitempty"`
}

// Validate validates this wfm read agent absences response
func (m *WfmReadAgentAbsencesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmReadAgentAbsencesResponse) validateItem(formats strfmt.Registry) error {
	if swag.IsZero(m.Item) { // not required
		return nil
	}

	if m.Item != nil {
		if err := m.Item.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this wfm read agent absences response based on the context it is used
func (m *WfmReadAgentAbsencesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmReadAgentAbsencesResponse) contextValidateItem(ctx context.Context, formats strfmt.Registry) error {

	if m.Item != nil {

		if swag.IsZero(m.Item) { // not required
			return nil
		}

		if err := m.Item.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WfmReadAgentAbsencesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmReadAgentAbsencesResponse) UnmarshalBinary(b []byte) error {
	var res WfmReadAgentAbsencesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
