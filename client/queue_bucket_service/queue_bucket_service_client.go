// Code generated by go-swagger; DO NOT EDIT.

package queue_bucket_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// New creates a new queue bucket service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for queue bucket service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateQueueBucket(queueID string, body *models.EngineCreateQueueBucketRequest, opts ...ClientOption) (*CreateQueueBucketOK, error)
	CreateQueueBucketWithParams(params *CreateQueueBucketParams, opts ...ClientOption) (*CreateQueueBucketOK, error)

	DeleteQueueBucket(queueID string, id string, opts ...ClientOption) (*DeleteQueueBucketOK, error)
	DeleteQueueBucketWithParams(params *DeleteQueueBucketParams, opts ...ClientOption) (*DeleteQueueBucketOK, error)

	PatchQueueBucket(params *PatchQueueBucketParams, opts ...ClientOption) (*PatchQueueBucketOK, error)

	ReadQueueBucket(queueID string, id string, opts ...ClientOption) (*ReadQueueBucketOK, error)
	ReadQueueBucketWithParams(params *ReadQueueBucketParams, opts ...ClientOption) (*ReadQueueBucketOK, error)

	SearchQueueBucket(params *SearchQueueBucketParams, opts ...ClientOption) (*SearchQueueBucketOK, error)

	UpdateQueueBucket(params *UpdateQueueBucketParams, opts ...ClientOption) (*UpdateQueueBucketOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateQueueBucket creates queue bucket
*/
func (a *Client) CreateQueueBucket(queueID string, body *models.EngineCreateQueueBucketRequest, opts ...ClientOption) (*CreateQueueBucketOK, error) {
	params := NewCreateQueueBucketParams().WithBody(body).WithQueueID(queueID)
	return a.CreateQueueBucketWithParams(params, opts...)
}

func (a *Client) CreateQueueBucketWithParams(params *CreateQueueBucketParams, opts ...ClientOption) (*CreateQueueBucketOK, error) {
	if params == nil {
		params = NewCreateQueueBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateQueueBucket",
		Method:             "POST",
		PathPattern:        "/call_center/queues/{queue_id}/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateQueueBucketReader{formats: a.formats},
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
	success, ok := result.(*CreateQueueBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateQueueBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteQueueBucket deletes queue routing
*/
func (a *Client) DeleteQueueBucket(queueID string, id string, opts ...ClientOption) (*DeleteQueueBucketOK, error) {
	params := NewDeleteQueueBucketParams().WithID(id).WithQueueID(queueID)
	return a.DeleteQueueBucketWithParams(params, opts...)
}

func (a *Client) DeleteQueueBucketWithParams(params *DeleteQueueBucketParams, opts ...ClientOption) (*DeleteQueueBucketOK, error) {
	if params == nil {
		params = NewDeleteQueueBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteQueueBucket",
		Method:             "DELETE",
		PathPattern:        "/call_center/queues/{queue_id}/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteQueueBucketReader{formats: a.formats},
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
	success, ok := result.(*DeleteQueueBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteQueueBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchQueueBucket patch queue bucket API
*/

func (a *Client) PatchQueueBucket(params *PatchQueueBucketParams, opts ...ClientOption) (*PatchQueueBucketOK, error) {
	if params == nil {
		params = NewPatchQueueBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchQueueBucket",
		Method:             "PATCH",
		PathPattern:        "/call_center/queues/{queue_id}/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchQueueBucketReader{formats: a.formats},
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
	success, ok := result.(*PatchQueueBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchQueueBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadQueueBucket reads queue routing
*/
func (a *Client) ReadQueueBucket(queueID string, id string, opts ...ClientOption) (*ReadQueueBucketOK, error) {
	params := NewReadQueueBucketParams().WithID(id).WithQueueID(queueID)
	return a.ReadQueueBucketWithParams(params, opts...)
}

func (a *Client) ReadQueueBucketWithParams(params *ReadQueueBucketParams, opts ...ClientOption) (*ReadQueueBucketOK, error) {
	if params == nil {
		params = NewReadQueueBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadQueueBucket",
		Method:             "GET",
		PathPattern:        "/call_center/queues/{queue_id}/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadQueueBucketReader{formats: a.formats},
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
	success, ok := result.(*ReadQueueBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadQueueBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchQueueBucket searches queue routing
*/

func (a *Client) SearchQueueBucket(params *SearchQueueBucketParams, opts ...ClientOption) (*SearchQueueBucketOK, error) {
	if params == nil {
		params = NewSearchQueueBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchQueueBucket",
		Method:             "GET",
		PathPattern:        "/call_center/queues/{queue_id}/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchQueueBucketReader{formats: a.formats},
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
	success, ok := result.(*SearchQueueBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchQueueBucketDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateQueueBucket updates queue bucket
*/

func (a *Client) UpdateQueueBucket(params *UpdateQueueBucketParams, opts ...ClientOption) (*UpdateQueueBucketOK, error) {
	if params == nil {
		params = NewUpdateQueueBucketParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateQueueBucket",
		Method:             "PUT",
		PathPattern:        "/call_center/queues/{queue_id}/buckets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateQueueBucketReader{formats: a.formats},
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
	success, ok := result.(*UpdateQueueBucketOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateQueueBucketDefault)
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
