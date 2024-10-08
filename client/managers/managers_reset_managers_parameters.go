// Code generated by go-swagger; DO NOT EDIT.

package managers

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

// NewManagersResetManagersParams creates a new ManagersResetManagersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManagersResetManagersParams() *ManagersResetManagersParams {
	return &ManagersResetManagersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManagersResetManagersParamsWithTimeout creates a new ManagersResetManagersParams object
// with the ability to set a timeout on a request.
func NewManagersResetManagersParamsWithTimeout(timeout time.Duration) *ManagersResetManagersParams {
	return &ManagersResetManagersParams{
		timeout: timeout,
	}
}

// NewManagersResetManagersParamsWithContext creates a new ManagersResetManagersParams object
// with the ability to set a context for a request.
func NewManagersResetManagersParamsWithContext(ctx context.Context) *ManagersResetManagersParams {
	return &ManagersResetManagersParams{
		Context: ctx,
	}
}

// NewManagersResetManagersParamsWithHTTPClient creates a new ManagersResetManagersParams object
// with the ability to set a custom HTTPClient for a request.
func NewManagersResetManagersParamsWithHTTPClient(client *http.Client) *ManagersResetManagersParams {
	return &ManagersResetManagersParams{
		HTTPClient: client,
	}
}

/*
ManagersResetManagersParams contains all the parameters to send to the API endpoint

	for the managers reset managers operation.

	Typically these are written to a http.Request.
*/
type ManagersResetManagersParams struct {

	/* ContactID.

	   Contact ID associated with.
	*/
	ContactID string

	/* Fields.

	   Fields to be retrieved as a result.
	*/
	Fields []string

	/* Input.

	     Final set of unique User(s) to be linked with the Contact.
	User(s) that are already linked with the Contact
	but not listed here will be removed.
	The first element will become `primary` if no other specified.
	*/
	Input []*models.WebitelContactsInputManager

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the managers reset managers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagersResetManagersParams) WithDefaults() *ManagersResetManagersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the managers reset managers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagersResetManagersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the managers reset managers params
func (o *ManagersResetManagersParams) WithTimeout(timeout time.Duration) *ManagersResetManagersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the managers reset managers params
func (o *ManagersResetManagersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the managers reset managers params
func (o *ManagersResetManagersParams) WithContext(ctx context.Context) *ManagersResetManagersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the managers reset managers params
func (o *ManagersResetManagersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the managers reset managers params
func (o *ManagersResetManagersParams) WithHTTPClient(client *http.Client) *ManagersResetManagersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the managers reset managers params
func (o *ManagersResetManagersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the managers reset managers params
func (o *ManagersResetManagersParams) WithContactID(contactID string) *ManagersResetManagersParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the managers reset managers params
func (o *ManagersResetManagersParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithFields adds the fields to the managers reset managers params
func (o *ManagersResetManagersParams) WithFields(fields []string) *ManagersResetManagersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the managers reset managers params
func (o *ManagersResetManagersParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInput adds the input to the managers reset managers params
func (o *ManagersResetManagersParams) WithInput(input []*models.WebitelContactsInputManager) *ManagersResetManagersParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the managers reset managers params
func (o *ManagersResetManagersParams) SetInput(input []*models.WebitelContactsInputManager) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ManagersResetManagersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
		return err
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamManagersResetManagers binds the parameter fields
func (o *ManagersResetManagersParams) bindParamFields(formats strfmt.Registry) []string {
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
