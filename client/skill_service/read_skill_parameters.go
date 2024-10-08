// Code generated by go-swagger; DO NOT EDIT.

package skill_service

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

// NewReadSkillParams creates a new ReadSkillParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadSkillParams() *ReadSkillParams {
	return &ReadSkillParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadSkillParamsWithTimeout creates a new ReadSkillParams object
// with the ability to set a timeout on a request.
func NewReadSkillParamsWithTimeout(timeout time.Duration) *ReadSkillParams {
	return &ReadSkillParams{
		timeout: timeout,
	}
}

// NewReadSkillParamsWithContext creates a new ReadSkillParams object
// with the ability to set a context for a request.
func NewReadSkillParamsWithContext(ctx context.Context) *ReadSkillParams {
	return &ReadSkillParams{
		Context: ctx,
	}
}

// NewReadSkillParamsWithHTTPClient creates a new ReadSkillParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadSkillParamsWithHTTPClient(client *http.Client) *ReadSkillParams {
	return &ReadSkillParams{
		HTTPClient: client,
	}
}

/*
ReadSkillParams contains all the parameters to send to the API endpoint

	for the read skill operation.

	Typically these are written to a http.Request.
*/
type ReadSkillParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadSkillParams) WithDefaults() *ReadSkillParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadSkillParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read skill params
func (o *ReadSkillParams) WithTimeout(timeout time.Duration) *ReadSkillParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read skill params
func (o *ReadSkillParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read skill params
func (o *ReadSkillParams) WithContext(ctx context.Context) *ReadSkillParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read skill params
func (o *ReadSkillParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read skill params
func (o *ReadSkillParams) WithHTTPClient(client *http.Client) *ReadSkillParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read skill params
func (o *ReadSkillParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the read skill params
func (o *ReadSkillParams) WithDomainID(domainID *string) *ReadSkillParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the read skill params
func (o *ReadSkillParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the read skill params
func (o *ReadSkillParams) WithID(id string) *ReadSkillParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read skill params
func (o *ReadSkillParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadSkillParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domain_id
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domain_id", qDomainID); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
