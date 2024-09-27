// Code generated by go-swagger; DO NOT EDIT.

package cognitive_profile_service

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

// CreateCognitiveProfileReader is a Reader for the CreateCognitiveProfile structure.
type CreateCognitiveProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCognitiveProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCognitiveProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /storage/cognitive_profiles] CreateCognitiveProfile", response, response.Code())
	}
}

// NewCreateCognitiveProfileOK creates a CreateCognitiveProfileOK with default headers values
func NewCreateCognitiveProfileOK() *CreateCognitiveProfileOK {
	return &CreateCognitiveProfileOK{}
}

/*
CreateCognitiveProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateCognitiveProfileOK struct {
	Payload *models.StorageCognitiveProfile
}

// IsSuccess returns true when this create cognitive profile Ok response has a 2xx status code
func (o *CreateCognitiveProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cognitive profile Ok response has a 3xx status code
func (o *CreateCognitiveProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cognitive profile Ok response has a 4xx status code
func (o *CreateCognitiveProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cognitive profile Ok response has a 5xx status code
func (o *CreateCognitiveProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create cognitive profile Ok response a status code equal to that given
func (o *CreateCognitiveProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create cognitive profile Ok response
func (o *CreateCognitiveProfileOK) Code() int {
	return 200
}

func (o *CreateCognitiveProfileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/cognitive_profiles][%d] createCognitiveProfileOk %s", 200, payload)
}

func (o *CreateCognitiveProfileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/cognitive_profiles][%d] createCognitiveProfileOk %s", 200, payload)
}

func (o *CreateCognitiveProfileOK) GetPayload() *models.StorageCognitiveProfile {
	return o.Payload
}

func (o *CreateCognitiveProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageCognitiveProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
