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

// UsersUpdateUserReader is a Reader for the UsersUpdateUser structure.
type UsersUpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /users/{user.id}] Users_UpdateUser", response, response.Code())
	}
}

// NewUsersUpdateUserOK creates a UsersUpdateUserOK with default headers values
func NewUsersUpdateUserOK() *UsersUpdateUserOK {
	return &UsersUpdateUserOK{}
}

/*
UsersUpdateUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type UsersUpdateUserOK struct {
	Payload *models.APIUpdateUserResponse
}

// IsSuccess returns true when this users update user Ok response has a 2xx status code
func (o *UsersUpdateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users update user Ok response has a 3xx status code
func (o *UsersUpdateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users update user Ok response has a 4xx status code
func (o *UsersUpdateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users update user Ok response has a 5xx status code
func (o *UsersUpdateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users update user Ok response a status code equal to that given
func (o *UsersUpdateUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users update user Ok response
func (o *UsersUpdateUserOK) Code() int {
	return 200
}

func (o *UsersUpdateUserOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{user.id}][%d] usersUpdateUserOk %s", 200, payload)
}

func (o *UsersUpdateUserOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{user.id}][%d] usersUpdateUserOk %s", 200, payload)
}

func (o *UsersUpdateUserOK) GetPayload() *models.APIUpdateUserResponse {
	return o.Payload
}

func (o *UsersUpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUpdateUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
