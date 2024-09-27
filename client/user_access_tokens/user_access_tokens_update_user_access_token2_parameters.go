// Code generated by go-swagger; DO NOT EDIT.

package user_access_tokens

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

// NewUserAccessTokensUpdateUserAccessToken2Params creates a new UserAccessTokensUpdateUserAccessToken2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserAccessTokensUpdateUserAccessToken2Params() *UserAccessTokensUpdateUserAccessToken2Params {
	return &UserAccessTokensUpdateUserAccessToken2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserAccessTokensUpdateUserAccessToken2ParamsWithTimeout creates a new UserAccessTokensUpdateUserAccessToken2Params object
// with the ability to set a timeout on a request.
func NewUserAccessTokensUpdateUserAccessToken2ParamsWithTimeout(timeout time.Duration) *UserAccessTokensUpdateUserAccessToken2Params {
	return &UserAccessTokensUpdateUserAccessToken2Params{
		timeout: timeout,
	}
}

// NewUserAccessTokensUpdateUserAccessToken2ParamsWithContext creates a new UserAccessTokensUpdateUserAccessToken2Params object
// with the ability to set a context for a request.
func NewUserAccessTokensUpdateUserAccessToken2ParamsWithContext(ctx context.Context) *UserAccessTokensUpdateUserAccessToken2Params {
	return &UserAccessTokensUpdateUserAccessToken2Params{
		Context: ctx,
	}
}

// NewUserAccessTokensUpdateUserAccessToken2ParamsWithHTTPClient creates a new UserAccessTokensUpdateUserAccessToken2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUserAccessTokensUpdateUserAccessToken2ParamsWithHTTPClient(client *http.Client) *UserAccessTokensUpdateUserAccessToken2Params {
	return &UserAccessTokensUpdateUserAccessToken2Params{
		HTTPClient: client,
	}
}

/*
UserAccessTokensUpdateUserAccessToken2Params contains all the parameters to send to the API endpoint

	for the user access tokens update user access token2 operation.

	Typically these are written to a http.Request.
*/
type UserAccessTokensUpdateUserAccessToken2Params struct {

	// Enable.
	Enable *bool

	/* Fields.

	   [optional] PATCH implementation
	*/
	Fields []string

	/* Update.

	   [required] identification + modifications(about)
	*/
	Update *models.UserAccessTokensUpdateUserAccessToken2ParamsBody

	/* UpdateID.

	   tokenKey::token_key
	*/
	UpdateID string

	/* UpdateUserID.

	   identifier

	   Format: int64
	*/
	UpdateUserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user access tokens update user access token2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithDefaults() *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user access tokens update user access token2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithTimeout(timeout time.Duration) *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithContext(ctx context.Context) *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithHTTPClient(client *http.Client) *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnable adds the enable to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithEnable(enable *bool) *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetEnable(enable)
	return o
}

// SetEnable adds the enable to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetEnable(enable *bool) {
	o.Enable = enable
}

// WithFields adds the fields to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithFields(fields []string) *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetFields(fields []string) {
	o.Fields = fields
}

// WithUpdate adds the update to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithUpdate(update *models.UserAccessTokensUpdateUserAccessToken2ParamsBody) *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetUpdate(update)
	return o
}

// SetUpdate adds the update to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetUpdate(update *models.UserAccessTokensUpdateUserAccessToken2ParamsBody) {
	o.Update = update
}

// WithUpdateID adds the updateID to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithUpdateID(updateID string) *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetUpdateID(updateID)
	return o
}

// SetUpdateID adds the updateId to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetUpdateID(updateID string) {
	o.UpdateID = updateID
}

// WithUpdateUserID adds the updateUserID to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) WithUpdateUserID(updateUserID string) *UserAccessTokensUpdateUserAccessToken2Params {
	o.SetUpdateUserID(updateUserID)
	return o
}

// SetUpdateUserID adds the updateUserId to the user access tokens update user access token2 params
func (o *UserAccessTokensUpdateUserAccessToken2Params) SetUpdateUserID(updateUserID string) {
	o.UpdateUserID = updateUserID
}

// WriteToRequest writes these params to a swagger request
func (o *UserAccessTokensUpdateUserAccessToken2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enable != nil {

		// query param enable
		var qrEnable bool

		if o.Enable != nil {
			qrEnable = *o.Enable
		}
		qEnable := swag.FormatBool(qrEnable)
		if qEnable != "" {

			if err := r.SetQueryParam("enable", qEnable); err != nil {
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
	if o.Update != nil {
		if err := r.SetBodyParam(o.Update); err != nil {
			return err
		}
	}

	// path param update.id
	if err := r.SetPathParam("update.id", o.UpdateID); err != nil {
		return err
	}

	// path param update.user.id
	if err := r.SetPathParam("update.user.id", o.UpdateUserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUserAccessTokensUpdateUserAccessToken2 binds the parameter fields
func (o *UserAccessTokensUpdateUserAccessToken2Params) bindParamFields(formats strfmt.Registry) []string {
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
