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

// PatchMemberReader is a Reader for the PatchMember structure.
type PatchMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchMemberOK creates a PatchMemberOK with default headers values
func NewPatchMemberOK() *PatchMemberOK {
	return &PatchMemberOK{}
}

/*
PatchMemberOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchMemberOK struct {
	Payload *models.EngineMemberInQueue
}

// IsSuccess returns true when this patch member Ok response has a 2xx status code
func (o *PatchMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch member Ok response has a 3xx status code
func (o *PatchMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch member Ok response has a 4xx status code
func (o *PatchMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch member Ok response has a 5xx status code
func (o *PatchMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch member Ok response a status code equal to that given
func (o *PatchMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch member Ok response
func (o *PatchMemberOK) Code() int {
	return 200
}

func (o *PatchMemberOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{id}][%d] patchMemberOk %s", 200, payload)
}

func (o *PatchMemberOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{id}][%d] patchMemberOk %s", 200, payload)
}

func (o *PatchMemberOK) GetPayload() *models.EngineMemberInQueue {
	return o.Payload
}

func (o *PatchMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineMemberInQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMemberDefault creates a PatchMemberDefault with default headers values
func NewPatchMemberDefault(code int) *PatchMemberDefault {
	return &PatchMemberDefault{
		_statusCode: code,
	}
}

/*
PatchMemberDefault describes a response with status code -1, with default header values.

Server error
*/
type PatchMemberDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this patch member default response has a 2xx status code
func (o *PatchMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch member default response has a 3xx status code
func (o *PatchMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch member default response has a 4xx status code
func (o *PatchMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch member default response has a 5xx status code
func (o *PatchMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch member default response a status code equal to that given
func (o *PatchMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch member default response
func (o *PatchMemberDefault) Code() int {
	return o._statusCode
}

func (o *PatchMemberDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{id}][%d] PatchMember default %s", o._statusCode, payload)
}

func (o *PatchMemberDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{id}][%d] PatchMember default %s", o._statusCode, payload)
}

func (o *PatchMemberDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *PatchMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
