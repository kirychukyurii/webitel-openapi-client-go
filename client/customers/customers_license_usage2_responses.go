// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// CustomersLicenseUsage2Reader is a Reader for the CustomersLicenseUsage2 structure.
type CustomersLicenseUsage2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomersLicenseUsage2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomersLicenseUsage2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /license] Customers_LicenseUsage2", response, response.Code())
	}
}

// NewCustomersLicenseUsage2OK creates a CustomersLicenseUsage2OK with default headers values
func NewCustomersLicenseUsage2OK() *CustomersLicenseUsage2OK {
	return &CustomersLicenseUsage2OK{}
}

/*
CustomersLicenseUsage2OK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomersLicenseUsage2OK struct {
	Payload *models.APILicenseUsageResponse
}

// IsSuccess returns true when this customers license usage2 Ok response has a 2xx status code
func (o *CustomersLicenseUsage2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customers license usage2 Ok response has a 3xx status code
func (o *CustomersLicenseUsage2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customers license usage2 Ok response has a 4xx status code
func (o *CustomersLicenseUsage2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customers license usage2 Ok response has a 5xx status code
func (o *CustomersLicenseUsage2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this customers license usage2 Ok response a status code equal to that given
func (o *CustomersLicenseUsage2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customers license usage2 Ok response
func (o *CustomersLicenseUsage2OK) Code() int {
	return 200
}

func (o *CustomersLicenseUsage2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /license][%d] customersLicenseUsage2Ok %s", 200, payload)
}

func (o *CustomersLicenseUsage2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /license][%d] customersLicenseUsage2Ok %s", 200, payload)
}

func (o *CustomersLicenseUsage2OK) GetPayload() *models.APILicenseUsageResponse {
	return o.Payload
}

func (o *CustomersLicenseUsage2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILicenseUsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
