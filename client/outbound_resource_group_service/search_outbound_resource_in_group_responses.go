// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_group_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kirychukyurii/webitel-openapi-client-go/models"
)

// SearchOutboundResourceInGroupReader is a Reader for the SearchOutboundResourceInGroup structure.
type SearchOutboundResourceInGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchOutboundResourceInGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOutboundResourceInGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchOutboundResourceInGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchOutboundResourceInGroupOK creates a SearchOutboundResourceInGroupOK with default headers values
func NewSearchOutboundResourceInGroupOK() *SearchOutboundResourceInGroupOK {
	return &SearchOutboundResourceInGroupOK{}
}

/*
SearchOutboundResourceInGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchOutboundResourceInGroupOK struct {
	Payload *models.EngineListOutboundResourceInGroup
}

// IsSuccess returns true when this search outbound resource in group Ok response has a 2xx status code
func (o *SearchOutboundResourceInGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search outbound resource in group Ok response has a 3xx status code
func (o *SearchOutboundResourceInGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search outbound resource in group Ok response has a 4xx status code
func (o *SearchOutboundResourceInGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search outbound resource in group Ok response has a 5xx status code
func (o *SearchOutboundResourceInGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search outbound resource in group Ok response a status code equal to that given
func (o *SearchOutboundResourceInGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search outbound resource in group Ok response
func (o *SearchOutboundResourceInGroupOK) Code() int {
	return 200
}

func (o *SearchOutboundResourceInGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resource_group/{group_id}/resource][%d] searchOutboundResourceInGroupOk %s", 200, payload)
}

func (o *SearchOutboundResourceInGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resource_group/{group_id}/resource][%d] searchOutboundResourceInGroupOk %s", 200, payload)
}

func (o *SearchOutboundResourceInGroupOK) GetPayload() *models.EngineListOutboundResourceInGroup {
	return o.Payload
}

func (o *SearchOutboundResourceInGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListOutboundResourceInGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOutboundResourceInGroupDefault creates a SearchOutboundResourceInGroupDefault with default headers values
func NewSearchOutboundResourceInGroupDefault(code int) *SearchOutboundResourceInGroupDefault {
	return &SearchOutboundResourceInGroupDefault{
		_statusCode: code,
	}
}

/*
SearchOutboundResourceInGroupDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchOutboundResourceInGroupDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search outbound resource in group default response has a 2xx status code
func (o *SearchOutboundResourceInGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search outbound resource in group default response has a 3xx status code
func (o *SearchOutboundResourceInGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search outbound resource in group default response has a 4xx status code
func (o *SearchOutboundResourceInGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search outbound resource in group default response has a 5xx status code
func (o *SearchOutboundResourceInGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search outbound resource in group default response a status code equal to that given
func (o *SearchOutboundResourceInGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search outbound resource in group default response
func (o *SearchOutboundResourceInGroupDefault) Code() int {
	return o._statusCode
}

func (o *SearchOutboundResourceInGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resource_group/{group_id}/resource][%d] SearchOutboundResourceInGroup default %s", o._statusCode, payload)
}

func (o *SearchOutboundResourceInGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resource_group/{group_id}/resource][%d] SearchOutboundResourceInGroup default %s", o._statusCode, payload)
}

func (o *SearchOutboundResourceInGroupDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchOutboundResourceInGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
