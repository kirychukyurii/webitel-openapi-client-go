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

// SearchAttemptsHistoryReader is a Reader for the SearchAttemptsHistory structure.
type SearchAttemptsHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAttemptsHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAttemptsHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchAttemptsHistoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchAttemptsHistoryOK creates a SearchAttemptsHistoryOK with default headers values
func NewSearchAttemptsHistoryOK() *SearchAttemptsHistoryOK {
	return &SearchAttemptsHistoryOK{}
}

/*
SearchAttemptsHistoryOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchAttemptsHistoryOK struct {
	Payload *models.EngineListHistoryAttempt
}

// IsSuccess returns true when this search attempts history Ok response has a 2xx status code
func (o *SearchAttemptsHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search attempts history Ok response has a 3xx status code
func (o *SearchAttemptsHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search attempts history Ok response has a 4xx status code
func (o *SearchAttemptsHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search attempts history Ok response has a 5xx status code
func (o *SearchAttemptsHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search attempts history Ok response a status code equal to that given
func (o *SearchAttemptsHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search attempts history Ok response
func (o *SearchAttemptsHistoryOK) Code() int {
	return 200
}

func (o *SearchAttemptsHistoryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/attempts/history][%d] searchAttemptsHistoryOk %s", 200, payload)
}

func (o *SearchAttemptsHistoryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/attempts/history][%d] searchAttemptsHistoryOk %s", 200, payload)
}

func (o *SearchAttemptsHistoryOK) GetPayload() *models.EngineListHistoryAttempt {
	return o.Payload
}

func (o *SearchAttemptsHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListHistoryAttempt)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAttemptsHistoryDefault creates a SearchAttemptsHistoryDefault with default headers values
func NewSearchAttemptsHistoryDefault(code int) *SearchAttemptsHistoryDefault {
	return &SearchAttemptsHistoryDefault{
		_statusCode: code,
	}
}

/*
SearchAttemptsHistoryDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchAttemptsHistoryDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search attempts history default response has a 2xx status code
func (o *SearchAttemptsHistoryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search attempts history default response has a 3xx status code
func (o *SearchAttemptsHistoryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search attempts history default response has a 4xx status code
func (o *SearchAttemptsHistoryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search attempts history default response has a 5xx status code
func (o *SearchAttemptsHistoryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search attempts history default response a status code equal to that given
func (o *SearchAttemptsHistoryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search attempts history default response
func (o *SearchAttemptsHistoryDefault) Code() int {
	return o._statusCode
}

func (o *SearchAttemptsHistoryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/attempts/history][%d] SearchAttemptsHistory default %s", o._statusCode, payload)
}

func (o *SearchAttemptsHistoryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/attempts/history][%d] SearchAttemptsHistory default %s", o._statusCode, payload)
}

func (o *SearchAttemptsHistoryDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchAttemptsHistoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
