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
)

// NewPhonesDeletePhonesParams creates a new PhonesDeletePhonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPhonesDeletePhonesParams() *PhonesDeletePhonesParams {
	return &PhonesDeletePhonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPhonesDeletePhonesParamsWithTimeout creates a new PhonesDeletePhonesParams object
// with the ability to set a timeout on a request.
func NewPhonesDeletePhonesParamsWithTimeout(timeout time.Duration) *PhonesDeletePhonesParams {
	return &PhonesDeletePhonesParams{
		timeout: timeout,
	}
}

// NewPhonesDeletePhonesParamsWithContext creates a new PhonesDeletePhonesParams object
// with the ability to set a context for a request.
func NewPhonesDeletePhonesParamsWithContext(ctx context.Context) *PhonesDeletePhonesParams {
	return &PhonesDeletePhonesParams{
		Context: ctx,
	}
}

// NewPhonesDeletePhonesParamsWithHTTPClient creates a new PhonesDeletePhonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPhonesDeletePhonesParamsWithHTTPClient(client *http.Client) *PhonesDeletePhonesParams {
	return &PhonesDeletePhonesParams{
		HTTPClient: client,
	}
}

/*
PhonesDeletePhonesParams contains all the parameters to send to the API endpoint

	for the phones delete phones operation.

	Typically these are written to a http.Request.
*/
type PhonesDeletePhonesParams struct {

	/* ContactID.

	   The Contact ID associated with.
	*/
	ContactID string

	/* Etag.

	   Set of linked ID(s) to be removed.
	*/
	Etag []string

	/* Fields.

	   Fields to be retrieved into result.
	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the phones delete phones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhonesDeletePhonesParams) WithDefaults() *PhonesDeletePhonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the phones delete phones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhonesDeletePhonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the phones delete phones params
func (o *PhonesDeletePhonesParams) WithTimeout(timeout time.Duration) *PhonesDeletePhonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the phones delete phones params
func (o *PhonesDeletePhonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the phones delete phones params
func (o *PhonesDeletePhonesParams) WithContext(ctx context.Context) *PhonesDeletePhonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the phones delete phones params
func (o *PhonesDeletePhonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the phones delete phones params
func (o *PhonesDeletePhonesParams) WithHTTPClient(client *http.Client) *PhonesDeletePhonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the phones delete phones params
func (o *PhonesDeletePhonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the phones delete phones params
func (o *PhonesDeletePhonesParams) WithContactID(contactID string) *PhonesDeletePhonesParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the phones delete phones params
func (o *PhonesDeletePhonesParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithEtag adds the etag to the phones delete phones params
func (o *PhonesDeletePhonesParams) WithEtag(etag []string) *PhonesDeletePhonesParams {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the phones delete phones params
func (o *PhonesDeletePhonesParams) SetEtag(etag []string) {
	o.Etag = etag
}

// WithFields adds the fields to the phones delete phones params
func (o *PhonesDeletePhonesParams) WithFields(fields []string) *PhonesDeletePhonesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the phones delete phones params
func (o *PhonesDeletePhonesParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *PhonesDeletePhonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
		return err
	}

	if o.Etag != nil {

		// binding items for etag
		joinedEtag := o.bindParamEtag(reg)

		// query array param etag
		if err := r.SetQueryParam("etag", joinedEtag...); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPhonesDeletePhones binds the parameter etag
func (o *PhonesDeletePhonesParams) bindParamEtag(formats strfmt.Registry) []string {
	etagIR := o.Etag

	var etagIC []string
	for _, etagIIR := range etagIR { // explode []string

		etagIIV := etagIIR // string as string
		etagIC = append(etagIC, etagIIV)
	}

	// items.CollectionFormat: "multi"
	etagIS := swag.JoinByFormat(etagIC, "multi")

	return etagIS
}

// bindParamPhonesDeletePhones binds the parameter fields
func (o *PhonesDeletePhonesParams) bindParamFields(formats strfmt.Registry) []string {
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
