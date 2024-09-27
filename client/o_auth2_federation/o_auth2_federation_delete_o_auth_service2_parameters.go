// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_federation

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

// NewOAuth2FederationDeleteOAuthService2Params creates a new OAuth2FederationDeleteOAuthService2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOAuth2FederationDeleteOAuthService2Params() *OAuth2FederationDeleteOAuthService2Params {
	return &OAuth2FederationDeleteOAuthService2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewOAuth2FederationDeleteOAuthService2ParamsWithTimeout creates a new OAuth2FederationDeleteOAuthService2Params object
// with the ability to set a timeout on a request.
func NewOAuth2FederationDeleteOAuthService2ParamsWithTimeout(timeout time.Duration) *OAuth2FederationDeleteOAuthService2Params {
	return &OAuth2FederationDeleteOAuthService2Params{
		timeout: timeout,
	}
}

// NewOAuth2FederationDeleteOAuthService2ParamsWithContext creates a new OAuth2FederationDeleteOAuthService2Params object
// with the ability to set a context for a request.
func NewOAuth2FederationDeleteOAuthService2ParamsWithContext(ctx context.Context) *OAuth2FederationDeleteOAuthService2Params {
	return &OAuth2FederationDeleteOAuthService2Params{
		Context: ctx,
	}
}

// NewOAuth2FederationDeleteOAuthService2ParamsWithHTTPClient creates a new OAuth2FederationDeleteOAuthService2Params object
// with the ability to set a custom HTTPClient for a request.
func NewOAuth2FederationDeleteOAuthService2ParamsWithHTTPClient(client *http.Client) *OAuth2FederationDeleteOAuthService2Params {
	return &OAuth2FederationDeleteOAuthService2Params{
		HTTPClient: client,
	}
}

/*
OAuth2FederationDeleteOAuthService2Params contains all the parameters to send to the API endpoint

	for the o auth2 federation delete o auth service2 operation.

	Typically these are written to a http.Request.
*/
type OAuth2FederationDeleteOAuthService2Params struct {

	// Body.
	Body *models.APIOAuth2FederationDeleteOAuthServiceBody

	// ID.
	ID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the o auth2 federation delete o auth service2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2FederationDeleteOAuthService2Params) WithDefaults() *OAuth2FederationDeleteOAuthService2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the o auth2 federation delete o auth service2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2FederationDeleteOAuthService2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) WithTimeout(timeout time.Duration) *OAuth2FederationDeleteOAuthService2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) WithContext(ctx context.Context) *OAuth2FederationDeleteOAuthService2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) WithHTTPClient(client *http.Client) *OAuth2FederationDeleteOAuthService2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) WithBody(body *models.APIOAuth2FederationDeleteOAuthServiceBody) *OAuth2FederationDeleteOAuthService2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) SetBody(body *models.APIOAuth2FederationDeleteOAuthServiceBody) {
	o.Body = body
}

// WithID adds the id to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) WithID(id []string) *OAuth2FederationDeleteOAuthService2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the o auth2 federation delete o auth service2 params
func (o *OAuth2FederationDeleteOAuthService2Params) SetID(id []string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OAuth2FederationDeleteOAuthService2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// path array param id
		// SetPathParam does not support variadic arguments, since we used JoinByFormat
		// we can send the first item in the array as it's all the items of the previous
		// array joined together
		if len(joinedID) > 0 {
			if err := r.SetPathParam("id", joinedID[0]); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamOAuth2FederationDeleteOAuthService2 binds the parameter id
func (o *OAuth2FederationDeleteOAuthService2Params) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "csv"
	iDIS := swag.JoinByFormat(iDIC, "csv")

	return iDIS
}
