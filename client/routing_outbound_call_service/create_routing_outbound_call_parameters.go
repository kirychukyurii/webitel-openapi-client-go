// Code generated by go-swagger; DO NOT EDIT.

package routing_outbound_call_service

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

// NewCreateRoutingOutboundCallParams creates a new CreateRoutingOutboundCallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRoutingOutboundCallParams() *CreateRoutingOutboundCallParams {
	return &CreateRoutingOutboundCallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoutingOutboundCallParamsWithTimeout creates a new CreateRoutingOutboundCallParams object
// with the ability to set a timeout on a request.
func NewCreateRoutingOutboundCallParamsWithTimeout(timeout time.Duration) *CreateRoutingOutboundCallParams {
	return &CreateRoutingOutboundCallParams{
		timeout: timeout,
	}
}

// NewCreateRoutingOutboundCallParamsWithContext creates a new CreateRoutingOutboundCallParams object
// with the ability to set a context for a request.
func NewCreateRoutingOutboundCallParamsWithContext(ctx context.Context) *CreateRoutingOutboundCallParams {
	return &CreateRoutingOutboundCallParams{
		Context: ctx,
	}
}

// NewCreateRoutingOutboundCallParamsWithHTTPClient creates a new CreateRoutingOutboundCallParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRoutingOutboundCallParamsWithHTTPClient(client *http.Client) *CreateRoutingOutboundCallParams {
	return &CreateRoutingOutboundCallParams{
		HTTPClient: client,
	}
}

/*
CreateRoutingOutboundCallParams contains all the parameters to send to the API endpoint

	for the create routing outbound call operation.

	Typically these are written to a http.Request.
*/
type CreateRoutingOutboundCallParams struct {

	// Body.
	Body *models.EngineCreateRoutingOutboundCallRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create routing outbound call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoutingOutboundCallParams) WithDefaults() *CreateRoutingOutboundCallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create routing outbound call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoutingOutboundCallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create routing outbound call params
func (o *CreateRoutingOutboundCallParams) WithTimeout(timeout time.Duration) *CreateRoutingOutboundCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create routing outbound call params
func (o *CreateRoutingOutboundCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create routing outbound call params
func (o *CreateRoutingOutboundCallParams) WithContext(ctx context.Context) *CreateRoutingOutboundCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create routing outbound call params
func (o *CreateRoutingOutboundCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create routing outbound call params
func (o *CreateRoutingOutboundCallParams) WithHTTPClient(client *http.Client) *CreateRoutingOutboundCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create routing outbound call params
func (o *CreateRoutingOutboundCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create routing outbound call params
func (o *CreateRoutingOutboundCallParams) WithBody(body *models.EngineCreateRoutingOutboundCallRequest) *CreateRoutingOutboundCallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create routing outbound call params
func (o *CreateRoutingOutboundCallParams) SetBody(body *models.EngineCreateRoutingOutboundCallRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoutingOutboundCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
