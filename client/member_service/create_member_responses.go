// Code generated by go-swagger; DO NOT EDIT.

package member_service

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

// CreateMemberReader is a Reader for the CreateMember structure.
type CreateMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMemberOK creates a CreateMemberOK with default headers values
func NewCreateMemberOK() *CreateMemberOK {
	return &CreateMemberOK{}
}

/*
CreateMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateMemberOK struct {
	Payload *models.EngineMemberInQueue
}

// IsSuccess returns true when this create member Ok response has a 2xx status code
func (o *CreateMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create member Ok response has a 3xx status code
func (o *CreateMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create member Ok response has a 4xx status code
func (o *CreateMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create member Ok response has a 5xx status code
func (o *CreateMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create member Ok response a status code equal to that given
func (o *CreateMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create member Ok response
func (o *CreateMemberOK) Code() int {
	return 200
}

func (o *CreateMemberOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/members][%d] createMemberOk %s", 200, payload)
}

func (o *CreateMemberOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/members][%d] createMemberOk %s", 200, payload)
}

func (o *CreateMemberOK) GetPayload() *models.EngineMemberInQueue {
	return o.Payload
}

func (o *CreateMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineMemberInQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMemberDefault creates a CreateMemberDefault with default headers values
func NewCreateMemberDefault(code int) *CreateMemberDefault {
	return &CreateMemberDefault{
		_statusCode: code,
	}
}

/*
CreateMemberDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateMemberDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create member default response has a 2xx status code
func (o *CreateMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create member default response has a 3xx status code
func (o *CreateMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create member default response has a 4xx status code
func (o *CreateMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create member default response has a 5xx status code
func (o *CreateMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create member default response a status code equal to that given
func (o *CreateMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create member default response
func (o *CreateMemberDefault) Code() int {
	return o._statusCode
}

func (o *CreateMemberDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/members][%d] CreateMember default %s", o._statusCode, payload)
}

func (o *CreateMemberDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/members][%d] CreateMember default %s", o._statusCode, payload)
}

func (o *CreateMemberDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
