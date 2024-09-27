// Code generated by go-swagger; DO NOT EDIT.

package users

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

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// NewUsersLogoutUserParams creates a new UsersLogoutUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersLogoutUserParams() *UsersLogoutUserParams {
	return &UsersLogoutUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersLogoutUserParamsWithTimeout creates a new UsersLogoutUserParams object
// with the ability to set a timeout on a request.
func NewUsersLogoutUserParamsWithTimeout(timeout time.Duration) *UsersLogoutUserParams {
	return &UsersLogoutUserParams{
		timeout: timeout,
	}
}

// NewUsersLogoutUserParamsWithContext creates a new UsersLogoutUserParams object
// with the ability to set a context for a request.
func NewUsersLogoutUserParamsWithContext(ctx context.Context) *UsersLogoutUserParams {
	return &UsersLogoutUserParams{
		Context: ctx,
	}
}

// NewUsersLogoutUserParamsWithHTTPClient creates a new UsersLogoutUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersLogoutUserParamsWithHTTPClient(client *http.Client) *UsersLogoutUserParams {
	return &UsersLogoutUserParams{
		HTTPClient: client,
	}
}

/*
UsersLogoutUserParams contains all the parameters to send to the API endpoint

	for the users logout user operation.

	Typically these are written to a http.Request.
*/
type UsersLogoutUserParams struct {

	// Body.
	Body models.APIUsersLogoutUserBody

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users logout user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersLogoutUserParams) WithDefaults() *UsersLogoutUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users logout user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersLogoutUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users logout user params
func (o *UsersLogoutUserParams) WithTimeout(timeout time.Duration) *UsersLogoutUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users logout user params
func (o *UsersLogoutUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users logout user params
func (o *UsersLogoutUserParams) WithContext(ctx context.Context) *UsersLogoutUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users logout user params
func (o *UsersLogoutUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users logout user params
func (o *UsersLogoutUserParams) WithHTTPClient(client *http.Client) *UsersLogoutUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users logout user params
func (o *UsersLogoutUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the users logout user params
func (o *UsersLogoutUserParams) WithBody(body models.APIUsersLogoutUserBody) *UsersLogoutUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the users logout user params
func (o *UsersLogoutUserParams) SetBody(body models.APIUsersLogoutUserBody) {
	o.Body = body
}

// WithID adds the id to the users logout user params
func (o *UsersLogoutUserParams) WithID(id string) *UsersLogoutUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users logout user params
func (o *UsersLogoutUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersLogoutUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
