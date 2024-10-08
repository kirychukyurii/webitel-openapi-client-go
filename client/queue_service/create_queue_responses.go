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

// CreateQueueReader is a Reader for the CreateQueue structure.
type CreateQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateQueueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateQueueOK creates a CreateQueueOK with default headers values
func NewCreateQueueOK() *CreateQueueOK {
	return &CreateQueueOK{}
}

/*
CreateQueueOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateQueueOK struct {
	Payload *models.EngineQueue
}

// IsSuccess returns true when this create queue Ok response has a 2xx status code
func (o *CreateQueueOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create queue Ok response has a 3xx status code
func (o *CreateQueueOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create queue Ok response has a 4xx status code
func (o *CreateQueueOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create queue Ok response has a 5xx status code
func (o *CreateQueueOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create queue Ok response a status code equal to that given
func (o *CreateQueueOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create queue Ok response
func (o *CreateQueueOK) Code() int {
	return 200
}

func (o *CreateQueueOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues][%d] createQueueOk %s", 200, payload)
}

func (o *CreateQueueOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues][%d] createQueueOk %s", 200, payload)
}

func (o *CreateQueueOK) GetPayload() *models.EngineQueue {
	return o.Payload
}

func (o *CreateQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueueDefault creates a CreateQueueDefault with default headers values
func NewCreateQueueDefault(code int) *CreateQueueDefault {
	return &CreateQueueDefault{
		_statusCode: code,
	}
}

/*
CreateQueueDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateQueueDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create queue default response has a 2xx status code
func (o *CreateQueueDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create queue default response has a 3xx status code
func (o *CreateQueueDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create queue default response has a 4xx status code
func (o *CreateQueueDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create queue default response has a 5xx status code
func (o *CreateQueueDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create queue default response a status code equal to that given
func (o *CreateQueueDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create queue default response
func (o *CreateQueueDefault) Code() int {
	return o._statusCode
}

func (o *CreateQueueDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues][%d] CreateQueue default %s", o._statusCode, payload)
}

func (o *CreateQueueDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues][%d] CreateQueue default %s", o._statusCode, payload)
}

func (o *CreateQueueDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateQueueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
