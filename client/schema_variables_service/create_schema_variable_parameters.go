// Code generated by go-swagger; DO NOT EDIT.

package schema_variables_service

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

// NewCreateSchemaVariableParams creates a new CreateSchemaVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSchemaVariableParams() *CreateSchemaVariableParams {
	return &CreateSchemaVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSchemaVariableParamsWithTimeout creates a new CreateSchemaVariableParams object
// with the ability to set a timeout on a request.
func NewCreateSchemaVariableParamsWithTimeout(timeout time.Duration) *CreateSchemaVariableParams {
	return &CreateSchemaVariableParams{
		timeout: timeout,
	}
}

// NewCreateSchemaVariableParamsWithContext creates a new CreateSchemaVariableParams object
// with the ability to set a context for a request.
func NewCreateSchemaVariableParamsWithContext(ctx context.Context) *CreateSchemaVariableParams {
	return &CreateSchemaVariableParams{
		Context: ctx,
	}
}

// NewCreateSchemaVariableParamsWithHTTPClient creates a new CreateSchemaVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSchemaVariableParamsWithHTTPClient(client *http.Client) *CreateSchemaVariableParams {
	return &CreateSchemaVariableParams{
		HTTPClient: client,
	}
}

/*
CreateSchemaVariableParams contains all the parameters to send to the API endpoint

	for the create schema variable operation.

	Typically these are written to a http.Request.
*/
type CreateSchemaVariableParams struct {

	// Body.
	Body *models.EngineCreateSchemaVariableRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create schema variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSchemaVariableParams) WithDefaults() *CreateSchemaVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create schema variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSchemaVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create schema variable params
func (o *CreateSchemaVariableParams) WithTimeout(timeout time.Duration) *CreateSchemaVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create schema variable params
func (o *CreateSchemaVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create schema variable params
func (o *CreateSchemaVariableParams) WithContext(ctx context.Context) *CreateSchemaVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create schema variable params
func (o *CreateSchemaVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create schema variable params
func (o *CreateSchemaVariableParams) WithHTTPClient(client *http.Client) *CreateSchemaVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create schema variable params
func (o *CreateSchemaVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create schema variable params
func (o *CreateSchemaVariableParams) WithBody(body *models.EngineCreateSchemaVariableRequest) *CreateSchemaVariableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create schema variable params
func (o *CreateSchemaVariableParams) SetBody(body *models.EngineCreateSchemaVariableRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSchemaVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
