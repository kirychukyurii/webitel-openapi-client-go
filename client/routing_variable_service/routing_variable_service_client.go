// Code generated by go-swagger; DO NOT EDIT.

package routing_variable_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new routing variable service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for routing variable service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRoutingVariable(body *models.EngineCreateRoutingVariableRequest, opts ...ClientOption) (*CreateRoutingVariableOK, error)
	CreateRoutingVariableWithParams(params *CreateRoutingVariableParams, opts ...ClientOption) (*CreateRoutingVariableOK, error)

	DeleteRoutingVariable(params *DeleteRoutingVariableParams, opts ...ClientOption) (*DeleteRoutingVariableOK, error)

	ReadRoutingVariable(params *ReadRoutingVariableParams, opts ...ClientOption) (*ReadRoutingVariableOK, error)

	SearchRoutingVariable(params *SearchRoutingVariableParams, opts ...ClientOption) (*SearchRoutingVariableOK, error)

	UpdateRoutingVariable(id string, body *models.EngineUpdateRoutingVariableRequest, opts ...ClientOption) (*UpdateRoutingVariableOK, error)
	UpdateRoutingVariableWithParams(params *UpdateRoutingVariableParams, opts ...ClientOption) (*UpdateRoutingVariableOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRoutingVariable creates routing variable
*/
func (a *Client) CreateRoutingVariable(body *models.EngineCreateRoutingVariableRequest, opts ...ClientOption) (*CreateRoutingVariableOK, error) {
	params := NewCreateRoutingVariableParams().WithBody(body)
	return a.CreateRoutingVariableWithParams(params, opts...)
}

func (a *Client) CreateRoutingVariableWithParams(params *CreateRoutingVariableParams, opts ...ClientOption) (*CreateRoutingVariableOK, error) {
	if params == nil {
		params = NewCreateRoutingVariableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRoutingVariable",
		Method:             "POST",
		PathPattern:        "/routing/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRoutingVariableReader{formats: a.formats},
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
	success, ok := result.(*CreateRoutingVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRoutingVariableDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteRoutingVariable removes routing variable
*/

func (a *Client) DeleteRoutingVariable(params *DeleteRoutingVariableParams, opts ...ClientOption) (*DeleteRoutingVariableOK, error) {
	if params == nil {
		params = NewDeleteRoutingVariableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRoutingVariable",
		Method:             "DELETE",
		PathPattern:        "/routing/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRoutingVariableReader{formats: a.formats},
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
	success, ok := result.(*DeleteRoutingVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRoutingVariableDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadRoutingVariable routings variable item
*/

func (a *Client) ReadRoutingVariable(params *ReadRoutingVariableParams, opts ...ClientOption) (*ReadRoutingVariableOK, error) {
	if params == nil {
		params = NewReadRoutingVariableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadRoutingVariable",
		Method:             "GET",
		PathPattern:        "/routing/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadRoutingVariableReader{formats: a.formats},
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
	success, ok := result.(*ReadRoutingVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadRoutingVariableDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchRoutingVariable lists of routing variable
*/

func (a *Client) SearchRoutingVariable(params *SearchRoutingVariableParams, opts ...ClientOption) (*SearchRoutingVariableOK, error) {
	if params == nil {
		params = NewSearchRoutingVariableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchRoutingVariable",
		Method:             "GET",
		PathPattern:        "/routing/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchRoutingVariableReader{formats: a.formats},
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
	success, ok := result.(*SearchRoutingVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchRoutingVariableDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateRoutingVariable updates routing variable
*/
func (a *Client) UpdateRoutingVariable(id string, body *models.EngineUpdateRoutingVariableRequest, opts ...ClientOption) (*UpdateRoutingVariableOK, error) {
	params := NewUpdateRoutingVariableParams().WithBody(body).WithID(id)
	return a.UpdateRoutingVariableWithParams(params, opts...)
}

func (a *Client) UpdateRoutingVariableWithParams(params *UpdateRoutingVariableParams, opts ...ClientOption) (*UpdateRoutingVariableOK, error) {
	if params == nil {
		params = NewUpdateRoutingVariableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRoutingVariable",
		Method:             "PUT",
		PathPattern:        "/routing/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRoutingVariableReader{formats: a.formats},
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
	success, ok := result.(*UpdateRoutingVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRoutingVariableDefault)
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
