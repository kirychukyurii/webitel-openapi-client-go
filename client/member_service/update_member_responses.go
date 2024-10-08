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

// UpdateMemberReader is a Reader for the UpdateMember structure.
type UpdateMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMemberOK creates a UpdateMemberOK with default headers values
func NewUpdateMemberOK() *UpdateMemberOK {
	return &UpdateMemberOK{}
}

/*
UpdateMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateMemberOK struct {
	Payload *models.EngineMemberInQueue
}

// IsSuccess returns true when this update member Ok response has a 2xx status code
func (o *UpdateMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update member Ok response has a 3xx status code
func (o *UpdateMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update member Ok response has a 4xx status code
func (o *UpdateMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update member Ok response has a 5xx status code
func (o *UpdateMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update member Ok response a status code equal to that given
func (o *UpdateMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update member Ok response
func (o *UpdateMemberOK) Code() int {
	return 200
}

func (o *UpdateMemberOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/members/{id}][%d] updateMemberOk %s", 200, payload)
}

func (o *UpdateMemberOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/members/{id}][%d] updateMemberOk %s", 200, payload)
}

func (o *UpdateMemberOK) GetPayload() *models.EngineMemberInQueue {
	return o.Payload
}

func (o *UpdateMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineMemberInQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMemberDefault creates a UpdateMemberDefault with default headers values
func NewUpdateMemberDefault(code int) *UpdateMemberDefault {
	return &UpdateMemberDefault{
		_statusCode: code,
	}
}

/*
UpdateMemberDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateMemberDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update member default response has a 2xx status code
func (o *UpdateMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update member default response has a 3xx status code
func (o *UpdateMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update member default response has a 4xx status code
func (o *UpdateMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update member default response has a 5xx status code
func (o *UpdateMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update member default response a status code equal to that given
func (o *UpdateMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update member default response
func (o *UpdateMemberDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMemberDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/members/{id}][%d] UpdateMember default %s", o._statusCode, payload)
}

func (o *UpdateMemberDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/members/{id}][%d] UpdateMember default %s", o._statusCode, payload)
}

func (o *UpdateMemberDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
