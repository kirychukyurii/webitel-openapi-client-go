// Code generated by go-swagger; DO NOT EDIT.

package member_service

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

// NewDeleteMembersParams creates a new DeleteMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMembersParams() *DeleteMembersParams {
	return &DeleteMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMembersParamsWithTimeout creates a new DeleteMembersParams object
// with the ability to set a timeout on a request.
func NewDeleteMembersParamsWithTimeout(timeout time.Duration) *DeleteMembersParams {
	return &DeleteMembersParams{
		timeout: timeout,
	}
}

// NewDeleteMembersParamsWithContext creates a new DeleteMembersParams object
// with the ability to set a context for a request.
func NewDeleteMembersParamsWithContext(ctx context.Context) *DeleteMembersParams {
	return &DeleteMembersParams{
		Context: ctx,
	}
}

// NewDeleteMembersParamsWithHTTPClient creates a new DeleteMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMembersParamsWithHTTPClient(client *http.Client) *DeleteMembersParams {
	return &DeleteMembersParams{
		HTTPClient: client,
	}
}

/*
DeleteMembersParams contains all the parameters to send to the API endpoint

	for the delete members operation.

	Typically these are written to a http.Request.
*/
type DeleteMembersParams struct {

	// Body.
	Body *models.EngineDeleteMembersRequest

	// QueueID.
	//
	// Format: int64
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMembersParams) WithDefaults() *DeleteMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete members params
func (o *DeleteMembersParams) WithTimeout(timeout time.Duration) *DeleteMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete members params
func (o *DeleteMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete members params
func (o *DeleteMembersParams) WithContext(ctx context.Context) *DeleteMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete members params
func (o *DeleteMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete members params
func (o *DeleteMembersParams) WithHTTPClient(client *http.Client) *DeleteMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete members params
func (o *DeleteMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete members params
func (o *DeleteMembersParams) WithBody(body *models.EngineDeleteMembersRequest) *DeleteMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete members params
func (o *DeleteMembersParams) SetBody(body *models.EngineDeleteMembersRequest) {
	o.Body = body
}

// WithQueueID adds the queueID to the delete members params
func (o *DeleteMembersParams) WithQueueID(queueID string) *DeleteMembersParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the delete members params
func (o *DeleteMembersParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
