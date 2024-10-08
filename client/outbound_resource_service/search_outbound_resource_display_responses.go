// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

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

// SearchOutboundResourceDisplayReader is a Reader for the SearchOutboundResourceDisplay structure.
type SearchOutboundResourceDisplayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchOutboundResourceDisplayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOutboundResourceDisplayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchOutboundResourceDisplayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchOutboundResourceDisplayOK creates a SearchOutboundResourceDisplayOK with default headers values
func NewSearchOutboundResourceDisplayOK() *SearchOutboundResourceDisplayOK {
	return &SearchOutboundResourceDisplayOK{}
}

/*
SearchOutboundResourceDisplayOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchOutboundResourceDisplayOK struct {
	Payload *models.EngineListOutboundResourceDisplay
}

// IsSuccess returns true when this search outbound resource display Ok response has a 2xx status code
func (o *SearchOutboundResourceDisplayOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search outbound resource display Ok response has a 3xx status code
func (o *SearchOutboundResourceDisplayOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search outbound resource display Ok response has a 4xx status code
func (o *SearchOutboundResourceDisplayOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search outbound resource display Ok response has a 5xx status code
func (o *SearchOutboundResourceDisplayOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search outbound resource display Ok response a status code equal to that given
func (o *SearchOutboundResourceDisplayOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search outbound resource display Ok response
func (o *SearchOutboundResourceDisplayOK) Code() int {
	return 200
}

func (o *SearchOutboundResourceDisplayOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resources/{resource_id}/display][%d] searchOutboundResourceDisplayOk %s", 200, payload)
}

func (o *SearchOutboundResourceDisplayOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resources/{resource_id}/display][%d] searchOutboundResourceDisplayOk %s", 200, payload)
}

func (o *SearchOutboundResourceDisplayOK) GetPayload() *models.EngineListOutboundResourceDisplay {
	return o.Payload
}

func (o *SearchOutboundResourceDisplayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListOutboundResourceDisplay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOutboundResourceDisplayDefault creates a SearchOutboundResourceDisplayDefault with default headers values
func NewSearchOutboundResourceDisplayDefault(code int) *SearchOutboundResourceDisplayDefault {
	return &SearchOutboundResourceDisplayDefault{
		_statusCode: code,
	}
}

/*
SearchOutboundResourceDisplayDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchOutboundResourceDisplayDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search outbound resource display default response has a 2xx status code
func (o *SearchOutboundResourceDisplayDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search outbound resource display default response has a 3xx status code
func (o *SearchOutboundResourceDisplayDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search outbound resource display default response has a 4xx status code
func (o *SearchOutboundResourceDisplayDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search outbound resource display default response has a 5xx status code
func (o *SearchOutboundResourceDisplayDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search outbound resource display default response a status code equal to that given
func (o *SearchOutboundResourceDisplayDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search outbound resource display default response
func (o *SearchOutboundResourceDisplayDefault) Code() int {
	return o._statusCode
}

func (o *SearchOutboundResourceDisplayDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resources/{resource_id}/display][%d] SearchOutboundResourceDisplay default %s", o._statusCode, payload)
}

func (o *SearchOutboundResourceDisplayDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resources/{resource_id}/display][%d] SearchOutboundResourceDisplay default %s", o._statusCode, payload)
}

func (o *SearchOutboundResourceDisplayDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchOutboundResourceDisplayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
