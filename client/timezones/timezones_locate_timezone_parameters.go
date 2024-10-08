// Code generated by go-swagger; DO NOT EDIT.

package timezones

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

// NewTimezonesLocateTimezoneParams creates a new TimezonesLocateTimezoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTimezonesLocateTimezoneParams() *TimezonesLocateTimezoneParams {
	return &TimezonesLocateTimezoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTimezonesLocateTimezoneParamsWithTimeout creates a new TimezonesLocateTimezoneParams object
// with the ability to set a timeout on a request.
func NewTimezonesLocateTimezoneParamsWithTimeout(timeout time.Duration) *TimezonesLocateTimezoneParams {
	return &TimezonesLocateTimezoneParams{
		timeout: timeout,
	}
}

// NewTimezonesLocateTimezoneParamsWithContext creates a new TimezonesLocateTimezoneParams object
// with the ability to set a context for a request.
func NewTimezonesLocateTimezoneParamsWithContext(ctx context.Context) *TimezonesLocateTimezoneParams {
	return &TimezonesLocateTimezoneParams{
		Context: ctx,
	}
}

// NewTimezonesLocateTimezoneParamsWithHTTPClient creates a new TimezonesLocateTimezoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewTimezonesLocateTimezoneParamsWithHTTPClient(client *http.Client) *TimezonesLocateTimezoneParams {
	return &TimezonesLocateTimezoneParams{
		HTTPClient: client,
	}
}

/*
TimezonesLocateTimezoneParams contains all the parameters to send to the API endpoint

	for the timezones locate timezone operation.

	Typically these are written to a http.Request.
*/
type TimezonesLocateTimezoneParams struct {

	/* ContactID.

	   Contact source ID.
	*/
	ContactID string

	/* Etag.

	     Unique timezone link IDentifier.
	Accept: `etag` (obsolete+) or `id`.

	     Format: \w+
	*/
	Etag string

	/* Fields.

	   Fields to be retrieved into result.
	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the timezones locate timezone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimezonesLocateTimezoneParams) WithDefaults() *TimezonesLocateTimezoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the timezones locate timezone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimezonesLocateTimezoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) WithTimeout(timeout time.Duration) *TimezonesLocateTimezoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) WithContext(ctx context.Context) *TimezonesLocateTimezoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) WithHTTPClient(client *http.Client) *TimezonesLocateTimezoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) WithContactID(contactID string) *TimezonesLocateTimezoneParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithEtag adds the etag to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) WithEtag(etag string) *TimezonesLocateTimezoneParams {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) SetEtag(etag string) {
	o.Etag = etag
}

// WithFields adds the fields to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) WithFields(fields []string) *TimezonesLocateTimezoneParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the timezones locate timezone params
func (o *TimezonesLocateTimezoneParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *TimezonesLocateTimezoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
		return err
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTimezonesLocateTimezone binds the parameter fields
func (o *TimezonesLocateTimezoneParams) bindParamFields(formats strfmt.Registry) []string {
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
