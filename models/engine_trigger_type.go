// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EngineTriggerType engine trigger type
//
// swagger:model engineTriggerType
type EngineTriggerType string

func NewEngineTriggerType(value EngineTriggerType) *EngineTriggerType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated EngineTriggerType.
func (m EngineTriggerType) Pointer() *EngineTriggerType {
	return &m
}

const (

	// EngineTriggerTypeDefaultTriggerType captures enum value "default_trigger_type"
	EngineTriggerTypeDefaultTriggerType EngineTriggerType = "default_trigger_type"

	// EngineTriggerTypeCron captures enum value "cron"
	EngineTriggerTypeCron EngineTriggerType = "cron"
)

// for schema
var engineTriggerTypeEnum []interface{}

func init() {
	var res []EngineTriggerType
	if err := json.Unmarshal([]byte(`["default_trigger_type","cron"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		engineTriggerTypeEnum = append(engineTriggerTypeEnum, v)
	}
}

func (m EngineTriggerType) validateEngineTriggerTypeEnum(path, location string, value EngineTriggerType) error {
	if err := validate.EnumCase(path, location, value, engineTriggerTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this engine trigger type
func (m EngineTriggerType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEngineTriggerTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this engine trigger type based on context it is used
func (m EngineTriggerType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
