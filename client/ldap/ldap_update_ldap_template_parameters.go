// Code generated by go-swagger; DO NOT EDIT.

package ldap

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

// NewLDAPUpdateLDAPTemplateParams creates a new LDAPUpdateLDAPTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLDAPUpdateLDAPTemplateParams() *LDAPUpdateLDAPTemplateParams {
	return &LDAPUpdateLDAPTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLDAPUpdateLDAPTemplateParamsWithTimeout creates a new LDAPUpdateLDAPTemplateParams object
// with the ability to set a timeout on a request.
func NewLDAPUpdateLDAPTemplateParamsWithTimeout(timeout time.Duration) *LDAPUpdateLDAPTemplateParams {
	return &LDAPUpdateLDAPTemplateParams{
		timeout: timeout,
	}
}

// NewLDAPUpdateLDAPTemplateParamsWithContext creates a new LDAPUpdateLDAPTemplateParams object
// with the ability to set a context for a request.
func NewLDAPUpdateLDAPTemplateParamsWithContext(ctx context.Context) *LDAPUpdateLDAPTemplateParams {
	return &LDAPUpdateLDAPTemplateParams{
		Context: ctx,
	}
}

// NewLDAPUpdateLDAPTemplateParamsWithHTTPClient creates a new LDAPUpdateLDAPTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLDAPUpdateLDAPTemplateParamsWithHTTPClient(client *http.Client) *LDAPUpdateLDAPTemplateParams {
	return &LDAPUpdateLDAPTemplateParams{
		HTTPClient: client,
	}
}

/*
LDAPUpdateLDAPTemplateParams contains all the parameters to send to the API endpoint

	for the LDAP update LDAP template operation.

	Typically these are written to a http.Request.
*/
type LDAPUpdateLDAPTemplateParams struct {

	// Body.
	Body *models.APILDAPUpdateLDAPTemplateBody

	/* TemplateCatalogID.

	   identifier

	   Format: int64
	*/
	TemplateCatalogID string

	/* TemplateID.

	   ID unique

	   Format: int64
	*/
	TemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the LDAP update LDAP template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPUpdateLDAPTemplateParams) WithDefaults() *LDAPUpdateLDAPTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the LDAP update LDAP template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPUpdateLDAPTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) WithTimeout(timeout time.Duration) *LDAPUpdateLDAPTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) WithContext(ctx context.Context) *LDAPUpdateLDAPTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) WithHTTPClient(client *http.Client) *LDAPUpdateLDAPTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) WithBody(body *models.APILDAPUpdateLDAPTemplateBody) *LDAPUpdateLDAPTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) SetBody(body *models.APILDAPUpdateLDAPTemplateBody) {
	o.Body = body
}

// WithTemplateCatalogID adds the templateCatalogID to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) WithTemplateCatalogID(templateCatalogID string) *LDAPUpdateLDAPTemplateParams {
	o.SetTemplateCatalogID(templateCatalogID)
	return o
}

// SetTemplateCatalogID adds the templateCatalogId to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) SetTemplateCatalogID(templateCatalogID string) {
	o.TemplateCatalogID = templateCatalogID
}

// WithTemplateID adds the templateID to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) WithTemplateID(templateID string) *LDAPUpdateLDAPTemplateParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the LDAP update LDAP template params
func (o *LDAPUpdateLDAPTemplateParams) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *LDAPUpdateLDAPTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param template.catalog.id
	if err := r.SetPathParam("template.catalog.id", o.TemplateCatalogID); err != nil {
		return err
	}

	// path param template.id
	if err := r.SetPathParam("template.id", o.TemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
