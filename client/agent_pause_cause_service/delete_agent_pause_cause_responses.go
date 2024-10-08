// Code generated by go-swagger; DO NOT EDIT.

package agent_pause_cause_service

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

// DeleteAgentPauseCauseReader is a Reader for the DeleteAgentPauseCause structure.
type DeleteAgentPauseCauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAgentPauseCauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAgentPauseCauseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAgentPauseCauseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAgentPauseCauseOK creates a DeleteAgentPauseCauseOK with default headers values
func NewDeleteAgentPauseCauseOK() *DeleteAgentPauseCauseOK {
	return &DeleteAgentPauseCauseOK{}
}

/*
DeleteAgentPauseCauseOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAgentPauseCauseOK struct {
	Payload *models.EngineAgentPauseCause
}

// IsSuccess returns true when this delete agent pause cause Ok response has a 2xx status code
func (o *DeleteAgentPauseCauseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete agent pause cause Ok response has a 3xx status code
func (o *DeleteAgentPauseCauseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete agent pause cause Ok response has a 4xx status code
func (o *DeleteAgentPauseCauseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete agent pause cause Ok response has a 5xx status code
func (o *DeleteAgentPauseCauseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete agent pause cause Ok response a status code equal to that given
func (o *DeleteAgentPauseCauseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete agent pause cause Ok response
func (o *DeleteAgentPauseCauseOK) Code() int {
	return 200
}

func (o *DeleteAgentPauseCauseOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/pause_causes/{id}][%d] deleteAgentPauseCauseOk %s", 200, payload)
}

func (o *DeleteAgentPauseCauseOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/pause_causes/{id}][%d] deleteAgentPauseCauseOk %s", 200, payload)
}

func (o *DeleteAgentPauseCauseOK) GetPayload() *models.EngineAgentPauseCause {
	return o.Payload
}

func (o *DeleteAgentPauseCauseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentPauseCause)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentPauseCauseDefault creates a DeleteAgentPauseCauseDefault with default headers values
func NewDeleteAgentPauseCauseDefault(code int) *DeleteAgentPauseCauseDefault {
	return &DeleteAgentPauseCauseDefault{
		_statusCode: code,
	}
}

/*
DeleteAgentPauseCauseDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteAgentPauseCauseDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete agent pause cause default response has a 2xx status code
func (o *DeleteAgentPauseCauseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete agent pause cause default response has a 3xx status code
func (o *DeleteAgentPauseCauseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete agent pause cause default response has a 4xx status code
func (o *DeleteAgentPauseCauseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete agent pause cause default response has a 5xx status code
func (o *DeleteAgentPauseCauseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete agent pause cause default response a status code equal to that given
func (o *DeleteAgentPauseCauseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete agent pause cause default response
func (o *DeleteAgentPauseCauseDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAgentPauseCauseDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/pause_causes/{id}][%d] DeleteAgentPauseCause default %s", o._statusCode, payload)
}

func (o *DeleteAgentPauseCauseDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/pause_causes/{id}][%d] DeleteAgentPauseCause default %s", o._statusCode, payload)
}

func (o *DeleteAgentPauseCauseDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteAgentPauseCauseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
