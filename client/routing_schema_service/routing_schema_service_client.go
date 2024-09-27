// Code generated by go-swagger; DO NOT EDIT.

package routing_schema_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new routing schema service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for routing schema service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRoutingSchema(body *models.EngineCreateRoutingSchemaRequest, opts ...ClientOption) (*CreateRoutingSchemaOK, error)
	CreateRoutingSchemaWithParams(params *CreateRoutingSchemaParams, opts ...ClientOption) (*CreateRoutingSchemaOK, error)

	DeleteRoutingSchema(params *DeleteRoutingSchemaParams, opts ...ClientOption) (*DeleteRoutingSchemaOK, error)

	PatchRoutingSchema(id string, body *models.EnginePatchRoutingSchemaRequest, opts ...ClientOption) (*PatchRoutingSchemaOK, error)
	PatchRoutingSchemaWithParams(params *PatchRoutingSchemaParams, opts ...ClientOption) (*PatchRoutingSchemaOK, error)

	ReadRoutingSchema(params *ReadRoutingSchemaParams, opts ...ClientOption) (*ReadRoutingSchemaOK, error)

	SearchRoutingSchema(params *SearchRoutingSchemaParams, opts ...ClientOption) (*SearchRoutingSchemaOK, error)

	SearchRoutingSchemaTags(params *SearchRoutingSchemaTagsParams, opts ...ClientOption) (*SearchRoutingSchemaTagsOK, error)

	UpdateRoutingSchema(id string, body *models.EngineUpdateRoutingSchemaRequest, opts ...ClientOption) (*UpdateRoutingSchemaOK, error)
	UpdateRoutingSchemaWithParams(params *UpdateRoutingSchemaParams, opts ...ClientOption) (*UpdateRoutingSchemaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRoutingSchema creates routing schema
*/
func (a *Client) CreateRoutingSchema(body *models.EngineCreateRoutingSchemaRequest, opts ...ClientOption) (*CreateRoutingSchemaOK, error) {
	params := NewCreateRoutingSchemaParams().WithBody(body)
	return a.CreateRoutingSchemaWithParams(params, opts...)
}

func (a *Client) CreateRoutingSchemaWithParams(params *CreateRoutingSchemaParams, opts ...ClientOption) (*CreateRoutingSchemaOK, error) {
	if params == nil {
		params = NewCreateRoutingSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRoutingSchema",
		Method:             "POST",
		PathPattern:        "/routing/schema",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRoutingSchemaReader{formats: a.formats},
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
	success, ok := result.(*CreateRoutingSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRoutingSchemaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteRoutingSchema removes routing schema
*/

func (a *Client) DeleteRoutingSchema(params *DeleteRoutingSchemaParams, opts ...ClientOption) (*DeleteRoutingSchemaOK, error) {
	if params == nil {
		params = NewDeleteRoutingSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRoutingSchema",
		Method:             "DELETE",
		PathPattern:        "/routing/schema/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRoutingSchemaReader{formats: a.formats},
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
	success, ok := result.(*DeleteRoutingSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRoutingSchemaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchRoutingSchema patches routing schema
*/
func (a *Client) PatchRoutingSchema(id string, body *models.EnginePatchRoutingSchemaRequest, opts ...ClientOption) (*PatchRoutingSchemaOK, error) {
	params := NewPatchRoutingSchemaParams().WithBody(body).WithID(id)
	return a.PatchRoutingSchemaWithParams(params, opts...)
}

func (a *Client) PatchRoutingSchemaWithParams(params *PatchRoutingSchemaParams, opts ...ClientOption) (*PatchRoutingSchemaOK, error) {
	if params == nil {
		params = NewPatchRoutingSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchRoutingSchema",
		Method:             "PATCH",
		PathPattern:        "/routing/schema/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchRoutingSchemaReader{formats: a.formats},
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
	success, ok := result.(*PatchRoutingSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchRoutingSchemaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadRoutingSchema routings schema item
*/

func (a *Client) ReadRoutingSchema(params *ReadRoutingSchemaParams, opts ...ClientOption) (*ReadRoutingSchemaOK, error) {
	if params == nil {
		params = NewReadRoutingSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadRoutingSchema",
		Method:             "GET",
		PathPattern:        "/routing/schema/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadRoutingSchemaReader{formats: a.formats},
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
	success, ok := result.(*ReadRoutingSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadRoutingSchemaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchRoutingSchema lists routing schema
*/

func (a *Client) SearchRoutingSchema(params *SearchRoutingSchemaParams, opts ...ClientOption) (*SearchRoutingSchemaOK, error) {
	if params == nil {
		params = NewSearchRoutingSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchRoutingSchema",
		Method:             "GET",
		PathPattern:        "/routing/schema",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchRoutingSchemaReader{formats: a.formats},
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
	success, ok := result.(*SearchRoutingSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchRoutingSchemaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchRoutingSchemaTags lists routing schema tags
*/

func (a *Client) SearchRoutingSchemaTags(params *SearchRoutingSchemaTagsParams, opts ...ClientOption) (*SearchRoutingSchemaTagsOK, error) {
	if params == nil {
		params = NewSearchRoutingSchemaTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchRoutingSchemaTags",
		Method:             "GET",
		PathPattern:        "/routing/schema/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchRoutingSchemaTagsReader{formats: a.formats},
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
	success, ok := result.(*SearchRoutingSchemaTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchRoutingSchemaTagsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateRoutingSchema updates routing schema
*/
func (a *Client) UpdateRoutingSchema(id string, body *models.EngineUpdateRoutingSchemaRequest, opts ...ClientOption) (*UpdateRoutingSchemaOK, error) {
	params := NewUpdateRoutingSchemaParams().WithBody(body).WithID(id)
	return a.UpdateRoutingSchemaWithParams(params, opts...)
}

func (a *Client) UpdateRoutingSchemaWithParams(params *UpdateRoutingSchemaParams, opts ...ClientOption) (*UpdateRoutingSchemaOK, error) {
	if params == nil {
		params = NewUpdateRoutingSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRoutingSchema",
		Method:             "PUT",
		PathPattern:        "/routing/schema/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRoutingSchemaReader{formats: a.formats},
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
	success, ok := result.(*UpdateRoutingSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRoutingSchemaDefault)
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
