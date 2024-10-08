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

// SearchAgentReader is a Reader for the SearchAgent structure.
type SearchAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchAgentOK creates a SearchAgentOK with default headers values
func NewSearchAgentOK() *SearchAgentOK {
	return &SearchAgentOK{}
}

/*
SearchAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchAgentOK struct {
	Payload *models.EngineListAgent
}

// IsSuccess returns true when this search agent Ok response has a 2xx status code
func (o *SearchAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search agent Ok response has a 3xx status code
func (o *SearchAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search agent Ok response has a 4xx status code
func (o *SearchAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search agent Ok response has a 5xx status code
func (o *SearchAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search agent Ok response a status code equal to that given
func (o *SearchAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search agent Ok response
func (o *SearchAgentOK) Code() int {
	return 200
}

func (o *SearchAgentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents][%d] searchAgentOk %s", 200, payload)
}

func (o *SearchAgentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents][%d] searchAgentOk %s", 200, payload)
}

func (o *SearchAgentOK) GetPayload() *models.EngineListAgent {
	return o.Payload
}

func (o *SearchAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAgentDefault creates a SearchAgentDefault with default headers values
func NewSearchAgentDefault(code int) *SearchAgentDefault {
	return &SearchAgentDefault{
		_statusCode: code,
	}
}

/*
SearchAgentDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchAgentDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search agent default response has a 2xx status code
func (o *SearchAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search agent default response has a 3xx status code
func (o *SearchAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search agent default response has a 4xx status code
func (o *SearchAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search agent default response has a 5xx status code
func (o *SearchAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search agent default response a status code equal to that given
func (o *SearchAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search agent default response
func (o *SearchAgentDefault) Code() int {
	return o._statusCode
}

func (o *SearchAgentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents][%d] SearchAgent default %s", o._statusCode, payload)
}

func (o *SearchAgentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents][%d] SearchAgent default %s", o._statusCode, payload)
}

func (o *SearchAgentDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
