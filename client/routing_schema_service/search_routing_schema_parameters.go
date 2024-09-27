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

// NewSearchRoutingSchemaParams creates a new SearchRoutingSchemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchRoutingSchemaParams() *SearchRoutingSchemaParams {
	return &SearchRoutingSchemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchRoutingSchemaParamsWithTimeout creates a new SearchRoutingSchemaParams object
// with the ability to set a timeout on a request.
func NewSearchRoutingSchemaParamsWithTimeout(timeout time.Duration) *SearchRoutingSchemaParams {
	return &SearchRoutingSchemaParams{
		timeout: timeout,
	}
}

// NewSearchRoutingSchemaParamsWithContext creates a new SearchRoutingSchemaParams object
// with the ability to set a context for a request.
func NewSearchRoutingSchemaParamsWithContext(ctx context.Context) *SearchRoutingSchemaParams {
	return &SearchRoutingSchemaParams{
		Context: ctx,
	}
}

// NewSearchRoutingSchemaParamsWithHTTPClient creates a new SearchRoutingSchemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchRoutingSchemaParamsWithHTTPClient(client *http.Client) *SearchRoutingSchemaParams {
	return &SearchRoutingSchemaParams{
		HTTPClient: client,
	}
}

/*
SearchRoutingSchemaParams contains all the parameters to send to the API endpoint

	for the search routing schema operation.

	Typically these are written to a http.Request.
*/
type SearchRoutingSchemaParams struct {

	// Editor.
	//
	// Format: boolean
	Editor *bool

	// Fields.
	Fields []string

	// ID.
	ID []int64

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

	// Tags.
	Tags []string

	// Type.
	Type []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search routing schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRoutingSchemaParams) WithDefaults() *SearchRoutingSchemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search routing schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRoutingSchemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search routing schema params
func (o *SearchRoutingSchemaParams) WithTimeout(timeout time.Duration) *SearchRoutingSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search routing schema params
func (o *SearchRoutingSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search routing schema params
func (o *SearchRoutingSchemaParams) WithContext(ctx context.Context) *SearchRoutingSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search routing schema params
func (o *SearchRoutingSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search routing schema params
func (o *SearchRoutingSchemaParams) WithHTTPClient(client *http.Client) *SearchRoutingSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search routing schema params
func (o *SearchRoutingSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEditor adds the editor to the search routing schema params
func (o *SearchRoutingSchemaParams) WithEditor(editor *bool) *SearchRoutingSchemaParams {
	o.SetEditor(editor)
	return o
}

// SetEditor adds the editor to the search routing schema params
func (o *SearchRoutingSchemaParams) SetEditor(editor *bool) {
	o.Editor = editor
}

// WithFields adds the fields to the search routing schema params
func (o *SearchRoutingSchemaParams) WithFields(fields []string) *SearchRoutingSchemaParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search routing schema params
func (o *SearchRoutingSchemaParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the search routing schema params
func (o *SearchRoutingSchemaParams) WithID(id []int64) *SearchRoutingSchemaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search routing schema params
func (o *SearchRoutingSchemaParams) SetID(id []int64) {
	o.ID = id
}

// WithName adds the name to the search routing schema params
func (o *SearchRoutingSchemaParams) WithName(name *string) *SearchRoutingSchemaParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the search routing schema params
func (o *SearchRoutingSchemaParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the search routing schema params
func (o *SearchRoutingSchemaParams) WithPage(page *int32) *SearchRoutingSchemaParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search routing schema params
func (o *SearchRoutingSchemaParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the search routing schema params
func (o *SearchRoutingSchemaParams) WithQ(q *string) *SearchRoutingSchemaParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search routing schema params
func (o *SearchRoutingSchemaParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the search routing schema params
func (o *SearchRoutingSchemaParams) WithSize(size *int32) *SearchRoutingSchemaParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search routing schema params
func (o *SearchRoutingSchemaParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the search routing schema params
func (o *SearchRoutingSchemaParams) WithSort(sort *string) *SearchRoutingSchemaParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search routing schema params
func (o *SearchRoutingSchemaParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTags adds the tags to the search routing schema params
func (o *SearchRoutingSchemaParams) WithTags(tags []string) *SearchRoutingSchemaParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the search routing schema params
func (o *SearchRoutingSchemaParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithType adds the typeVar to the search routing schema params
func (o *SearchRoutingSchemaParams) WithType(typeVar []string) *SearchRoutingSchemaParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the search routing schema params
func (o *SearchRoutingSchemaParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SearchRoutingSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Editor != nil {

		// query param editor
		var qrEditor bool

		if o.Editor != nil {
			qrEditor = *o.Editor
		}
		qEditor := swag.FormatBool(qrEditor)
		if qEditor != "" {

			if err := r.SetQueryParam("editor", qEditor); err != nil {
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

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// query array param tags
		if err := r.SetQueryParam("tags", joinedTags...); err != nil {
			return err
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

// bindParamSearchRoutingSchema binds the parameter fields
func (o *SearchRoutingSchemaParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSearchRoutingSchema binds the parameter id
func (o *SearchRoutingSchemaParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []int64

		iDIIV := swag.FormatInt64(iDIIR) // int64 as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}

// bindParamSearchRoutingSchema binds the parameter tags
func (o *SearchRoutingSchemaParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: "multi"
	tagsIS := swag.JoinByFormat(tagsIC, "multi")

	return tagsIS
}

// bindParamSearchRoutingSchema binds the parameter type
func (o *SearchRoutingSchemaParams) bindParamType(formats strfmt.Registry) []string {
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
