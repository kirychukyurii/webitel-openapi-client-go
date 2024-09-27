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

// DevicesSearchDeviceAuditReader is a Reader for the DevicesSearchDeviceAudit structure.
type DevicesSearchDeviceAuditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesSearchDeviceAuditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesSearchDeviceAuditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{device.id}/users/audit] Devices_SearchDeviceAudit", response, response.Code())
	}
}

// NewDevicesSearchDeviceAuditOK creates a DevicesSearchDeviceAuditOK with default headers values
func NewDevicesSearchDeviceAuditOK() *DevicesSearchDeviceAuditOK {
	return &DevicesSearchDeviceAuditOK{}
}

/*
DevicesSearchDeviceAuditOK describes a response with status code 200, with default header values.

A successful response.
*/
type DevicesSearchDeviceAuditOK struct {
	Payload *models.APIDeviceAuditResponse
}

// IsSuccess returns true when this devices search device audit Ok response has a 2xx status code
func (o *DevicesSearchDeviceAuditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this devices search device audit Ok response has a 3xx status code
func (o *DevicesSearchDeviceAuditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this devices search device audit Ok response has a 4xx status code
func (o *DevicesSearchDeviceAuditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this devices search device audit Ok response has a 5xx status code
func (o *DevicesSearchDeviceAuditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this devices search device audit Ok response a status code equal to that given
func (o *DevicesSearchDeviceAuditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the devices search device audit Ok response
func (o *DevicesSearchDeviceAuditOK) Code() int {
	return 200
}

func (o *DevicesSearchDeviceAuditOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /devices/{device.id}/users/audit][%d] devicesSearchDeviceAuditOk %s", 200, payload)
}

func (o *DevicesSearchDeviceAuditOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /devices/{device.id}/users/audit][%d] devicesSearchDeviceAuditOk %s", 200, payload)
}

func (o *DevicesSearchDeviceAuditOK) GetPayload() *models.APIDeviceAuditResponse {
	return o.Payload
}

func (o *DevicesSearchDeviceAuditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIDeviceAuditResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
