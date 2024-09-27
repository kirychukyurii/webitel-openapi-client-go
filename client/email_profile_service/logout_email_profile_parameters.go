// Code generated by go-swagger; DO NOT EDIT.

package email_profile_service

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

// NewLogoutEmailProfileParams creates a new LogoutEmailProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogoutEmailProfileParams() *LogoutEmailProfileParams {
	return &LogoutEmailProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogoutEmailProfileParamsWithTimeout creates a new LogoutEmailProfileParams object
// with the ability to set a timeout on a request.
func NewLogoutEmailProfileParamsWithTimeout(timeout time.Duration) *LogoutEmailProfileParams {
	return &LogoutEmailProfileParams{
		timeout: timeout,
	}
}

// NewLogoutEmailProfileParamsWithContext creates a new LogoutEmailProfileParams object
// with the ability to set a context for a request.
func NewLogoutEmailProfileParamsWithContext(ctx context.Context) *LogoutEmailProfileParams {
	return &LogoutEmailProfileParams{
		Context: ctx,
	}
}

// NewLogoutEmailProfileParamsWithHTTPClient creates a new LogoutEmailProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogoutEmailProfileParamsWithHTTPClient(client *http.Client) *LogoutEmailProfileParams {
	return &LogoutEmailProfileParams{
		HTTPClient: client,
	}
}

/*
LogoutEmailProfileParams contains all the parameters to send to the API endpoint

	for the logout email profile operation.

	Typically these are written to a http.Request.
*/
type LogoutEmailProfileParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the logout email profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogoutEmailProfileParams) WithDefaults() *LogoutEmailProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the logout email profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogoutEmailProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the logout email profile params
func (o *LogoutEmailProfileParams) WithTimeout(timeout time.Duration) *LogoutEmailProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logout email profile params
func (o *LogoutEmailProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logout email profile params
func (o *LogoutEmailProfileParams) WithContext(ctx context.Context) *LogoutEmailProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logout email profile params
func (o *LogoutEmailProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logout email profile params
func (o *LogoutEmailProfileParams) WithHTTPClient(client *http.Client) *LogoutEmailProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logout email profile params
func (o *LogoutEmailProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the logout email profile params
func (o *LogoutEmailProfileParams) WithID(id int32) *LogoutEmailProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the logout email profile params
func (o *LogoutEmailProfileParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LogoutEmailProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
