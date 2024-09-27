// Code generated by go-swagger; DO NOT EDIT.

package calendar_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new calendar service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for calendar service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCalendar(body *models.EngineCreateCalendarRequest, opts ...ClientOption) (*CreateCalendarOK, error)
	CreateCalendarWithParams(params *CreateCalendarParams, opts ...ClientOption) (*CreateCalendarOK, error)

	DeleteCalendar(params *DeleteCalendarParams, opts ...ClientOption) (*DeleteCalendarOK, error)

	ReadCalendar(params *ReadCalendarParams, opts ...ClientOption) (*ReadCalendarOK, error)

	SearchCalendar(params *SearchCalendarParams, opts ...ClientOption) (*SearchCalendarOK, error)

	SearchTimezones(params *SearchTimezonesParams, opts ...ClientOption) (*SearchTimezonesOK, error)

	UpdateCalendar(id string, body *models.EngineUpdateCalendarRequest, opts ...ClientOption) (*UpdateCalendarOK, error)
	UpdateCalendarWithParams(params *UpdateCalendarParams, opts ...ClientOption) (*UpdateCalendarOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateCalendar creates calendar
*/
func (a *Client) CreateCalendar(body *models.EngineCreateCalendarRequest, opts ...ClientOption) (*CreateCalendarOK, error) {
	params := NewCreateCalendarParams().WithBody(body)
	return a.CreateCalendarWithParams(params, opts...)
}

func (a *Client) CreateCalendarWithParams(params *CreateCalendarParams, opts ...ClientOption) (*CreateCalendarOK, error) {
	if params == nil {
		params = NewCreateCalendarParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCalendar",
		Method:             "POST",
		PathPattern:        "/calendars",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCalendarReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCalendarOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateCalendarDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteCalendar removes calendar
*/

func (a *Client) DeleteCalendar(params *DeleteCalendarParams, opts ...ClientOption) (*DeleteCalendarOK, error) {
	if params == nil {
		params = NewDeleteCalendarParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCalendar",
		Method:             "DELETE",
		PathPattern:        "/calendars/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCalendarReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCalendarOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCalendarDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadCalendar calendars item
*/

func (a *Client) ReadCalendar(params *ReadCalendarParams, opts ...ClientOption) (*ReadCalendarOK, error) {
	if params == nil {
		params = NewReadCalendarParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadCalendar",
		Method:             "GET",
		PathPattern:        "/calendars/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadCalendarReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadCalendarOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadCalendarDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchCalendar lists of calendar
*/

func (a *Client) SearchCalendar(params *SearchCalendarParams, opts ...ClientOption) (*SearchCalendarOK, error) {
	if params == nil {
		params = NewSearchCalendarParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchCalendar",
		Method:             "GET",
		PathPattern:        "/calendars",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchCalendarReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchCalendarOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchCalendarDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchTimezones lists timezones
*/

func (a *Client) SearchTimezones(params *SearchTimezonesParams, opts ...ClientOption) (*SearchTimezonesOK, error) {
	if params == nil {
		params = NewSearchTimezonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchTimezones",
		Method:             "GET",
		PathPattern:        "/calendars/timezones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchTimezonesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchTimezonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchTimezonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateCalendar updates calendar
*/
func (a *Client) UpdateCalendar(id string, body *models.EngineUpdateCalendarRequest, opts ...ClientOption) (*UpdateCalendarOK, error) {
	params := NewUpdateCalendarParams().WithBody(body).WithID(id)
	return a.UpdateCalendarWithParams(params, opts...)
}

func (a *Client) UpdateCalendarWithParams(params *UpdateCalendarParams, opts ...ClientOption) (*UpdateCalendarOK, error) {
	if params == nil {
		params = NewUpdateCalendarParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCalendar",
		Method:             "PUT",
		PathPattern:        "/calendars/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCalendarReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCalendarOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCalendarDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// WithAuthInfo changes the transport on the client
func WithAuthInfo(authInfo runtime.ClientAuthInfoWriter) ClientOption {
	return func(op *runtime.ClientOperation) {
		op.AuthInfo = authInfo
	}
}
