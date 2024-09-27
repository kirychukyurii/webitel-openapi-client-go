// Code generated by go-swagger; DO NOT EDIT.

package managers

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

// ManagersDeleteManagerReader is a Reader for the ManagersDeleteManager structure.
type ManagersDeleteManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManagersDeleteManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManagersDeleteManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /contacts/{contact_id}/managers/{etag}] Managers_DeleteManager", response, response.Code())
	}
}

// NewManagersDeleteManagerOK creates a ManagersDeleteManagerOK with default headers values
func NewManagersDeleteManagerOK() *ManagersDeleteManagerOK {
	return &ManagersDeleteManagerOK{}
}

/*
ManagersDeleteManagerOK describes a response with status code 200, with default header values.

A successful response.
*/
type ManagersDeleteManagerOK struct {
	Payload *models.WebitelContactsManager
}

// IsSuccess returns true when this managers delete manager Ok response has a 2xx status code
func (o *ManagersDeleteManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this managers delete manager Ok response has a 3xx status code
func (o *ManagersDeleteManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this managers delete manager Ok response has a 4xx status code
func (o *ManagersDeleteManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this managers delete manager Ok response has a 5xx status code
func (o *ManagersDeleteManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this managers delete manager Ok response a status code equal to that given
func (o *ManagersDeleteManagerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the managers delete manager Ok response
func (o *ManagersDeleteManagerOK) Code() int {
	return 200
}

func (o *ManagersDeleteManagerOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /contacts/{contact_id}/managers/{etag}][%d] managersDeleteManagerOk %s", 200, payload)
}

func (o *ManagersDeleteManagerOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /contacts/{contact_id}/managers/{etag}][%d] managersDeleteManagerOk %s", 200, payload)
}

func (o *ManagersDeleteManagerOK) GetPayload() *models.WebitelContactsManager {
	return o.Payload
}

func (o *ManagersDeleteManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsManager)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
