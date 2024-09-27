// Code generated by go-swagger; DO NOT EDIT.

package user_access_tokens

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

// UserAccessTokensUpdateUserAccessToken2Reader is a Reader for the UserAccessTokensUpdateUserAccessToken2 structure.
type UserAccessTokensUpdateUserAccessToken2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAccessTokensUpdateUserAccessToken2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserAccessTokensUpdateUserAccessToken2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /users/{update.user.id}/tokens/{update.id}] UserAccessTokens_UpdateUserAccessToken2", response, response.Code())
	}
}

// NewUserAccessTokensUpdateUserAccessToken2OK creates a UserAccessTokensUpdateUserAccessToken2OK with default headers values
func NewUserAccessTokensUpdateUserAccessToken2OK() *UserAccessTokensUpdateUserAccessToken2OK {
	return &UserAccessTokensUpdateUserAccessToken2OK{}
}

/*
UserAccessTokensUpdateUserAccessToken2OK describes a response with status code 200, with default header values.

A successful response.
*/
type UserAccessTokensUpdateUserAccessToken2OK struct {
	Payload *models.APIUserAccessToken
}

// IsSuccess returns true when this user access tokens update user access token2 Ok response has a 2xx status code
func (o *UserAccessTokensUpdateUserAccessToken2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user access tokens update user access token2 Ok response has a 3xx status code
func (o *UserAccessTokensUpdateUserAccessToken2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user access tokens update user access token2 Ok response has a 4xx status code
func (o *UserAccessTokensUpdateUserAccessToken2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user access tokens update user access token2 Ok response has a 5xx status code
func (o *UserAccessTokensUpdateUserAccessToken2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this user access tokens update user access token2 Ok response a status code equal to that given
func (o *UserAccessTokensUpdateUserAccessToken2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user access tokens update user access token2 Ok response
func (o *UserAccessTokensUpdateUserAccessToken2OK) Code() int {
	return 200
}

func (o *UserAccessTokensUpdateUserAccessToken2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/{update.user.id}/tokens/{update.id}][%d] userAccessTokensUpdateUserAccessToken2Ok %s", 200, payload)
}

func (o *UserAccessTokensUpdateUserAccessToken2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/{update.user.id}/tokens/{update.id}][%d] userAccessTokensUpdateUserAccessToken2Ok %s", 200, payload)
}

func (o *UserAccessTokensUpdateUserAccessToken2OK) GetPayload() *models.APIUserAccessToken {
	return o.Payload
}

func (o *UserAccessTokensUpdateUserAccessToken2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUserAccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
