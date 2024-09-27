// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewCustomersGetCustomerParams creates a new CustomersGetCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomersGetCustomerParams() *CustomersGetCustomerParams {
	return &CustomersGetCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomersGetCustomerParamsWithTimeout creates a new CustomersGetCustomerParams object
// with the ability to set a timeout on a request.
func NewCustomersGetCustomerParamsWithTimeout(timeout time.Duration) *CustomersGetCustomerParams {
	return &CustomersGetCustomerParams{
		timeout: timeout,
	}
}

// NewCustomersGetCustomerParamsWithContext creates a new CustomersGetCustomerParams object
// with the ability to set a context for a request.
func NewCustomersGetCustomerParamsWithContext(ctx context.Context) *CustomersGetCustomerParams {
	return &CustomersGetCustomerParams{
		Context: ctx,
	}
}

// NewCustomersGetCustomerParamsWithHTTPClient creates a new CustomersGetCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomersGetCustomerParamsWithHTTPClient(client *http.Client) *CustomersGetCustomerParams {
	return &CustomersGetCustomerParams{
		HTTPClient: client,
	}
}

/*
CustomersGetCustomerParams contains all the parameters to send to the API endpoint

	for the customers get customer operation.

	Typically these are written to a http.Request.
*/
type CustomersGetCustomerParams struct {

	/* DomainID.

	   identifier

	   Format: int64
	*/
	DomainID *string

	/* DomainName.

	   display name
	*/
	DomainName *string

	/* Fields.

	     Request Controls

	serial,
	*/
	Fields []string

	/* ID.

	     Available Filters

	show by customer id; serial number (uuid)
	*/
	ID *string

	// Sort.
	Sort []string

	/* Valid.

	   show if valid only!
	*/
	Valid *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customers get customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomersGetCustomerParams) WithDefaults() *CustomersGetCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customers get customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomersGetCustomerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customers get customer params
func (o *CustomersGetCustomerParams) WithTimeout(timeout time.Duration) *CustomersGetCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customers get customer params
func (o *CustomersGetCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customers get customer params
func (o *CustomersGetCustomerParams) WithContext(ctx context.Context) *CustomersGetCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customers get customer params
func (o *CustomersGetCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customers get customer params
func (o *CustomersGetCustomerParams) WithHTTPClient(client *http.Client) *CustomersGetCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customers get customer params
func (o *CustomersGetCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the customers get customer params
func (o *CustomersGetCustomerParams) WithDomainID(domainID *string) *CustomersGetCustomerParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the customers get customer params
func (o *CustomersGetCustomerParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithDomainName adds the domainName to the customers get customer params
func (o *CustomersGetCustomerParams) WithDomainName(domainName *string) *CustomersGetCustomerParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the customers get customer params
func (o *CustomersGetCustomerParams) SetDomainName(domainName *string) {
	o.DomainName = domainName
}

// WithFields adds the fields to the customers get customer params
func (o *CustomersGetCustomerParams) WithFields(fields []string) *CustomersGetCustomerParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the customers get customer params
func (o *CustomersGetCustomerParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the customers get customer params
func (o *CustomersGetCustomerParams) WithID(id *string) *CustomersGetCustomerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customers get customer params
func (o *CustomersGetCustomerParams) SetID(id *string) {
	o.ID = id
}

// WithSort adds the sort to the customers get customer params
func (o *CustomersGetCustomerParams) WithSort(sort []string) *CustomersGetCustomerParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the customers get customer params
func (o *CustomersGetCustomerParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithValid adds the valid to the customers get customer params
func (o *CustomersGetCustomerParams) WithValid(valid *bool) *CustomersGetCustomerParams {
	o.SetValid(valid)
	return o
}

// SetValid adds the valid to the customers get customer params
func (o *CustomersGetCustomerParams) SetValid(valid *bool) {
	o.Valid = valid
}

// WriteToRequest writes these params to a swagger request
func (o *CustomersGetCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domain.id
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domain.id", qDomainID); err != nil {
				return err
			}
		}
	}

	if o.DomainName != nil {

		// query param domain.name
		var qrDomainName string

		if o.DomainName != nil {
			qrDomainName = *o.DomainName
		}
		qDomainName := qrDomainName
		if qDomainName != "" {

			if err := r.SetQueryParam("domain.name", qDomainName); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// binding items for sort
		joinedSort := o.bindParamSort(reg)

		// query array param sort
		if err := r.SetQueryParam("sort", joinedSort...); err != nil {
			return err
		}
	}

	if o.Valid != nil {

		// query param valid
		var qrValid bool

		if o.Valid != nil {
			qrValid = *o.Valid
		}
		qValid := swag.FormatBool(qrValid)
		if qValid != "" {

			if err := r.SetQueryParam("valid", qValid); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCustomersGetCustomer binds the parameter fields
func (o *CustomersGetCustomerParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCustomersGetCustomer binds the parameter sort
func (o *CustomersGetCustomerParams) bindParamSort(formats strfmt.Registry) []string {
	sortIR := o.Sort

	var sortIC []string
	for _, sortIIR := range sortIR { // explode []string

		sortIIV := sortIIR // string as string
		sortIC = append(sortIC, sortIIV)
	}

	// items.CollectionFormat: "multi"
	sortIS := swag.JoinByFormat(sortIC, "multi")

	return sortIS
}
