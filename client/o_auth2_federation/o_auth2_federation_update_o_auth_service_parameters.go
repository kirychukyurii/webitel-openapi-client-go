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

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// NewOAuth2FederationUpdateOAuthServiceParams creates a new OAuth2FederationUpdateOAuthServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOAuth2FederationUpdateOAuthServiceParams() *OAuth2FederationUpdateOAuthServiceParams {
	return &OAuth2FederationUpdateOAuthServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOAuth2FederationUpdateOAuthServiceParamsWithTimeout creates a new OAuth2FederationUpdateOAuthServiceParams object
// with the ability to set a timeout on a request.
func NewOAuth2FederationUpdateOAuthServiceParamsWithTimeout(timeout time.Duration) *OAuth2FederationUpdateOAuthServiceParams {
	return &OAuth2FederationUpdateOAuthServiceParams{
		timeout: timeout,
	}
}

// NewOAuth2FederationUpdateOAuthServiceParamsWithContext creates a new OAuth2FederationUpdateOAuthServiceParams object
// with the ability to set a context for a request.
func NewOAuth2FederationUpdateOAuthServiceParamsWithContext(ctx context.Context) *OAuth2FederationUpdateOAuthServiceParams {
	return &OAuth2FederationUpdateOAuthServiceParams{
		Context: ctx,
	}
}

// NewOAuth2FederationUpdateOAuthServiceParamsWithHTTPClient creates a new OAuth2FederationUpdateOAuthServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewOAuth2FederationUpdateOAuthServiceParamsWithHTTPClient(client *http.Client) *OAuth2FederationUpdateOAuthServiceParams {
	return &OAuth2FederationUpdateOAuthServiceParams{
		HTTPClient: client,
	}
}

/*
OAuth2FederationUpdateOAuthServiceParams contains all the parameters to send to the API endpoint

	for the o auth2 federation update o auth service operation.

	Typically these are written to a http.Request.
*/
type OAuth2FederationUpdateOAuthServiceParams struct {

	// Body.
	Body *models.APIOAuth2FederationUpdateOAuthServiceBody

	// ChangesID.
	//
	// Format: int64
	ChangesID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the o auth2 federation update o auth service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2FederationUpdateOAuthServiceParams) WithDefaults() *OAuth2FederationUpdateOAuthServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the o auth2 federation update o auth service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2FederationUpdateOAuthServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) WithTimeout(timeout time.Duration) *OAuth2FederationUpdateOAuthServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) WithContext(ctx context.Context) *OAuth2FederationUpdateOAuthServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) WithHTTPClient(client *http.Client) *OAuth2FederationUpdateOAuthServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) WithBody(body *models.APIOAuth2FederationUpdateOAuthServiceBody) *OAuth2FederationUpdateOAuthServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) SetBody(body *models.APIOAuth2FederationUpdateOAuthServiceBody) {
	o.Body = body
}

// WithChangesID adds the changesID to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) WithChangesID(changesID string) *OAuth2FederationUpdateOAuthServiceParams {
	o.SetChangesID(changesID)
	return o
}

// SetChangesID adds the changesId to the o auth2 federation update o auth service params
func (o *OAuth2FederationUpdateOAuthServiceParams) SetChangesID(changesID string) {
	o.ChangesID = changesID
}

// WriteToRequest writes these params to a swagger request
func (o *OAuth2FederationUpdateOAuthServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param changes.id
	if err := r.SetPathParam("changes.id", o.ChangesID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
