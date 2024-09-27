// Code generated by go-swagger; DO NOT EDIT.

package queue_service

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

// ReadQueueReader is a Reader for the ReadQueue structure.
type ReadQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadQueueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadQueueOK creates a ReadQueueOK with default headers values
func NewReadQueueOK() *ReadQueueOK {
	return &ReadQueueOK{}
}

/*
ReadQueueOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadQueueOK struct {
	Payload *models.EngineQueue
}

// IsSuccess returns true when this read queue Ok response has a 2xx status code
func (o *ReadQueueOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read queue Ok response has a 3xx status code
func (o *ReadQueueOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read queue Ok response has a 4xx status code
func (o *ReadQueueOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read queue Ok response has a 5xx status code
func (o *ReadQueueOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read queue Ok response a status code equal to that given
func (o *ReadQueueOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read queue Ok response
func (o *ReadQueueOK) Code() int {
	return 200
}

func (o *ReadQueueOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{id}][%d] readQueueOk %s", 200, payload)
}

func (o *ReadQueueOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{id}][%d] readQueueOk %s", 200, payload)
}

func (o *ReadQueueOK) GetPayload() *models.EngineQueue {
	return o.Payload
}

func (o *ReadQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadQueueDefault creates a ReadQueueDefault with default headers values
func NewReadQueueDefault(code int) *ReadQueueDefault {
	return &ReadQueueDefault{
		_statusCode: code,
	}
}

/*
ReadQueueDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadQueueDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read queue default response has a 2xx status code
func (o *ReadQueueDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read queue default response has a 3xx status code
func (o *ReadQueueDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read queue default response has a 4xx status code
func (o *ReadQueueDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read queue default response has a 5xx status code
func (o *ReadQueueDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read queue default response a status code equal to that given
func (o *ReadQueueDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read queue default response
func (o *ReadQueueDefault) Code() int {
	return o._statusCode
}

func (o *ReadQueueDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{id}][%d] ReadQueue default %s", o._statusCode, payload)
}

func (o *ReadQueueDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{id}][%d] ReadQueue default %s", o._statusCode, payload)
}

func (o *ReadQueueDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadQueueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
