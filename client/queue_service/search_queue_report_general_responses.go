// Code generated by go-swagger; DO NOT EDIT.

package queue_service

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

// SearchQueueReportGeneralReader is a Reader for the SearchQueueReportGeneral structure.
type SearchQueueReportGeneralReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchQueueReportGeneralReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchQueueReportGeneralOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchQueueReportGeneralDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchQueueReportGeneralOK creates a SearchQueueReportGeneralOK with default headers values
func NewSearchQueueReportGeneralOK() *SearchQueueReportGeneralOK {
	return &SearchQueueReportGeneralOK{}
}

/*
SearchQueueReportGeneralOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchQueueReportGeneralOK struct {
	Payload *models.EngineListReportGeneral
}

// IsSuccess returns true when this search queue report general Ok response has a 2xx status code
func (o *SearchQueueReportGeneralOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search queue report general Ok response has a 3xx status code
func (o *SearchQueueReportGeneralOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search queue report general Ok response has a 4xx status code
func (o *SearchQueueReportGeneralOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search queue report general Ok response has a 5xx status code
func (o *SearchQueueReportGeneralOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search queue report general Ok response a status code equal to that given
func (o *SearchQueueReportGeneralOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search queue report general Ok response
func (o *SearchQueueReportGeneralOK) Code() int {
	return 200
}

func (o *SearchQueueReportGeneralOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/reports/general][%d] searchQueueReportGeneralOk %s", 200, payload)
}

func (o *SearchQueueReportGeneralOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/reports/general][%d] searchQueueReportGeneralOk %s", 200, payload)
}

func (o *SearchQueueReportGeneralOK) GetPayload() *models.EngineListReportGeneral {
	return o.Payload
}

func (o *SearchQueueReportGeneralOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListReportGeneral)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchQueueReportGeneralDefault creates a SearchQueueReportGeneralDefault with default headers values
func NewSearchQueueReportGeneralDefault(code int) *SearchQueueReportGeneralDefault {
	return &SearchQueueReportGeneralDefault{
		_statusCode: code,
	}
}

/*
SearchQueueReportGeneralDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchQueueReportGeneralDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search queue report general default response has a 2xx status code
func (o *SearchQueueReportGeneralDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search queue report general default response has a 3xx status code
func (o *SearchQueueReportGeneralDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search queue report general default response has a 4xx status code
func (o *SearchQueueReportGeneralDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search queue report general default response has a 5xx status code
func (o *SearchQueueReportGeneralDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search queue report general default response a status code equal to that given
func (o *SearchQueueReportGeneralDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search queue report general default response
func (o *SearchQueueReportGeneralDefault) Code() int {
	return o._statusCode
}

func (o *SearchQueueReportGeneralDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/reports/general][%d] SearchQueueReportGeneral default %s", o._statusCode, payload)
}

func (o *SearchQueueReportGeneralDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/reports/general][%d] SearchQueueReportGeneral default %s", o._statusCode, payload)
}

func (o *SearchQueueReportGeneralDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchQueueReportGeneralDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
