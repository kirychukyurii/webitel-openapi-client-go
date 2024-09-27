// Code generated by go-swagger; DO NOT EDIT.

package forecast_calculation_service

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

// NewForecastCalculationServiceCreateForecastCalculationParams creates a new ForecastCalculationServiceCreateForecastCalculationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForecastCalculationServiceCreateForecastCalculationParams() *ForecastCalculationServiceCreateForecastCalculationParams {
	return &ForecastCalculationServiceCreateForecastCalculationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForecastCalculationServiceCreateForecastCalculationParamsWithTimeout creates a new ForecastCalculationServiceCreateForecastCalculationParams object
// with the ability to set a timeout on a request.
func NewForecastCalculationServiceCreateForecastCalculationParamsWithTimeout(timeout time.Duration) *ForecastCalculationServiceCreateForecastCalculationParams {
	return &ForecastCalculationServiceCreateForecastCalculationParams{
		timeout: timeout,
	}
}

// NewForecastCalculationServiceCreateForecastCalculationParamsWithContext creates a new ForecastCalculationServiceCreateForecastCalculationParams object
// with the ability to set a context for a request.
func NewForecastCalculationServiceCreateForecastCalculationParamsWithContext(ctx context.Context) *ForecastCalculationServiceCreateForecastCalculationParams {
	return &ForecastCalculationServiceCreateForecastCalculationParams{
		Context: ctx,
	}
}

// NewForecastCalculationServiceCreateForecastCalculationParamsWithHTTPClient creates a new ForecastCalculationServiceCreateForecastCalculationParams object
// with the ability to set a custom HTTPClient for a request.
func NewForecastCalculationServiceCreateForecastCalculationParamsWithHTTPClient(client *http.Client) *ForecastCalculationServiceCreateForecastCalculationParams {
	return &ForecastCalculationServiceCreateForecastCalculationParams{
		HTTPClient: client,
	}
}

/*
ForecastCalculationServiceCreateForecastCalculationParams contains all the parameters to send to the API endpoint

	for the forecast calculation service create forecast calculation operation.

	Typically these are written to a http.Request.
*/
type ForecastCalculationServiceCreateForecastCalculationParams struct {

	// Body.
	Body *models.WfmCreateForecastCalculationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the forecast calculation service create forecast calculation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForecastCalculationServiceCreateForecastCalculationParams) WithDefaults() *ForecastCalculationServiceCreateForecastCalculationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the forecast calculation service create forecast calculation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForecastCalculationServiceCreateForecastCalculationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the forecast calculation service create forecast calculation params
func (o *ForecastCalculationServiceCreateForecastCalculationParams) WithTimeout(timeout time.Duration) *ForecastCalculationServiceCreateForecastCalculationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forecast calculation service create forecast calculation params
func (o *ForecastCalculationServiceCreateForecastCalculationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forecast calculation service create forecast calculation params
func (o *ForecastCalculationServiceCreateForecastCalculationParams) WithContext(ctx context.Context) *ForecastCalculationServiceCreateForecastCalculationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forecast calculation service create forecast calculation params
func (o *ForecastCalculationServiceCreateForecastCalculationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forecast calculation service create forecast calculation params
func (o *ForecastCalculationServiceCreateForecastCalculationParams) WithHTTPClient(client *http.Client) *ForecastCalculationServiceCreateForecastCalculationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forecast calculation service create forecast calculation params
func (o *ForecastCalculationServiceCreateForecastCalculationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the forecast calculation service create forecast calculation params
func (o *ForecastCalculationServiceCreateForecastCalculationParams) WithBody(body *models.WfmCreateForecastCalculationRequest) *ForecastCalculationServiceCreateForecastCalculationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the forecast calculation service create forecast calculation params
func (o *ForecastCalculationServiceCreateForecastCalculationParams) SetBody(body *models.WfmCreateForecastCalculationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ForecastCalculationServiceCreateForecastCalculationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
