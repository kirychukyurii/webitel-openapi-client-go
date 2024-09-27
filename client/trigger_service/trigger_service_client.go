// Code generated by go-swagger; DO NOT EDIT.

package trigger_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new trigger service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trigger service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTrigger(body *models.EngineCreateTriggerRequest, opts ...ClientOption) (*CreateTriggerOK, error)
	CreateTriggerWithParams(params *CreateTriggerParams, opts ...ClientOption) (*CreateTriggerOK, error)

	CreateTriggerJob(triggerID int32, body *models.EngineCreateTriggerJobRequest, opts ...ClientOption) (*CreateTriggerJobOK, error)
	CreateTriggerJobWithParams(params *CreateTriggerJobParams, opts ...ClientOption) (*CreateTriggerJobOK, error)

	DeleteTrigger(id int32, opts ...ClientOption) (*DeleteTriggerOK, error)
	DeleteTriggerWithParams(params *DeleteTriggerParams, opts ...ClientOption) (*DeleteTriggerOK, error)

	PatchTrigger(id int32, body *models.EnginePatchTriggerRequest, opts ...ClientOption) (*PatchTriggerOK, error)
	PatchTriggerWithParams(params *PatchTriggerParams, opts ...ClientOption) (*PatchTriggerOK, error)

	ReadTrigger(id int32, opts ...ClientOption) (*ReadTriggerOK, error)
	ReadTriggerWithParams(params *ReadTriggerParams, opts ...ClientOption) (*ReadTriggerOK, error)

	SearchTrigger(params *SearchTriggerParams, opts ...ClientOption) (*SearchTriggerOK, error)

	SearchTriggerJob(params *SearchTriggerJobParams, opts ...ClientOption) (*SearchTriggerJobOK, error)

	UpdateTrigger(id int32, body *models.EngineUpdateTriggerRequest, opts ...ClientOption) (*UpdateTriggerOK, error)
	UpdateTriggerWithParams(params *UpdateTriggerParams, opts ...ClientOption) (*UpdateTriggerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateTrigger creates trigger
*/
func (a *Client) CreateTrigger(body *models.EngineCreateTriggerRequest, opts ...ClientOption) (*CreateTriggerOK, error) {
	params := NewCreateTriggerParams().WithBody(body)
	return a.CreateTriggerWithParams(params, opts...)
}

func (a *Client) CreateTriggerWithParams(params *CreateTriggerParams, opts ...ClientOption) (*CreateTriggerOK, error) {
	if params == nil {
		params = NewCreateTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTrigger",
		Method:             "POST",
		PathPattern:        "/trigger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTriggerReader{formats: a.formats},
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
	success, ok := result.(*CreateTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateTriggerJob create trigger job API
*/
func (a *Client) CreateTriggerJob(triggerID int32, body *models.EngineCreateTriggerJobRequest, opts ...ClientOption) (*CreateTriggerJobOK, error) {
	params := NewCreateTriggerJobParams().WithBody(body).WithTriggerID(triggerID)
	return a.CreateTriggerJobWithParams(params, opts...)
}

func (a *Client) CreateTriggerJobWithParams(params *CreateTriggerJobParams, opts ...ClientOption) (*CreateTriggerJobOK, error) {
	if params == nil {
		params = NewCreateTriggerJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTriggerJob",
		Method:             "POST",
		PathPattern:        "/trigger/{trigger_id}/job",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTriggerJobReader{formats: a.formats},
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
	success, ok := result.(*CreateTriggerJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTriggerJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteTrigger removes trigger
*/
func (a *Client) DeleteTrigger(id int32, opts ...ClientOption) (*DeleteTriggerOK, error) {
	params := NewDeleteTriggerParams().WithID(id)
	return a.DeleteTriggerWithParams(params, opts...)
}

func (a *Client) DeleteTriggerWithParams(params *DeleteTriggerParams, opts ...ClientOption) (*DeleteTriggerOK, error) {
	if params == nil {
		params = NewDeleteTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTrigger",
		Method:             "DELETE",
		PathPattern:        "/trigger/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTriggerReader{formats: a.formats},
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
	success, ok := result.(*DeleteTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchTrigger patch trigger API
*/
func (a *Client) PatchTrigger(id int32, body *models.EnginePatchTriggerRequest, opts ...ClientOption) (*PatchTriggerOK, error) {
	params := NewPatchTriggerParams().WithBody(body).WithID(id)
	return a.PatchTriggerWithParams(params, opts...)
}

func (a *Client) PatchTriggerWithParams(params *PatchTriggerParams, opts ...ClientOption) (*PatchTriggerOK, error) {
	if params == nil {
		params = NewPatchTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchTrigger",
		Method:             "PATCH",
		PathPattern:        "/trigger/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchTriggerReader{formats: a.formats},
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
	success, ok := result.(*PatchTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadTrigger triggers item
*/
func (a *Client) ReadTrigger(id int32, opts ...ClientOption) (*ReadTriggerOK, error) {
	params := NewReadTriggerParams().WithID(id)
	return a.ReadTriggerWithParams(params, opts...)
}

func (a *Client) ReadTriggerWithParams(params *ReadTriggerParams, opts ...ClientOption) (*ReadTriggerOK, error) {
	if params == nil {
		params = NewReadTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadTrigger",
		Method:             "GET",
		PathPattern:        "/trigger/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadTriggerReader{formats: a.formats},
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
	success, ok := result.(*ReadTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchTrigger lists of trigger
*/

func (a *Client) SearchTrigger(params *SearchTriggerParams, opts ...ClientOption) (*SearchTriggerOK, error) {
	if params == nil {
		params = NewSearchTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchTrigger",
		Method:             "GET",
		PathPattern:        "/trigger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchTriggerReader{formats: a.formats},
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
	success, ok := result.(*SearchTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchTriggerJob search trigger job API
*/

func (a *Client) SearchTriggerJob(params *SearchTriggerJobParams, opts ...ClientOption) (*SearchTriggerJobOK, error) {
	if params == nil {
		params = NewSearchTriggerJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchTriggerJob",
		Method:             "GET",
		PathPattern:        "/trigger/{trigger_id}/job",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchTriggerJobReader{formats: a.formats},
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
	success, ok := result.(*SearchTriggerJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchTriggerJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateTrigger updates trigger
*/
func (a *Client) UpdateTrigger(id int32, body *models.EngineUpdateTriggerRequest, opts ...ClientOption) (*UpdateTriggerOK, error) {
	params := NewUpdateTriggerParams().WithBody(body).WithID(id)
	return a.UpdateTriggerWithParams(params, opts...)
}

func (a *Client) UpdateTriggerWithParams(params *UpdateTriggerParams, opts ...ClientOption) (*UpdateTriggerOK, error) {
	if params == nil {
		params = NewUpdateTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTrigger",
		Method:             "PUT",
		PathPattern:        "/trigger/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTriggerReader{formats: a.formats},
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
	success, ok := result.(*UpdateTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateTriggerDefault)
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
