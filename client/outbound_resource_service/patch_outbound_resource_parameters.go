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

// NewPatchOutboundResourceParams creates a new PatchOutboundResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchOutboundResourceParams() *PatchOutboundResourceParams {
	return &PatchOutboundResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchOutboundResourceParamsWithTimeout creates a new PatchOutboundResourceParams object
// with the ability to set a timeout on a request.
func NewPatchOutboundResourceParamsWithTimeout(timeout time.Duration) *PatchOutboundResourceParams {
	return &PatchOutboundResourceParams{
		timeout: timeout,
	}
}

// NewPatchOutboundResourceParamsWithContext creates a new PatchOutboundResourceParams object
// with the ability to set a context for a request.
func NewPatchOutboundResourceParamsWithContext(ctx context.Context) *PatchOutboundResourceParams {
	return &PatchOutboundResourceParams{
		Context: ctx,
	}
}

// NewPatchOutboundResourceParamsWithHTTPClient creates a new PatchOutboundResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchOutboundResourceParamsWithHTTPClient(client *http.Client) *PatchOutboundResourceParams {
	return &PatchOutboundResourceParams{
		HTTPClient: client,
	}
}

/*
PatchOutboundResourceParams contains all the parameters to send to the API endpoint

	for the patch outbound resource operation.

	Typically these are written to a http.Request.
*/
type PatchOutboundResourceParams struct {

	// Body.
	Body *models.EnginePatchOutboundResourceRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch outbound resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchOutboundResourceParams) WithDefaults() *PatchOutboundResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch outbound resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchOutboundResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch outbound resource params
func (o *PatchOutboundResourceParams) WithTimeout(timeout time.Duration) *PatchOutboundResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch outbound resource params
func (o *PatchOutboundResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch outbound resource params
func (o *PatchOutboundResourceParams) WithContext(ctx context.Context) *PatchOutboundResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch outbound resource params
func (o *PatchOutboundResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch outbound resource params
func (o *PatchOutboundResourceParams) WithHTTPClient(client *http.Client) *PatchOutboundResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch outbound resource params
func (o *PatchOutboundResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch outbound resource params
func (o *PatchOutboundResourceParams) WithBody(body *models.EnginePatchOutboundResourceRequest) *PatchOutboundResourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch outbound resource params
func (o *PatchOutboundResourceParams) SetBody(body *models.EnginePatchOutboundResourceRequest) {
	o.Body = body
}

// WithID adds the id to the patch outbound resource params
func (o *PatchOutboundResourceParams) WithID(id string) *PatchOutboundResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch outbound resource params
func (o *PatchOutboundResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchOutboundResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
