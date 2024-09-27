// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_federation

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

// OAuth2FederationCreateOAuthServiceReader is a Reader for the OAuth2FederationCreateOAuthService structure.
type OAuth2FederationCreateOAuthServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OAuth2FederationCreateOAuthServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOAuth2FederationCreateOAuthServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /oauth/apps] OAuth2Federation_CreateOAuthService", response, response.Code())
	}
}

// NewOAuth2FederationCreateOAuthServiceOK creates a OAuth2FederationCreateOAuthServiceOK with default headers values
func NewOAuth2FederationCreateOAuthServiceOK() *OAuth2FederationCreateOAuthServiceOK {
	return &OAuth2FederationCreateOAuthServiceOK{}
}

/*
OAuth2FederationCreateOAuthServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type OAuth2FederationCreateOAuthServiceOK struct {
	Payload *models.APIOAuthService
}

// IsSuccess returns true when this o auth2 federation create o auth service Ok response has a 2xx status code
func (o *OAuth2FederationCreateOAuthServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this o auth2 federation create o auth service Ok response has a 3xx status code
func (o *OAuth2FederationCreateOAuthServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this o auth2 federation create o auth service Ok response has a 4xx status code
func (o *OAuth2FederationCreateOAuthServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this o auth2 federation create o auth service Ok response has a 5xx status code
func (o *OAuth2FederationCreateOAuthServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this o auth2 federation create o auth service Ok response a status code equal to that given
func (o *OAuth2FederationCreateOAuthServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the o auth2 federation create o auth service Ok response
func (o *OAuth2FederationCreateOAuthServiceOK) Code() int {
	return 200
}

func (o *OAuth2FederationCreateOAuthServiceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /oauth/apps][%d] oAuth2FederationCreateOAuthServiceOk %s", 200, payload)
}

func (o *OAuth2FederationCreateOAuthServiceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /oauth/apps][%d] oAuth2FederationCreateOAuthServiceOk %s", 200, payload)
}

func (o *OAuth2FederationCreateOAuthServiceOK) GetPayload() *models.APIOAuthService {
	return o.Payload
}

func (o *OAuth2FederationCreateOAuthServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIOAuthService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
