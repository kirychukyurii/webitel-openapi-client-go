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

// EngineMemberAttempt engine member attempt
//
// swagger:model engineMemberAttempt
type EngineMemberAttempt struct {

	// active
	Active bool `json:"active,omitempty"`

	// agent
	Agent *EngineLookup `json:"agent,omitempty"`

	// answered at
	AnsweredAt string `json:"answered_at,omitempty"`

	// attempts
	Attempts int32 `json:"attempts,omitempty"`

	// bridged at
	BridgedAt string `json:"bridged_at,omitempty"`

	// bucket
	Bucket *EngineLookup `json:"bucket,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// destination
	Destination string `json:"destination,omitempty"`

	// hangup at
	HangupAt string `json:"hangup_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// leg a id
	LegaID string `json:"leg_a_id,omitempty"`

	// leg b id
	LegbID string `json:"leg_b_id,omitempty"`

	// logs
	Logs interface{} `json:"logs,omitempty"`

	// member
	Member *EngineLookup `json:"member,omitempty"`

	// node
	Node string `json:"node,omitempty"`

	// originate at
	OriginateAt string `json:"originate_at,omitempty"`

	// resource
	Resource *EngineLookup `json:"resource,omitempty"`

	// result
	Result string `json:"result,omitempty"`

	// weight
	Weight int32 `json:"weight,omitempty"`
}

// Validate validates this engine member attempt
func (m *EngineMemberAttempt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineMemberAttempt) validateAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberAttempt) validateBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if m.Bucket != nil {
		if err := m.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberAttempt) validateMember(formats strfmt.Registry) error {
	if swag.IsZero(m.Member) { // not required
		return nil
	}

	if m.Member != nil {
		if err := m.Member.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberAttempt) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine member attempt based on the context it is used
func (m *EngineMemberAttempt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBucket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMember(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineMemberAttempt) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.Agent != nil {

		if swag.IsZero(m.Agent) { // not required
			return nil
		}

		if err := m.Agent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberAttempt) contextValidateBucket(ctx context.Context, formats strfmt.Registry) error {

	if m.Bucket != nil {

		if swag.IsZero(m.Bucket) { // not required
			return nil
		}

		if err := m.Bucket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberAttempt) contextValidateMember(ctx context.Context, formats strfmt.Registry) error {

	if m.Member != nil {

		if swag.IsZero(m.Member) { // not required
			return nil
		}

		if err := m.Member.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberAttempt) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {

		if swag.IsZero(m.Resource) { // not required
			return nil
		}

		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineMemberAttempt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineMemberAttempt) UnmarshalBinary(b []byte) error {
	var res EngineMemberAttempt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
