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

// EngineAgentStatusStatisticItem engine agent status statistic item
//
// swagger:model engineAgentStatusStatisticItem
type EngineAgentStatusStatisticItem struct {

	// agent id
	AgentID int32 `json:"agent_id,omitempty"`

	// auditor
	Auditor []*EngineLookup `json:"auditor"`

	// chat count
	ChatCount int64 `json:"chat_count,omitempty"`

	// extension
	Extension string `json:"extension,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// offline
	Offline string `json:"offline,omitempty"`

	// online
	Online string `json:"online,omitempty"`

	// pause
	Pause string `json:"pause,omitempty"`

	// pause cause
	PauseCause string `json:"pause_cause,omitempty"`

	// progressive count
	ProgressiveCount int64 `json:"progressive_count,omitempty"`

	// region
	Region *EngineLookup `json:"region,omitempty"`

	// score count
	ScoreCount string `json:"score_count,omitempty"`

	// score optional avg
	ScoreOptionalAvg float32 `json:"score_optional_avg,omitempty"`

	// score required avg
	ScoreRequiredAvg float32 `json:"score_required_avg,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status duration
	StatusDuration string `json:"status_duration,omitempty"`

	// supervisor
	Supervisor []*EngineLookup `json:"supervisor"`

	// team
	Team *EngineLookup `json:"team,omitempty"`

	// user
	User *EngineLookup `json:"user,omitempty"`
}

// Validate validates this engine agent status statistic item
func (m *EngineAgentStatusStatisticItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupervisor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineAgentStatusStatisticItem) validateAuditor(formats strfmt.Registry) error {
	if swag.IsZero(m.Auditor) { // not required
		return nil
	}

	for i := 0; i < len(m.Auditor); i++ {
		if swag.IsZero(m.Auditor[i]) { // not required
			continue
		}

		if m.Auditor[i] != nil {
			if err := m.Auditor[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineAgentStatusStatisticItem) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *EngineAgentStatusStatisticItem) validateSupervisor(formats strfmt.Registry) error {
	if swag.IsZero(m.Supervisor) { // not required
		return nil
	}

	for i := 0; i < len(m.Supervisor); i++ {
		if swag.IsZero(m.Supervisor[i]) { // not required
			continue
		}

		if m.Supervisor[i] != nil {
			if err := m.Supervisor[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supervisor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supervisor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineAgentStatusStatisticItem) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *EngineAgentStatusStatisticItem) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this engine agent status statistic item based on the context it is used
func (m *EngineAgentStatusStatisticItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupervisor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EngineAgentStatusStatisticItem) contextValidateAuditor(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Auditor); i++ {

		if m.Auditor[i] != nil {

			if swag.IsZero(m.Auditor[i]) { // not required
				return nil
			}

			if err := m.Auditor[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineAgentStatusStatisticItem) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if swag.IsZero(m.Region) { // not required
			return nil
		}

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *EngineAgentStatusStatisticItem) contextValidateSupervisor(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Supervisor); i++ {

		if m.Supervisor[i] != nil {

			if swag.IsZero(m.Supervisor[i]) { // not required
				return nil
			}

			if err := m.Supervisor[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supervisor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supervisor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineAgentStatusStatisticItem) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.Team != nil {

		if swag.IsZero(m.Team) { // not required
			return nil
		}

		if err := m.Team.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *EngineAgentStatusStatisticItem) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EngineAgentStatusStatisticItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineAgentStatusStatisticItem) UnmarshalBinary(b []byte) error {
	var res EngineAgentStatusStatisticItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
