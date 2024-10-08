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

// EngineMemberInQueue engine member in queue
//
// swagger:model engineMemberInQueue
type EngineMemberInQueue struct {

	// agent
	Agent *EngineLookup `json:"agent,omitempty"`

	// attempts
	Attempts int32 `json:"attempts,omitempty"`

	// bucket
	Bucket *EngineLookup `json:"bucket,omitempty"`

	// communications
	Communications []*EngineMemberCommunication `json:"communications"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// expire at
	ExpireAt string `json:"expire_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last activity at
	LastActivityAt string `json:"last_activity_at,omitempty"`

	// min offering at
	MinOfferingAt string `json:"min_offering_at,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// queue
	Queue *EngineLookup `json:"queue,omitempty"`

	// reserved
	Reserved bool `json:"reserved,omitempty"`

	// skill
	Skill *EngineLookup `json:"skill,omitempty"`

	// stop cause
	StopCause string `json:"stop_cause,omitempty"`

	// timezone
	Timezone *EngineLookup `json:"timezone,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`
}

// Validate validates this engine member in queue
func (m *EngineMemberInQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommunications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkill(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineMemberInQueue) validateAgent(formats strfmt.Registry) error {
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

func (m *EngineMemberInQueue) validateBucket(formats strfmt.Registry) error {
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

func (m *EngineMemberInQueue) validateCommunications(formats strfmt.Registry) error {
	if swag.IsZero(m.Communications) { // not required
		return nil
	}

	for i := 0; i < len(m.Communications); i++ {
		if swag.IsZero(m.Communications[i]) { // not required
			continue
		}

		if m.Communications[i] != nil {
			if err := m.Communications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("communications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("communications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineMemberInQueue) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberInQueue) validateSkill(formats strfmt.Registry) error {
	if swag.IsZero(m.Skill) { // not required
		return nil
	}

	if m.Skill != nil {
		if err := m.Skill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("skill")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("skill")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberInQueue) validateTimezone(formats strfmt.Registry) error {
	if swag.IsZero(m.Timezone) { // not required
		return nil
	}

	if m.Timezone != nil {
		if err := m.Timezone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timezone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timezone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine member in queue based on the context it is used
func (m *EngineMemberInQueue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBucket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommunications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSkill(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimezone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineMemberInQueue) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineMemberInQueue) contextValidateBucket(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineMemberInQueue) contextValidateCommunications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Communications); i++ {

		if m.Communications[i] != nil {

			if swag.IsZero(m.Communications[i]) { // not required
				return nil
			}

			if err := m.Communications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("communications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("communications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineMemberInQueue) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {

		if swag.IsZero(m.Queue) { // not required
			return nil
		}

		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberInQueue) contextValidateSkill(ctx context.Context, formats strfmt.Registry) error {

	if m.Skill != nil {

		if swag.IsZero(m.Skill) { // not required
			return nil
		}

		if err := m.Skill.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("skill")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("skill")
			}
			return err
		}
	}

	return nil
}

func (m *EngineMemberInQueue) contextValidateTimezone(ctx context.Context, formats strfmt.Registry) error {

	if m.Timezone != nil {

		if swag.IsZero(m.Timezone) { // not required
			return nil
		}

		if err := m.Timezone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timezone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timezone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineMemberInQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineMemberInQueue) UnmarshalBinary(b []byte) error {
	var res EngineMemberInQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
