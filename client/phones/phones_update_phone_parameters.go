// Code generated by go-swagger; DO NOT EDIT.

package phones

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

// NewPhonesUpdatePhoneParams creates a new PhonesUpdatePhoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPhonesUpdatePhoneParams() *PhonesUpdatePhoneParams {
	return &PhonesUpdatePhoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPhonesUpdatePhoneParamsWithTimeout creates a new PhonesUpdatePhoneParams object
// with the ability to set a timeout on a request.
func NewPhonesUpdatePhoneParamsWithTimeout(timeout time.Duration) *PhonesUpdatePhoneParams {
	return &PhonesUpdatePhoneParams{
		timeout: timeout,
	}
}

// NewPhonesUpdatePhoneParamsWithContext creates a new PhonesUpdatePhoneParams object
// with the ability to set a context for a request.
func NewPhonesUpdatePhoneParamsWithContext(ctx context.Context) *PhonesUpdatePhoneParams {
	return &PhonesUpdatePhoneParams{
		Context: ctx,
	}
}

// NewPhonesUpdatePhoneParamsWithHTTPClient creates a new PhonesUpdatePhoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewPhonesUpdatePhoneParamsWithHTTPClient(client *http.Client) *PhonesUpdatePhoneParams {
	return &PhonesUpdatePhoneParams{
		HTTPClient: client,
	}
}

/*
PhonesUpdatePhoneParams contains all the parameters to send to the API endpoint

	for the phones update phone operation.

	Typically these are written to a http.Request.
*/
type PhonesUpdatePhoneParams struct {

	/* ContactID.

	   The Contact ID to be associated with.
	*/
	ContactID string

	/* Etag.

	   Unique ID of the latest version of an existing resorce.
	*/
	Etag string

	/* Fields.

	   Fields to be retrieved into result of changes.
	*/
	Fields []string

	/* Input.

	   NEW Update of the phone number details.
	*/
	Input *models.PhonesUpdatePhoneParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the phones update phone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhonesUpdatePhoneParams) WithDefaults() *PhonesUpdatePhoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the phones update phone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhonesUpdatePhoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the phones update phone params
func (o *PhonesUpdatePhoneParams) WithTimeout(timeout time.Duration) *PhonesUpdatePhoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the phones update phone params
func (o *PhonesUpdatePhoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the phones update phone params
func (o *PhonesUpdatePhoneParams) WithContext(ctx context.Context) *PhonesUpdatePhoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the phones update phone params
func (o *PhonesUpdatePhoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the phones update phone params
func (o *PhonesUpdatePhoneParams) WithHTTPClient(client *http.Client) *PhonesUpdatePhoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the phones update phone params
func (o *PhonesUpdatePhoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the phones update phone params
func (o *PhonesUpdatePhoneParams) WithContactID(contactID string) *PhonesUpdatePhoneParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the phones update phone params
func (o *PhonesUpdatePhoneParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithEtag adds the etag to the phones update phone params
func (o *PhonesUpdatePhoneParams) WithEtag(etag string) *PhonesUpdatePhoneParams {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the phones update phone params
func (o *PhonesUpdatePhoneParams) SetEtag(etag string) {
	o.Etag = etag
}

// WithFields adds the fields to the phones update phone params
func (o *PhonesUpdatePhoneParams) WithFields(fields []string) *PhonesUpdatePhoneParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the phones update phone params
func (o *PhonesUpdatePhoneParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInput adds the input to the phones update phone params
func (o *PhonesUpdatePhoneParams) WithInput(input *models.PhonesUpdatePhoneParamsBody) *PhonesUpdatePhoneParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the phones update phone params
func (o *PhonesUpdatePhoneParams) SetInput(input *models.PhonesUpdatePhoneParamsBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PhonesUpdatePhoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamPhonesUpdatePhone binds the parameter fields
func (o *PhonesUpdatePhoneParams) bindParamFields(formats strfmt.Registry) []string {
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
