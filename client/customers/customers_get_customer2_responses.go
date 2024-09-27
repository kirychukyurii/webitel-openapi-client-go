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

// CustomersGetCustomer2Reader is a Reader for the CustomersGetCustomer2 structure.
type CustomersGetCustomer2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomersGetCustomer2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomersGetCustomer2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /customer/{id}] Customers_GetCustomer2", response, response.Code())
	}
}

// NewCustomersGetCustomer2OK creates a CustomersGetCustomer2OK with default headers values
func NewCustomersGetCustomer2OK() *CustomersGetCustomer2OK {
	return &CustomersGetCustomer2OK{}
}

/*
CustomersGetCustomer2OK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomersGetCustomer2OK struct {
	Payload *models.APIGetCustomerResponse
}

// IsSuccess returns true when this customers get customer2 Ok response has a 2xx status code
func (o *CustomersGetCustomer2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customers get customer2 Ok response has a 3xx status code
func (o *CustomersGetCustomer2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customers get customer2 Ok response has a 4xx status code
func (o *CustomersGetCustomer2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customers get customer2 Ok response has a 5xx status code
func (o *CustomersGetCustomer2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this customers get customer2 Ok response a status code equal to that given
func (o *CustomersGetCustomer2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customers get customer2 Ok response
func (o *CustomersGetCustomer2OK) Code() int {
	return 200
}

func (o *CustomersGetCustomer2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /customer/{id}][%d] customersGetCustomer2Ok %s", 200, payload)
}

func (o *CustomersGetCustomer2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /customer/{id}][%d] customersGetCustomer2Ok %s", 200, payload)
}

func (o *CustomersGetCustomer2OK) GetPayload() *models.APIGetCustomerResponse {
	return o.Payload
}

func (o *CustomersGetCustomer2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIGetCustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
