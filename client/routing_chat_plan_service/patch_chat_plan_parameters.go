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

// NewPatchChatPlanParams creates a new PatchChatPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchChatPlanParams() *PatchChatPlanParams {
	return &PatchChatPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchChatPlanParamsWithTimeout creates a new PatchChatPlanParams object
// with the ability to set a timeout on a request.
func NewPatchChatPlanParamsWithTimeout(timeout time.Duration) *PatchChatPlanParams {
	return &PatchChatPlanParams{
		timeout: timeout,
	}
}

// NewPatchChatPlanParamsWithContext creates a new PatchChatPlanParams object
// with the ability to set a context for a request.
func NewPatchChatPlanParamsWithContext(ctx context.Context) *PatchChatPlanParams {
	return &PatchChatPlanParams{
		Context: ctx,
	}
}

// NewPatchChatPlanParamsWithHTTPClient creates a new PatchChatPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchChatPlanParamsWithHTTPClient(client *http.Client) *PatchChatPlanParams {
	return &PatchChatPlanParams{
		HTTPClient: client,
	}
}

/*
PatchChatPlanParams contains all the parameters to send to the API endpoint

	for the patch chat plan operation.

	Typically these are written to a http.Request.
*/
type PatchChatPlanParams struct {

	// Body.
	Body *models.EnginePatchChatPlanRequest

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch chat plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchChatPlanParams) WithDefaults() *PatchChatPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch chat plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchChatPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch chat plan params
func (o *PatchChatPlanParams) WithTimeout(timeout time.Duration) *PatchChatPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch chat plan params
func (o *PatchChatPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch chat plan params
func (o *PatchChatPlanParams) WithContext(ctx context.Context) *PatchChatPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch chat plan params
func (o *PatchChatPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch chat plan params
func (o *PatchChatPlanParams) WithHTTPClient(client *http.Client) *PatchChatPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch chat plan params
func (o *PatchChatPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch chat plan params
func (o *PatchChatPlanParams) WithBody(body *models.EnginePatchChatPlanRequest) *PatchChatPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch chat plan params
func (o *PatchChatPlanParams) SetBody(body *models.EnginePatchChatPlanRequest) {
	o.Body = body
}

// WithID adds the id to the patch chat plan params
func (o *PatchChatPlanParams) WithID(id int32) *PatchChatPlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch chat plan params
func (o *PatchChatPlanParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchChatPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
