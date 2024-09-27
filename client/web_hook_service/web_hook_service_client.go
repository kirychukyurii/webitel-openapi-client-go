// Code generated by go-swagger; DO NOT EDIT.

package web_hook_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new web hook service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for web hook service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateWebHook(body *models.EngineCreateWebHookRequest, opts ...ClientOption) (*CreateWebHookOK, error)
	CreateWebHookWithParams(params *CreateWebHookParams, opts ...ClientOption) (*CreateWebHookOK, error)

	DeleteWebHook(id int32, opts ...ClientOption) (*DeleteWebHookOK, error)
	DeleteWebHookWithParams(params *DeleteWebHookParams, opts ...ClientOption) (*DeleteWebHookOK, error)

	PatchWebHook(id int32, body *models.EnginePatchWebHookRequest, opts ...ClientOption) (*PatchWebHookOK, error)
	PatchWebHookWithParams(params *PatchWebHookParams, opts ...ClientOption) (*PatchWebHookOK, error)

	ReadWebHook(id int32, opts ...ClientOption) (*ReadWebHookOK, error)
	ReadWebHookWithParams(params *ReadWebHookParams, opts ...ClientOption) (*ReadWebHookOK, error)

	SearchWebHook(params *SearchWebHookParams, opts ...ClientOption) (*SearchWebHookOK, error)

	UpdateWebHook(id int32, body *models.EngineUpdateWebHookRequest, opts ...ClientOption) (*UpdateWebHookOK, error)
	UpdateWebHookWithParams(params *UpdateWebHookParams, opts ...ClientOption) (*UpdateWebHookOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateWebHook creates web hook
*/
func (a *Client) CreateWebHook(body *models.EngineCreateWebHookRequest, opts ...ClientOption) (*CreateWebHookOK, error) {
	params := NewCreateWebHookParams().WithBody(body)
	return a.CreateWebHookWithParams(params, opts...)
}

func (a *Client) CreateWebHookWithParams(params *CreateWebHookParams, opts ...ClientOption) (*CreateWebHookOK, error) {
	if params == nil {
		params = NewCreateWebHookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateWebHook",
		Method:             "POST",
		PathPattern:        "/hook",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateWebHookReader{formats: a.formats},
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
	success, ok := result.(*CreateWebHookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateWebHookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteWebHook removes web hook
*/
func (a *Client) DeleteWebHook(id int32, opts ...ClientOption) (*DeleteWebHookOK, error) {
	params := NewDeleteWebHookParams().WithID(id)
	return a.DeleteWebHookWithParams(params, opts...)
}

func (a *Client) DeleteWebHookWithParams(params *DeleteWebHookParams, opts ...ClientOption) (*DeleteWebHookOK, error) {
	if params == nil {
		params = NewDeleteWebHookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteWebHook",
		Method:             "DELETE",
		PathPattern:        "/hook/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWebHookReader{formats: a.formats},
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
	success, ok := result.(*DeleteWebHookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteWebHookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchWebHook patches web hook
*/
func (a *Client) PatchWebHook(id int32, body *models.EnginePatchWebHookRequest, opts ...ClientOption) (*PatchWebHookOK, error) {
	params := NewPatchWebHookParams().WithBody(body).WithID(id)
	return a.PatchWebHookWithParams(params, opts...)
}

func (a *Client) PatchWebHookWithParams(params *PatchWebHookParams, opts ...ClientOption) (*PatchWebHookOK, error) {
	if params == nil {
		params = NewPatchWebHookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchWebHook",
		Method:             "PATCH",
		PathPattern:        "/hook/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchWebHookReader{formats: a.formats},
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
	success, ok := result.(*PatchWebHookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchWebHookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadWebHook webs hook item
*/
func (a *Client) ReadWebHook(id int32, opts ...ClientOption) (*ReadWebHookOK, error) {
	params := NewReadWebHookParams().WithID(id)
	return a.ReadWebHookWithParams(params, opts...)
}

func (a *Client) ReadWebHookWithParams(params *ReadWebHookParams, opts ...ClientOption) (*ReadWebHookOK, error) {
	if params == nil {
		params = NewReadWebHookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadWebHook",
		Method:             "GET",
		PathPattern:        "/hook/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadWebHookReader{formats: a.formats},
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
	success, ok := result.(*ReadWebHookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadWebHookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchWebHook lists of web hook
*/

func (a *Client) SearchWebHook(params *SearchWebHookParams, opts ...ClientOption) (*SearchWebHookOK, error) {
	if params == nil {
		params = NewSearchWebHookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchWebHook",
		Method:             "GET",
		PathPattern:        "/hook",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchWebHookReader{formats: a.formats},
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
	success, ok := result.(*SearchWebHookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchWebHookDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateWebHook updates web hook
*/
func (a *Client) UpdateWebHook(id int32, body *models.EngineUpdateWebHookRequest, opts ...ClientOption) (*UpdateWebHookOK, error) {
	params := NewUpdateWebHookParams().WithBody(body).WithID(id)
	return a.UpdateWebHookWithParams(params, opts...)
}

func (a *Client) UpdateWebHookWithParams(params *UpdateWebHookParams, opts ...ClientOption) (*UpdateWebHookOK, error) {
	if params == nil {
		params = NewUpdateWebHookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateWebHook",
		Method:             "PUT",
		PathPattern:        "/hook/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWebHookReader{formats: a.formats},
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
	success, ok := result.(*UpdateWebHookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateWebHookDefault)
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
