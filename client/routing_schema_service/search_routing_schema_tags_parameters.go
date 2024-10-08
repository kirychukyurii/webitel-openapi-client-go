// Code generated by go-swagger; DO NOT EDIT.

package routing_schema_service

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

// NewSearchRoutingSchemaTagsParams creates a new SearchRoutingSchemaTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchRoutingSchemaTagsParams() *SearchRoutingSchemaTagsParams {
	return &SearchRoutingSchemaTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchRoutingSchemaTagsParamsWithTimeout creates a new SearchRoutingSchemaTagsParams object
// with the ability to set a timeout on a request.
func NewSearchRoutingSchemaTagsParamsWithTimeout(timeout time.Duration) *SearchRoutingSchemaTagsParams {
	return &SearchRoutingSchemaTagsParams{
		timeout: timeout,
	}
}

// NewSearchRoutingSchemaTagsParamsWithContext creates a new SearchRoutingSchemaTagsParams object
// with the ability to set a context for a request.
func NewSearchRoutingSchemaTagsParamsWithContext(ctx context.Context) *SearchRoutingSchemaTagsParams {
	return &SearchRoutingSchemaTagsParams{
		Context: ctx,
	}
}

// NewSearchRoutingSchemaTagsParamsWithHTTPClient creates a new SearchRoutingSchemaTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchRoutingSchemaTagsParamsWithHTTPClient(client *http.Client) *SearchRoutingSchemaTagsParams {
	return &SearchRoutingSchemaTagsParams{
		HTTPClient: client,
	}
}

/*
SearchRoutingSchemaTagsParams contains all the parameters to send to the API endpoint

	for the search routing schema tags operation.

	Typically these are written to a http.Request.
*/
type SearchRoutingSchemaTagsParams struct {

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

	// Type.
	Type []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search routing schema tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRoutingSchemaTagsParams) WithDefaults() *SearchRoutingSchemaTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search routing schema tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRoutingSchemaTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithTimeout(timeout time.Duration) *SearchRoutingSchemaTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithContext(ctx context.Context) *SearchRoutingSchemaTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithHTTPClient(client *http.Client) *SearchRoutingSchemaTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithFields(fields []string) *SearchRoutingSchemaTagsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithPage adds the page to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithPage(page *int32) *SearchRoutingSchemaTagsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithQ(q *string) *SearchRoutingSchemaTagsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithSize(size *int32) *SearchRoutingSchemaTagsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithSort(sort *string) *SearchRoutingSchemaTagsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithType adds the typeVar to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) WithType(typeVar []string) *SearchRoutingSchemaTagsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the search routing schema tags params
func (o *SearchRoutingSchemaTagsParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SearchRoutingSchemaTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Type != nil {

		// binding items for type
		joinedType := o.bindParamType(reg)

		// query array param type
		if err := r.SetQueryParam("type", joinedType...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchRoutingSchemaTags binds the parameter fields
func (o *SearchRoutingSchemaTagsParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSearchRoutingSchemaTags binds the parameter type
func (o *SearchRoutingSchemaTagsParams) bindParamType(formats strfmt.Registry) []string {
	typeIR := o.Type

	var typeIC []string
	for _, typeIIR := range typeIR { // explode []string

		typeIIV := typeIIR // string as string
		typeIC = append(typeIC, typeIIV)
	}

	// items.CollectionFormat: "multi"
	typeIS := swag.JoinByFormat(typeIC, "multi")

	return typeIS
}
