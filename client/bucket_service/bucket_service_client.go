// Code generated by go-swagger; DO NOT EDIT.

package bucket_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new bucket service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bucket service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBucket(body *models.EngineCreateBucketRequest, opts ...ClientOption) (*CreateBucketOK, error)
	CreateBucketWithParams(params *CreateBucketParams, opts ...ClientOption) (*CreateBucketOK, error)

	DeleteBucket(params *DeleteBucketParams, opts ...ClientOption) (*DeleteBucketOK, error)

	ReadBucket(params *ReadBucketParams, opts ...ClientOption) (*ReadBucketOK, error)

	SearchBucket(params *SearchBucketParams, opts ...ClientOption) (*SearchBucketOK, error)

	UpdateBucket(id string, body *models.EngineUpdateBucketRequest, opts ...ClientOption) (*UpdateBucketOK, error)
	UpdateBucketWithParams(params *UpdateBucketParams, opts ...ClientOption) (*UpdateBucketOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateBucket creates bucket
*/
func (a *Client) CreateBucket(body *models.EngineCreateBucketRequest, opts ...ClientOption) (*CreateBucketOK, error) {
	params := NewCreateBucketParams().WithBody(body)
	return a.CreateBucketWithParams(params, opts...)
}

func (a *Client) CreateBucketWithParams(params *CreateBucketParams, opts ...ClientOption) (*CreateBucketOK, error) {
	if params == nil {
		params = NewCreateBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateBucket",
		Method:             "POST",
		PathPattern:        "/call_center/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBucketReader{formats: a.formats},
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
	success, ok := result.(*CreateBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteBucket removes bucket
*/

func (a *Client) DeleteBucket(params *DeleteBucketParams, opts ...ClientOption) (*DeleteBucketOK, error) {
	if params == nil {
		params = NewDeleteBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBucket",
		Method:             "DELETE",
		PathPattern:        "/call_center/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBucketReader{formats: a.formats},
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
	success, ok := result.(*DeleteBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadBucket buckets item
*/

func (a *Client) ReadBucket(params *ReadBucketParams, opts ...ClientOption) (*ReadBucketOK, error) {
	if params == nil {
		params = NewReadBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadBucket",
		Method:             "GET",
		PathPattern:        "/call_center/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadBucketReader{formats: a.formats},
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
	success, ok := result.(*ReadBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchBucket lists of bucket
*/

func (a *Client) SearchBucket(params *SearchBucketParams, opts ...ClientOption) (*SearchBucketOK, error) {
	if params == nil {
		params = NewSearchBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchBucket",
		Method:             "GET",
		PathPattern:        "/call_center/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchBucketReader{formats: a.formats},
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
	success, ok := result.(*SearchBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateBucket updates bucket
*/
func (a *Client) UpdateBucket(id string, body *models.EngineUpdateBucketRequest, opts ...ClientOption) (*UpdateBucketOK, error) {
	params := NewUpdateBucketParams().WithBody(body).WithID(id)
	return a.UpdateBucketWithParams(params, opts...)
}

func (a *Client) UpdateBucketWithParams(params *UpdateBucketParams, opts ...ClientOption) (*UpdateBucketOK, error) {
	if params == nil {
		params = NewUpdateBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateBucket",
		Method:             "PUT",
		PathPattern:        "/call_center/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBucketReader{formats: a.formats},
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
	success, ok := result.(*UpdateBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateBucketDefault)
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
