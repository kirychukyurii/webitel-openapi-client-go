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

// EngineEmailAuthType engine email auth type
//
// swagger:model engineEmailAuthType
type EngineEmailAuthType string

func NewEngineEmailAuthType(value EngineEmailAuthType) *EngineEmailAuthType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated EngineEmailAuthType.
func (m EngineEmailAuthType) Pointer() *EngineEmailAuthType {
	return &m
}

const (

	// EngineEmailAuthTypeEmailAuthTypeUndefined captures enum value "EmailAuthTypeUndefined"
	EngineEmailAuthTypeEmailAuthTypeUndefined EngineEmailAuthType = "EmailAuthTypeUndefined"

	// EngineEmailAuthTypePlain captures enum value "Plain"
	EngineEmailAuthTypePlain EngineEmailAuthType = "Plain"

	// EngineEmailAuthTypeOAuth2 captures enum value "OAuth2"
	EngineEmailAuthTypeOAuth2 EngineEmailAuthType = "OAuth2"
)

// for schema
var engineEmailAuthTypeEnum []interface{}

func init() {
	var res []EngineEmailAuthType
	if err := json.Unmarshal([]byte(`["EmailAuthTypeUndefined","Plain","OAuth2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		engineEmailAuthTypeEnum = append(engineEmailAuthTypeEnum, v)
	}
}

func (m EngineEmailAuthType) validateEngineEmailAuthTypeEnum(path, location string, value EngineEmailAuthType) error {
	if err := validate.EnumCase(path, location, value, engineEmailAuthTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this engine email auth type
func (m EngineEmailAuthType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEngineEmailAuthTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this engine email auth type based on context it is used
func (m EngineEmailAuthType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
