// Code generated by go-swagger; DO NOT EDIT.

package system_setting_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new system setting service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system setting service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSystemSetting(body *models.EngineCreateSystemSettingRequest, opts ...ClientOption) (*CreateSystemSettingOK, error)
	CreateSystemSettingWithParams(params *CreateSystemSettingParams, opts ...ClientOption) (*CreateSystemSettingOK, error)

	DeleteSystemSetting(id int32, opts ...ClientOption) (*DeleteSystemSettingOK, error)
	DeleteSystemSettingWithParams(params *DeleteSystemSettingParams, opts ...ClientOption) (*DeleteSystemSettingOK, error)

	PatchSystemSetting(id int32, body *models.EnginePatchSystemSettingRequest, opts ...ClientOption) (*PatchSystemSettingOK, error)
	PatchSystemSettingWithParams(params *PatchSystemSettingParams, opts ...ClientOption) (*PatchSystemSettingOK, error)

	ReadSystemSetting(id int32, opts ...ClientOption) (*ReadSystemSettingOK, error)
	ReadSystemSettingWithParams(params *ReadSystemSettingParams, opts ...ClientOption) (*ReadSystemSettingOK, error)

	SearchAvailableSystemSetting(params *SearchAvailableSystemSettingParams, opts ...ClientOption) (*SearchAvailableSystemSettingOK, error)

	SearchSystemSetting(params *SearchSystemSettingParams, opts ...ClientOption) (*SearchSystemSettingOK, error)

	UpdateSystemSetting(id int32, body *models.EngineUpdateSystemSettingRequest, opts ...ClientOption) (*UpdateSystemSettingOK, error)
	UpdateSystemSettingWithParams(params *UpdateSystemSettingParams, opts ...ClientOption) (*UpdateSystemSettingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSystemSetting create system setting API
*/
func (a *Client) CreateSystemSetting(body *models.EngineCreateSystemSettingRequest, opts ...ClientOption) (*CreateSystemSettingOK, error) {
	params := NewCreateSystemSettingParams().WithBody(body)
	return a.CreateSystemSettingWithParams(params, opts...)
}

func (a *Client) CreateSystemSettingWithParams(params *CreateSystemSettingParams, opts ...ClientOption) (*CreateSystemSettingOK, error) {
	if params == nil {
		params = NewCreateSystemSettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSystemSetting",
		Method:             "POST",
		PathPattern:        "/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSystemSettingReader{formats: a.formats},
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
	success, ok := result.(*CreateSystemSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSystemSettingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSystemSetting delete system setting API
*/
func (a *Client) DeleteSystemSetting(id int32, opts ...ClientOption) (*DeleteSystemSettingOK, error) {
	params := NewDeleteSystemSettingParams().WithID(id)
	return a.DeleteSystemSettingWithParams(params, opts...)
}

func (a *Client) DeleteSystemSettingWithParams(params *DeleteSystemSettingParams, opts ...ClientOption) (*DeleteSystemSettingOK, error) {
	if params == nil {
		params = NewDeleteSystemSettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSystemSetting",
		Method:             "DELETE",
		PathPattern:        "/settings/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSystemSettingReader{formats: a.formats},
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
	success, ok := result.(*DeleteSystemSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSystemSettingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchSystemSetting patch system setting API
*/
func (a *Client) PatchSystemSetting(id int32, body *models.EnginePatchSystemSettingRequest, opts ...ClientOption) (*PatchSystemSettingOK, error) {
	params := NewPatchSystemSettingParams().WithBody(body).WithID(id)
	return a.PatchSystemSettingWithParams(params, opts...)
}

func (a *Client) PatchSystemSettingWithParams(params *PatchSystemSettingParams, opts ...ClientOption) (*PatchSystemSettingOK, error) {
	if params == nil {
		params = NewPatchSystemSettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchSystemSetting",
		Method:             "PATCH",
		PathPattern:        "/settings/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchSystemSettingReader{formats: a.formats},
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
	success, ok := result.(*PatchSystemSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchSystemSettingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadSystemSetting read system setting API
*/
func (a *Client) ReadSystemSetting(id int32, opts ...ClientOption) (*ReadSystemSettingOK, error) {
	params := NewReadSystemSettingParams().WithID(id)
	return a.ReadSystemSettingWithParams(params, opts...)
}

func (a *Client) ReadSystemSettingWithParams(params *ReadSystemSettingParams, opts ...ClientOption) (*ReadSystemSettingOK, error) {
	if params == nil {
		params = NewReadSystemSettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadSystemSetting",
		Method:             "GET",
		PathPattern:        "/settings/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadSystemSettingReader{formats: a.formats},
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
	success, ok := result.(*ReadSystemSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadSystemSettingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchAvailableSystemSetting search available system setting API
*/

func (a *Client) SearchAvailableSystemSetting(params *SearchAvailableSystemSettingParams, opts ...ClientOption) (*SearchAvailableSystemSettingOK, error) {
	if params == nil {
		params = NewSearchAvailableSystemSettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchAvailableSystemSetting",
		Method:             "GET",
		PathPattern:        "/settings/available",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchAvailableSystemSettingReader{formats: a.formats},
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
	success, ok := result.(*SearchAvailableSystemSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchAvailableSystemSettingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchSystemSetting search system setting API
*/

func (a *Client) SearchSystemSetting(params *SearchSystemSettingParams, opts ...ClientOption) (*SearchSystemSettingOK, error) {
	if params == nil {
		params = NewSearchSystemSettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchSystemSetting",
		Method:             "GET",
		PathPattern:        "/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchSystemSettingReader{formats: a.formats},
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
	success, ok := result.(*SearchSystemSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchSystemSettingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSystemSetting update system setting API
*/
func (a *Client) UpdateSystemSetting(id int32, body *models.EngineUpdateSystemSettingRequest, opts ...ClientOption) (*UpdateSystemSettingOK, error) {
	params := NewUpdateSystemSettingParams().WithBody(body).WithID(id)
	return a.UpdateSystemSettingWithParams(params, opts...)
}

func (a *Client) UpdateSystemSettingWithParams(params *UpdateSystemSettingParams, opts ...ClientOption) (*UpdateSystemSettingOK, error) {
	if params == nil {
		params = NewUpdateSystemSettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSystemSetting",
		Method:             "PUT",
		PathPattern:        "/settings/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSystemSettingReader{formats: a.formats},
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
	success, ok := result.(*UpdateSystemSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSystemSettingDefault)
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
