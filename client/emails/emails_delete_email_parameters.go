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
)

// NewEmailsDeleteEmailParams creates a new EmailsDeleteEmailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmailsDeleteEmailParams() *EmailsDeleteEmailParams {
	return &EmailsDeleteEmailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmailsDeleteEmailParamsWithTimeout creates a new EmailsDeleteEmailParams object
// with the ability to set a timeout on a request.
func NewEmailsDeleteEmailParamsWithTimeout(timeout time.Duration) *EmailsDeleteEmailParams {
	return &EmailsDeleteEmailParams{
		timeout: timeout,
	}
}

// NewEmailsDeleteEmailParamsWithContext creates a new EmailsDeleteEmailParams object
// with the ability to set a context for a request.
func NewEmailsDeleteEmailParamsWithContext(ctx context.Context) *EmailsDeleteEmailParams {
	return &EmailsDeleteEmailParams{
		Context: ctx,
	}
}

// NewEmailsDeleteEmailParamsWithHTTPClient creates a new EmailsDeleteEmailParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmailsDeleteEmailParamsWithHTTPClient(client *http.Client) *EmailsDeleteEmailParams {
	return &EmailsDeleteEmailParams{
		HTTPClient: client,
	}
}

/*
EmailsDeleteEmailParams contains all the parameters to send to the API endpoint

	for the emails delete email operation.

	Typically these are written to a http.Request.
*/
type EmailsDeleteEmailParams struct {

	/* ContactID.

	   Contact ID associated with.
	*/
	ContactID string

	/* Etag.

	   Unique ID to remove.
	*/
	Etag string

	/* Fields.

	   Fields to be retrieved as a result.
	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the emails delete email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmailsDeleteEmailParams) WithDefaults() *EmailsDeleteEmailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the emails delete email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmailsDeleteEmailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the emails delete email params
func (o *EmailsDeleteEmailParams) WithTimeout(timeout time.Duration) *EmailsDeleteEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emails delete email params
func (o *EmailsDeleteEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emails delete email params
func (o *EmailsDeleteEmailParams) WithContext(ctx context.Context) *EmailsDeleteEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emails delete email params
func (o *EmailsDeleteEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emails delete email params
func (o *EmailsDeleteEmailParams) WithHTTPClient(client *http.Client) *EmailsDeleteEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emails delete email params
func (o *EmailsDeleteEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the emails delete email params
func (o *EmailsDeleteEmailParams) WithContactID(contactID string) *EmailsDeleteEmailParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the emails delete email params
func (o *EmailsDeleteEmailParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithEtag adds the etag to the emails delete email params
func (o *EmailsDeleteEmailParams) WithEtag(etag string) *EmailsDeleteEmailParams {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the emails delete email params
func (o *EmailsDeleteEmailParams) SetEtag(etag string) {
	o.Etag = etag
}

// WithFields adds the fields to the emails delete email params
func (o *EmailsDeleteEmailParams) WithFields(fields []string) *EmailsDeleteEmailParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the emails delete email params
func (o *EmailsDeleteEmailParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *EmailsDeleteEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamEmailsDeleteEmail binds the parameter fields
func (o *EmailsDeleteEmailParams) bindParamFields(formats strfmt.Registry) []string {
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
