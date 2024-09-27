// Code generated by go-swagger; DO NOT EDIT.

package queue_resources_service

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

// NewDeleteQueueResourceGroupParams creates a new DeleteQueueResourceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteQueueResourceGroupParams() *DeleteQueueResourceGroupParams {
	return &DeleteQueueResourceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteQueueResourceGroupParamsWithTimeout creates a new DeleteQueueResourceGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteQueueResourceGroupParamsWithTimeout(timeout time.Duration) *DeleteQueueResourceGroupParams {
	return &DeleteQueueResourceGroupParams{
		timeout: timeout,
	}
}

// NewDeleteQueueResourceGroupParamsWithContext creates a new DeleteQueueResourceGroupParams object
// with the ability to set a context for a request.
func NewDeleteQueueResourceGroupParamsWithContext(ctx context.Context) *DeleteQueueResourceGroupParams {
	return &DeleteQueueResourceGroupParams{
		Context: ctx,
	}
}

// NewDeleteQueueResourceGroupParamsWithHTTPClient creates a new DeleteQueueResourceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteQueueResourceGroupParamsWithHTTPClient(client *http.Client) *DeleteQueueResourceGroupParams {
	return &DeleteQueueResourceGroupParams{
		HTTPClient: client,
	}
}

/*
DeleteQueueResourceGroupParams contains all the parameters to send to the API endpoint

	for the delete queue resource group operation.

	Typically these are written to a http.Request.
*/
type DeleteQueueResourceGroupParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

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

// WithDefaults hydrates default values in the delete queue resource group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQueueResourceGroupParams) WithDefaults() *DeleteQueueResourceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete queue resource group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQueueResourceGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) WithTimeout(timeout time.Duration) *DeleteQueueResourceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) WithContext(ctx context.Context) *DeleteQueueResourceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) WithHTTPClient(client *http.Client) *DeleteQueueResourceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) WithDomainID(domainID *string) *DeleteQueueResourceGroupParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) WithID(id string) *DeleteQueueResourceGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) SetID(id string) {
	o.ID = id
}

// WithQueueID adds the queueID to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) WithQueueID(queueID string) *DeleteQueueResourceGroupParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the delete queue resource group params
func (o *DeleteQueueResourceGroupParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteQueueResourceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param queue_id
	if err := r.SetPathParam("queue_id", o.QueueID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
