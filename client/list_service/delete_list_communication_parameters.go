// Code generated by go-swagger; DO NOT EDIT.

package list_service

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

// NewDeleteListCommunicationParams creates a new DeleteListCommunicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteListCommunicationParams() *DeleteListCommunicationParams {
	return &DeleteListCommunicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteListCommunicationParamsWithTimeout creates a new DeleteListCommunicationParams object
// with the ability to set a timeout on a request.
func NewDeleteListCommunicationParamsWithTimeout(timeout time.Duration) *DeleteListCommunicationParams {
	return &DeleteListCommunicationParams{
		timeout: timeout,
	}
}

// NewDeleteListCommunicationParamsWithContext creates a new DeleteListCommunicationParams object
// with the ability to set a context for a request.
func NewDeleteListCommunicationParamsWithContext(ctx context.Context) *DeleteListCommunicationParams {
	return &DeleteListCommunicationParams{
		Context: ctx,
	}
}

// NewDeleteListCommunicationParamsWithHTTPClient creates a new DeleteListCommunicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteListCommunicationParamsWithHTTPClient(client *http.Client) *DeleteListCommunicationParams {
	return &DeleteListCommunicationParams{
		HTTPClient: client,
	}
}

/*
DeleteListCommunicationParams contains all the parameters to send to the API endpoint

	for the delete list communication operation.

	Typically these are written to a http.Request.
*/
type DeleteListCommunicationParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID string

	// ListID.
	//
	// Format: int64
	ListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete list communication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteListCommunicationParams) WithDefaults() *DeleteListCommunicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete list communication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteListCommunicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete list communication params
func (o *DeleteListCommunicationParams) WithTimeout(timeout time.Duration) *DeleteListCommunicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete list communication params
func (o *DeleteListCommunicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete list communication params
func (o *DeleteListCommunicationParams) WithContext(ctx context.Context) *DeleteListCommunicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete list communication params
func (o *DeleteListCommunicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete list communication params
func (o *DeleteListCommunicationParams) WithHTTPClient(client *http.Client) *DeleteListCommunicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete list communication params
func (o *DeleteListCommunicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete list communication params
func (o *DeleteListCommunicationParams) WithDomainID(domainID *string) *DeleteListCommunicationParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete list communication params
func (o *DeleteListCommunicationParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the delete list communication params
func (o *DeleteListCommunicationParams) WithID(id string) *DeleteListCommunicationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete list communication params
func (o *DeleteListCommunicationParams) SetID(id string) {
	o.ID = id
}

// WithListID adds the listID to the delete list communication params
func (o *DeleteListCommunicationParams) WithListID(listID string) *DeleteListCommunicationParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the delete list communication params
func (o *DeleteListCommunicationParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteListCommunicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param list_id
	if err := r.SetPathParam("list_id", o.ListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
