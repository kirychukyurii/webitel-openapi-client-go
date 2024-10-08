// Code generated by go-swagger; DO NOT EDIT.

package timeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new timeline API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for timeline API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TimelineGetTimeline(params *TimelineGetTimelineParams, opts ...ClientOption) (*TimelineGetTimelineOK, error)

	TimelineGetTimelineCounter(contactID string, opts ...ClientOption) (*TimelineGetTimelineCounterOK, error)
	TimelineGetTimelineCounterWithParams(params *TimelineGetTimelineCounterParams, opts ...ClientOption) (*TimelineGetTimelineCounterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
TimelineGetTimeline timeline get timeline API
*/

func (a *Client) TimelineGetTimeline(params *TimelineGetTimelineParams, opts ...ClientOption) (*TimelineGetTimelineOK, error) {
	if params == nil {
		params = NewTimelineGetTimelineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Timeline_GetTimeline",
		Method:             "GET",
		PathPattern:        "/contacts/{contact_id}/timeline",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimelineGetTimelineReader{formats: a.formats},
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
	success, ok := result.(*TimelineGetTimelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Timeline_GetTimeline: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TimelineGetTimelineCounter timeline get timeline counter API
*/
func (a *Client) TimelineGetTimelineCounter(contactID string, opts ...ClientOption) (*TimelineGetTimelineCounterOK, error) {
	params := NewTimelineGetTimelineCounterParams().WithContactID(contactID)
	return a.TimelineGetTimelineCounterWithParams(params, opts...)
}

func (a *Client) TimelineGetTimelineCounterWithParams(params *TimelineGetTimelineCounterParams, opts ...ClientOption) (*TimelineGetTimelineCounterOK, error) {
	if params == nil {
		params = NewTimelineGetTimelineCounterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Timeline_GetTimelineCounter",
		Method:             "GET",
		PathPattern:        "/contacts/{contact_id}/timeline/counter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimelineGetTimelineCounterReader{formats: a.formats},
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
	success, ok := result.(*TimelineGetTimelineCounterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Timeline_GetTimelineCounter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
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
