// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// DevicesListRegistrationsReader is a Reader for the DevicesListRegistrations structure.
type DevicesListRegistrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesListRegistrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesListRegistrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{device.id}/registered] Devices_ListRegistrations", response, response.Code())
	}
}

// NewDevicesListRegistrationsOK creates a DevicesListRegistrationsOK with default headers values
func NewDevicesListRegistrationsOK() *DevicesListRegistrationsOK {
	return &DevicesListRegistrationsOK{}
}

/*
DevicesListRegistrationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type DevicesListRegistrationsOK struct {
	Payload *models.APIListRegistrationsResponse
}

// IsSuccess returns true when this devices list registrations Ok response has a 2xx status code
func (o *DevicesListRegistrationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this devices list registrations Ok response has a 3xx status code
func (o *DevicesListRegistrationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this devices list registrations Ok response has a 4xx status code
func (o *DevicesListRegistrationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this devices list registrations Ok response has a 5xx status code
func (o *DevicesListRegistrationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this devices list registrations Ok response a status code equal to that given
func (o *DevicesListRegistrationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the devices list registrations Ok response
func (o *DevicesListRegistrationsOK) Code() int {
	return 200
}

func (o *DevicesListRegistrationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /devices/{device.id}/registered][%d] devicesListRegistrationsOk %s", 200, payload)
}

func (o *DevicesListRegistrationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /devices/{device.id}/registered][%d] devicesListRegistrationsOk %s", 200, payload)
}

func (o *DevicesListRegistrationsOK) GetPayload() *models.APIListRegistrationsResponse {
	return o.Payload
}

func (o *DevicesListRegistrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIListRegistrationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
