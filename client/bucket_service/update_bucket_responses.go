// Code generated by go-swagger; DO NOT EDIT.

package bucket_service

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

// UpdateBucketReader is a Reader for the UpdateBucket structure.
type UpdateBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBucketOK creates a UpdateBucketOK with default headers values
func NewUpdateBucketOK() *UpdateBucketOK {
	return &UpdateBucketOK{}
}

/*
UpdateBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateBucketOK struct {
	Payload *models.EngineBucket
}

// IsSuccess returns true when this update bucket Ok response has a 2xx status code
func (o *UpdateBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update bucket Ok response has a 3xx status code
func (o *UpdateBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update bucket Ok response has a 4xx status code
func (o *UpdateBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update bucket Ok response has a 5xx status code
func (o *UpdateBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update bucket Ok response a status code equal to that given
func (o *UpdateBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update bucket Ok response
func (o *UpdateBucketOK) Code() int {
	return 200
}

func (o *UpdateBucketOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/buckets/{id}][%d] updateBucketOk %s", 200, payload)
}

func (o *UpdateBucketOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/buckets/{id}][%d] updateBucketOk %s", 200, payload)
}

func (o *UpdateBucketOK) GetPayload() *models.EngineBucket {
	return o.Payload
}

func (o *UpdateBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineBucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBucketDefault creates a UpdateBucketDefault with default headers values
func NewUpdateBucketDefault(code int) *UpdateBucketDefault {
	return &UpdateBucketDefault{
		_statusCode: code,
	}
}

/*
UpdateBucketDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateBucketDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update bucket default response has a 2xx status code
func (o *UpdateBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update bucket default response has a 3xx status code
func (o *UpdateBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update bucket default response has a 4xx status code
func (o *UpdateBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update bucket default response has a 5xx status code
func (o *UpdateBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update bucket default response a status code equal to that given
func (o *UpdateBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update bucket default response
func (o *UpdateBucketDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBucketDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/buckets/{id}][%d] UpdateBucket default %s", o._statusCode, payload)
}

func (o *UpdateBucketDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/buckets/{id}][%d] UpdateBucket default %s", o._statusCode, payload)
}

func (o *UpdateBucketDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
