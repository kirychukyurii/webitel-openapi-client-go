// Code generated by go-swagger; DO NOT EDIT.

package agent_working_conditions_service

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
)

// NewAgentWorkingConditionsServiceReadAgentWorkingConditionsParams creates a new AgentWorkingConditionsServiceReadAgentWorkingConditionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAgentWorkingConditionsServiceReadAgentWorkingConditionsParams() *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	return &AgentWorkingConditionsServiceReadAgentWorkingConditionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAgentWorkingConditionsServiceReadAgentWorkingConditionsParamsWithTimeout creates a new AgentWorkingConditionsServiceReadAgentWorkingConditionsParams object
// with the ability to set a timeout on a request.
func NewAgentWorkingConditionsServiceReadAgentWorkingConditionsParamsWithTimeout(timeout time.Duration) *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	return &AgentWorkingConditionsServiceReadAgentWorkingConditionsParams{
		timeout: timeout,
	}
}

// NewAgentWorkingConditionsServiceReadAgentWorkingConditionsParamsWithContext creates a new AgentWorkingConditionsServiceReadAgentWorkingConditionsParams object
// with the ability to set a context for a request.
func NewAgentWorkingConditionsServiceReadAgentWorkingConditionsParamsWithContext(ctx context.Context) *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	return &AgentWorkingConditionsServiceReadAgentWorkingConditionsParams{
		Context: ctx,
	}
}

// NewAgentWorkingConditionsServiceReadAgentWorkingConditionsParamsWithHTTPClient creates a new AgentWorkingConditionsServiceReadAgentWorkingConditionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAgentWorkingConditionsServiceReadAgentWorkingConditionsParamsWithHTTPClient(client *http.Client) *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	return &AgentWorkingConditionsServiceReadAgentWorkingConditionsParams{
		HTTPClient: client,
	}
}

/*
AgentWorkingConditionsServiceReadAgentWorkingConditionsParams contains all the parameters to send to the API endpoint

	for the agent working conditions service read agent working conditions operation.

	Typically these are written to a http.Request.
*/
type AgentWorkingConditionsServiceReadAgentWorkingConditionsParams struct {

	// AgentID.
	//
	// Format: int64
	AgentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the agent working conditions service read agent working conditions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) WithDefaults() *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the agent working conditions service read agent working conditions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the agent working conditions service read agent working conditions params
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) WithTimeout(timeout time.Duration) *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the agent working conditions service read agent working conditions params
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the agent working conditions service read agent working conditions params
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) WithContext(ctx context.Context) *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the agent working conditions service read agent working conditions params
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the agent working conditions service read agent working conditions params
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) WithHTTPClient(client *http.Client) *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the agent working conditions service read agent working conditions params
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the agent working conditions service read agent working conditions params
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) WithAgentID(agentID string) *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the agent working conditions service read agent working conditions params
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WriteToRequest writes these params to a swagger request
func (o *AgentWorkingConditionsServiceReadAgentWorkingConditionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent_id
	if err := r.SetPathParam("agent_id", o.AgentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
