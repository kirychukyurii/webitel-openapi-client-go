// Code generated by go-swagger; DO NOT EDIT.

package pause_template_service

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

// PauseTemplateServiceCreatePauseTemplateReader is a Reader for the PauseTemplateServiceCreatePauseTemplate structure.
type PauseTemplateServiceCreatePauseTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseTemplateServiceCreatePauseTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPauseTemplateServiceCreatePauseTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPauseTemplateServiceCreatePauseTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPauseTemplateServiceCreatePauseTemplateOK creates a PauseTemplateServiceCreatePauseTemplateOK with default headers values
func NewPauseTemplateServiceCreatePauseTemplateOK() *PauseTemplateServiceCreatePauseTemplateOK {
	return &PauseTemplateServiceCreatePauseTemplateOK{}
}

/*
PauseTemplateServiceCreatePauseTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type PauseTemplateServiceCreatePauseTemplateOK struct {
	Payload *models.WfmCreatePauseTemplateResponse
}

// IsSuccess returns true when this pause template service create pause template Ok response has a 2xx status code
func (o *PauseTemplateServiceCreatePauseTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pause template service create pause template Ok response has a 3xx status code
func (o *PauseTemplateServiceCreatePauseTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause template service create pause template Ok response has a 4xx status code
func (o *PauseTemplateServiceCreatePauseTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pause template service create pause template Ok response has a 5xx status code
func (o *PauseTemplateServiceCreatePauseTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pause template service create pause template Ok response a status code equal to that given
func (o *PauseTemplateServiceCreatePauseTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pause template service create pause template Ok response
func (o *PauseTemplateServiceCreatePauseTemplateOK) Code() int {
	return 200
}

func (o *PauseTemplateServiceCreatePauseTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /wfm/lookups/pause_templates][%d] pauseTemplateServiceCreatePauseTemplateOk %s", 200, payload)
}

func (o *PauseTemplateServiceCreatePauseTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /wfm/lookups/pause_templates][%d] pauseTemplateServiceCreatePauseTemplateOk %s", 200, payload)
}

func (o *PauseTemplateServiceCreatePauseTemplateOK) GetPayload() *models.WfmCreatePauseTemplateResponse {
	return o.Payload
}

func (o *PauseTemplateServiceCreatePauseTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmCreatePauseTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseTemplateServiceCreatePauseTemplateDefault creates a PauseTemplateServiceCreatePauseTemplateDefault with default headers values
func NewPauseTemplateServiceCreatePauseTemplateDefault(code int) *PauseTemplateServiceCreatePauseTemplateDefault {
	return &PauseTemplateServiceCreatePauseTemplateDefault{
		_statusCode: code,
	}
}

/*
PauseTemplateServiceCreatePauseTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PauseTemplateServiceCreatePauseTemplateDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this pause template service create pause template default response has a 2xx status code
func (o *PauseTemplateServiceCreatePauseTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pause template service create pause template default response has a 3xx status code
func (o *PauseTemplateServiceCreatePauseTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pause template service create pause template default response has a 4xx status code
func (o *PauseTemplateServiceCreatePauseTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pause template service create pause template default response has a 5xx status code
func (o *PauseTemplateServiceCreatePauseTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pause template service create pause template default response a status code equal to that given
func (o *PauseTemplateServiceCreatePauseTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the pause template service create pause template default response
func (o *PauseTemplateServiceCreatePauseTemplateDefault) Code() int {
	return o._statusCode
}

func (o *PauseTemplateServiceCreatePauseTemplateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /wfm/lookups/pause_templates][%d] PauseTemplateService_CreatePauseTemplate default %s", o._statusCode, payload)
}

func (o *PauseTemplateServiceCreatePauseTemplateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /wfm/lookups/pause_templates][%d] PauseTemplateService_CreatePauseTemplate default %s", o._statusCode, payload)
}

func (o *PauseTemplateServiceCreatePauseTemplateDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PauseTemplateServiceCreatePauseTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
