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

// DevicesReadDeviceReader is a Reader for the DevicesReadDevice structure.
type DevicesReadDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesReadDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesReadDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{id}] Devices_ReadDevice", response, response.Code())
	}
}

// NewDevicesReadDeviceOK creates a DevicesReadDeviceOK with default headers values
func NewDevicesReadDeviceOK() *DevicesReadDeviceOK {
	return &DevicesReadDeviceOK{}
}

/*
DevicesReadDeviceOK describes a response with status code 200, with default header values.

A successful response.
*/
type DevicesReadDeviceOK struct {
	Payload *models.APIReadDeviceResponse
}

// IsSuccess returns true when this devices read device Ok response has a 2xx status code
func (o *DevicesReadDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this devices read device Ok response has a 3xx status code
func (o *DevicesReadDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this devices read device Ok response has a 4xx status code
func (o *DevicesReadDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this devices read device Ok response has a 5xx status code
func (o *DevicesReadDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this devices read device Ok response a status code equal to that given
func (o *DevicesReadDeviceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the devices read device Ok response
func (o *DevicesReadDeviceOK) Code() int {
	return 200
}

func (o *DevicesReadDeviceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /devices/{id}][%d] devicesReadDeviceOk %s", 200, payload)
}

func (o *DevicesReadDeviceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /devices/{id}][%d] devicesReadDeviceOk %s", 200, payload)
}

func (o *DevicesReadDeviceOK) GetPayload() *models.APIReadDeviceResponse {
	return o.Payload
}

func (o *DevicesReadDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIReadDeviceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
