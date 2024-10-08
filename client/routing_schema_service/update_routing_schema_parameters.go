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

// NewUpdateRoutingSchemaParams creates a new UpdateRoutingSchemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRoutingSchemaParams() *UpdateRoutingSchemaParams {
	return &UpdateRoutingSchemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoutingSchemaParamsWithTimeout creates a new UpdateRoutingSchemaParams object
// with the ability to set a timeout on a request.
func NewUpdateRoutingSchemaParamsWithTimeout(timeout time.Duration) *UpdateRoutingSchemaParams {
	return &UpdateRoutingSchemaParams{
		timeout: timeout,
	}
}

// NewUpdateRoutingSchemaParamsWithContext creates a new UpdateRoutingSchemaParams object
// with the ability to set a context for a request.
func NewUpdateRoutingSchemaParamsWithContext(ctx context.Context) *UpdateRoutingSchemaParams {
	return &UpdateRoutingSchemaParams{
		Context: ctx,
	}
}

// NewUpdateRoutingSchemaParamsWithHTTPClient creates a new UpdateRoutingSchemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRoutingSchemaParamsWithHTTPClient(client *http.Client) *UpdateRoutingSchemaParams {
	return &UpdateRoutingSchemaParams{
		HTTPClient: client,
	}
}

/*
UpdateRoutingSchemaParams contains all the parameters to send to the API endpoint

	for the update routing schema operation.

	Typically these are written to a http.Request.
*/
type UpdateRoutingSchemaParams struct {

	// Body.
	Body *models.EngineUpdateRoutingSchemaRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update routing schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoutingSchemaParams) WithDefaults() *UpdateRoutingSchemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update routing schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoutingSchemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update routing schema params
func (o *UpdateRoutingSchemaParams) WithTimeout(timeout time.Duration) *UpdateRoutingSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update routing schema params
func (o *UpdateRoutingSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update routing schema params
func (o *UpdateRoutingSchemaParams) WithContext(ctx context.Context) *UpdateRoutingSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update routing schema params
func (o *UpdateRoutingSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update routing schema params
func (o *UpdateRoutingSchemaParams) WithHTTPClient(client *http.Client) *UpdateRoutingSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update routing schema params
func (o *UpdateRoutingSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update routing schema params
func (o *UpdateRoutingSchemaParams) WithBody(body *models.EngineUpdateRoutingSchemaRequest) *UpdateRoutingSchemaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update routing schema params
func (o *UpdateRoutingSchemaParams) SetBody(body *models.EngineUpdateRoutingSchemaRequest) {
	o.Body = body
}

// WithID adds the id to the update routing schema params
func (o *UpdateRoutingSchemaParams) WithID(id string) *UpdateRoutingSchemaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update routing schema params
func (o *UpdateRoutingSchemaParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoutingSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
