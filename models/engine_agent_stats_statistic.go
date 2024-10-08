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

// EngineAgentStatsStatistic engine agent stats statistic
//
// swagger:model engineAgentStatsStatistic
type EngineAgentStatsStatistic struct {

	// active call id
	ActiveCallID string `json:"active_call_id,omitempty"`

	// agent id
	AgentID int32 `json:"agent_id,omitempty"`

	// auditor
	Auditor []*EngineLookup `json:"auditor"`

	// call time
	CallTime string `json:"call_time,omitempty"`

	// chat count
	ChatCount int32 `json:"chat_count,omitempty"`

	// extension
	Extension string `json:"extension,omitempty"`

	// handles
	Handles int32 `json:"handles,omitempty"`

	// max bridged at
	MaxBridgedAt string `json:"max_bridged_at,omitempty"`

	// max offering at
	MaxOfferingAt string `json:"max_offering_at,omitempty"`

	// missed
	Missed int32 `json:"missed,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// occupancy
	Occupancy float32 `json:"occupancy,omitempty"`

	// offline
	Offline string `json:"offline,omitempty"`

	// online
	Online string `json:"online,omitempty"`

	// pause
	Pause string `json:"pause,omitempty"`

	// pause cause
	PauseCause string `json:"pause_cause,omitempty"`

	// queues
	Queues []*EngineLookup `json:"queues"`

	// skills
	Skills []*EngineLookup `json:"skills"`

	// status
	Status string `json:"status,omitempty"`

	// status duration
	StatusDuration string `json:"status_duration,omitempty"`

	// supervisor
	Supervisor []*EngineLookup `json:"supervisor"`

	// team
	Team *EngineLookup `json:"team,omitempty"`

	// transferred
	Transferred int64 `json:"transferred,omitempty"`

	// user
	User *EngineLookup `json:"user,omitempty"`

	// utilization
	Utilization float32 `json:"utilization,omitempty"`
}

// Validate validates this engine agent stats statistic
func (m *EngineAgentStatsStatistic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkills(formats); err != nil {
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

func (m *EngineAgentStatsStatistic) validateAuditor(formats strfmt.Registry) error {
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

func (m *EngineAgentStatsStatistic) validateQueues(formats strfmt.Registry) error {
	if swag.IsZero(m.Queues) { // not required
		return nil
	}

	for i := 0; i < len(m.Queues); i++ {
		if swag.IsZero(m.Queues[i]) { // not required
			continue
		}

		if m.Queues[i] != nil {
			if err := m.Queues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineAgentStatsStatistic) validateSkills(formats strfmt.Registry) error {
	if swag.IsZero(m.Skills) { // not required
		return nil
	}

	for i := 0; i < len(m.Skills); i++ {
		if swag.IsZero(m.Skills[i]) { // not required
			continue
		}

		if m.Skills[i] != nil {
			if err := m.Skills[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skills" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineAgentStatsStatistic) validateSupervisor(formats strfmt.Registry) error {
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

func (m *EngineAgentStatsStatistic) validateTeam(formats strfmt.Registry) error {
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

func (m *EngineAgentStatsStatistic) validateUser(formats strfmt.Registry) error {
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

// ContextValidate validate this engine agent stats statistic based on the context it is used
func (m *EngineAgentStatsStatistic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSkills(ctx, formats); err != nil {
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

func (m *EngineAgentStatsStatistic) contextValidateAuditor(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineAgentStatsStatistic) contextValidateQueues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Queues); i++ {

		if m.Queues[i] != nil {

			if swag.IsZero(m.Queues[i]) { // not required
				return nil
			}

			if err := m.Queues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineAgentStatsStatistic) contextValidateSkills(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Skills); i++ {

		if m.Skills[i] != nil {

			if swag.IsZero(m.Skills[i]) { // not required
				return nil
			}

			if err := m.Skills[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skills" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EngineAgentStatsStatistic) contextValidateSupervisor(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineAgentStatsStatistic) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EngineAgentStatsStatistic) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EngineAgentStatsStatistic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineAgentStatsStatistic) UnmarshalBinary(b []byte) error {
	var res EngineAgentStatsStatistic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
