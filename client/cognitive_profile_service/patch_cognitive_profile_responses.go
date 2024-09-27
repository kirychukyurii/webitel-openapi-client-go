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

// PatchCognitiveProfileReader is a Reader for the PatchCognitiveProfile structure.
type PatchCognitiveProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCognitiveProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCognitiveProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /storage/cognitive_profiles/{id}] PatchCognitiveProfile", response, response.Code())
	}
}

// NewPatchCognitiveProfileOK creates a PatchCognitiveProfileOK with default headers values
func NewPatchCognitiveProfileOK() *PatchCognitiveProfileOK {
	return &PatchCognitiveProfileOK{}
}

/*
PatchCognitiveProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchCognitiveProfileOK struct {
	Payload *models.StorageCognitiveProfile
}

// IsSuccess returns true when this patch cognitive profile Ok response has a 2xx status code
func (o *PatchCognitiveProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch cognitive profile Ok response has a 3xx status code
func (o *PatchCognitiveProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch cognitive profile Ok response has a 4xx status code
func (o *PatchCognitiveProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch cognitive profile Ok response has a 5xx status code
func (o *PatchCognitiveProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch cognitive profile Ok response a status code equal to that given
func (o *PatchCognitiveProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch cognitive profile Ok response
func (o *PatchCognitiveProfileOK) Code() int {
	return 200
}

func (o *PatchCognitiveProfileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/cognitive_profiles/{id}][%d] patchCognitiveProfileOk %s", 200, payload)
}

func (o *PatchCognitiveProfileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/cognitive_profiles/{id}][%d] patchCognitiveProfileOk %s", 200, payload)
}

func (o *PatchCognitiveProfileOK) GetPayload() *models.StorageCognitiveProfile {
	return o.Payload
}

func (o *PatchCognitiveProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageCognitiveProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
