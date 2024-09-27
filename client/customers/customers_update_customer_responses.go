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

// CustomersUpdateCustomerReader is a Reader for the CustomersUpdateCustomer structure.
type CustomersUpdateCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomersUpdateCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomersUpdateCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /customer] Customers_UpdateCustomer", response, response.Code())
	}
}

// NewCustomersUpdateCustomerOK creates a CustomersUpdateCustomerOK with default headers values
func NewCustomersUpdateCustomerOK() *CustomersUpdateCustomerOK {
	return &CustomersUpdateCustomerOK{}
}

/*
CustomersUpdateCustomerOK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomersUpdateCustomerOK struct {
	Payload *models.APIUpdateCustomerResponse
}

// IsSuccess returns true when this customers update customer Ok response has a 2xx status code
func (o *CustomersUpdateCustomerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customers update customer Ok response has a 3xx status code
func (o *CustomersUpdateCustomerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customers update customer Ok response has a 4xx status code
func (o *CustomersUpdateCustomerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customers update customer Ok response has a 5xx status code
func (o *CustomersUpdateCustomerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customers update customer Ok response a status code equal to that given
func (o *CustomersUpdateCustomerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customers update customer Ok response
func (o *CustomersUpdateCustomerOK) Code() int {
	return 200
}

func (o *CustomersUpdateCustomerOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /customer][%d] customersUpdateCustomerOk %s", 200, payload)
}

func (o *CustomersUpdateCustomerOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /customer][%d] customersUpdateCustomerOk %s", 200, payload)
}

func (o *CustomersUpdateCustomerOK) GetPayload() *models.APIUpdateCustomerResponse {
	return o.Payload
}

func (o *CustomersUpdateCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUpdateCustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
