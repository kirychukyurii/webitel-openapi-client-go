// Code generated by go-swagger; DO NOT EDIT.

package users

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

// UsersDeleteUsersReader is a Reader for the UsersDeleteUsers structure.
type UsersDeleteUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersDeleteUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersDeleteUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /users/{id}] Users_DeleteUsers", response, response.Code())
	}
}

// NewUsersDeleteUsersOK creates a UsersDeleteUsersOK with default headers values
func NewUsersDeleteUsersOK() *UsersDeleteUsersOK {
	return &UsersDeleteUsersOK{}
}

/*
UsersDeleteUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type UsersDeleteUsersOK struct {
	Payload *models.APIDeleteUsersResponse
}

// IsSuccess returns true when this users delete users Ok response has a 2xx status code
func (o *UsersDeleteUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users delete users Ok response has a 3xx status code
func (o *UsersDeleteUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users delete users Ok response has a 4xx status code
func (o *UsersDeleteUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users delete users Ok response has a 5xx status code
func (o *UsersDeleteUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users delete users Ok response a status code equal to that given
func (o *UsersDeleteUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users delete users Ok response
func (o *UsersDeleteUsersOK) Code() int {
	return 200
}

func (o *UsersDeleteUsersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/{id}][%d] usersDeleteUsersOk %s", 200, payload)
}

func (o *UsersDeleteUsersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/{id}][%d] usersDeleteUsersOk %s", 200, payload)
}

func (o *UsersDeleteUsersOK) GetPayload() *models.APIDeleteUsersResponse {
	return o.Payload
}

func (o *UsersDeleteUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIDeleteUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
