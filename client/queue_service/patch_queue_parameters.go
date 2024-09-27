// Code generated by go-swagger; DO NOT EDIT.

package queue_service

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

// NewPatchQueueParams creates a new PatchQueueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchQueueParams() *PatchQueueParams {
	return &PatchQueueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchQueueParamsWithTimeout creates a new PatchQueueParams object
// with the ability to set a timeout on a request.
func NewPatchQueueParamsWithTimeout(timeout time.Duration) *PatchQueueParams {
	return &PatchQueueParams{
		timeout: timeout,
	}
}

// NewPatchQueueParamsWithContext creates a new PatchQueueParams object
// with the ability to set a context for a request.
func NewPatchQueueParamsWithContext(ctx context.Context) *PatchQueueParams {
	return &PatchQueueParams{
		Context: ctx,
	}
}

// NewPatchQueueParamsWithHTTPClient creates a new PatchQueueParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchQueueParamsWithHTTPClient(client *http.Client) *PatchQueueParams {
	return &PatchQueueParams{
		HTTPClient: client,
	}
}

/*
PatchQueueParams contains all the parameters to send to the API endpoint

	for the patch queue operation.

	Typically these are written to a http.Request.
*/
type PatchQueueParams struct {

	// Body.
	Body *models.EnginePatchQueueRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchQueueParams) WithDefaults() *PatchQueueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchQueueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch queue params
func (o *PatchQueueParams) WithTimeout(timeout time.Duration) *PatchQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch queue params
func (o *PatchQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch queue params
func (o *PatchQueueParams) WithContext(ctx context.Context) *PatchQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch queue params
func (o *PatchQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch queue params
func (o *PatchQueueParams) WithHTTPClient(client *http.Client) *PatchQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch queue params
func (o *PatchQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch queue params
func (o *PatchQueueParams) WithBody(body *models.EnginePatchQueueRequest) *PatchQueueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch queue params
func (o *PatchQueueParams) SetBody(body *models.EnginePatchQueueRequest) {
	o.Body = body
}

// WithID adds the id to the patch queue params
func (o *PatchQueueParams) WithID(id string) *PatchQueueParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch queue params
func (o *PatchQueueParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
