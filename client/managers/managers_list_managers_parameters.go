// Code generated by go-swagger; DO NOT EDIT.

package managers

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

// NewManagersListManagersParams creates a new ManagersListManagersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManagersListManagersParams() *ManagersListManagersParams {
	return &ManagersListManagersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManagersListManagersParamsWithTimeout creates a new ManagersListManagersParams object
// with the ability to set a timeout on a request.
func NewManagersListManagersParamsWithTimeout(timeout time.Duration) *ManagersListManagersParams {
	return &ManagersListManagersParams{
		timeout: timeout,
	}
}

// NewManagersListManagersParamsWithContext creates a new ManagersListManagersParams object
// with the ability to set a context for a request.
func NewManagersListManagersParamsWithContext(ctx context.Context) *ManagersListManagersParams {
	return &ManagersListManagersParams{
		Context: ctx,
	}
}

// NewManagersListManagersParamsWithHTTPClient creates a new ManagersListManagersParams object
// with the ability to set a custom HTTPClient for a request.
func NewManagersListManagersParamsWithHTTPClient(client *http.Client) *ManagersListManagersParams {
	return &ManagersListManagersParams{
		HTTPClient: client,
	}
}

/*
ManagersListManagersParams contains all the parameters to send to the API endpoint

	for the managers list managers operation.

	Typically these are written to a http.Request.
*/
type ManagersListManagersParams struct {

	/* ContactID.

	   Contact ID associated with.
	*/
	ContactID string

	/* Fields.

	   Fields to be retrieved as a result.
	*/
	Fields []string

	/* ID.

	   Record(s) with unique ID only.
	*/
	ID []string

	/* Page.

	   Page number of result. offset = ((page-1)*size)

	   Format: int32
	*/
	Page *int32

	/* Q.

	     Search term: user name;
	`?` - matches any one character
	`*` - matches 0 or more characters
	*/
	Q *string

	/* Size.

	   Size of result page. limit = (size++)

	   Format: int32
	*/
	Size *int32

	/* Sort.

	   Sort the result according to fields.
	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the managers list managers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagersListManagersParams) WithDefaults() *ManagersListManagersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the managers list managers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagersListManagersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the managers list managers params
func (o *ManagersListManagersParams) WithTimeout(timeout time.Duration) *ManagersListManagersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the managers list managers params
func (o *ManagersListManagersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the managers list managers params
func (o *ManagersListManagersParams) WithContext(ctx context.Context) *ManagersListManagersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the managers list managers params
func (o *ManagersListManagersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the managers list managers params
func (o *ManagersListManagersParams) WithHTTPClient(client *http.Client) *ManagersListManagersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the managers list managers params
func (o *ManagersListManagersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the managers list managers params
func (o *ManagersListManagersParams) WithContactID(contactID string) *ManagersListManagersParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the managers list managers params
func (o *ManagersListManagersParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithFields adds the fields to the managers list managers params
func (o *ManagersListManagersParams) WithFields(fields []string) *ManagersListManagersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the managers list managers params
func (o *ManagersListManagersParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the managers list managers params
func (o *ManagersListManagersParams) WithID(id []string) *ManagersListManagersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the managers list managers params
func (o *ManagersListManagersParams) SetID(id []string) {
	o.ID = id
}

// WithPage adds the page to the managers list managers params
func (o *ManagersListManagersParams) WithPage(page *int32) *ManagersListManagersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the managers list managers params
func (o *ManagersListManagersParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the managers list managers params
func (o *ManagersListManagersParams) WithQ(q *string) *ManagersListManagersParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the managers list managers params
func (o *ManagersListManagersParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the managers list managers params
func (o *ManagersListManagersParams) WithSize(size *int32) *ManagersListManagersParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the managers list managers params
func (o *ManagersListManagersParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the managers list managers params
func (o *ManagersListManagersParams) WithSort(sort []string) *ManagersListManagersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the managers list managers params
func (o *ManagersListManagersParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ManagersListManagersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
		return err
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

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamManagersListManagers binds the parameter fields
func (o *ManagersListManagersParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamManagersListManagers binds the parameter id
func (o *ManagersListManagersParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}

// bindParamManagersListManagers binds the parameter sort
func (o *ManagersListManagersParams) bindParamSort(formats strfmt.Registry) []string {
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
