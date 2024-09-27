// Code generated by go-swagger; DO NOT EDIT.

package queue_resources_service

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

// UpdateQueueResourceGroupReader is a Reader for the UpdateQueueResourceGroup structure.
type UpdateQueueResourceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateQueueResourceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateQueueResourceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateQueueResourceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateQueueResourceGroupOK creates a UpdateQueueResourceGroupOK with default headers values
func NewUpdateQueueResourceGroupOK() *UpdateQueueResourceGroupOK {
	return &UpdateQueueResourceGroupOK{}
}

/*
UpdateQueueResourceGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateQueueResourceGroupOK struct {
	Payload *models.EngineQueueResourceGroup
}

// IsSuccess returns true when this update queue resource group Ok response has a 2xx status code
func (o *UpdateQueueResourceGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update queue resource group Ok response has a 3xx status code
func (o *UpdateQueueResourceGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update queue resource group Ok response has a 4xx status code
func (o *UpdateQueueResourceGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update queue resource group Ok response has a 5xx status code
func (o *UpdateQueueResourceGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update queue resource group Ok response a status code equal to that given
func (o *UpdateQueueResourceGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update queue resource group Ok response
func (o *UpdateQueueResourceGroupOK) Code() int {
	return 200
}

func (o *UpdateQueueResourceGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/resource_groups/{id}][%d] updateQueueResourceGroupOk %s", 200, payload)
}

func (o *UpdateQueueResourceGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/resource_groups/{id}][%d] updateQueueResourceGroupOk %s", 200, payload)
}

func (o *UpdateQueueResourceGroupOK) GetPayload() *models.EngineQueueResourceGroup {
	return o.Payload
}

func (o *UpdateQueueResourceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueResourceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQueueResourceGroupDefault creates a UpdateQueueResourceGroupDefault with default headers values
func NewUpdateQueueResourceGroupDefault(code int) *UpdateQueueResourceGroupDefault {
	return &UpdateQueueResourceGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateQueueResourceGroupDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateQueueResourceGroupDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update queue resource group default response has a 2xx status code
func (o *UpdateQueueResourceGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update queue resource group default response has a 3xx status code
func (o *UpdateQueueResourceGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update queue resource group default response has a 4xx status code
func (o *UpdateQueueResourceGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update queue resource group default response has a 5xx status code
func (o *UpdateQueueResourceGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update queue resource group default response a status code equal to that given
func (o *UpdateQueueResourceGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update queue resource group default response
func (o *UpdateQueueResourceGroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateQueueResourceGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/resource_groups/{id}][%d] UpdateQueueResourceGroup default %s", o._statusCode, payload)
}

func (o *UpdateQueueResourceGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/queues/{queue_id}/resource_groups/{id}][%d] UpdateQueueResourceGroup default %s", o._statusCode, payload)
}

func (o *UpdateQueueResourceGroupDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateQueueResourceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
