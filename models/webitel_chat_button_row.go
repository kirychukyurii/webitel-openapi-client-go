// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebitelChatButtonRow webitel chat button row
//
// swagger:model webitel.chat.ButtonRow
type WebitelChatButtonRow struct {

	// Button(s) in a row
	Row []*WebitelChatButton `json:"row"`
}

// Validate validates this webitel chat button row
func (m *WebitelChatButtonRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelChatButtonRow) validateRow(formats strfmt.Registry) error {
	if swag.IsZero(m.Row) { // not required
		return nil
	}

	for i := 0; i < len(m.Row); i++ {
		if swag.IsZero(m.Row[i]) { // not required
			continue
		}

		if m.Row[i] != nil {
			if err := m.Row[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("row" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("row" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this webitel chat button row based on the context it is used
func (m *WebitelChatButtonRow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelChatButtonRow) contextValidateRow(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Row); i++ {

		if m.Row[i] != nil {

			if swag.IsZero(m.Row[i]) { // not required
				return nil
			}

			if err := m.Row[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("row" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("row" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebitelChatButtonRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelChatButtonRow) UnmarshalBinary(b []byte) error {
	var res WebitelChatButtonRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
