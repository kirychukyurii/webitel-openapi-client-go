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
)

// NewReadAgentPauseCauseParams creates a new ReadAgentPauseCauseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadAgentPauseCauseParams() *ReadAgentPauseCauseParams {
	return &ReadAgentPauseCauseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadAgentPauseCauseParamsWithTimeout creates a new ReadAgentPauseCauseParams object
// with the ability to set a timeout on a request.
func NewReadAgentPauseCauseParamsWithTimeout(timeout time.Duration) *ReadAgentPauseCauseParams {
	return &ReadAgentPauseCauseParams{
		timeout: timeout,
	}
}

// NewReadAgentPauseCauseParamsWithContext creates a new ReadAgentPauseCauseParams object
// with the ability to set a context for a request.
func NewReadAgentPauseCauseParamsWithContext(ctx context.Context) *ReadAgentPauseCauseParams {
	return &ReadAgentPauseCauseParams{
		Context: ctx,
	}
}

// NewReadAgentPauseCauseParamsWithHTTPClient creates a new ReadAgentPauseCauseParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadAgentPauseCauseParamsWithHTTPClient(client *http.Client) *ReadAgentPauseCauseParams {
	return &ReadAgentPauseCauseParams{
		HTTPClient: client,
	}
}

/*
ReadAgentPauseCauseParams contains all the parameters to send to the API endpoint

	for the read agent pause cause operation.

	Typically these are written to a http.Request.
*/
type ReadAgentPauseCauseParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read agent pause cause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadAgentPauseCauseParams) WithDefaults() *ReadAgentPauseCauseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read agent pause cause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadAgentPauseCauseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read agent pause cause params
func (o *ReadAgentPauseCauseParams) WithTimeout(timeout time.Duration) *ReadAgentPauseCauseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read agent pause cause params
func (o *ReadAgentPauseCauseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read agent pause cause params
func (o *ReadAgentPauseCauseParams) WithContext(ctx context.Context) *ReadAgentPauseCauseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read agent pause cause params
func (o *ReadAgentPauseCauseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read agent pause cause params
func (o *ReadAgentPauseCauseParams) WithHTTPClient(client *http.Client) *ReadAgentPauseCauseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read agent pause cause params
func (o *ReadAgentPauseCauseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the read agent pause cause params
func (o *ReadAgentPauseCauseParams) WithID(id int64) *ReadAgentPauseCauseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read agent pause cause params
func (o *ReadAgentPauseCauseParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadAgentPauseCauseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
