// Code generated by go-swagger; DO NOT EDIT.

package classes

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

// NewClassesUpdateClassParams creates a new ClassesUpdateClassParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClassesUpdateClassParams() *ClassesUpdateClassParams {
	return &ClassesUpdateClassParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClassesUpdateClassParamsWithTimeout creates a new ClassesUpdateClassParams object
// with the ability to set a timeout on a request.
func NewClassesUpdateClassParamsWithTimeout(timeout time.Duration) *ClassesUpdateClassParams {
	return &ClassesUpdateClassParams{
		timeout: timeout,
	}
}

// NewClassesUpdateClassParamsWithContext creates a new ClassesUpdateClassParams object
// with the ability to set a context for a request.
func NewClassesUpdateClassParamsWithContext(ctx context.Context) *ClassesUpdateClassParams {
	return &ClassesUpdateClassParams{
		Context: ctx,
	}
}

// NewClassesUpdateClassParamsWithHTTPClient creates a new ClassesUpdateClassParams object
// with the ability to set a custom HTTPClient for a request.
func NewClassesUpdateClassParamsWithHTTPClient(client *http.Client) *ClassesUpdateClassParams {
	return &ClassesUpdateClassParams{
		HTTPClient: client,
	}
}

/*
ClassesUpdateClassParams contains all the parameters to send to the API endpoint

	for the classes update class operation.

	Typically these are written to a http.Request.
*/
type ClassesUpdateClassParams struct {

	// Body.
	Body *models.APIClassesUpdateClassBody

	/* ClassID.

	   (class::object).id

	   Format: int64
	*/
	ClassID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the classes update class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClassesUpdateClassParams) WithDefaults() *ClassesUpdateClassParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the classes update class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClassesUpdateClassParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the classes update class params
func (o *ClassesUpdateClassParams) WithTimeout(timeout time.Duration) *ClassesUpdateClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the classes update class params
func (o *ClassesUpdateClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the classes update class params
func (o *ClassesUpdateClassParams) WithContext(ctx context.Context) *ClassesUpdateClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the classes update class params
func (o *ClassesUpdateClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the classes update class params
func (o *ClassesUpdateClassParams) WithHTTPClient(client *http.Client) *ClassesUpdateClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the classes update class params
func (o *ClassesUpdateClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the classes update class params
func (o *ClassesUpdateClassParams) WithBody(body *models.APIClassesUpdateClassBody) *ClassesUpdateClassParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the classes update class params
func (o *ClassesUpdateClassParams) SetBody(body *models.APIClassesUpdateClassBody) {
	o.Body = body
}

// WithClassID adds the classID to the classes update class params
func (o *ClassesUpdateClassParams) WithClassID(classID string) *ClassesUpdateClassParams {
	o.SetClassID(classID)
	return o
}

// SetClassID adds the classId to the classes update class params
func (o *ClassesUpdateClassParams) SetClassID(classID string) {
	o.ClassID = classID
}

// WriteToRequest writes these params to a swagger request
func (o *ClassesUpdateClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param class.id
	if err := r.SetPathParam("class.id", o.ClassID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
