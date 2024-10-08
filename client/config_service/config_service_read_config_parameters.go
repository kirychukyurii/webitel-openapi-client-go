// Code generated by go-swagger; DO NOT EDIT.

package config_service

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

// NewConfigServiceReadConfigParams creates a new ConfigServiceReadConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfigServiceReadConfigParams() *ConfigServiceReadConfigParams {
	return &ConfigServiceReadConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfigServiceReadConfigParamsWithTimeout creates a new ConfigServiceReadConfigParams object
// with the ability to set a timeout on a request.
func NewConfigServiceReadConfigParamsWithTimeout(timeout time.Duration) *ConfigServiceReadConfigParams {
	return &ConfigServiceReadConfigParams{
		timeout: timeout,
	}
}

// NewConfigServiceReadConfigParamsWithContext creates a new ConfigServiceReadConfigParams object
// with the ability to set a context for a request.
func NewConfigServiceReadConfigParamsWithContext(ctx context.Context) *ConfigServiceReadConfigParams {
	return &ConfigServiceReadConfigParams{
		Context: ctx,
	}
}

// NewConfigServiceReadConfigParamsWithHTTPClient creates a new ConfigServiceReadConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfigServiceReadConfigParamsWithHTTPClient(client *http.Client) *ConfigServiceReadConfigParams {
	return &ConfigServiceReadConfigParams{
		HTTPClient: client,
	}
}

/*
ConfigServiceReadConfigParams contains all the parameters to send to the API endpoint

	for the config service read config operation.

	Typically these are written to a http.Request.
*/
type ConfigServiceReadConfigParams struct {

	// ConfigID.
	//
	// Format: int32
	ConfigID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the config service read config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigServiceReadConfigParams) WithDefaults() *ConfigServiceReadConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the config service read config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigServiceReadConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the config service read config params
func (o *ConfigServiceReadConfigParams) WithTimeout(timeout time.Duration) *ConfigServiceReadConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the config service read config params
func (o *ConfigServiceReadConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the config service read config params
func (o *ConfigServiceReadConfigParams) WithContext(ctx context.Context) *ConfigServiceReadConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the config service read config params
func (o *ConfigServiceReadConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the config service read config params
func (o *ConfigServiceReadConfigParams) WithHTTPClient(client *http.Client) *ConfigServiceReadConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the config service read config params
func (o *ConfigServiceReadConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigID adds the configID to the config service read config params
func (o *ConfigServiceReadConfigParams) WithConfigID(configID int32) *ConfigServiceReadConfigParams {
	o.SetConfigID(configID)
	return o
}

// SetConfigID adds the configId to the config service read config params
func (o *ConfigServiceReadConfigParams) SetConfigID(configID int32) {
	o.ConfigID = configID
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigServiceReadConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param config_id
	if err := r.SetPathParam("config_id", swag.FormatInt32(o.ConfigID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
