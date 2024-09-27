// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

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

// CreateOutboundResourceDisplayReader is a Reader for the CreateOutboundResourceDisplay structure.
type CreateOutboundResourceDisplayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOutboundResourceDisplayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOutboundResourceDisplayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateOutboundResourceDisplayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOutboundResourceDisplayOK creates a CreateOutboundResourceDisplayOK with default headers values
func NewCreateOutboundResourceDisplayOK() *CreateOutboundResourceDisplayOK {
	return &CreateOutboundResourceDisplayOK{}
}

/*
CreateOutboundResourceDisplayOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateOutboundResourceDisplayOK struct {
	Payload *models.EngineResourceDisplay
}

// IsSuccess returns true when this create outbound resource display Ok response has a 2xx status code
func (o *CreateOutboundResourceDisplayOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create outbound resource display Ok response has a 3xx status code
func (o *CreateOutboundResourceDisplayOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create outbound resource display Ok response has a 4xx status code
func (o *CreateOutboundResourceDisplayOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create outbound resource display Ok response has a 5xx status code
func (o *CreateOutboundResourceDisplayOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create outbound resource display Ok response a status code equal to that given
func (o *CreateOutboundResourceDisplayOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create outbound resource display Ok response
func (o *CreateOutboundResourceDisplayOK) Code() int {
	return 200
}

func (o *CreateOutboundResourceDisplayOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/resources/{resource_id}/display][%d] createOutboundResourceDisplayOk %s", 200, payload)
}

func (o *CreateOutboundResourceDisplayOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/resources/{resource_id}/display][%d] createOutboundResourceDisplayOk %s", 200, payload)
}

func (o *CreateOutboundResourceDisplayOK) GetPayload() *models.EngineResourceDisplay {
	return o.Payload
}

func (o *CreateOutboundResourceDisplayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineResourceDisplay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOutboundResourceDisplayDefault creates a CreateOutboundResourceDisplayDefault with default headers values
func NewCreateOutboundResourceDisplayDefault(code int) *CreateOutboundResourceDisplayDefault {
	return &CreateOutboundResourceDisplayDefault{
		_statusCode: code,
	}
}

/*
CreateOutboundResourceDisplayDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateOutboundResourceDisplayDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create outbound resource display default response has a 2xx status code
func (o *CreateOutboundResourceDisplayDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create outbound resource display default response has a 3xx status code
func (o *CreateOutboundResourceDisplayDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create outbound resource display default response has a 4xx status code
func (o *CreateOutboundResourceDisplayDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create outbound resource display default response has a 5xx status code
func (o *CreateOutboundResourceDisplayDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create outbound resource display default response a status code equal to that given
func (o *CreateOutboundResourceDisplayDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create outbound resource display default response
func (o *CreateOutboundResourceDisplayDefault) Code() int {
	return o._statusCode
}

func (o *CreateOutboundResourceDisplayDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/resources/{resource_id}/display][%d] CreateOutboundResourceDisplay default %s", o._statusCode, payload)
}

func (o *CreateOutboundResourceDisplayDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/resources/{resource_id}/display][%d] CreateOutboundResourceDisplay default %s", o._statusCode, payload)
}

func (o *CreateOutboundResourceDisplayDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateOutboundResourceDisplayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
