// Code generated by go-swagger; DO NOT EDIT.

package ldap

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

// NewLDAPSearchLDAPCatalogParams creates a new LDAPSearchLDAPCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLDAPSearchLDAPCatalogParams() *LDAPSearchLDAPCatalogParams {
	return &LDAPSearchLDAPCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLDAPSearchLDAPCatalogParamsWithTimeout creates a new LDAPSearchLDAPCatalogParams object
// with the ability to set a timeout on a request.
func NewLDAPSearchLDAPCatalogParamsWithTimeout(timeout time.Duration) *LDAPSearchLDAPCatalogParams {
	return &LDAPSearchLDAPCatalogParams{
		timeout: timeout,
	}
}

// NewLDAPSearchLDAPCatalogParamsWithContext creates a new LDAPSearchLDAPCatalogParams object
// with the ability to set a context for a request.
func NewLDAPSearchLDAPCatalogParamsWithContext(ctx context.Context) *LDAPSearchLDAPCatalogParams {
	return &LDAPSearchLDAPCatalogParams{
		Context: ctx,
	}
}

// NewLDAPSearchLDAPCatalogParamsWithHTTPClient creates a new LDAPSearchLDAPCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewLDAPSearchLDAPCatalogParamsWithHTTPClient(client *http.Client) *LDAPSearchLDAPCatalogParams {
	return &LDAPSearchLDAPCatalogParams{
		HTTPClient: client,
	}
}

/*
LDAPSearchLDAPCatalogParams contains all the parameters to send to the API endpoint

	for the LDAP search LDAP catalog operation.

	Typically these are written to a http.Request.
*/
type LDAPSearchLDAPCatalogParams struct {

	/* Access.

	   [M]andatory[A]ccess[C]ontrol: with access mode (action) granted!
	*/
	Access *string

	/* Fields.

	   attributes list
	*/
	Fields []string

	/* ID.

	     ----- Search Basic Filters ---------------------------

	selection: by unique identifier
	*/
	ID []string

	/* Name.

	   case-ignore substring match: ILIKE '*' - any; '?' - one
	*/
	Name *string

	/* Page.

	     ----- Select Options -------------------------

	default: 1

	     Format: int32
	*/
	Page *int32

	/* Q.

	   term-of-search: lookup[name]
	*/
	Q *string

	/* Size.

	   default: 16

	   Format: int32
	*/
	Size *int32

	/* Sort.

	   e.g.: "updated_at" - ASC; "!updated_at" - DESC;
	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the LDAP search LDAP catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPSearchLDAPCatalogParams) WithDefaults() *LDAPSearchLDAPCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the LDAP search LDAP catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPSearchLDAPCatalogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithTimeout(timeout time.Duration) *LDAPSearchLDAPCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithContext(ctx context.Context) *LDAPSearchLDAPCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithHTTPClient(client *http.Client) *LDAPSearchLDAPCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccess adds the access to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithAccess(access *string) *LDAPSearchLDAPCatalogParams {
	o.SetAccess(access)
	return o
}

// SetAccess adds the access to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetAccess(access *string) {
	o.Access = access
}

// WithFields adds the fields to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithFields(fields []string) *LDAPSearchLDAPCatalogParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithID(id []string) *LDAPSearchLDAPCatalogParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithName(name *string) *LDAPSearchLDAPCatalogParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithPage(page *int32) *LDAPSearchLDAPCatalogParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithQ(q *string) *LDAPSearchLDAPCatalogParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithSize(size *int32) *LDAPSearchLDAPCatalogParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) WithSort(sort []string) *LDAPSearchLDAPCatalogParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the LDAP search LDAP catalog params
func (o *LDAPSearchLDAPCatalogParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *LDAPSearchLDAPCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Access != nil {

		// query param access
		var qrAccess string

		if o.Access != nil {
			qrAccess = *o.Access
		}
		qAccess := qrAccess
		if qAccess != "" {

			if err := r.SetQueryParam("access", qAccess); err != nil {
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

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
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

// bindParamLDAPSearchLDAPCatalog binds the parameter fields
func (o *LDAPSearchLDAPCatalogParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamLDAPSearchLDAPCatalog binds the parameter id
func (o *LDAPSearchLDAPCatalogParams) bindParamID(formats strfmt.Registry) []string {
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

// bindParamLDAPSearchLDAPCatalog binds the parameter sort
func (o *LDAPSearchLDAPCatalogParams) bindParamSort(formats strfmt.Registry) []string {
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
