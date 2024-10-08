// Code generated by go-swagger; DO NOT EDIT.

package call_service

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

// SearchActiveCallReader is a Reader for the SearchActiveCall structure.
type SearchActiveCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchActiveCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchActiveCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchActiveCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchActiveCallOK creates a SearchActiveCallOK with default headers values
func NewSearchActiveCallOK() *SearchActiveCallOK {
	return &SearchActiveCallOK{}
}

/*
SearchActiveCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchActiveCallOK struct {
	Payload *models.EngineListCall
}

// IsSuccess returns true when this search active call Ok response has a 2xx status code
func (o *SearchActiveCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search active call Ok response has a 3xx status code
func (o *SearchActiveCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search active call Ok response has a 4xx status code
func (o *SearchActiveCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search active call Ok response has a 5xx status code
func (o *SearchActiveCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search active call Ok response a status code equal to that given
func (o *SearchActiveCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search active call Ok response
func (o *SearchActiveCallOK) Code() int {
	return 200
}

func (o *SearchActiveCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /calls/active][%d] searchActiveCallOk %s", 200, payload)
}

func (o *SearchActiveCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /calls/active][%d] searchActiveCallOk %s", 200, payload)
}

func (o *SearchActiveCallOK) GetPayload() *models.EngineListCall {
	return o.Payload
}

func (o *SearchActiveCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListCall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchActiveCallDefault creates a SearchActiveCallDefault with default headers values
func NewSearchActiveCallDefault(code int) *SearchActiveCallDefault {
	return &SearchActiveCallDefault{
		_statusCode: code,
	}
}

/*
SearchActiveCallDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchActiveCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search active call default response has a 2xx status code
func (o *SearchActiveCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search active call default response has a 3xx status code
func (o *SearchActiveCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search active call default response has a 4xx status code
func (o *SearchActiveCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search active call default response has a 5xx status code
func (o *SearchActiveCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search active call default response a status code equal to that given
func (o *SearchActiveCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search active call default response
func (o *SearchActiveCallDefault) Code() int {
	return o._statusCode
}

func (o *SearchActiveCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /calls/active][%d] SearchActiveCall default %s", o._statusCode, payload)
}

func (o *SearchActiveCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /calls/active][%d] SearchActiveCall default %s", o._statusCode, payload)
}

func (o *SearchActiveCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchActiveCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
