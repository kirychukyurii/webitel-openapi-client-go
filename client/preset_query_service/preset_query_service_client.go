// Code generated by go-swagger; DO NOT EDIT.

package preset_query_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new preset query service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for preset query service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePresetQuery(body *models.EngineCreatePresetQueryRequest, opts ...ClientOption) (*CreatePresetQueryOK, error)
	CreatePresetQueryWithParams(params *CreatePresetQueryParams, opts ...ClientOption) (*CreatePresetQueryOK, error)

	DeletePresetQuery(id int32, opts ...ClientOption) (*DeletePresetQueryOK, error)
	DeletePresetQueryWithParams(params *DeletePresetQueryParams, opts ...ClientOption) (*DeletePresetQueryOK, error)

	PatchPresetQuery(id int32, body *models.EnginePatchPresetQueryRequest, opts ...ClientOption) (*PatchPresetQueryOK, error)
	PatchPresetQueryWithParams(params *PatchPresetQueryParams, opts ...ClientOption) (*PatchPresetQueryOK, error)

	ReadPresetQuery(id int32, opts ...ClientOption) (*ReadPresetQueryOK, error)
	ReadPresetQueryWithParams(params *ReadPresetQueryParams, opts ...ClientOption) (*ReadPresetQueryOK, error)

	SearchPresetQuery(params *SearchPresetQueryParams, opts ...ClientOption) (*SearchPresetQueryOK, error)

	UpdatePresetQuery(id int32, body *models.EngineUpdatePresetQueryRequest, opts ...ClientOption) (*UpdatePresetQueryOK, error)
	UpdatePresetQueryWithParams(params *UpdatePresetQueryParams, opts ...ClientOption) (*UpdatePresetQueryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePresetQuery create preset query API
*/
func (a *Client) CreatePresetQuery(body *models.EngineCreatePresetQueryRequest, opts ...ClientOption) (*CreatePresetQueryOK, error) {
	params := NewCreatePresetQueryParams().WithBody(body)
	return a.CreatePresetQueryWithParams(params, opts...)
}

func (a *Client) CreatePresetQueryWithParams(params *CreatePresetQueryParams, opts ...ClientOption) (*CreatePresetQueryOK, error) {
	if params == nil {
		params = NewCreatePresetQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreatePresetQuery",
		Method:             "POST",
		PathPattern:        "/call_center/preset/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePresetQueryReader{formats: a.formats},
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
	success, ok := result.(*CreatePresetQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreatePresetQueryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeletePresetQuery delete preset query API
*/
func (a *Client) DeletePresetQuery(id int32, opts ...ClientOption) (*DeletePresetQueryOK, error) {
	params := NewDeletePresetQueryParams().WithID(id)
	return a.DeletePresetQueryWithParams(params, opts...)
}

func (a *Client) DeletePresetQueryWithParams(params *DeletePresetQueryParams, opts ...ClientOption) (*DeletePresetQueryOK, error) {
	if params == nil {
		params = NewDeletePresetQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePresetQuery",
		Method:             "DELETE",
		PathPattern:        "/call_center/preset/query/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePresetQueryReader{formats: a.formats},
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
	success, ok := result.(*DeletePresetQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePresetQueryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchPresetQuery patch preset query API
*/
func (a *Client) PatchPresetQuery(id int32, body *models.EnginePatchPresetQueryRequest, opts ...ClientOption) (*PatchPresetQueryOK, error) {
	params := NewPatchPresetQueryParams().WithBody(body).WithID(id)
	return a.PatchPresetQueryWithParams(params, opts...)
}

func (a *Client) PatchPresetQueryWithParams(params *PatchPresetQueryParams, opts ...ClientOption) (*PatchPresetQueryOK, error) {
	if params == nil {
		params = NewPatchPresetQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchPresetQuery",
		Method:             "PATCH",
		PathPattern:        "/call_center/preset/query/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPresetQueryReader{formats: a.formats},
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
	success, ok := result.(*PatchPresetQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchPresetQueryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadPresetQuery read preset query API
*/
func (a *Client) ReadPresetQuery(id int32, opts ...ClientOption) (*ReadPresetQueryOK, error) {
	params := NewReadPresetQueryParams().WithID(id)
	return a.ReadPresetQueryWithParams(params, opts...)
}

func (a *Client) ReadPresetQueryWithParams(params *ReadPresetQueryParams, opts ...ClientOption) (*ReadPresetQueryOK, error) {
	if params == nil {
		params = NewReadPresetQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadPresetQuery",
		Method:             "GET",
		PathPattern:        "/call_center/preset/query/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadPresetQueryReader{formats: a.formats},
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
	success, ok := result.(*ReadPresetQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadPresetQueryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchPresetQuery search preset query API
*/

func (a *Client) SearchPresetQuery(params *SearchPresetQueryParams, opts ...ClientOption) (*SearchPresetQueryOK, error) {
	if params == nil {
		params = NewSearchPresetQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchPresetQuery",
		Method:             "GET",
		PathPattern:        "/call_center/preset/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchPresetQueryReader{formats: a.formats},
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
	success, ok := result.(*SearchPresetQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchPresetQueryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdatePresetQuery update preset query API
*/
func (a *Client) UpdatePresetQuery(id int32, body *models.EngineUpdatePresetQueryRequest, opts ...ClientOption) (*UpdatePresetQueryOK, error) {
	params := NewUpdatePresetQueryParams().WithBody(body).WithID(id)
	return a.UpdatePresetQueryWithParams(params, opts...)
}

func (a *Client) UpdatePresetQueryWithParams(params *UpdatePresetQueryParams, opts ...ClientOption) (*UpdatePresetQueryOK, error) {
	if params == nil {
		params = NewUpdatePresetQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdatePresetQuery",
		Method:             "PUT",
		PathPattern:        "/call_center/preset/query/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePresetQueryReader{formats: a.formats},
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
	success, ok := result.(*UpdatePresetQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePresetQueryDefault)
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
