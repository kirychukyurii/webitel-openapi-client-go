// Code generated by go-swagger; DO NOT EDIT.

package routing_schema_service

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

// NewPatchRoutingSchemaParams creates a new PatchRoutingSchemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchRoutingSchemaParams() *PatchRoutingSchemaParams {
	return &PatchRoutingSchemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRoutingSchemaParamsWithTimeout creates a new PatchRoutingSchemaParams object
// with the ability to set a timeout on a request.
func NewPatchRoutingSchemaParamsWithTimeout(timeout time.Duration) *PatchRoutingSchemaParams {
	return &PatchRoutingSchemaParams{
		timeout: timeout,
	}
}

// NewPatchRoutingSchemaParamsWithContext creates a new PatchRoutingSchemaParams object
// with the ability to set a context for a request.
func NewPatchRoutingSchemaParamsWithContext(ctx context.Context) *PatchRoutingSchemaParams {
	return &PatchRoutingSchemaParams{
		Context: ctx,
	}
}

// NewPatchRoutingSchemaParamsWithHTTPClient creates a new PatchRoutingSchemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchRoutingSchemaParamsWithHTTPClient(client *http.Client) *PatchRoutingSchemaParams {
	return &PatchRoutingSchemaParams{
		HTTPClient: client,
	}
}

/*
PatchRoutingSchemaParams contains all the parameters to send to the API endpoint

	for the patch routing schema operation.

	Typically these are written to a http.Request.
*/
type PatchRoutingSchemaParams struct {

	// Body.
	Body *models.EnginePatchRoutingSchemaRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch routing schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRoutingSchemaParams) WithDefaults() *PatchRoutingSchemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch routing schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRoutingSchemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch routing schema params
func (o *PatchRoutingSchemaParams) WithTimeout(timeout time.Duration) *PatchRoutingSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch routing schema params
func (o *PatchRoutingSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch routing schema params
func (o *PatchRoutingSchemaParams) WithContext(ctx context.Context) *PatchRoutingSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch routing schema params
func (o *PatchRoutingSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch routing schema params
func (o *PatchRoutingSchemaParams) WithHTTPClient(client *http.Client) *PatchRoutingSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch routing schema params
func (o *PatchRoutingSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch routing schema params
func (o *PatchRoutingSchemaParams) WithBody(body *models.EnginePatchRoutingSchemaRequest) *PatchRoutingSchemaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch routing schema params
func (o *PatchRoutingSchemaParams) SetBody(body *models.EnginePatchRoutingSchemaRequest) {
	o.Body = body
}

// WithID adds the id to the patch routing schema params
func (o *PatchRoutingSchemaParams) WithID(id string) *PatchRoutingSchemaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch routing schema params
func (o *PatchRoutingSchemaParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRoutingSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
