// Code generated by go-swagger; DO NOT EDIT.

package emails

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

// NewEmailsUpdateEmail2Params creates a new EmailsUpdateEmail2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmailsUpdateEmail2Params() *EmailsUpdateEmail2Params {
	return &EmailsUpdateEmail2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmailsUpdateEmail2ParamsWithTimeout creates a new EmailsUpdateEmail2Params object
// with the ability to set a timeout on a request.
func NewEmailsUpdateEmail2ParamsWithTimeout(timeout time.Duration) *EmailsUpdateEmail2Params {
	return &EmailsUpdateEmail2Params{
		timeout: timeout,
	}
}

// NewEmailsUpdateEmail2ParamsWithContext creates a new EmailsUpdateEmail2Params object
// with the ability to set a context for a request.
func NewEmailsUpdateEmail2ParamsWithContext(ctx context.Context) *EmailsUpdateEmail2Params {
	return &EmailsUpdateEmail2Params{
		Context: ctx,
	}
}

// NewEmailsUpdateEmail2ParamsWithHTTPClient creates a new EmailsUpdateEmail2Params object
// with the ability to set a custom HTTPClient for a request.
func NewEmailsUpdateEmail2ParamsWithHTTPClient(client *http.Client) *EmailsUpdateEmail2Params {
	return &EmailsUpdateEmail2Params{
		HTTPClient: client,
	}
}

/*
EmailsUpdateEmail2Params contains all the parameters to send to the API endpoint

	for the emails update email2 operation.

	Typically these are written to a http.Request.
*/
type EmailsUpdateEmail2Params struct {

	/* ContactID.

	   Link contact ID.
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

	   NEW Update of the email address link.
	*/
	Input *models.EmailsUpdateEmail2ParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the emails update email2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmailsUpdateEmail2Params) WithDefaults() *EmailsUpdateEmail2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the emails update email2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmailsUpdateEmail2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the emails update email2 params
func (o *EmailsUpdateEmail2Params) WithTimeout(timeout time.Duration) *EmailsUpdateEmail2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emails update email2 params
func (o *EmailsUpdateEmail2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emails update email2 params
func (o *EmailsUpdateEmail2Params) WithContext(ctx context.Context) *EmailsUpdateEmail2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emails update email2 params
func (o *EmailsUpdateEmail2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emails update email2 params
func (o *EmailsUpdateEmail2Params) WithHTTPClient(client *http.Client) *EmailsUpdateEmail2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emails update email2 params
func (o *EmailsUpdateEmail2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the emails update email2 params
func (o *EmailsUpdateEmail2Params) WithContactID(contactID string) *EmailsUpdateEmail2Params {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the emails update email2 params
func (o *EmailsUpdateEmail2Params) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithEtag adds the etag to the emails update email2 params
func (o *EmailsUpdateEmail2Params) WithEtag(etag string) *EmailsUpdateEmail2Params {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the emails update email2 params
func (o *EmailsUpdateEmail2Params) SetEtag(etag string) {
	o.Etag = etag
}

// WithFields adds the fields to the emails update email2 params
func (o *EmailsUpdateEmail2Params) WithFields(fields []string) *EmailsUpdateEmail2Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the emails update email2 params
func (o *EmailsUpdateEmail2Params) SetFields(fields []string) {
	o.Fields = fields
}

// WithInput adds the input to the emails update email2 params
func (o *EmailsUpdateEmail2Params) WithInput(input *models.EmailsUpdateEmail2ParamsBody) *EmailsUpdateEmail2Params {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the emails update email2 params
func (o *EmailsUpdateEmail2Params) SetInput(input *models.EmailsUpdateEmail2ParamsBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *EmailsUpdateEmail2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamEmailsUpdateEmail2 binds the parameter fields
func (o *EmailsUpdateEmail2Params) bindParamFields(formats strfmt.Registry) []string {
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
