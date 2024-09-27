// Code generated by go-swagger; DO NOT EDIT.

package contacts_chat_catalog

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

// NewContactsChatCatalogGetContactChatHistory2Params creates a new ContactsChatCatalogGetContactChatHistory2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactsChatCatalogGetContactChatHistory2Params() *ContactsChatCatalogGetContactChatHistory2Params {
	return &ContactsChatCatalogGetContactChatHistory2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactsChatCatalogGetContactChatHistory2ParamsWithTimeout creates a new ContactsChatCatalogGetContactChatHistory2Params object
// with the ability to set a timeout on a request.
func NewContactsChatCatalogGetContactChatHistory2ParamsWithTimeout(timeout time.Duration) *ContactsChatCatalogGetContactChatHistory2Params {
	return &ContactsChatCatalogGetContactChatHistory2Params{
		timeout: timeout,
	}
}

// NewContactsChatCatalogGetContactChatHistory2ParamsWithContext creates a new ContactsChatCatalogGetContactChatHistory2Params object
// with the ability to set a context for a request.
func NewContactsChatCatalogGetContactChatHistory2ParamsWithContext(ctx context.Context) *ContactsChatCatalogGetContactChatHistory2Params {
	return &ContactsChatCatalogGetContactChatHistory2Params{
		Context: ctx,
	}
}

// NewContactsChatCatalogGetContactChatHistory2ParamsWithHTTPClient creates a new ContactsChatCatalogGetContactChatHistory2Params object
// with the ability to set a custom HTTPClient for a request.
func NewContactsChatCatalogGetContactChatHistory2ParamsWithHTTPClient(client *http.Client) *ContactsChatCatalogGetContactChatHistory2Params {
	return &ContactsChatCatalogGetContactChatHistory2Params{
		HTTPClient: client,
	}
}

/*
ContactsChatCatalogGetContactChatHistory2Params contains all the parameters to send to the API endpoint

	for the contacts chat catalog get contact chat history2 operation.

	Typically these are written to a http.Request.
*/
type ContactsChatCatalogGetContactChatHistory2Params struct {

	/* ChatID.

	   Unique chat dialog
	*/
	ChatID *string

	/* ContactID.

	   Contact identificator
	*/
	ContactID string

	/* Fields.

	   Fields to return into result.
	*/
	Fields []string

	/* GroupStringString.

	   This is a request variable of the map type. The query format is "map_name[key]=value", e.g. If the map name is Age, the key type is string, and the value type is integer, the query parameter is expressed as Age["bob"]=18
	*/
	GroupStringString *string

	/* OffsetDate.

	   Messages ONLY been sent before the specified epochtime(milli).

	   Format: int64
	*/
	OffsetDate *string

	/* OffsetID.

	   Messages ONLY starting from the specified message ID

	   Format: int64
	*/
	OffsetID *string

	// Page.
	//
	// Format: int32
	Page *int32

	/* Q.

	   Search term: message.text
	*/
	Q *string

	/* Size.

	   Number of messages to return.

	   Format: int32
	*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contacts chat catalog get contact chat history2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithDefaults() *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contacts chat catalog get contact chat history2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithTimeout(timeout time.Duration) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithContext(ctx context.Context) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithHTTPClient(client *http.Client) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChatID adds the chatID to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithChatID(chatID *string) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetChatID(chatID)
	return o
}

// SetChatID adds the chatId to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetChatID(chatID *string) {
	o.ChatID = chatID
}

// WithContactID adds the contactID to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithContactID(contactID string) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithFields adds the fields to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithFields(fields []string) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetFields(fields []string) {
	o.Fields = fields
}

// WithGroupStringString adds the groupStringString to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithGroupStringString(groupStringString *string) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetGroupStringString(groupStringString)
	return o
}

// SetGroupStringString adds the groupStringString to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetGroupStringString(groupStringString *string) {
	o.GroupStringString = groupStringString
}

// WithOffsetDate adds the offsetDate to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithOffsetDate(offsetDate *string) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetOffsetDate(offsetDate)
	return o
}

// SetOffsetDate adds the offsetDate to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetOffsetDate(offsetDate *string) {
	o.OffsetDate = offsetDate
}

// WithOffsetID adds the offsetID to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithOffsetID(offsetID *string) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetOffsetID(offsetID)
	return o
}

// SetOffsetID adds the offsetId to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetOffsetID(offsetID *string) {
	o.OffsetID = offsetID
}

// WithPage adds the page to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithPage(page *int32) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithQ(q *string) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) WithSize(size *int32) *ContactsChatCatalogGetContactChatHistory2Params {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the contacts chat catalog get contact chat history2 params
func (o *ContactsChatCatalogGetContactChatHistory2Params) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *ContactsChatCatalogGetContactChatHistory2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChatID != nil {

		// query param chat_id
		var qrChatID string

		if o.ChatID != nil {
			qrChatID = *o.ChatID
		}
		qChatID := qrChatID
		if qChatID != "" {

			if err := r.SetQueryParam("chat_id", qChatID); err != nil {
				return err
			}
		}
	}

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

	if o.GroupStringString != nil {

		// query param group[string][string]
		var qrGroupStringString string

		if o.GroupStringString != nil {
			qrGroupStringString = *o.GroupStringString
		}
		qGroupStringString := qrGroupStringString
		if qGroupStringString != "" {

			if err := r.SetQueryParam("group[string][string]", qGroupStringString); err != nil {
				return err
			}
		}
	}

	if o.OffsetDate != nil {

		// query param offset.date
		var qrOffsetDate string

		if o.OffsetDate != nil {
			qrOffsetDate = *o.OffsetDate
		}
		qOffsetDate := qrOffsetDate
		if qOffsetDate != "" {

			if err := r.SetQueryParam("offset.date", qOffsetDate); err != nil {
				return err
			}
		}
	}

	if o.OffsetID != nil {

		// query param offset.id
		var qrOffsetID string

		if o.OffsetID != nil {
			qrOffsetID = *o.OffsetID
		}
		qOffsetID := qrOffsetID
		if qOffsetID != "" {

			if err := r.SetQueryParam("offset.id", qOffsetID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamContactsChatCatalogGetContactChatHistory2 binds the parameter fields
func (o *ContactsChatCatalogGetContactChatHistory2Params) bindParamFields(formats strfmt.Registry) []string {
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
