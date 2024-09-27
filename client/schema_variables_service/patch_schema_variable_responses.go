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

// PatchSchemaVariableReader is a Reader for the PatchSchemaVariable structure.
type PatchSchemaVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSchemaVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchSchemaVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchSchemaVariableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSchemaVariableOK creates a PatchSchemaVariableOK with default headers values
func NewPatchSchemaVariableOK() *PatchSchemaVariableOK {
	return &PatchSchemaVariableOK{}
}

/*
PatchSchemaVariableOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchSchemaVariableOK struct {
	Payload *models.EngineSchemaVariable
}

// IsSuccess returns true when this patch schema variable Ok response has a 2xx status code
func (o *PatchSchemaVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch schema variable Ok response has a 3xx status code
func (o *PatchSchemaVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch schema variable Ok response has a 4xx status code
func (o *PatchSchemaVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch schema variable Ok response has a 5xx status code
func (o *PatchSchemaVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch schema variable Ok response a status code equal to that given
func (o *PatchSchemaVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch schema variable Ok response
func (o *PatchSchemaVariableOK) Code() int {
	return 200
}

func (o *PatchSchemaVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /routing/schema/variables/{id}][%d] patchSchemaVariableOk %s", 200, payload)
}

func (o *PatchSchemaVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /routing/schema/variables/{id}][%d] patchSchemaVariableOk %s", 200, payload)
}

func (o *PatchSchemaVariableOK) GetPayload() *models.EngineSchemaVariable {
	return o.Payload
}

func (o *PatchSchemaVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineSchemaVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSchemaVariableDefault creates a PatchSchemaVariableDefault with default headers values
func NewPatchSchemaVariableDefault(code int) *PatchSchemaVariableDefault {
	return &PatchSchemaVariableDefault{
		_statusCode: code,
	}
}

/*
PatchSchemaVariableDefault describes a response with status code -1, with default header values.

Server error
*/
type PatchSchemaVariableDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this patch schema variable default response has a 2xx status code
func (o *PatchSchemaVariableDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch schema variable default response has a 3xx status code
func (o *PatchSchemaVariableDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch schema variable default response has a 4xx status code
func (o *PatchSchemaVariableDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch schema variable default response has a 5xx status code
func (o *PatchSchemaVariableDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch schema variable default response a status code equal to that given
func (o *PatchSchemaVariableDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch schema variable default response
func (o *PatchSchemaVariableDefault) Code() int {
	return o._statusCode
}

func (o *PatchSchemaVariableDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /routing/schema/variables/{id}][%d] PatchSchemaVariable default %s", o._statusCode, payload)
}

func (o *PatchSchemaVariableDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /routing/schema/variables/{id}][%d] PatchSchemaVariable default %s", o._statusCode, payload)
}

func (o *PatchSchemaVariableDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *PatchSchemaVariableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
