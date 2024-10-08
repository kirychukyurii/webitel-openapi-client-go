// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// VariablesResetVariablesReader is a Reader for the VariablesResetVariables structure.
type VariablesResetVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VariablesResetVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVariablesResetVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /contacts/{contact_id}/variables] Variables_ResetVariables", response, response.Code())
	}
}

// NewVariablesResetVariablesOK creates a VariablesResetVariablesOK with default headers values
func NewVariablesResetVariablesOK() *VariablesResetVariablesOK {
	return &VariablesResetVariablesOK{}
}

/*
VariablesResetVariablesOK describes a response with status code 200, with default header values.

A successful response.
*/
type VariablesResetVariablesOK struct {
	Payload *models.WebitelContactsVariableList
}

// IsSuccess returns true when this variables reset variables Ok response has a 2xx status code
func (o *VariablesResetVariablesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this variables reset variables Ok response has a 3xx status code
func (o *VariablesResetVariablesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this variables reset variables Ok response has a 4xx status code
func (o *VariablesResetVariablesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this variables reset variables Ok response has a 5xx status code
func (o *VariablesResetVariablesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this variables reset variables Ok response a status code equal to that given
func (o *VariablesResetVariablesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the variables reset variables Ok response
func (o *VariablesResetVariablesOK) Code() int {
	return 200
}

func (o *VariablesResetVariablesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contacts/{contact_id}/variables][%d] variablesResetVariablesOk %s", 200, payload)
}

func (o *VariablesResetVariablesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contacts/{contact_id}/variables][%d] variablesResetVariablesOk %s", 200, payload)
}

func (o *VariablesResetVariablesOK) GetPayload() *models.WebitelContactsVariableList {
	return o.Payload
}

func (o *VariablesResetVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsVariableList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
