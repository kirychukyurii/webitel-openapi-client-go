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

// EngineUpdateAgentTeamRequest engine update agent team request
//
// swagger:model engineUpdateAgentTeamRequest
type EngineUpdateAgentTeamRequest struct {

	// admin
	Admin []*EngineLookup `json:"admin"`

	// call timeout
	CallTimeout int32 `json:"call_timeout,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invite chat timeout
	InviteChatTimeout int32 `json:"invite_chat_timeout,omitempty"`

	// max no answer
	MaxNoAnswer int32 `json:"max_no_answer,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// no answer delay time
	NoAnswerDelayTime int32 `json:"no_answer_delay_time,omitempty"`

	// strategy
	Strategy string `json:"strategy,omitempty"`

	// task accept timeout
	TaskAcceptTimeout int32 `json:"task_accept_timeout,omitempty"`

	// wrap up time
	WrapUpTime int32 `json:"wrap_up_time,omitempty"`
}

// Validate validates this engine update agent team request
func (m *EngineUpdateAgentTeamRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineUpdateAgentTeamRequest) validateAdmin(formats strfmt.Registry) error {
	if swag.IsZero(m.Admin) { // not required
		return nil
	}

	for i := 0; i < len(m.Admin); i++ {
		if swag.IsZero(m.Admin[i]) { // not required
			continue
		}

		if m.Admin[i] != nil {
			if err := m.Admin[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("admin" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("admin" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this engine update agent team request based on the context it is used
func (m *EngineUpdateAgentTeamRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdmin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineUpdateAgentTeamRequest) contextValidateAdmin(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Admin); i++ {

		if m.Admin[i] != nil {

			if swag.IsZero(m.Admin[i]) { // not required
				return nil
			}

			if err := m.Admin[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("admin" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("admin" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineUpdateAgentTeamRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineUpdateAgentTeamRequest) UnmarshalBinary(b []byte) error {
	var res EngineUpdateAgentTeamRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
