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

// SearchAgentStatusStatisticItemReader is a Reader for the SearchAgentStatusStatisticItem structure.
type SearchAgentStatusStatisticItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAgentStatusStatisticItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAgentStatusStatisticItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchAgentStatusStatisticItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchAgentStatusStatisticItemOK creates a SearchAgentStatusStatisticItemOK with default headers values
func NewSearchAgentStatusStatisticItemOK() *SearchAgentStatusStatisticItemOK {
	return &SearchAgentStatusStatisticItemOK{}
}

/*
SearchAgentStatusStatisticItemOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchAgentStatusStatisticItemOK struct {
	Payload *models.EngineAgentStatusStatisticItem
}

// IsSuccess returns true when this search agent status statistic item Ok response has a 2xx status code
func (o *SearchAgentStatusStatisticItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search agent status statistic item Ok response has a 3xx status code
func (o *SearchAgentStatusStatisticItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search agent status statistic item Ok response has a 4xx status code
func (o *SearchAgentStatusStatisticItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search agent status statistic item Ok response has a 5xx status code
func (o *SearchAgentStatusStatisticItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search agent status statistic item Ok response a status code equal to that given
func (o *SearchAgentStatusStatisticItemOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search agent status statistic item Ok response
func (o *SearchAgentStatusStatisticItemOK) Code() int {
	return 200
}

func (o *SearchAgentStatusStatisticItemOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/reports/status/{agent_id}][%d] searchAgentStatusStatisticItemOk %s", 200, payload)
}

func (o *SearchAgentStatusStatisticItemOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/reports/status/{agent_id}][%d] searchAgentStatusStatisticItemOk %s", 200, payload)
}

func (o *SearchAgentStatusStatisticItemOK) GetPayload() *models.EngineAgentStatusStatisticItem {
	return o.Payload
}

func (o *SearchAgentStatusStatisticItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentStatusStatisticItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAgentStatusStatisticItemDefault creates a SearchAgentStatusStatisticItemDefault with default headers values
func NewSearchAgentStatusStatisticItemDefault(code int) *SearchAgentStatusStatisticItemDefault {
	return &SearchAgentStatusStatisticItemDefault{
		_statusCode: code,
	}
}

/*
SearchAgentStatusStatisticItemDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchAgentStatusStatisticItemDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search agent status statistic item default response has a 2xx status code
func (o *SearchAgentStatusStatisticItemDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search agent status statistic item default response has a 3xx status code
func (o *SearchAgentStatusStatisticItemDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search agent status statistic item default response has a 4xx status code
func (o *SearchAgentStatusStatisticItemDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search agent status statistic item default response has a 5xx status code
func (o *SearchAgentStatusStatisticItemDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search agent status statistic item default response a status code equal to that given
func (o *SearchAgentStatusStatisticItemDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search agent status statistic item default response
func (o *SearchAgentStatusStatisticItemDefault) Code() int {
	return o._statusCode
}

func (o *SearchAgentStatusStatisticItemDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/reports/status/{agent_id}][%d] SearchAgentStatusStatisticItem default %s", o._statusCode, payload)
}

func (o *SearchAgentStatusStatisticItemDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/reports/status/{agent_id}][%d] SearchAgentStatusStatisticItem default %s", o._statusCode, payload)
}

func (o *SearchAgentStatusStatisticItemDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchAgentStatusStatisticItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
