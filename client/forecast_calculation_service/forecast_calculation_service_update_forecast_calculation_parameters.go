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

// NewForecastCalculationServiceUpdateForecastCalculationParams creates a new ForecastCalculationServiceUpdateForecastCalculationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForecastCalculationServiceUpdateForecastCalculationParams() *ForecastCalculationServiceUpdateForecastCalculationParams {
	return &ForecastCalculationServiceUpdateForecastCalculationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForecastCalculationServiceUpdateForecastCalculationParamsWithTimeout creates a new ForecastCalculationServiceUpdateForecastCalculationParams object
// with the ability to set a timeout on a request.
func NewForecastCalculationServiceUpdateForecastCalculationParamsWithTimeout(timeout time.Duration) *ForecastCalculationServiceUpdateForecastCalculationParams {
	return &ForecastCalculationServiceUpdateForecastCalculationParams{
		timeout: timeout,
	}
}

// NewForecastCalculationServiceUpdateForecastCalculationParamsWithContext creates a new ForecastCalculationServiceUpdateForecastCalculationParams object
// with the ability to set a context for a request.
func NewForecastCalculationServiceUpdateForecastCalculationParamsWithContext(ctx context.Context) *ForecastCalculationServiceUpdateForecastCalculationParams {
	return &ForecastCalculationServiceUpdateForecastCalculationParams{
		Context: ctx,
	}
}

// NewForecastCalculationServiceUpdateForecastCalculationParamsWithHTTPClient creates a new ForecastCalculationServiceUpdateForecastCalculationParams object
// with the ability to set a custom HTTPClient for a request.
func NewForecastCalculationServiceUpdateForecastCalculationParamsWithHTTPClient(client *http.Client) *ForecastCalculationServiceUpdateForecastCalculationParams {
	return &ForecastCalculationServiceUpdateForecastCalculationParams{
		HTTPClient: client,
	}
}

/*
ForecastCalculationServiceUpdateForecastCalculationParams contains all the parameters to send to the API endpoint

	for the forecast calculation service update forecast calculation operation.

	Typically these are written to a http.Request.
*/
type ForecastCalculationServiceUpdateForecastCalculationParams struct {

	// Body.
	Body *models.ForecastCalculationServiceUpdateForecastCalculationParamsBody

	// ItemID.
	//
	// Format: int64
	ItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the forecast calculation service update forecast calculation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) WithDefaults() *ForecastCalculationServiceUpdateForecastCalculationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the forecast calculation service update forecast calculation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) WithTimeout(timeout time.Duration) *ForecastCalculationServiceUpdateForecastCalculationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) WithContext(ctx context.Context) *ForecastCalculationServiceUpdateForecastCalculationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) WithHTTPClient(client *http.Client) *ForecastCalculationServiceUpdateForecastCalculationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) WithBody(body *models.ForecastCalculationServiceUpdateForecastCalculationParamsBody) *ForecastCalculationServiceUpdateForecastCalculationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) SetBody(body *models.ForecastCalculationServiceUpdateForecastCalculationParamsBody) {
	o.Body = body
}

// WithItemID adds the itemID to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) WithItemID(itemID string) *ForecastCalculationServiceUpdateForecastCalculationParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the forecast calculation service update forecast calculation params
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *ForecastCalculationServiceUpdateForecastCalculationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param item.id
	if err := r.SetPathParam("item.id", o.ItemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
