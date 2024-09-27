// Code generated by go-swagger; DO NOT EDIT.

package schema_variables_service

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

// ReadSchemaVariableReader is a Reader for the ReadSchemaVariable structure.
type ReadSchemaVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadSchemaVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadSchemaVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadSchemaVariableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadSchemaVariableOK creates a ReadSchemaVariableOK with default headers values
func NewReadSchemaVariableOK() *ReadSchemaVariableOK {
	return &ReadSchemaVariableOK{}
}

/*
ReadSchemaVariableOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadSchemaVariableOK struct {
	Payload *models.EngineSchemaVariable
}

// IsSuccess returns true when this read schema variable Ok response has a 2xx status code
func (o *ReadSchemaVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read schema variable Ok response has a 3xx status code
func (o *ReadSchemaVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read schema variable Ok response has a 4xx status code
func (o *ReadSchemaVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read schema variable Ok response has a 5xx status code
func (o *ReadSchemaVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read schema variable Ok response a status code equal to that given
func (o *ReadSchemaVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read schema variable Ok response
func (o *ReadSchemaVariableOK) Code() int {
	return 200
}

func (o *ReadSchemaVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/schema/variables/{id}][%d] readSchemaVariableOk %s", 200, payload)
}

func (o *ReadSchemaVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/schema/variables/{id}][%d] readSchemaVariableOk %s", 200, payload)
}

func (o *ReadSchemaVariableOK) GetPayload() *models.EngineSchemaVariable {
	return o.Payload
}

func (o *ReadSchemaVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineSchemaVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadSchemaVariableDefault creates a ReadSchemaVariableDefault with default headers values
func NewReadSchemaVariableDefault(code int) *ReadSchemaVariableDefault {
	return &ReadSchemaVariableDefault{
		_statusCode: code,
	}
}

/*
ReadSchemaVariableDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadSchemaVariableDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read schema variable default response has a 2xx status code
func (o *ReadSchemaVariableDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read schema variable default response has a 3xx status code
func (o *ReadSchemaVariableDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read schema variable default response has a 4xx status code
func (o *ReadSchemaVariableDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read schema variable default response has a 5xx status code
func (o *ReadSchemaVariableDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read schema variable default response a status code equal to that given
func (o *ReadSchemaVariableDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read schema variable default response
func (o *ReadSchemaVariableDefault) Code() int {
	return o._statusCode
}

func (o *ReadSchemaVariableDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/schema/variables/{id}][%d] ReadSchemaVariable default %s", o._statusCode, payload)
}

func (o *ReadSchemaVariableDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/schema/variables/{id}][%d] ReadSchemaVariable default %s", o._statusCode, payload)
}

func (o *ReadSchemaVariableDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadSchemaVariableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
