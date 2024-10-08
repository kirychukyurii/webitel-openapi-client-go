// Code generated by go-swagger; DO NOT EDIT.

package agent_skill_service

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

// UpdateAgentSkillReader is a Reader for the UpdateAgentSkill structure.
type UpdateAgentSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAgentSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAgentSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAgentSkillDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAgentSkillOK creates a UpdateAgentSkillOK with default headers values
func NewUpdateAgentSkillOK() *UpdateAgentSkillOK {
	return &UpdateAgentSkillOK{}
}

/*
UpdateAgentSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateAgentSkillOK struct {
	Payload *models.EngineAgentSkill
}

// IsSuccess returns true when this update agent skill Ok response has a 2xx status code
func (o *UpdateAgentSkillOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update agent skill Ok response has a 3xx status code
func (o *UpdateAgentSkillOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update agent skill Ok response has a 4xx status code
func (o *UpdateAgentSkillOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update agent skill Ok response has a 5xx status code
func (o *UpdateAgentSkillOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update agent skill Ok response a status code equal to that given
func (o *UpdateAgentSkillOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update agent skill Ok response
func (o *UpdateAgentSkillOK) Code() int {
	return 200
}

func (o *UpdateAgentSkillOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/agents/{agent_id}/skills/{id}][%d] updateAgentSkillOk %s", 200, payload)
}

func (o *UpdateAgentSkillOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/agents/{agent_id}/skills/{id}][%d] updateAgentSkillOk %s", 200, payload)
}

func (o *UpdateAgentSkillOK) GetPayload() *models.EngineAgentSkill {
	return o.Payload
}

func (o *UpdateAgentSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAgentSkillDefault creates a UpdateAgentSkillDefault with default headers values
func NewUpdateAgentSkillDefault(code int) *UpdateAgentSkillDefault {
	return &UpdateAgentSkillDefault{
		_statusCode: code,
	}
}

/*
UpdateAgentSkillDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateAgentSkillDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update agent skill default response has a 2xx status code
func (o *UpdateAgentSkillDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update agent skill default response has a 3xx status code
func (o *UpdateAgentSkillDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update agent skill default response has a 4xx status code
func (o *UpdateAgentSkillDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update agent skill default response has a 5xx status code
func (o *UpdateAgentSkillDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update agent skill default response a status code equal to that given
func (o *UpdateAgentSkillDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update agent skill default response
func (o *UpdateAgentSkillDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAgentSkillDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/agents/{agent_id}/skills/{id}][%d] UpdateAgentSkill default %s", o._statusCode, payload)
}

func (o *UpdateAgentSkillDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/agents/{agent_id}/skills/{id}][%d] UpdateAgentSkill default %s", o._statusCode, payload)
}

func (o *UpdateAgentSkillDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateAgentSkillDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
