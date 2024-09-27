// Code generated by go-swagger; DO NOT EDIT.

package queue_skill_service

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

// UpdateQueueSkillReader is a Reader for the UpdateQueueSkill structure.
type UpdateQueueSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateQueueSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateQueueSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateQueueSkillDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateQueueSkillOK creates a UpdateQueueSkillOK with default headers values
func NewUpdateQueueSkillOK() *UpdateQueueSkillOK {
	return &UpdateQueueSkillOK{}
}

/*
UpdateQueueSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateQueueSkillOK struct {
	Payload *models.EngineQueueSkill
}

// IsSuccess returns true when this update queue skill Ok response has a 2xx status code
func (o *UpdateQueueSkillOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update queue skill Ok response has a 3xx status code
func (o *UpdateQueueSkillOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update queue skill Ok response has a 4xx status code
func (o *UpdateQueueSkillOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update queue skill Ok response has a 5xx status code
func (o *UpdateQueueSkillOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update queue skill Ok response a status code equal to that given
func (o *UpdateQueueSkillOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update queue skill Ok response
func (o *UpdateQueueSkillOK) Code() int {
	return 200
}

func (o *UpdateQueueSkillOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/skills/{id}][%d] updateQueueSkillOk %s", 200, payload)
}

func (o *UpdateQueueSkillOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/skills/{id}][%d] updateQueueSkillOk %s", 200, payload)
}

func (o *UpdateQueueSkillOK) GetPayload() *models.EngineQueueSkill {
	return o.Payload
}

func (o *UpdateQueueSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQueueSkillDefault creates a UpdateQueueSkillDefault with default headers values
func NewUpdateQueueSkillDefault(code int) *UpdateQueueSkillDefault {
	return &UpdateQueueSkillDefault{
		_statusCode: code,
	}
}

/*
UpdateQueueSkillDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateQueueSkillDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update queue skill default response has a 2xx status code
func (o *UpdateQueueSkillDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update queue skill default response has a 3xx status code
func (o *UpdateQueueSkillDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update queue skill default response has a 4xx status code
func (o *UpdateQueueSkillDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update queue skill default response has a 5xx status code
func (o *UpdateQueueSkillDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update queue skill default response a status code equal to that given
func (o *UpdateQueueSkillDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update queue skill default response
func (o *UpdateQueueSkillDefault) Code() int {
	return o._statusCode
}

func (o *UpdateQueueSkillDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/skills/{id}][%d] UpdateQueueSkill default %s", o._statusCode, payload)
}

func (o *UpdateQueueSkillDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/skills/{id}][%d] UpdateQueueSkill default %s", o._statusCode, payload)
}

func (o *UpdateQueueSkillDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateQueueSkillDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
