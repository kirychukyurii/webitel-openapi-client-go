// Code generated by go-swagger; DO NOT EDIT.

package routing_chat_plan_service

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

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// NewUpdateChatPlanParams creates a new UpdateChatPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateChatPlanParams() *UpdateChatPlanParams {
	return &UpdateChatPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateChatPlanParamsWithTimeout creates a new UpdateChatPlanParams object
// with the ability to set a timeout on a request.
func NewUpdateChatPlanParamsWithTimeout(timeout time.Duration) *UpdateChatPlanParams {
	return &UpdateChatPlanParams{
		timeout: timeout,
	}
}

// NewUpdateChatPlanParamsWithContext creates a new UpdateChatPlanParams object
// with the ability to set a context for a request.
func NewUpdateChatPlanParamsWithContext(ctx context.Context) *UpdateChatPlanParams {
	return &UpdateChatPlanParams{
		Context: ctx,
	}
}

// NewUpdateChatPlanParamsWithHTTPClient creates a new UpdateChatPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateChatPlanParamsWithHTTPClient(client *http.Client) *UpdateChatPlanParams {
	return &UpdateChatPlanParams{
		HTTPClient: client,
	}
}

/*
UpdateChatPlanParams contains all the parameters to send to the API endpoint

	for the update chat plan operation.

	Typically these are written to a http.Request.
*/
type UpdateChatPlanParams struct {

	// Body.
	Body *models.EngineUpdateChatPlanRequest

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update chat plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateChatPlanParams) WithDefaults() *UpdateChatPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update chat plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateChatPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update chat plan params
func (o *UpdateChatPlanParams) WithTimeout(timeout time.Duration) *UpdateChatPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update chat plan params
func (o *UpdateChatPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update chat plan params
func (o *UpdateChatPlanParams) WithContext(ctx context.Context) *UpdateChatPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update chat plan params
func (o *UpdateChatPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update chat plan params
func (o *UpdateChatPlanParams) WithHTTPClient(client *http.Client) *UpdateChatPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update chat plan params
func (o *UpdateChatPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update chat plan params
func (o *UpdateChatPlanParams) WithBody(body *models.EngineUpdateChatPlanRequest) *UpdateChatPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update chat plan params
func (o *UpdateChatPlanParams) SetBody(body *models.EngineUpdateChatPlanRequest) {
	o.Body = body
}

// WithID adds the id to the update chat plan params
func (o *UpdateChatPlanParams) WithID(id int32) *UpdateChatPlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update chat plan params
func (o *UpdateChatPlanParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateChatPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
