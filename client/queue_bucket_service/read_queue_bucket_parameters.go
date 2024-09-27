// Code generated by go-swagger; DO NOT EDIT.

package queue_bucket_service

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

// NewReadQueueBucketParams creates a new ReadQueueBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadQueueBucketParams() *ReadQueueBucketParams {
	return &ReadQueueBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadQueueBucketParamsWithTimeout creates a new ReadQueueBucketParams object
// with the ability to set a timeout on a request.
func NewReadQueueBucketParamsWithTimeout(timeout time.Duration) *ReadQueueBucketParams {
	return &ReadQueueBucketParams{
		timeout: timeout,
	}
}

// NewReadQueueBucketParamsWithContext creates a new ReadQueueBucketParams object
// with the ability to set a context for a request.
func NewReadQueueBucketParamsWithContext(ctx context.Context) *ReadQueueBucketParams {
	return &ReadQueueBucketParams{
		Context: ctx,
	}
}

// NewReadQueueBucketParamsWithHTTPClient creates a new ReadQueueBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadQueueBucketParamsWithHTTPClient(client *http.Client) *ReadQueueBucketParams {
	return &ReadQueueBucketParams{
		HTTPClient: client,
	}
}

/*
ReadQueueBucketParams contains all the parameters to send to the API endpoint

	for the read queue bucket operation.

	Typically these are written to a http.Request.
*/
type ReadQueueBucketParams struct {

	// ID.
	//
	// Format: int64
	ID string

	// QueueID.
	//
	// Format: int64
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read queue bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadQueueBucketParams) WithDefaults() *ReadQueueBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read queue bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadQueueBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read queue bucket params
func (o *ReadQueueBucketParams) WithTimeout(timeout time.Duration) *ReadQueueBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read queue bucket params
func (o *ReadQueueBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read queue bucket params
func (o *ReadQueueBucketParams) WithContext(ctx context.Context) *ReadQueueBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read queue bucket params
func (o *ReadQueueBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read queue bucket params
func (o *ReadQueueBucketParams) WithHTTPClient(client *http.Client) *ReadQueueBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read queue bucket params
func (o *ReadQueueBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the read queue bucket params
func (o *ReadQueueBucketParams) WithID(id string) *ReadQueueBucketParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read queue bucket params
func (o *ReadQueueBucketParams) SetID(id string) {
	o.ID = id
}

// WithQueueID adds the queueID to the read queue bucket params
func (o *ReadQueueBucketParams) WithQueueID(queueID string) *ReadQueueBucketParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the read queue bucket params
func (o *ReadQueueBucketParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *ReadQueueBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param queue_id
	if err := r.SetPathParam("queue_id", o.QueueID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
