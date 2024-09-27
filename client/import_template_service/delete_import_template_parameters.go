// Code generated by go-swagger; DO NOT EDIT.

package import_template_service

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

// NewDeleteImportTemplateParams creates a new DeleteImportTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteImportTemplateParams() *DeleteImportTemplateParams {
	return &DeleteImportTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteImportTemplateParamsWithTimeout creates a new DeleteImportTemplateParams object
// with the ability to set a timeout on a request.
func NewDeleteImportTemplateParamsWithTimeout(timeout time.Duration) *DeleteImportTemplateParams {
	return &DeleteImportTemplateParams{
		timeout: timeout,
	}
}

// NewDeleteImportTemplateParamsWithContext creates a new DeleteImportTemplateParams object
// with the ability to set a context for a request.
func NewDeleteImportTemplateParamsWithContext(ctx context.Context) *DeleteImportTemplateParams {
	return &DeleteImportTemplateParams{
		Context: ctx,
	}
}

// NewDeleteImportTemplateParamsWithHTTPClient creates a new DeleteImportTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteImportTemplateParamsWithHTTPClient(client *http.Client) *DeleteImportTemplateParams {
	return &DeleteImportTemplateParams{
		HTTPClient: client,
	}
}

/*
DeleteImportTemplateParams contains all the parameters to send to the API endpoint

	for the delete import template operation.

	Typically these are written to a http.Request.
*/
type DeleteImportTemplateParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete import template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteImportTemplateParams) WithDefaults() *DeleteImportTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete import template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteImportTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete import template params
func (o *DeleteImportTemplateParams) WithTimeout(timeout time.Duration) *DeleteImportTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete import template params
func (o *DeleteImportTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete import template params
func (o *DeleteImportTemplateParams) WithContext(ctx context.Context) *DeleteImportTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete import template params
func (o *DeleteImportTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete import template params
func (o *DeleteImportTemplateParams) WithHTTPClient(client *http.Client) *DeleteImportTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete import template params
func (o *DeleteImportTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete import template params
func (o *DeleteImportTemplateParams) WithID(id int32) *DeleteImportTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete import template params
func (o *DeleteImportTemplateParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteImportTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
