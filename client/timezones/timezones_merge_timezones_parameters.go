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

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// NewTimezonesMergeTimezonesParams creates a new TimezonesMergeTimezonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTimezonesMergeTimezonesParams() *TimezonesMergeTimezonesParams {
	return &TimezonesMergeTimezonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTimezonesMergeTimezonesParamsWithTimeout creates a new TimezonesMergeTimezonesParams object
// with the ability to set a timeout on a request.
func NewTimezonesMergeTimezonesParamsWithTimeout(timeout time.Duration) *TimezonesMergeTimezonesParams {
	return &TimezonesMergeTimezonesParams{
		timeout: timeout,
	}
}

// NewTimezonesMergeTimezonesParamsWithContext creates a new TimezonesMergeTimezonesParams object
// with the ability to set a context for a request.
func NewTimezonesMergeTimezonesParamsWithContext(ctx context.Context) *TimezonesMergeTimezonesParams {
	return &TimezonesMergeTimezonesParams{
		Context: ctx,
	}
}

// NewTimezonesMergeTimezonesParamsWithHTTPClient creates a new TimezonesMergeTimezonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewTimezonesMergeTimezonesParamsWithHTTPClient(client *http.Client) *TimezonesMergeTimezonesParams {
	return &TimezonesMergeTimezonesParams{
		HTTPClient: client,
	}
}

/*
TimezonesMergeTimezonesParams contains all the parameters to send to the API endpoint

	for the timezones merge timezones operation.

	Typically these are written to a http.Request.
*/
type TimezonesMergeTimezonesParams struct {

	/* ContactID.

	   Link contact ID.
	*/
	ContactID string

	/* Fields.

	   Fields to be retrieved as a result.
	*/
	Fields []string

	/* Input.

	     Array of the unique User(s) to associate with the Contact.
	Any duplicate of an already linked user{id} will result in an error.
	*/
	Input []*models.WebitelContactsInputTimezone

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the timezones merge timezones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimezonesMergeTimezonesParams) WithDefaults() *TimezonesMergeTimezonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the timezones merge timezones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimezonesMergeTimezonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) WithTimeout(timeout time.Duration) *TimezonesMergeTimezonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) WithContext(ctx context.Context) *TimezonesMergeTimezonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) WithHTTPClient(client *http.Client) *TimezonesMergeTimezonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) WithContactID(contactID string) *TimezonesMergeTimezonesParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithFields adds the fields to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) WithFields(fields []string) *TimezonesMergeTimezonesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInput adds the input to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) WithInput(input []*models.WebitelContactsInputTimezone) *TimezonesMergeTimezonesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the timezones merge timezones params
func (o *TimezonesMergeTimezonesParams) SetInput(input []*models.WebitelContactsInputTimezone) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *TimezonesMergeTimezonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
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
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTimezonesMergeTimezones binds the parameter fields
func (o *TimezonesMergeTimezonesParams) bindParamFields(formats strfmt.Registry) []string {
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
