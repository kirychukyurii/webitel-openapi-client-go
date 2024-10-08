// Code generated by go-swagger; DO NOT EDIT.

package team_trigger_service

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

// NewCreateTeamTriggerParams creates a new CreateTeamTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTeamTriggerParams() *CreateTeamTriggerParams {
	return &CreateTeamTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTeamTriggerParamsWithTimeout creates a new CreateTeamTriggerParams object
// with the ability to set a timeout on a request.
func NewCreateTeamTriggerParamsWithTimeout(timeout time.Duration) *CreateTeamTriggerParams {
	return &CreateTeamTriggerParams{
		timeout: timeout,
	}
}

// NewCreateTeamTriggerParamsWithContext creates a new CreateTeamTriggerParams object
// with the ability to set a context for a request.
func NewCreateTeamTriggerParamsWithContext(ctx context.Context) *CreateTeamTriggerParams {
	return &CreateTeamTriggerParams{
		Context: ctx,
	}
}

// NewCreateTeamTriggerParamsWithHTTPClient creates a new CreateTeamTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTeamTriggerParamsWithHTTPClient(client *http.Client) *CreateTeamTriggerParams {
	return &CreateTeamTriggerParams{
		HTTPClient: client,
	}
}

/*
CreateTeamTriggerParams contains all the parameters to send to the API endpoint

	for the create team trigger operation.

	Typically these are written to a http.Request.
*/
type CreateTeamTriggerParams struct {

	// Body.
	Body *models.EngineCreateTeamTriggerRequest

	// TeamID.
	//
	// Format: int64
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create team trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTeamTriggerParams) WithDefaults() *CreateTeamTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create team trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTeamTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create team trigger params
func (o *CreateTeamTriggerParams) WithTimeout(timeout time.Duration) *CreateTeamTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create team trigger params
func (o *CreateTeamTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create team trigger params
func (o *CreateTeamTriggerParams) WithContext(ctx context.Context) *CreateTeamTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create team trigger params
func (o *CreateTeamTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create team trigger params
func (o *CreateTeamTriggerParams) WithHTTPClient(client *http.Client) *CreateTeamTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create team trigger params
func (o *CreateTeamTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create team trigger params
func (o *CreateTeamTriggerParams) WithBody(body *models.EngineCreateTeamTriggerRequest) *CreateTeamTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create team trigger params
func (o *CreateTeamTriggerParams) SetBody(body *models.EngineCreateTeamTriggerRequest) {
	o.Body = body
}

// WithTeamID adds the teamID to the create team trigger params
func (o *CreateTeamTriggerParams) WithTeamID(teamID string) *CreateTeamTriggerParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the create team trigger params
func (o *CreateTeamTriggerParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTeamTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
