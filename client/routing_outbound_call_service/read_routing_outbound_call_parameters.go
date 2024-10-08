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
)

// NewReadRoutingOutboundCallParams creates a new ReadRoutingOutboundCallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadRoutingOutboundCallParams() *ReadRoutingOutboundCallParams {
	return &ReadRoutingOutboundCallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadRoutingOutboundCallParamsWithTimeout creates a new ReadRoutingOutboundCallParams object
// with the ability to set a timeout on a request.
func NewReadRoutingOutboundCallParamsWithTimeout(timeout time.Duration) *ReadRoutingOutboundCallParams {
	return &ReadRoutingOutboundCallParams{
		timeout: timeout,
	}
}

// NewReadRoutingOutboundCallParamsWithContext creates a new ReadRoutingOutboundCallParams object
// with the ability to set a context for a request.
func NewReadRoutingOutboundCallParamsWithContext(ctx context.Context) *ReadRoutingOutboundCallParams {
	return &ReadRoutingOutboundCallParams{
		Context: ctx,
	}
}

// NewReadRoutingOutboundCallParamsWithHTTPClient creates a new ReadRoutingOutboundCallParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadRoutingOutboundCallParamsWithHTTPClient(client *http.Client) *ReadRoutingOutboundCallParams {
	return &ReadRoutingOutboundCallParams{
		HTTPClient: client,
	}
}

/*
ReadRoutingOutboundCallParams contains all the parameters to send to the API endpoint

	for the read routing outbound call operation.

	Typically these are written to a http.Request.
*/
type ReadRoutingOutboundCallParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read routing outbound call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadRoutingOutboundCallParams) WithDefaults() *ReadRoutingOutboundCallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read routing outbound call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadRoutingOutboundCallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) WithTimeout(timeout time.Duration) *ReadRoutingOutboundCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) WithContext(ctx context.Context) *ReadRoutingOutboundCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) WithHTTPClient(client *http.Client) *ReadRoutingOutboundCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) WithDomainID(domainID *string) *ReadRoutingOutboundCallParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) WithID(id string) *ReadRoutingOutboundCallParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read routing outbound call params
func (o *ReadRoutingOutboundCallParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadRoutingOutboundCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domain_id
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domain_id", qDomainID); err != nil {
				return err
			}
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
