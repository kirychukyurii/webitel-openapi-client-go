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

// HangupCallReader is a Reader for the HangupCall structure.
type HangupCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HangupCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHangupCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHangupCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHangupCallOK creates a HangupCallOK with default headers values
func NewHangupCallOK() *HangupCallOK {
	return &HangupCallOK{}
}

/*
HangupCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type HangupCallOK struct {
	Payload models.EngineHangupCallResponse
}

// IsSuccess returns true when this hangup call Ok response has a 2xx status code
func (o *HangupCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hangup call Ok response has a 3xx status code
func (o *HangupCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hangup call Ok response has a 4xx status code
func (o *HangupCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hangup call Ok response has a 5xx status code
func (o *HangupCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hangup call Ok response a status code equal to that given
func (o *HangupCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hangup call Ok response
func (o *HangupCallOK) Code() int {
	return 200
}

func (o *HangupCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calls/active/{id}][%d] hangupCallOk %s", 200, payload)
}

func (o *HangupCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calls/active/{id}][%d] hangupCallOk %s", 200, payload)
}

func (o *HangupCallOK) GetPayload() models.EngineHangupCallResponse {
	return o.Payload
}

func (o *HangupCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHangupCallDefault creates a HangupCallDefault with default headers values
func NewHangupCallDefault(code int) *HangupCallDefault {
	return &HangupCallDefault{
		_statusCode: code,
	}
}

/*
HangupCallDefault describes a response with status code -1, with default header values.

Server error
*/
type HangupCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this hangup call default response has a 2xx status code
func (o *HangupCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hangup call default response has a 3xx status code
func (o *HangupCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hangup call default response has a 4xx status code
func (o *HangupCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hangup call default response has a 5xx status code
func (o *HangupCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hangup call default response a status code equal to that given
func (o *HangupCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hangup call default response
func (o *HangupCallDefault) Code() int {
	return o._statusCode
}

func (o *HangupCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calls/active/{id}][%d] HangupCall default %s", o._statusCode, payload)
}

func (o *HangupCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calls/active/{id}][%d] HangupCall default %s", o._statusCode, payload)
}

func (o *HangupCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *HangupCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
