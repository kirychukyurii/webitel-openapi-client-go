// Code generated by go-swagger; DO NOT EDIT.

package skill_service

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

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// NewCreateSkillAgentParams creates a new CreateSkillAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSkillAgentParams() *CreateSkillAgentParams {
	return &CreateSkillAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSkillAgentParamsWithTimeout creates a new CreateSkillAgentParams object
// with the ability to set a timeout on a request.
func NewCreateSkillAgentParamsWithTimeout(timeout time.Duration) *CreateSkillAgentParams {
	return &CreateSkillAgentParams{
		timeout: timeout,
	}
}

// NewCreateSkillAgentParamsWithContext creates a new CreateSkillAgentParams object
// with the ability to set a context for a request.
func NewCreateSkillAgentParamsWithContext(ctx context.Context) *CreateSkillAgentParams {
	return &CreateSkillAgentParams{
		Context: ctx,
	}
}

// NewCreateSkillAgentParamsWithHTTPClient creates a new CreateSkillAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSkillAgentParamsWithHTTPClient(client *http.Client) *CreateSkillAgentParams {
	return &CreateSkillAgentParams{
		HTTPClient: client,
	}
}

/*
CreateSkillAgentParams contains all the parameters to send to the API endpoint

	for the create skill agent operation.

	Typically these are written to a http.Request.
*/
type CreateSkillAgentParams struct {

	// Body.
	Body *models.EngineCreateSkillAgentRequest

	// SkillID.
	//
	// Format: int64
	SkillID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create skill agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSkillAgentParams) WithDefaults() *CreateSkillAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create skill agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSkillAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create skill agent params
func (o *CreateSkillAgentParams) WithTimeout(timeout time.Duration) *CreateSkillAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create skill agent params
func (o *CreateSkillAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create skill agent params
func (o *CreateSkillAgentParams) WithContext(ctx context.Context) *CreateSkillAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create skill agent params
func (o *CreateSkillAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create skill agent params
func (o *CreateSkillAgentParams) WithHTTPClient(client *http.Client) *CreateSkillAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create skill agent params
func (o *CreateSkillAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create skill agent params
func (o *CreateSkillAgentParams) WithBody(body *models.EngineCreateSkillAgentRequest) *CreateSkillAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create skill agent params
func (o *CreateSkillAgentParams) SetBody(body *models.EngineCreateSkillAgentRequest) {
	o.Body = body
}

// WithSkillID adds the skillID to the create skill agent params
func (o *CreateSkillAgentParams) WithSkillID(skillID string) *CreateSkillAgentParams {
	o.SetSkillID(skillID)
	return o
}

// SetSkillID adds the skillId to the create skill agent params
func (o *CreateSkillAgentParams) SetSkillID(skillID string) {
	o.SkillID = skillID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSkillAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param skill_id
	if err := r.SetPathParam("skill_id", o.SkillID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
