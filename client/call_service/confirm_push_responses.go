// Code generated by go-swagger; DO NOT EDIT.

package call_service

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

// ConfirmPushReader is a Reader for the ConfirmPush structure.
type ConfirmPushReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfirmPushReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfirmPushOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfirmPushDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfirmPushOK creates a ConfirmPushOK with default headers values
func NewConfirmPushOK() *ConfirmPushOK {
	return &ConfirmPushOK{}
}

/*
ConfirmPushOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConfirmPushOK struct {
	Payload models.EngineConfirmPushResponse
}

// IsSuccess returns true when this confirm push Ok response has a 2xx status code
func (o *ConfirmPushOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this confirm push Ok response has a 3xx status code
func (o *ConfirmPushOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm push Ok response has a 4xx status code
func (o *ConfirmPushOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this confirm push Ok response has a 5xx status code
func (o *ConfirmPushOK) IsServerError() bool {
	return false
}

// IsCode returns true when this confirm push Ok response a status code equal to that given
func (o *ConfirmPushOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the confirm push Ok response
func (o *ConfirmPushOK) Code() int {
	return 200
}

func (o *ConfirmPushOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/confirm_push][%d] confirmPushOk %s", 200, payload)
}

func (o *ConfirmPushOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/confirm_push][%d] confirmPushOk %s", 200, payload)
}

func (o *ConfirmPushOK) GetPayload() models.EngineConfirmPushResponse {
	return o.Payload
}

func (o *ConfirmPushOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmPushDefault creates a ConfirmPushDefault with default headers values
func NewConfirmPushDefault(code int) *ConfirmPushDefault {
	return &ConfirmPushDefault{
		_statusCode: code,
	}
}

/*
ConfirmPushDefault describes a response with status code -1, with default header values.

Server error
*/
type ConfirmPushDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this confirm push default response has a 2xx status code
func (o *ConfirmPushDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this confirm push default response has a 3xx status code
func (o *ConfirmPushDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this confirm push default response has a 4xx status code
func (o *ConfirmPushDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this confirm push default response has a 5xx status code
func (o *ConfirmPushDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this confirm push default response a status code equal to that given
func (o *ConfirmPushDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the confirm push default response
func (o *ConfirmPushDefault) Code() int {
	return o._statusCode
}

func (o *ConfirmPushDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/confirm_push][%d] ConfirmPush default %s", o._statusCode, payload)
}

func (o *ConfirmPushDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/confirm_push][%d] ConfirmPush default %s", o._statusCode, payload)
}

func (o *ConfirmPushDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ConfirmPushDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
