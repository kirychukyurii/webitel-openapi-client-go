// Code generated by go-swagger; DO NOT EDIT.

package calendar_service

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

// DeleteCalendarReader is a Reader for the DeleteCalendar structure.
type DeleteCalendarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCalendarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCalendarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCalendarDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCalendarOK creates a DeleteCalendarOK with default headers values
func NewDeleteCalendarOK() *DeleteCalendarOK {
	return &DeleteCalendarOK{}
}

/*
DeleteCalendarOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteCalendarOK struct {
	Payload *models.EngineCalendar
}

// IsSuccess returns true when this delete calendar Ok response has a 2xx status code
func (o *DeleteCalendarOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete calendar Ok response has a 3xx status code
func (o *DeleteCalendarOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete calendar Ok response has a 4xx status code
func (o *DeleteCalendarOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete calendar Ok response has a 5xx status code
func (o *DeleteCalendarOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete calendar Ok response a status code equal to that given
func (o *DeleteCalendarOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete calendar Ok response
func (o *DeleteCalendarOK) Code() int {
	return 200
}

func (o *DeleteCalendarOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calendars/{id}][%d] deleteCalendarOk %s", 200, payload)
}

func (o *DeleteCalendarOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calendars/{id}][%d] deleteCalendarOk %s", 200, payload)
}

func (o *DeleteCalendarOK) GetPayload() *models.EngineCalendar {
	return o.Payload
}

func (o *DeleteCalendarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCalendar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCalendarDefault creates a DeleteCalendarDefault with default headers values
func NewDeleteCalendarDefault(code int) *DeleteCalendarDefault {
	return &DeleteCalendarDefault{
		_statusCode: code,
	}
}

/*
DeleteCalendarDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteCalendarDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete calendar default response has a 2xx status code
func (o *DeleteCalendarDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete calendar default response has a 3xx status code
func (o *DeleteCalendarDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete calendar default response has a 4xx status code
func (o *DeleteCalendarDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete calendar default response has a 5xx status code
func (o *DeleteCalendarDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete calendar default response a status code equal to that given
func (o *DeleteCalendarDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete calendar default response
func (o *DeleteCalendarDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCalendarDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calendars/{id}][%d] DeleteCalendar default %s", o._statusCode, payload)
}

func (o *DeleteCalendarDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calendars/{id}][%d] DeleteCalendar default %s", o._statusCode, payload)
}

func (o *DeleteCalendarDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteCalendarDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
