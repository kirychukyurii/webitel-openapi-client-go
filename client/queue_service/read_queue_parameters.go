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
)

// NewReadQueueParams creates a new ReadQueueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadQueueParams() *ReadQueueParams {
	return &ReadQueueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadQueueParamsWithTimeout creates a new ReadQueueParams object
// with the ability to set a timeout on a request.
func NewReadQueueParamsWithTimeout(timeout time.Duration) *ReadQueueParams {
	return &ReadQueueParams{
		timeout: timeout,
	}
}

// NewReadQueueParamsWithContext creates a new ReadQueueParams object
// with the ability to set a context for a request.
func NewReadQueueParamsWithContext(ctx context.Context) *ReadQueueParams {
	return &ReadQueueParams{
		Context: ctx,
	}
}

// NewReadQueueParamsWithHTTPClient creates a new ReadQueueParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadQueueParamsWithHTTPClient(client *http.Client) *ReadQueueParams {
	return &ReadQueueParams{
		HTTPClient: client,
	}
}

/*
ReadQueueParams contains all the parameters to send to the API endpoint

	for the read queue operation.

	Typically these are written to a http.Request.
*/
type ReadQueueParams struct {

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

// WithDefaults hydrates default values in the read queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadQueueParams) WithDefaults() *ReadQueueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadQueueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read queue params
func (o *ReadQueueParams) WithTimeout(timeout time.Duration) *ReadQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read queue params
func (o *ReadQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read queue params
func (o *ReadQueueParams) WithContext(ctx context.Context) *ReadQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read queue params
func (o *ReadQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read queue params
func (o *ReadQueueParams) WithHTTPClient(client *http.Client) *ReadQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read queue params
func (o *ReadQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the read queue params
func (o *ReadQueueParams) WithDomainID(domainID *string) *ReadQueueParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the read queue params
func (o *ReadQueueParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the read queue params
func (o *ReadQueueParams) WithID(id string) *ReadQueueParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read queue params
func (o *ReadQueueParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
