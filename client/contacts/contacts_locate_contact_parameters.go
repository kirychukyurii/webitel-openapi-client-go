// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewContactsLocateContactParams creates a new ContactsLocateContactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactsLocateContactParams() *ContactsLocateContactParams {
	return &ContactsLocateContactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactsLocateContactParamsWithTimeout creates a new ContactsLocateContactParams object
// with the ability to set a timeout on a request.
func NewContactsLocateContactParamsWithTimeout(timeout time.Duration) *ContactsLocateContactParams {
	return &ContactsLocateContactParams{
		timeout: timeout,
	}
}

// NewContactsLocateContactParamsWithContext creates a new ContactsLocateContactParams object
// with the ability to set a context for a request.
func NewContactsLocateContactParamsWithContext(ctx context.Context) *ContactsLocateContactParams {
	return &ContactsLocateContactParams{
		Context: ctx,
	}
}

// NewContactsLocateContactParamsWithHTTPClient creates a new ContactsLocateContactParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactsLocateContactParamsWithHTTPClient(client *http.Client) *ContactsLocateContactParams {
	return &ContactsLocateContactParams{
		HTTPClient: client,
	}
}

/*
ContactsLocateContactParams contains all the parameters to send to the API endpoint

	for the contacts locate contact operation.

	Typically these are written to a http.Request.
*/
type ContactsLocateContactParams struct {

	/* Etag.

	     The Contact source IDentifier.
	Accept: `etag` (obsolete+) or `id`.
	*/
	Etag string

	/* Fields.

	   Source Fields to return into result.
	*/
	Fields []string

	/* Mode.

	    The requirement of a special access mode to the Source.

	- READ: Can `fetch` record. [GET]
	- WRITE: Can `update` record. [PUT|PATCH]
	- DELETE: Can `delete` record. [DELETE]

	    Default: "READ"
	*/
	Mode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contacts locate contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactsLocateContactParams) WithDefaults() *ContactsLocateContactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contacts locate contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactsLocateContactParams) SetDefaults() {
	var (
		modeDefault = string("READ")
	)

	val := ContactsLocateContactParams{
		Mode: &modeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the contacts locate contact params
func (o *ContactsLocateContactParams) WithTimeout(timeout time.Duration) *ContactsLocateContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contacts locate contact params
func (o *ContactsLocateContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contacts locate contact params
func (o *ContactsLocateContactParams) WithContext(ctx context.Context) *ContactsLocateContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contacts locate contact params
func (o *ContactsLocateContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contacts locate contact params
func (o *ContactsLocateContactParams) WithHTTPClient(client *http.Client) *ContactsLocateContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contacts locate contact params
func (o *ContactsLocateContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEtag adds the etag to the contacts locate contact params
func (o *ContactsLocateContactParams) WithEtag(etag string) *ContactsLocateContactParams {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the contacts locate contact params
func (o *ContactsLocateContactParams) SetEtag(etag string) {
	o.Etag = etag
}

// WithFields adds the fields to the contacts locate contact params
func (o *ContactsLocateContactParams) WithFields(fields []string) *ContactsLocateContactParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the contacts locate contact params
func (o *ContactsLocateContactParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMode adds the mode to the contacts locate contact params
func (o *ContactsLocateContactParams) WithMode(mode *string) *ContactsLocateContactParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the contacts locate contact params
func (o *ContactsLocateContactParams) SetMode(mode *string) {
	o.Mode = mode
}

// WriteToRequest writes these params to a swagger request
func (o *ContactsLocateContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param etag
	if err := r.SetPathParam("etag", o.Etag); err != nil {
		return err
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.Mode != nil {

		// query param mode
		var qrMode string

		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamContactsLocateContact binds the parameter fields
func (o *ContactsLocateContactParams) bindParamFields(formats strfmt.Registry) []string {
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
