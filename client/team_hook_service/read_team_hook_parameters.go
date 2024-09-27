// Code generated by go-swagger; DO NOT EDIT.

package team_hook_service

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

// NewReadTeamHookParams creates a new ReadTeamHookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadTeamHookParams() *ReadTeamHookParams {
	return &ReadTeamHookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadTeamHookParamsWithTimeout creates a new ReadTeamHookParams object
// with the ability to set a timeout on a request.
func NewReadTeamHookParamsWithTimeout(timeout time.Duration) *ReadTeamHookParams {
	return &ReadTeamHookParams{
		timeout: timeout,
	}
}

// NewReadTeamHookParamsWithContext creates a new ReadTeamHookParams object
// with the ability to set a context for a request.
func NewReadTeamHookParamsWithContext(ctx context.Context) *ReadTeamHookParams {
	return &ReadTeamHookParams{
		Context: ctx,
	}
}

// NewReadTeamHookParamsWithHTTPClient creates a new ReadTeamHookParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadTeamHookParamsWithHTTPClient(client *http.Client) *ReadTeamHookParams {
	return &ReadTeamHookParams{
		HTTPClient: client,
	}
}

/*
ReadTeamHookParams contains all the parameters to send to the API endpoint

	for the read team hook operation.

	Typically these are written to a http.Request.
*/
type ReadTeamHookParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	// TeamID.
	//
	// Format: int64
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read team hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadTeamHookParams) WithDefaults() *ReadTeamHookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read team hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadTeamHookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read team hook params
func (o *ReadTeamHookParams) WithTimeout(timeout time.Duration) *ReadTeamHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read team hook params
func (o *ReadTeamHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read team hook params
func (o *ReadTeamHookParams) WithContext(ctx context.Context) *ReadTeamHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read team hook params
func (o *ReadTeamHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read team hook params
func (o *ReadTeamHookParams) WithHTTPClient(client *http.Client) *ReadTeamHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read team hook params
func (o *ReadTeamHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the read team hook params
func (o *ReadTeamHookParams) WithID(id int64) *ReadTeamHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read team hook params
func (o *ReadTeamHookParams) SetID(id int64) {
	o.ID = id
}

// WithTeamID adds the teamID to the read team hook params
func (o *ReadTeamHookParams) WithTeamID(teamID string) *ReadTeamHookParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the read team hook params
func (o *ReadTeamHookParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *ReadTeamHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
