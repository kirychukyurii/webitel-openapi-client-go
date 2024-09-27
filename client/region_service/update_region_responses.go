// Code generated by go-swagger; DO NOT EDIT.

package region_service

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

// UpdateRegionReader is a Reader for the UpdateRegion structure.
type UpdateRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRegionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateRegionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRegionOK creates a UpdateRegionOK with default headers values
func NewUpdateRegionOK() *UpdateRegionOK {
	return &UpdateRegionOK{}
}

/*
UpdateRegionOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRegionOK struct {
	Payload *models.EngineRegion
}

// IsSuccess returns true when this update region Ok response has a 2xx status code
func (o *UpdateRegionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update region Ok response has a 3xx status code
func (o *UpdateRegionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update region Ok response has a 4xx status code
func (o *UpdateRegionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update region Ok response has a 5xx status code
func (o *UpdateRegionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update region Ok response a status code equal to that given
func (o *UpdateRegionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update region Ok response
func (o *UpdateRegionOK) Code() int {
	return 200
}

func (o *UpdateRegionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /regions/{id}][%d] updateRegionOk %s", 200, payload)
}

func (o *UpdateRegionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /regions/{id}][%d] updateRegionOk %s", 200, payload)
}

func (o *UpdateRegionOK) GetPayload() *models.EngineRegion {
	return o.Payload
}

func (o *UpdateRegionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRegionDefault creates a UpdateRegionDefault with default headers values
func NewUpdateRegionDefault(code int) *UpdateRegionDefault {
	return &UpdateRegionDefault{
		_statusCode: code,
	}
}

/*
UpdateRegionDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateRegionDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update region default response has a 2xx status code
func (o *UpdateRegionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update region default response has a 3xx status code
func (o *UpdateRegionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update region default response has a 4xx status code
func (o *UpdateRegionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update region default response has a 5xx status code
func (o *UpdateRegionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update region default response a status code equal to that given
func (o *UpdateRegionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update region default response
func (o *UpdateRegionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRegionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /regions/{id}][%d] UpdateRegion default %s", o._statusCode, payload)
}

func (o *UpdateRegionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /regions/{id}][%d] UpdateRegion default %s", o._statusCode, payload)
}

func (o *UpdateRegionDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateRegionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
