// Code generated by go-swagger; DO NOT EDIT.

package routing_variable_service

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
)

// NewDeleteRoutingVariableParams creates a new DeleteRoutingVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRoutingVariableParams() *DeleteRoutingVariableParams {
	return &DeleteRoutingVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingVariableParamsWithTimeout creates a new DeleteRoutingVariableParams object
// with the ability to set a timeout on a request.
func NewDeleteRoutingVariableParamsWithTimeout(timeout time.Duration) *DeleteRoutingVariableParams {
	return &DeleteRoutingVariableParams{
		timeout: timeout,
	}
}

// NewDeleteRoutingVariableParamsWithContext creates a new DeleteRoutingVariableParams object
// with the ability to set a context for a request.
func NewDeleteRoutingVariableParamsWithContext(ctx context.Context) *DeleteRoutingVariableParams {
	return &DeleteRoutingVariableParams{
		Context: ctx,
	}
}

// NewDeleteRoutingVariableParamsWithHTTPClient creates a new DeleteRoutingVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRoutingVariableParamsWithHTTPClient(client *http.Client) *DeleteRoutingVariableParams {
	return &DeleteRoutingVariableParams{
		HTTPClient: client,
	}
}

/*
DeleteRoutingVariableParams contains all the parameters to send to the API endpoint

	for the delete routing variable operation.

	Typically these are written to a http.Request.
*/
type DeleteRoutingVariableParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete routing variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoutingVariableParams) WithDefaults() *DeleteRoutingVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete routing variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoutingVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete routing variable params
func (o *DeleteRoutingVariableParams) WithTimeout(timeout time.Duration) *DeleteRoutingVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing variable params
func (o *DeleteRoutingVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing variable params
func (o *DeleteRoutingVariableParams) WithContext(ctx context.Context) *DeleteRoutingVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing variable params
func (o *DeleteRoutingVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing variable params
func (o *DeleteRoutingVariableParams) WithHTTPClient(client *http.Client) *DeleteRoutingVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing variable params
func (o *DeleteRoutingVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete routing variable params
func (o *DeleteRoutingVariableParams) WithDomainID(domainID *string) *DeleteRoutingVariableParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete routing variable params
func (o *DeleteRoutingVariableParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the delete routing variable params
func (o *DeleteRoutingVariableParams) WithID(id string) *DeleteRoutingVariableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete routing variable params
func (o *DeleteRoutingVariableParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domain_id
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domain_id", qDomainID); err != nil {
				return err
			}
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
