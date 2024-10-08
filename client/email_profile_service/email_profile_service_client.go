// Code generated by go-swagger; DO NOT EDIT.

package email_profile_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new email profile service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for email profile service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEmailProfile(body *models.EngineCreateEmailProfileRequest, opts ...ClientOption) (*CreateEmailProfileOK, error)
	CreateEmailProfileWithParams(params *CreateEmailProfileParams, opts ...ClientOption) (*CreateEmailProfileOK, error)

	DeleteEmailProfile(id string, opts ...ClientOption) (*DeleteEmailProfileOK, error)
	DeleteEmailProfileWithParams(params *DeleteEmailProfileParams, opts ...ClientOption) (*DeleteEmailProfileOK, error)

	LoginEmailProfile(id int32, opts ...ClientOption) (*LoginEmailProfileOK, error)
	LoginEmailProfileWithParams(params *LoginEmailProfileParams, opts ...ClientOption) (*LoginEmailProfileOK, error)

	LogoutEmailProfile(id int32, opts ...ClientOption) (*LogoutEmailProfileOK, error)
	LogoutEmailProfileWithParams(params *LogoutEmailProfileParams, opts ...ClientOption) (*LogoutEmailProfileOK, error)

	PatchEmailProfile(id string, body *models.EnginePatchEmailProfileRequest, opts ...ClientOption) (*PatchEmailProfileOK, error)
	PatchEmailProfileWithParams(params *PatchEmailProfileParams, opts ...ClientOption) (*PatchEmailProfileOK, error)

	ReadEmailProfile(id string, opts ...ClientOption) (*ReadEmailProfileOK, error)
	ReadEmailProfileWithParams(params *ReadEmailProfileParams, opts ...ClientOption) (*ReadEmailProfileOK, error)

	SearchEmailProfile(params *SearchEmailProfileParams, opts ...ClientOption) (*SearchEmailProfileOK, error)

	TestEmailProfile(id int32, opts ...ClientOption) (*TestEmailProfileOK, error)
	TestEmailProfileWithParams(params *TestEmailProfileParams, opts ...ClientOption) (*TestEmailProfileOK, error)

	UpdateEmailProfile(id string, body *models.EngineUpdateEmailProfileRequest, opts ...ClientOption) (*UpdateEmailProfileOK, error)
	UpdateEmailProfileWithParams(params *UpdateEmailProfileParams, opts ...ClientOption) (*UpdateEmailProfileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateEmailProfile creates email profile
*/
func (a *Client) CreateEmailProfile(body *models.EngineCreateEmailProfileRequest, opts ...ClientOption) (*CreateEmailProfileOK, error) {
	params := NewCreateEmailProfileParams().WithBody(body)
	return a.CreateEmailProfileWithParams(params, opts...)
}

func (a *Client) CreateEmailProfileWithParams(params *CreateEmailProfileParams, opts ...ClientOption) (*CreateEmailProfileOK, error) {
	if params == nil {
		params = NewCreateEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateEmailProfile",
		Method:             "POST",
		PathPattern:        "/email/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*CreateEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateEmailProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteEmailProfile removes email profile
*/
func (a *Client) DeleteEmailProfile(id string, opts ...ClientOption) (*DeleteEmailProfileOK, error) {
	params := NewDeleteEmailProfileParams().WithID(id)
	return a.DeleteEmailProfileWithParams(params, opts...)
}

func (a *Client) DeleteEmailProfileWithParams(params *DeleteEmailProfileParams, opts ...ClientOption) (*DeleteEmailProfileOK, error) {
	if params == nil {
		params = NewDeleteEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEmailProfile",
		Method:             "DELETE",
		PathPattern:        "/email/profile/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*DeleteEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteEmailProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
LoginEmailProfile login email profile API
*/
func (a *Client) LoginEmailProfile(id int32, opts ...ClientOption) (*LoginEmailProfileOK, error) {
	params := NewLoginEmailProfileParams().WithID(id)
	return a.LoginEmailProfileWithParams(params, opts...)
}

func (a *Client) LoginEmailProfileWithParams(params *LoginEmailProfileParams, opts ...ClientOption) (*LoginEmailProfileOK, error) {
	if params == nil {
		params = NewLoginEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LoginEmailProfile",
		Method:             "GET",
		PathPattern:        "/email/profile/{id}/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LoginEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*LoginEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LoginEmailProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
LogoutEmailProfile logout email profile API
*/
func (a *Client) LogoutEmailProfile(id int32, opts ...ClientOption) (*LogoutEmailProfileOK, error) {
	params := NewLogoutEmailProfileParams().WithID(id)
	return a.LogoutEmailProfileWithParams(params, opts...)
}

func (a *Client) LogoutEmailProfileWithParams(params *LogoutEmailProfileParams, opts ...ClientOption) (*LogoutEmailProfileOK, error) {
	if params == nil {
		params = NewLogoutEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LogoutEmailProfile",
		Method:             "PATCH",
		PathPattern:        "/email/profile/{id}/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LogoutEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*LogoutEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LogoutEmailProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchEmailProfile patch email profile API
*/
func (a *Client) PatchEmailProfile(id string, body *models.EnginePatchEmailProfileRequest, opts ...ClientOption) (*PatchEmailProfileOK, error) {
	params := NewPatchEmailProfileParams().WithBody(body).WithID(id)
	return a.PatchEmailProfileWithParams(params, opts...)
}

func (a *Client) PatchEmailProfileWithParams(params *PatchEmailProfileParams, opts ...ClientOption) (*PatchEmailProfileOK, error) {
	if params == nil {
		params = NewPatchEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchEmailProfile",
		Method:             "PATCH",
		PathPattern:        "/email/profile/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*PatchEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchEmailProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadEmailProfile emails profile item
*/
func (a *Client) ReadEmailProfile(id string, opts ...ClientOption) (*ReadEmailProfileOK, error) {
	params := NewReadEmailProfileParams().WithID(id)
	return a.ReadEmailProfileWithParams(params, opts...)
}

func (a *Client) ReadEmailProfileWithParams(params *ReadEmailProfileParams, opts ...ClientOption) (*ReadEmailProfileOK, error) {
	if params == nil {
		params = NewReadEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadEmailProfile",
		Method:             "GET",
		PathPattern:        "/email/profile/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*ReadEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadEmailProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchEmailProfile searches email profile
*/

func (a *Client) SearchEmailProfile(params *SearchEmailProfileParams, opts ...ClientOption) (*SearchEmailProfileOK, error) {
	if params == nil {
		params = NewSearchEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchEmailProfile",
		Method:             "GET",
		PathPattern:        "/email/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*SearchEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchEmailProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TestEmailProfile emails profile check login
*/
func (a *Client) TestEmailProfile(id int32, opts ...ClientOption) (*TestEmailProfileOK, error) {
	params := NewTestEmailProfileParams().WithID(id)
	return a.TestEmailProfileWithParams(params, opts...)
}

func (a *Client) TestEmailProfileWithParams(params *TestEmailProfileParams, opts ...ClientOption) (*TestEmailProfileOK, error) {
	if params == nil {
		params = NewTestEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TestEmailProfile",
		Method:             "GET",
		PathPattern:        "/email/profile/{id}/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*TestEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TestEmailProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateEmailProfile updates email profile
*/
func (a *Client) UpdateEmailProfile(id string, body *models.EngineUpdateEmailProfileRequest, opts ...ClientOption) (*UpdateEmailProfileOK, error) {
	params := NewUpdateEmailProfileParams().WithBody(body).WithID(id)
	return a.UpdateEmailProfileWithParams(params, opts...)
}

func (a *Client) UpdateEmailProfileWithParams(params *UpdateEmailProfileParams, opts ...ClientOption) (*UpdateEmailProfileOK, error) {
	if params == nil {
		params = NewUpdateEmailProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateEmailProfile",
		Method:             "PUT",
		PathPattern:        "/email/profile/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEmailProfileReader{formats: a.formats},
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
	success, ok := result.(*UpdateEmailProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateEmailProfileDefault)
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
