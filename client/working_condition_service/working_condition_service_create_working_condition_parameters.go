// Code generated by go-swagger; DO NOT EDIT.

package working_condition_service

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

// NewWorkingConditionServiceCreateWorkingConditionParams creates a new WorkingConditionServiceCreateWorkingConditionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkingConditionServiceCreateWorkingConditionParams() *WorkingConditionServiceCreateWorkingConditionParams {
	return &WorkingConditionServiceCreateWorkingConditionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkingConditionServiceCreateWorkingConditionParamsWithTimeout creates a new WorkingConditionServiceCreateWorkingConditionParams object
// with the ability to set a timeout on a request.
func NewWorkingConditionServiceCreateWorkingConditionParamsWithTimeout(timeout time.Duration) *WorkingConditionServiceCreateWorkingConditionParams {
	return &WorkingConditionServiceCreateWorkingConditionParams{
		timeout: timeout,
	}
}

// NewWorkingConditionServiceCreateWorkingConditionParamsWithContext creates a new WorkingConditionServiceCreateWorkingConditionParams object
// with the ability to set a context for a request.
func NewWorkingConditionServiceCreateWorkingConditionParamsWithContext(ctx context.Context) *WorkingConditionServiceCreateWorkingConditionParams {
	return &WorkingConditionServiceCreateWorkingConditionParams{
		Context: ctx,
	}
}

// NewWorkingConditionServiceCreateWorkingConditionParamsWithHTTPClient creates a new WorkingConditionServiceCreateWorkingConditionParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkingConditionServiceCreateWorkingConditionParamsWithHTTPClient(client *http.Client) *WorkingConditionServiceCreateWorkingConditionParams {
	return &WorkingConditionServiceCreateWorkingConditionParams{
		HTTPClient: client,
	}
}

/*
WorkingConditionServiceCreateWorkingConditionParams contains all the parameters to send to the API endpoint

	for the working condition service create working condition operation.

	Typically these are written to a http.Request.
*/
type WorkingConditionServiceCreateWorkingConditionParams struct {

	// Body.
	Body *models.WfmCreateWorkingConditionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the working condition service create working condition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkingConditionServiceCreateWorkingConditionParams) WithDefaults() *WorkingConditionServiceCreateWorkingConditionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the working condition service create working condition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkingConditionServiceCreateWorkingConditionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the working condition service create working condition params
func (o *WorkingConditionServiceCreateWorkingConditionParams) WithTimeout(timeout time.Duration) *WorkingConditionServiceCreateWorkingConditionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the working condition service create working condition params
func (o *WorkingConditionServiceCreateWorkingConditionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the working condition service create working condition params
func (o *WorkingConditionServiceCreateWorkingConditionParams) WithContext(ctx context.Context) *WorkingConditionServiceCreateWorkingConditionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the working condition service create working condition params
func (o *WorkingConditionServiceCreateWorkingConditionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the working condition service create working condition params
func (o *WorkingConditionServiceCreateWorkingConditionParams) WithHTTPClient(client *http.Client) *WorkingConditionServiceCreateWorkingConditionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the working condition service create working condition params
func (o *WorkingConditionServiceCreateWorkingConditionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the working condition service create working condition params
func (o *WorkingConditionServiceCreateWorkingConditionParams) WithBody(body *models.WfmCreateWorkingConditionRequest) *WorkingConditionServiceCreateWorkingConditionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the working condition service create working condition params
func (o *WorkingConditionServiceCreateWorkingConditionParams) SetBody(body *models.WfmCreateWorkingConditionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WorkingConditionServiceCreateWorkingConditionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
