// Code generated by go-swagger; DO NOT EDIT.

package trigger_service

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

// NewSearchTriggerJobParams creates a new SearchTriggerJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchTriggerJobParams() *SearchTriggerJobParams {
	return &SearchTriggerJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchTriggerJobParamsWithTimeout creates a new SearchTriggerJobParams object
// with the ability to set a timeout on a request.
func NewSearchTriggerJobParamsWithTimeout(timeout time.Duration) *SearchTriggerJobParams {
	return &SearchTriggerJobParams{
		timeout: timeout,
	}
}

// NewSearchTriggerJobParamsWithContext creates a new SearchTriggerJobParams object
// with the ability to set a context for a request.
func NewSearchTriggerJobParamsWithContext(ctx context.Context) *SearchTriggerJobParams {
	return &SearchTriggerJobParams{
		Context: ctx,
	}
}

// NewSearchTriggerJobParamsWithHTTPClient creates a new SearchTriggerJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchTriggerJobParamsWithHTTPClient(client *http.Client) *SearchTriggerJobParams {
	return &SearchTriggerJobParams{
		HTTPClient: client,
	}
}

/*
SearchTriggerJobParams contains all the parameters to send to the API endpoint

	for the search trigger job operation.

	Typically these are written to a http.Request.
*/
type SearchTriggerJobParams struct {

	// CreatedAtFrom.
	//
	// Format: int64
	CreatedAtFrom *string

	// CreatedAtTo.
	//
	// Format: int64
	CreatedAtTo *string

	// DurationFrom.
	//
	// Format: int64
	DurationFrom *string

	// DurationTo.
	//
	// Format: int64
	DurationTo *string

	// Fields.
	Fields []string

	// Page.
	//
	// Format: int32
	Page *int32

	// Q.
	Q *string

	// Size.
	//
	// Format: int32
	Size *int32

	// Sort.
	Sort *string

	// StartedAtFrom.
	//
	// Format: int64
	StartedAtFrom *string

	// StartedAtTo.
	//
	// Format: int64
	StartedAtTo *string

	// State.
	State []string

	// TriggerID.
	//
	// Format: int32
	TriggerID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search trigger job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTriggerJobParams) WithDefaults() *SearchTriggerJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search trigger job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTriggerJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search trigger job params
func (o *SearchTriggerJobParams) WithTimeout(timeout time.Duration) *SearchTriggerJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search trigger job params
func (o *SearchTriggerJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search trigger job params
func (o *SearchTriggerJobParams) WithContext(ctx context.Context) *SearchTriggerJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search trigger job params
func (o *SearchTriggerJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search trigger job params
func (o *SearchTriggerJobParams) WithHTTPClient(client *http.Client) *SearchTriggerJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search trigger job params
func (o *SearchTriggerJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAtFrom adds the createdAtFrom to the search trigger job params
func (o *SearchTriggerJobParams) WithCreatedAtFrom(createdAtFrom *string) *SearchTriggerJobParams {
	o.SetCreatedAtFrom(createdAtFrom)
	return o
}

// SetCreatedAtFrom adds the createdAtFrom to the search trigger job params
func (o *SearchTriggerJobParams) SetCreatedAtFrom(createdAtFrom *string) {
	o.CreatedAtFrom = createdAtFrom
}

// WithCreatedAtTo adds the createdAtTo to the search trigger job params
func (o *SearchTriggerJobParams) WithCreatedAtTo(createdAtTo *string) *SearchTriggerJobParams {
	o.SetCreatedAtTo(createdAtTo)
	return o
}

// SetCreatedAtTo adds the createdAtTo to the search trigger job params
func (o *SearchTriggerJobParams) SetCreatedAtTo(createdAtTo *string) {
	o.CreatedAtTo = createdAtTo
}

// WithDurationFrom adds the durationFrom to the search trigger job params
func (o *SearchTriggerJobParams) WithDurationFrom(durationFrom *string) *SearchTriggerJobParams {
	o.SetDurationFrom(durationFrom)
	return o
}

// SetDurationFrom adds the durationFrom to the search trigger job params
func (o *SearchTriggerJobParams) SetDurationFrom(durationFrom *string) {
	o.DurationFrom = durationFrom
}

// WithDurationTo adds the durationTo to the search trigger job params
func (o *SearchTriggerJobParams) WithDurationTo(durationTo *string) *SearchTriggerJobParams {
	o.SetDurationTo(durationTo)
	return o
}

// SetDurationTo adds the durationTo to the search trigger job params
func (o *SearchTriggerJobParams) SetDurationTo(durationTo *string) {
	o.DurationTo = durationTo
}

// WithFields adds the fields to the search trigger job params
func (o *SearchTriggerJobParams) WithFields(fields []string) *SearchTriggerJobParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search trigger job params
func (o *SearchTriggerJobParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithPage adds the page to the search trigger job params
func (o *SearchTriggerJobParams) WithPage(page *int32) *SearchTriggerJobParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search trigger job params
func (o *SearchTriggerJobParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the search trigger job params
func (o *SearchTriggerJobParams) WithQ(q *string) *SearchTriggerJobParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search trigger job params
func (o *SearchTriggerJobParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the search trigger job params
func (o *SearchTriggerJobParams) WithSize(size *int32) *SearchTriggerJobParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search trigger job params
func (o *SearchTriggerJobParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the search trigger job params
func (o *SearchTriggerJobParams) WithSort(sort *string) *SearchTriggerJobParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search trigger job params
func (o *SearchTriggerJobParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStartedAtFrom adds the startedAtFrom to the search trigger job params
func (o *SearchTriggerJobParams) WithStartedAtFrom(startedAtFrom *string) *SearchTriggerJobParams {
	o.SetStartedAtFrom(startedAtFrom)
	return o
}

// SetStartedAtFrom adds the startedAtFrom to the search trigger job params
func (o *SearchTriggerJobParams) SetStartedAtFrom(startedAtFrom *string) {
	o.StartedAtFrom = startedAtFrom
}

// WithStartedAtTo adds the startedAtTo to the search trigger job params
func (o *SearchTriggerJobParams) WithStartedAtTo(startedAtTo *string) *SearchTriggerJobParams {
	o.SetStartedAtTo(startedAtTo)
	return o
}

// SetStartedAtTo adds the startedAtTo to the search trigger job params
func (o *SearchTriggerJobParams) SetStartedAtTo(startedAtTo *string) {
	o.StartedAtTo = startedAtTo
}

// WithState adds the state to the search trigger job params
func (o *SearchTriggerJobParams) WithState(state []string) *SearchTriggerJobParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the search trigger job params
func (o *SearchTriggerJobParams) SetState(state []string) {
	o.State = state
}

// WithTriggerID adds the triggerID to the search trigger job params
func (o *SearchTriggerJobParams) WithTriggerID(triggerID int32) *SearchTriggerJobParams {
	o.SetTriggerID(triggerID)
	return o
}

// SetTriggerID adds the triggerId to the search trigger job params
func (o *SearchTriggerJobParams) SetTriggerID(triggerID int32) {
	o.TriggerID = triggerID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchTriggerJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatedAtFrom != nil {

		// query param created_at.from
		var qrCreatedAtFrom string

		if o.CreatedAtFrom != nil {
			qrCreatedAtFrom = *o.CreatedAtFrom
		}
		qCreatedAtFrom := qrCreatedAtFrom
		if qCreatedAtFrom != "" {

			if err := r.SetQueryParam("created_at.from", qCreatedAtFrom); err != nil {
				return err
			}
		}
	}

	if o.CreatedAtTo != nil {

		// query param created_at.to
		var qrCreatedAtTo string

		if o.CreatedAtTo != nil {
			qrCreatedAtTo = *o.CreatedAtTo
		}
		qCreatedAtTo := qrCreatedAtTo
		if qCreatedAtTo != "" {

			if err := r.SetQueryParam("created_at.to", qCreatedAtTo); err != nil {
				return err
			}
		}
	}

	if o.DurationFrom != nil {

		// query param duration.from
		var qrDurationFrom string

		if o.DurationFrom != nil {
			qrDurationFrom = *o.DurationFrom
		}
		qDurationFrom := qrDurationFrom
		if qDurationFrom != "" {

			if err := r.SetQueryParam("duration.from", qDurationFrom); err != nil {
				return err
			}
		}
	}

	if o.DurationTo != nil {

		// query param duration.to
		var qrDurationTo string

		if o.DurationTo != nil {
			qrDurationTo = *o.DurationTo
		}
		qDurationTo := qrDurationTo
		if qDurationTo != "" {

			if err := r.SetQueryParam("duration.to", qDurationTo); err != nil {
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

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.StartedAtFrom != nil {

		// query param started_at.from
		var qrStartedAtFrom string

		if o.StartedAtFrom != nil {
			qrStartedAtFrom = *o.StartedAtFrom
		}
		qStartedAtFrom := qrStartedAtFrom
		if qStartedAtFrom != "" {

			if err := r.SetQueryParam("started_at.from", qStartedAtFrom); err != nil {
				return err
			}
		}
	}

	if o.StartedAtTo != nil {

		// query param started_at.to
		var qrStartedAtTo string

		if o.StartedAtTo != nil {
			qrStartedAtTo = *o.StartedAtTo
		}
		qStartedAtTo := qrStartedAtTo
		if qStartedAtTo != "" {

			if err := r.SetQueryParam("started_at.to", qStartedAtTo); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// binding items for state
		joinedState := o.bindParamState(reg)

		// query array param state
		if err := r.SetQueryParam("state", joinedState...); err != nil {
			return err
		}
	}

	// path param trigger_id
	if err := r.SetPathParam("trigger_id", swag.FormatInt32(o.TriggerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchTriggerJob binds the parameter fields
func (o *SearchTriggerJobParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSearchTriggerJob binds the parameter state
func (o *SearchTriggerJobParams) bindParamState(formats strfmt.Registry) []string {
	stateIR := o.State

	var stateIC []string
	for _, stateIIR := range stateIR { // explode []string

		stateIIV := stateIIR // string as string
		stateIC = append(stateIC, stateIIV)
	}

	// items.CollectionFormat: "multi"
	stateIS := swag.JoinByFormat(stateIC, "multi")

	return stateIS
}
