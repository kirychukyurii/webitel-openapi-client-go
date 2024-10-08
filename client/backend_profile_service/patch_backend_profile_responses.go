// Code generated by go-swagger; DO NOT EDIT.

package backend_profile_service

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

// PatchBackendProfileReader is a Reader for the PatchBackendProfile structure.
type PatchBackendProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBackendProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchBackendProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /storage/backend_profiles/{id}] PatchBackendProfile", response, response.Code())
	}
}

// NewPatchBackendProfileOK creates a PatchBackendProfileOK with default headers values
func NewPatchBackendProfileOK() *PatchBackendProfileOK {
	return &PatchBackendProfileOK{}
}

/*
PatchBackendProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchBackendProfileOK struct {
	Payload *models.StorageBackendProfile
}

// IsSuccess returns true when this patch backend profile Ok response has a 2xx status code
func (o *PatchBackendProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch backend profile Ok response has a 3xx status code
func (o *PatchBackendProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch backend profile Ok response has a 4xx status code
func (o *PatchBackendProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch backend profile Ok response has a 5xx status code
func (o *PatchBackendProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch backend profile Ok response a status code equal to that given
func (o *PatchBackendProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch backend profile Ok response
func (o *PatchBackendProfileOK) Code() int {
	return 200
}

func (o *PatchBackendProfileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/backend_profiles/{id}][%d] patchBackendProfileOk %s", 200, payload)
}

func (o *PatchBackendProfileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/backend_profiles/{id}][%d] patchBackendProfileOk %s", 200, payload)
}

func (o *PatchBackendProfileOK) GetPayload() *models.StorageBackendProfile {
	return o.Payload
}

func (o *PatchBackendProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageBackendProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
