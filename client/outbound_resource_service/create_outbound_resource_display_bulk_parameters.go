// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

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

// NewCreateOutboundResourceDisplayBulkParams creates a new CreateOutboundResourceDisplayBulkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOutboundResourceDisplayBulkParams() *CreateOutboundResourceDisplayBulkParams {
	return &CreateOutboundResourceDisplayBulkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOutboundResourceDisplayBulkParamsWithTimeout creates a new CreateOutboundResourceDisplayBulkParams object
// with the ability to set a timeout on a request.
func NewCreateOutboundResourceDisplayBulkParamsWithTimeout(timeout time.Duration) *CreateOutboundResourceDisplayBulkParams {
	return &CreateOutboundResourceDisplayBulkParams{
		timeout: timeout,
	}
}

// NewCreateOutboundResourceDisplayBulkParamsWithContext creates a new CreateOutboundResourceDisplayBulkParams object
// with the ability to set a context for a request.
func NewCreateOutboundResourceDisplayBulkParamsWithContext(ctx context.Context) *CreateOutboundResourceDisplayBulkParams {
	return &CreateOutboundResourceDisplayBulkParams{
		Context: ctx,
	}
}

// NewCreateOutboundResourceDisplayBulkParamsWithHTTPClient creates a new CreateOutboundResourceDisplayBulkParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOutboundResourceDisplayBulkParamsWithHTTPClient(client *http.Client) *CreateOutboundResourceDisplayBulkParams {
	return &CreateOutboundResourceDisplayBulkParams{
		HTTPClient: client,
	}
}

/*
CreateOutboundResourceDisplayBulkParams contains all the parameters to send to the API endpoint

	for the create outbound resource display bulk operation.

	Typically these are written to a http.Request.
*/
type CreateOutboundResourceDisplayBulkParams struct {

	// Body.
	Body *models.EngineCreateOutboundResourceDisplayBulkRequest

	// ResourceID.
	//
	// Format: int64
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create outbound resource display bulk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOutboundResourceDisplayBulkParams) WithDefaults() *CreateOutboundResourceDisplayBulkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create outbound resource display bulk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOutboundResourceDisplayBulkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) WithTimeout(timeout time.Duration) *CreateOutboundResourceDisplayBulkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) WithContext(ctx context.Context) *CreateOutboundResourceDisplayBulkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) WithHTTPClient(client *http.Client) *CreateOutboundResourceDisplayBulkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) WithBody(body *models.EngineCreateOutboundResourceDisplayBulkRequest) *CreateOutboundResourceDisplayBulkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) SetBody(body *models.EngineCreateOutboundResourceDisplayBulkRequest) {
	o.Body = body
}

// WithResourceID adds the resourceID to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) WithResourceID(resourceID string) *CreateOutboundResourceDisplayBulkParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the create outbound resource display bulk params
func (o *CreateOutboundResourceDisplayBulkParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOutboundResourceDisplayBulkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param resource_id
	if err := r.SetPathParam("resource_id", o.ResourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
