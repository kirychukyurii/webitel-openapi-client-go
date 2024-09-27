// Code generated by go-swagger; DO NOT EDIT.

package access_store

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

// NewAccessStoreToggleDefaultAccessParams creates a new AccessStoreToggleDefaultAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccessStoreToggleDefaultAccessParams() *AccessStoreToggleDefaultAccessParams {
	return &AccessStoreToggleDefaultAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccessStoreToggleDefaultAccessParamsWithTimeout creates a new AccessStoreToggleDefaultAccessParams object
// with the ability to set a timeout on a request.
func NewAccessStoreToggleDefaultAccessParamsWithTimeout(timeout time.Duration) *AccessStoreToggleDefaultAccessParams {
	return &AccessStoreToggleDefaultAccessParams{
		timeout: timeout,
	}
}

// NewAccessStoreToggleDefaultAccessParamsWithContext creates a new AccessStoreToggleDefaultAccessParams object
// with the ability to set a context for a request.
func NewAccessStoreToggleDefaultAccessParamsWithContext(ctx context.Context) *AccessStoreToggleDefaultAccessParams {
	return &AccessStoreToggleDefaultAccessParams{
		Context: ctx,
	}
}

// NewAccessStoreToggleDefaultAccessParamsWithHTTPClient creates a new AccessStoreToggleDefaultAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccessStoreToggleDefaultAccessParamsWithHTTPClient(client *http.Client) *AccessStoreToggleDefaultAccessParams {
	return &AccessStoreToggleDefaultAccessParams{
		HTTPClient: client,
	}
}

/*
AccessStoreToggleDefaultAccessParams contains all the parameters to send to the API endpoint

	for the access store toggle default access operation.

	Typically these are written to a http.Request.
*/
type AccessStoreToggleDefaultAccessParams struct {

	// Body.
	Body *models.APIAccessStoreToggleDefaultAccessBody

	/* Grantor.

	   [FOR] creator user/role

	   Format: int64
	*/
	Grantor string

	/* ObjectID.

	   identifier

	   Format: int64
	*/
	ObjectID string

	/* ObjectName.

	   display name
	*/
	ObjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the access store toggle default access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessStoreToggleDefaultAccessParams) WithDefaults() *AccessStoreToggleDefaultAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the access store toggle default access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessStoreToggleDefaultAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) WithTimeout(timeout time.Duration) *AccessStoreToggleDefaultAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) WithContext(ctx context.Context) *AccessStoreToggleDefaultAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) WithHTTPClient(client *http.Client) *AccessStoreToggleDefaultAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) WithBody(body *models.APIAccessStoreToggleDefaultAccessBody) *AccessStoreToggleDefaultAccessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) SetBody(body *models.APIAccessStoreToggleDefaultAccessBody) {
	o.Body = body
}

// WithGrantor adds the grantor to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) WithGrantor(grantor string) *AccessStoreToggleDefaultAccessParams {
	o.SetGrantor(grantor)
	return o
}

// SetGrantor adds the grantor to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) SetGrantor(grantor string) {
	o.Grantor = grantor
}

// WithObjectID adds the objectID to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) WithObjectID(objectID string) *AccessStoreToggleDefaultAccessParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) SetObjectID(objectID string) {
	o.ObjectID = objectID
}

// WithObjectName adds the objectName to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) WithObjectName(objectName string) *AccessStoreToggleDefaultAccessParams {
	o.SetObjectName(objectName)
	return o
}

// SetObjectName adds the objectName to the access store toggle default access params
func (o *AccessStoreToggleDefaultAccessParams) SetObjectName(objectName string) {
	o.ObjectName = objectName
}

// WriteToRequest writes these params to a swagger request
func (o *AccessStoreToggleDefaultAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param grantor
	if err := r.SetPathParam("grantor", o.Grantor); err != nil {
		return err
	}

	// path param object.id
	if err := r.SetPathParam("object.id", o.ObjectID); err != nil {
		return err
	}

	// path param object.name
	if err := r.SetPathParam("object.name", o.ObjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
