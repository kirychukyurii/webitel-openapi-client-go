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

// SearchAgentInTeamReader is a Reader for the SearchAgentInTeam structure.
type SearchAgentInTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAgentInTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAgentInTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchAgentInTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchAgentInTeamOK creates a SearchAgentInTeamOK with default headers values
func NewSearchAgentInTeamOK() *SearchAgentInTeamOK {
	return &SearchAgentInTeamOK{}
}

/*
SearchAgentInTeamOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchAgentInTeamOK struct {
	Payload *models.EngineListAgentInTeam
}

// IsSuccess returns true when this search agent in team Ok response has a 2xx status code
func (o *SearchAgentInTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search agent in team Ok response has a 3xx status code
func (o *SearchAgentInTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search agent in team Ok response has a 4xx status code
func (o *SearchAgentInTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search agent in team Ok response has a 5xx status code
func (o *SearchAgentInTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search agent in team Ok response a status code equal to that given
func (o *SearchAgentInTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search agent in team Ok response
func (o *SearchAgentInTeamOK) Code() int {
	return 200
}

func (o *SearchAgentInTeamOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{id}/teams][%d] searchAgentInTeamOk %s", 200, payload)
}

func (o *SearchAgentInTeamOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{id}/teams][%d] searchAgentInTeamOk %s", 200, payload)
}

func (o *SearchAgentInTeamOK) GetPayload() *models.EngineListAgentInTeam {
	return o.Payload
}

func (o *SearchAgentInTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAgentInTeam)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAgentInTeamDefault creates a SearchAgentInTeamDefault with default headers values
func NewSearchAgentInTeamDefault(code int) *SearchAgentInTeamDefault {
	return &SearchAgentInTeamDefault{
		_statusCode: code,
	}
}

/*
SearchAgentInTeamDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchAgentInTeamDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search agent in team default response has a 2xx status code
func (o *SearchAgentInTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search agent in team default response has a 3xx status code
func (o *SearchAgentInTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search agent in team default response has a 4xx status code
func (o *SearchAgentInTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search agent in team default response has a 5xx status code
func (o *SearchAgentInTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search agent in team default response a status code equal to that given
func (o *SearchAgentInTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search agent in team default response
func (o *SearchAgentInTeamDefault) Code() int {
	return o._statusCode
}

func (o *SearchAgentInTeamDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{id}/teams][%d] SearchAgentInTeam default %s", o._statusCode, payload)
}

func (o *SearchAgentInTeamDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{id}/teams][%d] SearchAgentInTeam default %s", o._statusCode, payload)
}

func (o *SearchAgentInTeamDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchAgentInTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
