// Code generated by go-swagger; DO NOT EDIT.

package cognitive_profile_service

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

// NewPatchCognitiveProfileParams creates a new PatchCognitiveProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchCognitiveProfileParams() *PatchCognitiveProfileParams {
	return &PatchCognitiveProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCognitiveProfileParamsWithTimeout creates a new PatchCognitiveProfileParams object
// with the ability to set a timeout on a request.
func NewPatchCognitiveProfileParamsWithTimeout(timeout time.Duration) *PatchCognitiveProfileParams {
	return &PatchCognitiveProfileParams{
		timeout: timeout,
	}
}

// NewPatchCognitiveProfileParamsWithContext creates a new PatchCognitiveProfileParams object
// with the ability to set a context for a request.
func NewPatchCognitiveProfileParamsWithContext(ctx context.Context) *PatchCognitiveProfileParams {
	return &PatchCognitiveProfileParams{
		Context: ctx,
	}
}

// NewPatchCognitiveProfileParamsWithHTTPClient creates a new PatchCognitiveProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchCognitiveProfileParamsWithHTTPClient(client *http.Client) *PatchCognitiveProfileParams {
	return &PatchCognitiveProfileParams{
		HTTPClient: client,
	}
}

/*
PatchCognitiveProfileParams contains all the parameters to send to the API endpoint

	for the patch cognitive profile operation.

	Typically these are written to a http.Request.
*/
type PatchCognitiveProfileParams struct {

	// Body.
	Body *models.StoragePatchCognitiveProfileRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch cognitive profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCognitiveProfileParams) WithDefaults() *PatchCognitiveProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch cognitive profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCognitiveProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) WithTimeout(timeout time.Duration) *PatchCognitiveProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) WithContext(ctx context.Context) *PatchCognitiveProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) WithHTTPClient(client *http.Client) *PatchCognitiveProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) WithBody(body *models.StoragePatchCognitiveProfileRequest) *PatchCognitiveProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) SetBody(body *models.StoragePatchCognitiveProfileRequest) {
	o.Body = body
}

// WithID adds the id to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) WithID(id string) *PatchCognitiveProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch cognitive profile params
func (o *PatchCognitiveProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCognitiveProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
