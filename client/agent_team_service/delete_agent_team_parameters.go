// Code generated by go-swagger; DO NOT EDIT.

package agent_team_service

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

// NewDeleteAgentTeamParams creates a new DeleteAgentTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAgentTeamParams() *DeleteAgentTeamParams {
	return &DeleteAgentTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAgentTeamParamsWithTimeout creates a new DeleteAgentTeamParams object
// with the ability to set a timeout on a request.
func NewDeleteAgentTeamParamsWithTimeout(timeout time.Duration) *DeleteAgentTeamParams {
	return &DeleteAgentTeamParams{
		timeout: timeout,
	}
}

// NewDeleteAgentTeamParamsWithContext creates a new DeleteAgentTeamParams object
// with the ability to set a context for a request.
func NewDeleteAgentTeamParamsWithContext(ctx context.Context) *DeleteAgentTeamParams {
	return &DeleteAgentTeamParams{
		Context: ctx,
	}
}

// NewDeleteAgentTeamParamsWithHTTPClient creates a new DeleteAgentTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAgentTeamParamsWithHTTPClient(client *http.Client) *DeleteAgentTeamParams {
	return &DeleteAgentTeamParams{
		HTTPClient: client,
	}
}

/*
DeleteAgentTeamParams contains all the parameters to send to the API endpoint

	for the delete agent team operation.

	Typically these are written to a http.Request.
*/
type DeleteAgentTeamParams struct {

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

// WithDefaults hydrates default values in the delete agent team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentTeamParams) WithDefaults() *DeleteAgentTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete agent team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete agent team params
func (o *DeleteAgentTeamParams) WithTimeout(timeout time.Duration) *DeleteAgentTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete agent team params
func (o *DeleteAgentTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete agent team params
func (o *DeleteAgentTeamParams) WithContext(ctx context.Context) *DeleteAgentTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete agent team params
func (o *DeleteAgentTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete agent team params
func (o *DeleteAgentTeamParams) WithHTTPClient(client *http.Client) *DeleteAgentTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete agent team params
func (o *DeleteAgentTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete agent team params
func (o *DeleteAgentTeamParams) WithDomainID(domainID *string) *DeleteAgentTeamParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete agent team params
func (o *DeleteAgentTeamParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the delete agent team params
func (o *DeleteAgentTeamParams) WithID(id string) *DeleteAgentTeamParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete agent team params
func (o *DeleteAgentTeamParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAgentTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
