// Code generated by go-swagger; DO NOT EDIT.

package i_m_clients

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

// NewIMClientsDeleteIMClientParams creates a new IMClientsDeleteIMClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIMClientsDeleteIMClientParams() *IMClientsDeleteIMClientParams {
	return &IMClientsDeleteIMClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIMClientsDeleteIMClientParamsWithTimeout creates a new IMClientsDeleteIMClientParams object
// with the ability to set a timeout on a request.
func NewIMClientsDeleteIMClientParamsWithTimeout(timeout time.Duration) *IMClientsDeleteIMClientParams {
	return &IMClientsDeleteIMClientParams{
		timeout: timeout,
	}
}

// NewIMClientsDeleteIMClientParamsWithContext creates a new IMClientsDeleteIMClientParams object
// with the ability to set a context for a request.
func NewIMClientsDeleteIMClientParamsWithContext(ctx context.Context) *IMClientsDeleteIMClientParams {
	return &IMClientsDeleteIMClientParams{
		Context: ctx,
	}
}

// NewIMClientsDeleteIMClientParamsWithHTTPClient creates a new IMClientsDeleteIMClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewIMClientsDeleteIMClientParamsWithHTTPClient(client *http.Client) *IMClientsDeleteIMClientParams {
	return &IMClientsDeleteIMClientParams{
		HTTPClient: client,
	}
}

/*
IMClientsDeleteIMClientParams contains all the parameters to send to the API endpoint

	for the i m clients delete i m client operation.

	Typically these are written to a http.Request.
*/
type IMClientsDeleteIMClientParams struct {

	// ContactID.
	ContactID string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the i m clients delete i m client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IMClientsDeleteIMClientParams) WithDefaults() *IMClientsDeleteIMClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the i m clients delete i m client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IMClientsDeleteIMClientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) WithTimeout(timeout time.Duration) *IMClientsDeleteIMClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) WithContext(ctx context.Context) *IMClientsDeleteIMClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) WithHTTPClient(client *http.Client) *IMClientsDeleteIMClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) WithContactID(contactID string) *IMClientsDeleteIMClientParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithID adds the id to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) WithID(id string) *IMClientsDeleteIMClientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the i m clients delete i m client params
func (o *IMClientsDeleteIMClientParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IMClientsDeleteIMClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
		return err
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
