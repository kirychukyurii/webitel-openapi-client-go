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

// BlindTransferCallReader is a Reader for the BlindTransferCall structure.
type BlindTransferCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BlindTransferCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBlindTransferCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBlindTransferCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBlindTransferCallOK creates a BlindTransferCallOK with default headers values
func NewBlindTransferCallOK() *BlindTransferCallOK {
	return &BlindTransferCallOK{}
}

/*
BlindTransferCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type BlindTransferCallOK struct {
	Payload models.EngineBlindTransferCallResponse
}

// IsSuccess returns true when this blind transfer call Ok response has a 2xx status code
func (o *BlindTransferCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this blind transfer call Ok response has a 3xx status code
func (o *BlindTransferCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this blind transfer call Ok response has a 4xx status code
func (o *BlindTransferCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this blind transfer call Ok response has a 5xx status code
func (o *BlindTransferCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this blind transfer call Ok response a status code equal to that given
func (o *BlindTransferCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the blind transfer call Ok response
func (o *BlindTransferCallOK) Code() int {
	return 200
}

func (o *BlindTransferCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/transfer][%d] blindTransferCallOk %s", 200, payload)
}

func (o *BlindTransferCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/transfer][%d] blindTransferCallOk %s", 200, payload)
}

func (o *BlindTransferCallOK) GetPayload() models.EngineBlindTransferCallResponse {
	return o.Payload
}

func (o *BlindTransferCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBlindTransferCallDefault creates a BlindTransferCallDefault with default headers values
func NewBlindTransferCallDefault(code int) *BlindTransferCallDefault {
	return &BlindTransferCallDefault{
		_statusCode: code,
	}
}

/*
BlindTransferCallDefault describes a response with status code -1, with default header values.

Server error
*/
type BlindTransferCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this blind transfer call default response has a 2xx status code
func (o *BlindTransferCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this blind transfer call default response has a 3xx status code
func (o *BlindTransferCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this blind transfer call default response has a 4xx status code
func (o *BlindTransferCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this blind transfer call default response has a 5xx status code
func (o *BlindTransferCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this blind transfer call default response a status code equal to that given
func (o *BlindTransferCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the blind transfer call default response
func (o *BlindTransferCallDefault) Code() int {
	return o._statusCode
}

func (o *BlindTransferCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/transfer][%d] BlindTransferCall default %s", o._statusCode, payload)
}

func (o *BlindTransferCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/transfer][%d] BlindTransferCall default %s", o._statusCode, payload)
}

func (o *BlindTransferCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *BlindTransferCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
