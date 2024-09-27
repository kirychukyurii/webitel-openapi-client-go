// Code generated by go-swagger; DO NOT EDIT.

package agent_service

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

// NewSearchAgentCallStatisticsParams creates a new SearchAgentCallStatisticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchAgentCallStatisticsParams() *SearchAgentCallStatisticsParams {
	return &SearchAgentCallStatisticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAgentCallStatisticsParamsWithTimeout creates a new SearchAgentCallStatisticsParams object
// with the ability to set a timeout on a request.
func NewSearchAgentCallStatisticsParamsWithTimeout(timeout time.Duration) *SearchAgentCallStatisticsParams {
	return &SearchAgentCallStatisticsParams{
		timeout: timeout,
	}
}

// NewSearchAgentCallStatisticsParamsWithContext creates a new SearchAgentCallStatisticsParams object
// with the ability to set a context for a request.
func NewSearchAgentCallStatisticsParamsWithContext(ctx context.Context) *SearchAgentCallStatisticsParams {
	return &SearchAgentCallStatisticsParams{
		Context: ctx,
	}
}

// NewSearchAgentCallStatisticsParamsWithHTTPClient creates a new SearchAgentCallStatisticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchAgentCallStatisticsParamsWithHTTPClient(client *http.Client) *SearchAgentCallStatisticsParams {
	return &SearchAgentCallStatisticsParams{
		HTTPClient: client,
	}
}

/*
SearchAgentCallStatisticsParams contains all the parameters to send to the API endpoint

	for the search agent call statistics operation.

	Typically these are written to a http.Request.
*/
type SearchAgentCallStatisticsParams struct {

	// AgentID.
	AgentID []int32

	// DomainID.
	//
	// Format: int64
	DomainID *string

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

	// TimeFrom.
	//
	// Format: int64
	TimeFrom *string

	// TimeTo.
	//
	// Format: int64
	TimeTo *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search agent call statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAgentCallStatisticsParams) WithDefaults() *SearchAgentCallStatisticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search agent call statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAgentCallStatisticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithTimeout(timeout time.Duration) *SearchAgentCallStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithContext(ctx context.Context) *SearchAgentCallStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithHTTPClient(client *http.Client) *SearchAgentCallStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithAgentID(agentID []int32) *SearchAgentCallStatisticsParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetAgentID(agentID []int32) {
	o.AgentID = agentID
}

// WithDomainID adds the domainID to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithDomainID(domainID *string) *SearchAgentCallStatisticsParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithFields adds the fields to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithFields(fields []string) *SearchAgentCallStatisticsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithPage adds the page to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithPage(page *int32) *SearchAgentCallStatisticsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithQ(q *string) *SearchAgentCallStatisticsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithSize(size *int32) *SearchAgentCallStatisticsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithSort(sort *string) *SearchAgentCallStatisticsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTimeFrom adds the timeFrom to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithTimeFrom(timeFrom *string) *SearchAgentCallStatisticsParams {
	o.SetTimeFrom(timeFrom)
	return o
}

// SetTimeFrom adds the timeFrom to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetTimeFrom(timeFrom *string) {
	o.TimeFrom = timeFrom
}

// WithTimeTo adds the timeTo to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) WithTimeTo(timeTo *string) *SearchAgentCallStatisticsParams {
	o.SetTimeTo(timeTo)
	return o
}

// SetTimeTo adds the timeTo to the search agent call statistics params
func (o *SearchAgentCallStatisticsParams) SetTimeTo(timeTo *string) {
	o.TimeTo = timeTo
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAgentCallStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AgentID != nil {

		// binding items for agent_id
		joinedAgentID := o.bindParamAgentID(reg)

		// query array param agent_id
		if err := r.SetQueryParam("agent_id", joinedAgentID...); err != nil {
			return err
		}
	}

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

	if o.TimeFrom != nil {

		// query param time.from
		var qrTimeFrom string

		if o.TimeFrom != nil {
			qrTimeFrom = *o.TimeFrom
		}
		qTimeFrom := qrTimeFrom
		if qTimeFrom != "" {

			if err := r.SetQueryParam("time.from", qTimeFrom); err != nil {
				return err
			}
		}
	}

	if o.TimeTo != nil {

		// query param time.to
		var qrTimeTo string

		if o.TimeTo != nil {
			qrTimeTo = *o.TimeTo
		}
		qTimeTo := qrTimeTo
		if qTimeTo != "" {

			if err := r.SetQueryParam("time.to", qTimeTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchAgentCallStatistics binds the parameter agent_id
func (o *SearchAgentCallStatisticsParams) bindParamAgentID(formats strfmt.Registry) []string {
	agentIDIR := o.AgentID

	var agentIDIC []string
	for _, agentIDIIR := range agentIDIR { // explode []int32

		agentIDIIV := swag.FormatInt32(agentIDIIR) // int32 as string
		agentIDIC = append(agentIDIC, agentIDIIV)
	}

	// items.CollectionFormat: "multi"
	agentIDIS := swag.JoinByFormat(agentIDIC, "multi")

	return agentIDIS
}

// bindParamSearchAgentCallStatistics binds the parameter fields
func (o *SearchAgentCallStatisticsParams) bindParamFields(formats strfmt.Registry) []string {
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
