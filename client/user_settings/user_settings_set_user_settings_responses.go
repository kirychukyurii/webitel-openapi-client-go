// Code generated by go-swagger; DO NOT EDIT.

package user_settings

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

// UserSettingsSetUserSettingsReader is a Reader for the UserSettingsSetUserSettings structure.
type UserSettingsSetUserSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSettingsSetUserSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserSettingsSetUserSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /user/settings/{key}] UserSettings_SetUserSettings", response, response.Code())
	}
}

// NewUserSettingsSetUserSettingsOK creates a UserSettingsSetUserSettingsOK with default headers values
func NewUserSettingsSetUserSettingsOK() *UserSettingsSetUserSettingsOK {
	return &UserSettingsSetUserSettingsOK{}
}

/*
UserSettingsSetUserSettingsOK describes a response with status code 200, with default header values.

A successful response.
*/
type UserSettingsSetUserSettingsOK struct {
	Payload *models.APIUserSetting
}

// IsSuccess returns true when this user settings set user settings Ok response has a 2xx status code
func (o *UserSettingsSetUserSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user settings set user settings Ok response has a 3xx status code
func (o *UserSettingsSetUserSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user settings set user settings Ok response has a 4xx status code
func (o *UserSettingsSetUserSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user settings set user settings Ok response has a 5xx status code
func (o *UserSettingsSetUserSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user settings set user settings Ok response a status code equal to that given
func (o *UserSettingsSetUserSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user settings set user settings Ok response
func (o *UserSettingsSetUserSettingsOK) Code() int {
	return 200
}

func (o *UserSettingsSetUserSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user/settings/{key}][%d] userSettingsSetUserSettingsOk %s", 200, payload)
}

func (o *UserSettingsSetUserSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user/settings/{key}][%d] userSettingsSetUserSettingsOk %s", 200, payload)
}

func (o *UserSettingsSetUserSettingsOK) GetPayload() *models.APIUserSetting {
	return o.Payload
}

func (o *UserSettingsSetUserSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUserSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
