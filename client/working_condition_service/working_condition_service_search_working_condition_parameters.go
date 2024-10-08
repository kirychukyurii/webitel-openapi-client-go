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

// NewWorkingConditionServiceSearchWorkingConditionParams creates a new WorkingConditionServiceSearchWorkingConditionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkingConditionServiceSearchWorkingConditionParams() *WorkingConditionServiceSearchWorkingConditionParams {
	return &WorkingConditionServiceSearchWorkingConditionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkingConditionServiceSearchWorkingConditionParamsWithTimeout creates a new WorkingConditionServiceSearchWorkingConditionParams object
// with the ability to set a timeout on a request.
func NewWorkingConditionServiceSearchWorkingConditionParamsWithTimeout(timeout time.Duration) *WorkingConditionServiceSearchWorkingConditionParams {
	return &WorkingConditionServiceSearchWorkingConditionParams{
		timeout: timeout,
	}
}

// NewWorkingConditionServiceSearchWorkingConditionParamsWithContext creates a new WorkingConditionServiceSearchWorkingConditionParams object
// with the ability to set a context for a request.
func NewWorkingConditionServiceSearchWorkingConditionParamsWithContext(ctx context.Context) *WorkingConditionServiceSearchWorkingConditionParams {
	return &WorkingConditionServiceSearchWorkingConditionParams{
		Context: ctx,
	}
}

// NewWorkingConditionServiceSearchWorkingConditionParamsWithHTTPClient creates a new WorkingConditionServiceSearchWorkingConditionParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkingConditionServiceSearchWorkingConditionParamsWithHTTPClient(client *http.Client) *WorkingConditionServiceSearchWorkingConditionParams {
	return &WorkingConditionServiceSearchWorkingConditionParams{
		HTTPClient: client,
	}
}

/*
WorkingConditionServiceSearchWorkingConditionParams contains all the parameters to send to the API endpoint

	for the working condition service search working condition operation.

	Typically these are written to a http.Request.
*/
type WorkingConditionServiceSearchWorkingConditionParams struct {

	// Fields.
	Fields []string

	// Page.
	//
	// Format: int32
	Page *int32

	// Q.
	Q *string

	// Size.
	//
	// Format: int32
	Size *int32

	// Sort.
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the working condition service search working condition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithDefaults() *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the working condition service search working condition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithTimeout(timeout time.Duration) *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithContext(ctx context.Context) *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithHTTPClient(client *http.Client) *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithFields(fields []string) *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithPage adds the page to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithPage(page *int32) *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithQ(q *string) *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithSize(size *int32) *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) WithSort(sort *string) *WorkingConditionServiceSearchWorkingConditionParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the working condition service search working condition params
func (o *WorkingConditionServiceSearchWorkingConditionParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *WorkingConditionServiceSearchWorkingConditionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamWorkingConditionServiceSearchWorkingCondition binds the parameter fields
func (o *WorkingConditionServiceSearchWorkingConditionParams) bindParamFields(formats strfmt.Registry) []string {
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
