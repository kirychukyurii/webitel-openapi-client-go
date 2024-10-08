// Code generated by go-swagger; DO NOT EDIT.

package domains

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

// NewDomainsSearchDomainsParams creates a new DomainsSearchDomainsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainsSearchDomainsParams() *DomainsSearchDomainsParams {
	return &DomainsSearchDomainsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainsSearchDomainsParamsWithTimeout creates a new DomainsSearchDomainsParams object
// with the ability to set a timeout on a request.
func NewDomainsSearchDomainsParamsWithTimeout(timeout time.Duration) *DomainsSearchDomainsParams {
	return &DomainsSearchDomainsParams{
		timeout: timeout,
	}
}

// NewDomainsSearchDomainsParamsWithContext creates a new DomainsSearchDomainsParams object
// with the ability to set a context for a request.
func NewDomainsSearchDomainsParamsWithContext(ctx context.Context) *DomainsSearchDomainsParams {
	return &DomainsSearchDomainsParams{
		Context: ctx,
	}
}

// NewDomainsSearchDomainsParamsWithHTTPClient creates a new DomainsSearchDomainsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainsSearchDomainsParamsWithHTTPClient(client *http.Client) *DomainsSearchDomainsParams {
	return &DomainsSearchDomainsParams{
		HTTPClient: client,
	}
}

/*
DomainsSearchDomainsParams contains all the parameters to send to the API endpoint

	for the domains search domains operation.

	Typically these are written to a http.Request.
*/
type DomainsSearchDomainsParams struct {

	/* Domain.

	   like
	*/
	Domain *string

	/* Fields.

	   attrs
	*/
	Fields []string

	// Page.
	//
	// Format: int64
	Page *string

	// Size.
	//
	// Format: int64
	Size *string

	// Sort.
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domains search domains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainsSearchDomainsParams) WithDefaults() *DomainsSearchDomainsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domains search domains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainsSearchDomainsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the domains search domains params
func (o *DomainsSearchDomainsParams) WithTimeout(timeout time.Duration) *DomainsSearchDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domains search domains params
func (o *DomainsSearchDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domains search domains params
func (o *DomainsSearchDomainsParams) WithContext(ctx context.Context) *DomainsSearchDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domains search domains params
func (o *DomainsSearchDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domains search domains params
func (o *DomainsSearchDomainsParams) WithHTTPClient(client *http.Client) *DomainsSearchDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domains search domains params
func (o *DomainsSearchDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the domains search domains params
func (o *DomainsSearchDomainsParams) WithDomain(domain *string) *DomainsSearchDomainsParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the domains search domains params
func (o *DomainsSearchDomainsParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithFields adds the fields to the domains search domains params
func (o *DomainsSearchDomainsParams) WithFields(fields []string) *DomainsSearchDomainsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the domains search domains params
func (o *DomainsSearchDomainsParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithPage adds the page to the domains search domains params
func (o *DomainsSearchDomainsParams) WithPage(page *string) *DomainsSearchDomainsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the domains search domains params
func (o *DomainsSearchDomainsParams) SetPage(page *string) {
	o.Page = page
}

// WithSize adds the size to the domains search domains params
func (o *DomainsSearchDomainsParams) WithSize(size *string) *DomainsSearchDomainsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the domains search domains params
func (o *DomainsSearchDomainsParams) SetSize(size *string) {
	o.Size = size
}

// WithSort adds the sort to the domains search domains params
func (o *DomainsSearchDomainsParams) WithSort(sort []string) *DomainsSearchDomainsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the domains search domains params
func (o *DomainsSearchDomainsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *DomainsSearchDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Domain != nil {

		// query param domain
		var qrDomain string

		if o.Domain != nil {
			qrDomain = *o.Domain
		}
		qDomain := qrDomain
		if qDomain != "" {

			if err := r.SetQueryParam("domain", qDomain); err != nil {
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

	if o.Page != nil {

		// query param page
		var qrPage string

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize string

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := qrSize
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// binding items for sort
		joinedSort := o.bindParamSort(reg)

		// query array param sort
		if err := r.SetQueryParam("sort", joinedSort...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDomainsSearchDomains binds the parameter fields
func (o *DomainsSearchDomainsParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamDomainsSearchDomains binds the parameter sort
func (o *DomainsSearchDomainsParams) bindParamSort(formats strfmt.Registry) []string {
	sortIR := o.Sort

	var sortIC []string
	for _, sortIIR := range sortIR { // explode []string

		sortIIV := sortIIR // string as string
		sortIC = append(sortIC, sortIIV)
	}

	// items.CollectionFormat: "multi"
	sortIS := swag.JoinByFormat(sortIC, "multi")

	return sortIS
}
