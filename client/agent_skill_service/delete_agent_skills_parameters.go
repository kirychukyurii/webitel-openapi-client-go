// Code generated by go-swagger; DO NOT EDIT.

package agent_skill_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteAgentSkillsParams creates a new DeleteAgentSkillsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAgentSkillsParams() *DeleteAgentSkillsParams {
	return &DeleteAgentSkillsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAgentSkillsParamsWithTimeout creates a new DeleteAgentSkillsParams object
// with the ability to set a timeout on a request.
func NewDeleteAgentSkillsParamsWithTimeout(timeout time.Duration) *DeleteAgentSkillsParams {
	return &DeleteAgentSkillsParams{
		timeout: timeout,
	}
}

// NewDeleteAgentSkillsParamsWithContext creates a new DeleteAgentSkillsParams object
// with the ability to set a context for a request.
func NewDeleteAgentSkillsParamsWithContext(ctx context.Context) *DeleteAgentSkillsParams {
	return &DeleteAgentSkillsParams{
		Context: ctx,
	}
}

// NewDeleteAgentSkillsParamsWithHTTPClient creates a new DeleteAgentSkillsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAgentSkillsParamsWithHTTPClient(client *http.Client) *DeleteAgentSkillsParams {
	return &DeleteAgentSkillsParams{
		HTTPClient: client,
	}
}

/*
DeleteAgentSkillsParams contains all the parameters to send to the API endpoint

	for the delete agent skills operation.

	Typically these are written to a http.Request.
*/
type DeleteAgentSkillsParams struct {

	// AgentID.
	//
	// Format: int64
	AgentID string

	// ID.
	ID []string

	// SkillID.
	SkillID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete agent skills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentSkillsParams) WithDefaults() *DeleteAgentSkillsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete agent skills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentSkillsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete agent skills params
func (o *DeleteAgentSkillsParams) WithTimeout(timeout time.Duration) *DeleteAgentSkillsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete agent skills params
func (o *DeleteAgentSkillsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete agent skills params
func (o *DeleteAgentSkillsParams) WithContext(ctx context.Context) *DeleteAgentSkillsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete agent skills params
func (o *DeleteAgentSkillsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete agent skills params
func (o *DeleteAgentSkillsParams) WithHTTPClient(client *http.Client) *DeleteAgentSkillsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete agent skills params
func (o *DeleteAgentSkillsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the delete agent skills params
func (o *DeleteAgentSkillsParams) WithAgentID(agentID string) *DeleteAgentSkillsParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the delete agent skills params
func (o *DeleteAgentSkillsParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithID adds the id to the delete agent skills params
func (o *DeleteAgentSkillsParams) WithID(id []string) *DeleteAgentSkillsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete agent skills params
func (o *DeleteAgentSkillsParams) SetID(id []string) {
	o.ID = id
}

// WithSkillID adds the skillID to the delete agent skills params
func (o *DeleteAgentSkillsParams) WithSkillID(skillID []string) *DeleteAgentSkillsParams {
	o.SetSkillID(skillID)
	return o
}

// SetSkillID adds the skillId to the delete agent skills params
func (o *DeleteAgentSkillsParams) SetSkillID(skillID []string) {
	o.SkillID = skillID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAgentSkillsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent_id
	if err := r.SetPathParam("agent_id", o.AgentID); err != nil {
		return err
	}

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.SkillID != nil {

		// binding items for skill_id
		joinedSkillID := o.bindParamSkillID(reg)

		// query array param skill_id
		if err := r.SetQueryParam("skill_id", joinedSkillID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteAgentSkills binds the parameter id
func (o *DeleteAgentSkillsParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}

// bindParamDeleteAgentSkills binds the parameter skill_id
func (o *DeleteAgentSkillsParams) bindParamSkillID(formats strfmt.Registry) []string {
	skillIDIR := o.SkillID

	var skillIDIC []string
	for _, skillIDIIR := range skillIDIR { // explode []string

		skillIDIIV := skillIDIIR // string as string
		skillIDIC = append(skillIDIC, skillIDIIV)
	}

	// items.CollectionFormat: "multi"
	skillIDIS := swag.JoinByFormat(skillIDIC, "multi")

	return skillIDIS
}
