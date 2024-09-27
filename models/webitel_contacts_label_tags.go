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

// WebitelContactsLabelTags LabelTags dataset list of label's info.
//
// swagger:model webitel.contacts.LabelTags
type WebitelContactsLabelTags struct {

	// Labels info dataset
	Labels []*WebitelContactsLabelInfo `json:"labels"`

	// Next page available ?
	Next bool `json:"next,omitempty"`

	// Page number of partial result.
	Page int32 `json:"page,omitempty"`

	// Contact sources dataset
	Sources []*WebitelContactsLookup `json:"sources"`
}

// Validate validates this webitel contacts label tags
func (m *WebitelContactsLabelTags) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsLabelTags) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebitelContactsLabelTags) validateSources(formats strfmt.Registry) error {
	if swag.IsZero(m.Sources) { // not required
		return nil
	}

	for i := 0; i < len(m.Sources); i++ {
		if swag.IsZero(m.Sources[i]) { // not required
			continue
		}

		if m.Sources[i] != nil {
			if err := m.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this webitel contacts label tags based on the context it is used
func (m *WebitelContactsLabelTags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsLabelTags) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {

			if swag.IsZero(m.Labels[i]) { // not required
				return nil
			}

			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebitelContactsLabelTags) contextValidateSources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sources); i++ {

		if m.Sources[i] != nil {

			if swag.IsZero(m.Sources[i]) { // not required
				return nil
			}

			if err := m.Sources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebitelContactsLabelTags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsLabelTags) UnmarshalBinary(b []byte) error {
	var res WebitelContactsLabelTags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
