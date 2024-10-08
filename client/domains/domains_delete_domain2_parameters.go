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
)

// NewDomainsDeleteDomain2Params creates a new DomainsDeleteDomain2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainsDeleteDomain2Params() *DomainsDeleteDomain2Params {
	return &DomainsDeleteDomain2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainsDeleteDomain2ParamsWithTimeout creates a new DomainsDeleteDomain2Params object
// with the ability to set a timeout on a request.
func NewDomainsDeleteDomain2ParamsWithTimeout(timeout time.Duration) *DomainsDeleteDomain2Params {
	return &DomainsDeleteDomain2Params{
		timeout: timeout,
	}
}

// NewDomainsDeleteDomain2ParamsWithContext creates a new DomainsDeleteDomain2Params object
// with the ability to set a context for a request.
func NewDomainsDeleteDomain2ParamsWithContext(ctx context.Context) *DomainsDeleteDomain2Params {
	return &DomainsDeleteDomain2Params{
		Context: ctx,
	}
}

// NewDomainsDeleteDomain2ParamsWithHTTPClient creates a new DomainsDeleteDomain2Params object
// with the ability to set a custom HTTPClient for a request.
func NewDomainsDeleteDomain2ParamsWithHTTPClient(client *http.Client) *DomainsDeleteDomain2Params {
	return &DomainsDeleteDomain2Params{
		HTTPClient: client,
	}
}

/*
DomainsDeleteDomain2Params contains all the parameters to send to the API endpoint

	for the domains delete domain2 operation.

	Typically these are written to a http.Request.
*/
type DomainsDeleteDomain2Params struct {

	// Dc.
	//
	// Format: int64
	Dc string

	// Domain.
	Domain *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domains delete domain2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainsDeleteDomain2Params) WithDefaults() *DomainsDeleteDomain2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domains delete domain2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainsDeleteDomain2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) WithTimeout(timeout time.Duration) *DomainsDeleteDomain2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) WithContext(ctx context.Context) *DomainsDeleteDomain2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) WithHTTPClient(client *http.Client) *DomainsDeleteDomain2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDc adds the dc to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) WithDc(dc string) *DomainsDeleteDomain2Params {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) SetDc(dc string) {
	o.Dc = dc
}

// WithDomain adds the domain to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) WithDomain(domain *string) *DomainsDeleteDomain2Params {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the domains delete domain2 params
func (o *DomainsDeleteDomain2Params) SetDomain(domain *string) {
	o.Domain = domain
}

// WriteToRequest writes these params to a swagger request
func (o *DomainsDeleteDomain2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dc
	if err := r.SetPathParam("dc", o.Dc); err != nil {
		return err
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
