// Code generated by go-swagger; DO NOT EDIT.

package agent_service

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

// SearchPauseCauseForAgentReader is a Reader for the SearchPauseCauseForAgent structure.
type SearchPauseCauseForAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPauseCauseForAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchPauseCauseForAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchPauseCauseForAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchPauseCauseForAgentOK creates a SearchPauseCauseForAgentOK with default headers values
func NewSearchPauseCauseForAgentOK() *SearchPauseCauseForAgentOK {
	return &SearchPauseCauseForAgentOK{}
}

/*
SearchPauseCauseForAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchPauseCauseForAgentOK struct {
	Payload *models.EngineForAgentPauseCauseList
}

// IsSuccess returns true when this search pause cause for agent Ok response has a 2xx status code
func (o *SearchPauseCauseForAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search pause cause for agent Ok response has a 3xx status code
func (o *SearchPauseCauseForAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search pause cause for agent Ok response has a 4xx status code
func (o *SearchPauseCauseForAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search pause cause for agent Ok response has a 5xx status code
func (o *SearchPauseCauseForAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search pause cause for agent Ok response a status code equal to that given
func (o *SearchPauseCauseForAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search pause cause for agent Ok response
func (o *SearchPauseCauseForAgentOK) Code() int {
	return 200
}

func (o *SearchPauseCauseForAgentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/pause_causes][%d] searchPauseCauseForAgentOk %s", 200, payload)
}

func (o *SearchPauseCauseForAgentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/pause_causes][%d] searchPauseCauseForAgentOk %s", 200, payload)
}

func (o *SearchPauseCauseForAgentOK) GetPayload() *models.EngineForAgentPauseCauseList {
	return o.Payload
}

func (o *SearchPauseCauseForAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineForAgentPauseCauseList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPauseCauseForAgentDefault creates a SearchPauseCauseForAgentDefault with default headers values
func NewSearchPauseCauseForAgentDefault(code int) *SearchPauseCauseForAgentDefault {
	return &SearchPauseCauseForAgentDefault{
		_statusCode: code,
	}
}

/*
SearchPauseCauseForAgentDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchPauseCauseForAgentDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search pause cause for agent default response has a 2xx status code
func (o *SearchPauseCauseForAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search pause cause for agent default response has a 3xx status code
func (o *SearchPauseCauseForAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search pause cause for agent default response has a 4xx status code
func (o *SearchPauseCauseForAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search pause cause for agent default response has a 5xx status code
func (o *SearchPauseCauseForAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search pause cause for agent default response a status code equal to that given
func (o *SearchPauseCauseForAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search pause cause for agent default response
func (o *SearchPauseCauseForAgentDefault) Code() int {
	return o._statusCode
}

func (o *SearchPauseCauseForAgentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/pause_causes][%d] SearchPauseCauseForAgent default %s", o._statusCode, payload)
}

func (o *SearchPauseCauseForAgentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/pause_causes][%d] SearchPauseCauseForAgent default %s", o._statusCode, payload)
}

func (o *SearchPauseCauseForAgentDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchPauseCauseForAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
