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
	"github.com/go-openapi/swag"
)

// NewWorkingConditionServiceReadWorkingConditionParams creates a new WorkingConditionServiceReadWorkingConditionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkingConditionServiceReadWorkingConditionParams() *WorkingConditionServiceReadWorkingConditionParams {
	return &WorkingConditionServiceReadWorkingConditionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkingConditionServiceReadWorkingConditionParamsWithTimeout creates a new WorkingConditionServiceReadWorkingConditionParams object
// with the ability to set a timeout on a request.
func NewWorkingConditionServiceReadWorkingConditionParamsWithTimeout(timeout time.Duration) *WorkingConditionServiceReadWorkingConditionParams {
	return &WorkingConditionServiceReadWorkingConditionParams{
		timeout: timeout,
	}
}

// NewWorkingConditionServiceReadWorkingConditionParamsWithContext creates a new WorkingConditionServiceReadWorkingConditionParams object
// with the ability to set a context for a request.
func NewWorkingConditionServiceReadWorkingConditionParamsWithContext(ctx context.Context) *WorkingConditionServiceReadWorkingConditionParams {
	return &WorkingConditionServiceReadWorkingConditionParams{
		Context: ctx,
	}
}

// NewWorkingConditionServiceReadWorkingConditionParamsWithHTTPClient creates a new WorkingConditionServiceReadWorkingConditionParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkingConditionServiceReadWorkingConditionParamsWithHTTPClient(client *http.Client) *WorkingConditionServiceReadWorkingConditionParams {
	return &WorkingConditionServiceReadWorkingConditionParams{
		HTTPClient: client,
	}
}

/*
WorkingConditionServiceReadWorkingConditionParams contains all the parameters to send to the API endpoint

	for the working condition service read working condition operation.

	Typically these are written to a http.Request.
*/
type WorkingConditionServiceReadWorkingConditionParams struct {

	// Fields.
	Fields []string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the working condition service read working condition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkingConditionServiceReadWorkingConditionParams) WithDefaults() *WorkingConditionServiceReadWorkingConditionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the working condition service read working condition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkingConditionServiceReadWorkingConditionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) WithTimeout(timeout time.Duration) *WorkingConditionServiceReadWorkingConditionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) WithContext(ctx context.Context) *WorkingConditionServiceReadWorkingConditionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) WithHTTPClient(client *http.Client) *WorkingConditionServiceReadWorkingConditionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) WithFields(fields []string) *WorkingConditionServiceReadWorkingConditionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) WithID(id string) *WorkingConditionServiceReadWorkingConditionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the working condition service read working condition params
func (o *WorkingConditionServiceReadWorkingConditionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WorkingConditionServiceReadWorkingConditionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
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

// bindParamWorkingConditionServiceReadWorkingCondition binds the parameter fields
func (o *WorkingConditionServiceReadWorkingConditionParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "multi"
	fieldsIS := swag.JoinByFormat(fieldsIC, "multi")

	return fieldsIS
}
