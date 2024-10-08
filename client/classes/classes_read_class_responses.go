// Code generated by go-swagger; DO NOT EDIT.

package classes

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

// ClassesReadClassReader is a Reader for the ClassesReadClass structure.
type ClassesReadClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClassesReadClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClassesReadClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /objclass/{id}] Classes_ReadClass", response, response.Code())
	}
}

// NewClassesReadClassOK creates a ClassesReadClassOK with default headers values
func NewClassesReadClassOK() *ClassesReadClassOK {
	return &ClassesReadClassOK{}
}

/*
ClassesReadClassOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClassesReadClassOK struct {
	Payload *models.APIReadClassResponse
}

// IsSuccess returns true when this classes read class Ok response has a 2xx status code
func (o *ClassesReadClassOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this classes read class Ok response has a 3xx status code
func (o *ClassesReadClassOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this classes read class Ok response has a 4xx status code
func (o *ClassesReadClassOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this classes read class Ok response has a 5xx status code
func (o *ClassesReadClassOK) IsServerError() bool {
	return false
}

// IsCode returns true when this classes read class Ok response a status code equal to that given
func (o *ClassesReadClassOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the classes read class Ok response
func (o *ClassesReadClassOK) Code() int {
	return 200
}

func (o *ClassesReadClassOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /objclass/{id}][%d] classesReadClassOk %s", 200, payload)
}

func (o *ClassesReadClassOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /objclass/{id}][%d] classesReadClassOk %s", 200, payload)
}

func (o *ClassesReadClassOK) GetPayload() *models.APIReadClassResponse {
	return o.Payload
}

func (o *ClassesReadClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIReadClassResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
