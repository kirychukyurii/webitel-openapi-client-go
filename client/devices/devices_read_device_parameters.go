// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewDevicesReadDeviceParams creates a new DevicesReadDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDevicesReadDeviceParams() *DevicesReadDeviceParams {
	return &DevicesReadDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesReadDeviceParamsWithTimeout creates a new DevicesReadDeviceParams object
// with the ability to set a timeout on a request.
func NewDevicesReadDeviceParamsWithTimeout(timeout time.Duration) *DevicesReadDeviceParams {
	return &DevicesReadDeviceParams{
		timeout: timeout,
	}
}

// NewDevicesReadDeviceParamsWithContext creates a new DevicesReadDeviceParams object
// with the ability to set a context for a request.
func NewDevicesReadDeviceParamsWithContext(ctx context.Context) *DevicesReadDeviceParams {
	return &DevicesReadDeviceParams{
		Context: ctx,
	}
}

// NewDevicesReadDeviceParamsWithHTTPClient creates a new DevicesReadDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDevicesReadDeviceParamsWithHTTPClient(client *http.Client) *DevicesReadDeviceParams {
	return &DevicesReadDeviceParams{
		HTTPClient: client,
	}
}

/*
DevicesReadDeviceParams contains all the parameters to send to the API endpoint

	for the devices read device operation.

	Typically these are written to a http.Request.
*/
type DevicesReadDeviceParams struct {

	/* Fields.

	   output selection
	*/
	Fields []string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the devices read device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DevicesReadDeviceParams) WithDefaults() *DevicesReadDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the devices read device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DevicesReadDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the devices read device params
func (o *DevicesReadDeviceParams) WithTimeout(timeout time.Duration) *DevicesReadDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices read device params
func (o *DevicesReadDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices read device params
func (o *DevicesReadDeviceParams) WithContext(ctx context.Context) *DevicesReadDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices read device params
func (o *DevicesReadDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices read device params
func (o *DevicesReadDeviceParams) WithHTTPClient(client *http.Client) *DevicesReadDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices read device params
func (o *DevicesReadDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the devices read device params
func (o *DevicesReadDeviceParams) WithFields(fields []string) *DevicesReadDeviceParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the devices read device params
func (o *DevicesReadDeviceParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the devices read device params
func (o *DevicesReadDeviceParams) WithID(id string) *DevicesReadDeviceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the devices read device params
func (o *DevicesReadDeviceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesReadDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
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

// bindParamDevicesReadDevice binds the parameter fields
func (o *DevicesReadDeviceParams) bindParamFields(formats strfmt.Registry) []string {
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
