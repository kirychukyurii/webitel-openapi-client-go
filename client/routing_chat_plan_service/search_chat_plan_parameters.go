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
)

// NewSearchChatPlanParams creates a new SearchChatPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchChatPlanParams() *SearchChatPlanParams {
	return &SearchChatPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchChatPlanParamsWithTimeout creates a new SearchChatPlanParams object
// with the ability to set a timeout on a request.
func NewSearchChatPlanParamsWithTimeout(timeout time.Duration) *SearchChatPlanParams {
	return &SearchChatPlanParams{
		timeout: timeout,
	}
}

// NewSearchChatPlanParamsWithContext creates a new SearchChatPlanParams object
// with the ability to set a context for a request.
func NewSearchChatPlanParamsWithContext(ctx context.Context) *SearchChatPlanParams {
	return &SearchChatPlanParams{
		Context: ctx,
	}
}

// NewSearchChatPlanParamsWithHTTPClient creates a new SearchChatPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchChatPlanParamsWithHTTPClient(client *http.Client) *SearchChatPlanParams {
	return &SearchChatPlanParams{
		HTTPClient: client,
	}
}

/*
SearchChatPlanParams contains all the parameters to send to the API endpoint

	for the search chat plan operation.

	Typically these are written to a http.Request.
*/
type SearchChatPlanParams struct {

	// Enabled.
	//
	// Format: boolean
	Enabled *bool

	// Fields.
	Fields []string

	// ID.
	ID []int32

	// Name.
	Name *string

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

// WithDefaults hydrates default values in the search chat plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchChatPlanParams) WithDefaults() *SearchChatPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search chat plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchChatPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search chat plan params
func (o *SearchChatPlanParams) WithTimeout(timeout time.Duration) *SearchChatPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search chat plan params
func (o *SearchChatPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search chat plan params
func (o *SearchChatPlanParams) WithContext(ctx context.Context) *SearchChatPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search chat plan params
func (o *SearchChatPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search chat plan params
func (o *SearchChatPlanParams) WithHTTPClient(client *http.Client) *SearchChatPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search chat plan params
func (o *SearchChatPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the search chat plan params
func (o *SearchChatPlanParams) WithEnabled(enabled *bool) *SearchChatPlanParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the search chat plan params
func (o *SearchChatPlanParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithFields adds the fields to the search chat plan params
func (o *SearchChatPlanParams) WithFields(fields []string) *SearchChatPlanParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search chat plan params
func (o *SearchChatPlanParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the search chat plan params
func (o *SearchChatPlanParams) WithID(id []int32) *SearchChatPlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search chat plan params
func (o *SearchChatPlanParams) SetID(id []int32) {
	o.ID = id
}

// WithName adds the name to the search chat plan params
func (o *SearchChatPlanParams) WithName(name *string) *SearchChatPlanParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the search chat plan params
func (o *SearchChatPlanParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the search chat plan params
func (o *SearchChatPlanParams) WithPage(page *int32) *SearchChatPlanParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search chat plan params
func (o *SearchChatPlanParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the search chat plan params
func (o *SearchChatPlanParams) WithQ(q *string) *SearchChatPlanParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search chat plan params
func (o *SearchChatPlanParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the search chat plan params
func (o *SearchChatPlanParams) WithSize(size *int32) *SearchChatPlanParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search chat plan params
func (o *SearchChatPlanParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the search chat plan params
func (o *SearchChatPlanParams) WithSort(sort *string) *SearchChatPlanParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search chat plan params
func (o *SearchChatPlanParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchChatPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
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

// bindParamSearchChatPlan binds the parameter fields
func (o *SearchChatPlanParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSearchChatPlan binds the parameter id
func (o *SearchChatPlanParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []int32

		iDIIV := swag.FormatInt32(iDIIR) // int32 as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}
