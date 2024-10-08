// Code generated by go-swagger; DO NOT EDIT.

package communication_type_service

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

// DeleteCommunicationTypeReader is a Reader for the DeleteCommunicationType structure.
type DeleteCommunicationTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCommunicationTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCommunicationTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCommunicationTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCommunicationTypeOK creates a DeleteCommunicationTypeOK with default headers values
func NewDeleteCommunicationTypeOK() *DeleteCommunicationTypeOK {
	return &DeleteCommunicationTypeOK{}
}

/*
DeleteCommunicationTypeOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteCommunicationTypeOK struct {
	Payload *models.EngineCommunicationType
}

// IsSuccess returns true when this delete communication type Ok response has a 2xx status code
func (o *DeleteCommunicationTypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete communication type Ok response has a 3xx status code
func (o *DeleteCommunicationTypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete communication type Ok response has a 4xx status code
func (o *DeleteCommunicationTypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete communication type Ok response has a 5xx status code
func (o *DeleteCommunicationTypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete communication type Ok response a status code equal to that given
func (o *DeleteCommunicationTypeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete communication type Ok response
func (o *DeleteCommunicationTypeOK) Code() int {
	return 200
}

func (o *DeleteCommunicationTypeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/communication_type/{id}][%d] deleteCommunicationTypeOk %s", 200, payload)
}

func (o *DeleteCommunicationTypeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/communication_type/{id}][%d] deleteCommunicationTypeOk %s", 200, payload)
}

func (o *DeleteCommunicationTypeOK) GetPayload() *models.EngineCommunicationType {
	return o.Payload
}

func (o *DeleteCommunicationTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCommunicationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCommunicationTypeDefault creates a DeleteCommunicationTypeDefault with default headers values
func NewDeleteCommunicationTypeDefault(code int) *DeleteCommunicationTypeDefault {
	return &DeleteCommunicationTypeDefault{
		_statusCode: code,
	}
}

/*
DeleteCommunicationTypeDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteCommunicationTypeDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete communication type default response has a 2xx status code
func (o *DeleteCommunicationTypeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete communication type default response has a 3xx status code
func (o *DeleteCommunicationTypeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete communication type default response has a 4xx status code
func (o *DeleteCommunicationTypeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete communication type default response has a 5xx status code
func (o *DeleteCommunicationTypeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete communication type default response a status code equal to that given
func (o *DeleteCommunicationTypeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete communication type default response
func (o *DeleteCommunicationTypeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCommunicationTypeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/communication_type/{id}][%d] DeleteCommunicationType default %s", o._statusCode, payload)
}

func (o *DeleteCommunicationTypeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/communication_type/{id}][%d] DeleteCommunicationType default %s", o._statusCode, payload)
}

func (o *DeleteCommunicationTypeDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteCommunicationTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
