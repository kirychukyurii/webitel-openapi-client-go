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

// NewUpdateMemberParams creates a new UpdateMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMemberParams() *UpdateMemberParams {
	return &UpdateMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMemberParamsWithTimeout creates a new UpdateMemberParams object
// with the ability to set a timeout on a request.
func NewUpdateMemberParamsWithTimeout(timeout time.Duration) *UpdateMemberParams {
	return &UpdateMemberParams{
		timeout: timeout,
	}
}

// NewUpdateMemberParamsWithContext creates a new UpdateMemberParams object
// with the ability to set a context for a request.
func NewUpdateMemberParamsWithContext(ctx context.Context) *UpdateMemberParams {
	return &UpdateMemberParams{
		Context: ctx,
	}
}

// NewUpdateMemberParamsWithHTTPClient creates a new UpdateMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMemberParamsWithHTTPClient(client *http.Client) *UpdateMemberParams {
	return &UpdateMemberParams{
		HTTPClient: client,
	}
}

/*
UpdateMemberParams contains all the parameters to send to the API endpoint

	for the update member operation.

	Typically these are written to a http.Request.
*/
type UpdateMemberParams struct {

	// Body.
	Body *models.EngineUpdateMemberRequest

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

// WithDefaults hydrates default values in the update member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMemberParams) WithDefaults() *UpdateMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update member params
func (o *UpdateMemberParams) WithTimeout(timeout time.Duration) *UpdateMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update member params
func (o *UpdateMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update member params
func (o *UpdateMemberParams) WithContext(ctx context.Context) *UpdateMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update member params
func (o *UpdateMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update member params
func (o *UpdateMemberParams) WithHTTPClient(client *http.Client) *UpdateMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update member params
func (o *UpdateMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update member params
func (o *UpdateMemberParams) WithBody(body *models.EngineUpdateMemberRequest) *UpdateMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update member params
func (o *UpdateMemberParams) SetBody(body *models.EngineUpdateMemberRequest) {
	o.Body = body
}

// WithID adds the id to the update member params
func (o *UpdateMemberParams) WithID(id string) *UpdateMemberParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update member params
func (o *UpdateMemberParams) SetID(id string) {
	o.ID = id
}

// WithQueueID adds the queueID to the update member params
func (o *UpdateMemberParams) WithQueueID(queueID string) *UpdateMemberParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the update member params
func (o *UpdateMemberParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param queue_id
	if err := r.SetPathParam("queue_id", o.QueueID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
