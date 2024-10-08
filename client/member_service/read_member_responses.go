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

// ReadMemberReader is a Reader for the ReadMember structure.
type ReadMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadMemberOK creates a ReadMemberOK with default headers values
func NewReadMemberOK() *ReadMemberOK {
	return &ReadMemberOK{}
}

/*
ReadMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadMemberOK struct {
	Payload *models.EngineMemberInQueue
}

// IsSuccess returns true when this read member Ok response has a 2xx status code
func (o *ReadMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read member Ok response has a 3xx status code
func (o *ReadMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read member Ok response has a 4xx status code
func (o *ReadMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read member Ok response has a 5xx status code
func (o *ReadMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read member Ok response a status code equal to that given
func (o *ReadMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read member Ok response
func (o *ReadMemberOK) Code() int {
	return 200
}

func (o *ReadMemberOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/members/{id}][%d] readMemberOk %s", 200, payload)
}

func (o *ReadMemberOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/members/{id}][%d] readMemberOk %s", 200, payload)
}

func (o *ReadMemberOK) GetPayload() *models.EngineMemberInQueue {
	return o.Payload
}

func (o *ReadMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineMemberInQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadMemberDefault creates a ReadMemberDefault with default headers values
func NewReadMemberDefault(code int) *ReadMemberDefault {
	return &ReadMemberDefault{
		_statusCode: code,
	}
}

/*
ReadMemberDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadMemberDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read member default response has a 2xx status code
func (o *ReadMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read member default response has a 3xx status code
func (o *ReadMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read member default response has a 4xx status code
func (o *ReadMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read member default response has a 5xx status code
func (o *ReadMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read member default response a status code equal to that given
func (o *ReadMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read member default response
func (o *ReadMemberDefault) Code() int {
	return o._statusCode
}

func (o *ReadMemberDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/members/{id}][%d] ReadMember default %s", o._statusCode, payload)
}

func (o *ReadMemberDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/members/{id}][%d] ReadMember default %s", o._statusCode, payload)
}

func (o *ReadMemberDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
