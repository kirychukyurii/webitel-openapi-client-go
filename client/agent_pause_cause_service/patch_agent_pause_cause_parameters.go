// Code generated by go-swagger; DO NOT EDIT.

package agent_pause_cause_service

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

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// NewPatchAgentPauseCauseParams creates a new PatchAgentPauseCauseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAgentPauseCauseParams() *PatchAgentPauseCauseParams {
	return &PatchAgentPauseCauseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAgentPauseCauseParamsWithTimeout creates a new PatchAgentPauseCauseParams object
// with the ability to set a timeout on a request.
func NewPatchAgentPauseCauseParamsWithTimeout(timeout time.Duration) *PatchAgentPauseCauseParams {
	return &PatchAgentPauseCauseParams{
		timeout: timeout,
	}
}

// NewPatchAgentPauseCauseParamsWithContext creates a new PatchAgentPauseCauseParams object
// with the ability to set a context for a request.
func NewPatchAgentPauseCauseParamsWithContext(ctx context.Context) *PatchAgentPauseCauseParams {
	return &PatchAgentPauseCauseParams{
		Context: ctx,
	}
}

// NewPatchAgentPauseCauseParamsWithHTTPClient creates a new PatchAgentPauseCauseParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAgentPauseCauseParamsWithHTTPClient(client *http.Client) *PatchAgentPauseCauseParams {
	return &PatchAgentPauseCauseParams{
		HTTPClient: client,
	}
}

/*
PatchAgentPauseCauseParams contains all the parameters to send to the API endpoint

	for the patch agent pause cause operation.

	Typically these are written to a http.Request.
*/
type PatchAgentPauseCauseParams struct {

	// Body.
	Body *models.EnginePatchAgentPauseCauseRequest

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch agent pause cause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAgentPauseCauseParams) WithDefaults() *PatchAgentPauseCauseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch agent pause cause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAgentPauseCauseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) WithTimeout(timeout time.Duration) *PatchAgentPauseCauseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) WithContext(ctx context.Context) *PatchAgentPauseCauseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) WithHTTPClient(client *http.Client) *PatchAgentPauseCauseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) WithBody(body *models.EnginePatchAgentPauseCauseRequest) *PatchAgentPauseCauseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) SetBody(body *models.EnginePatchAgentPauseCauseRequest) {
	o.Body = body
}

// WithID adds the id to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) WithID(id int64) *PatchAgentPauseCauseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch agent pause cause params
func (o *PatchAgentPauseCauseParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAgentPauseCauseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
