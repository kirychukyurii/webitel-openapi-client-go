// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// ContactsUpdateContactReader is a Reader for the ContactsUpdateContact structure.
type ContactsUpdateContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactsUpdateContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactsUpdateContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /contacts/{etag}] Contacts_UpdateContact", response, response.Code())
	}
}

// NewContactsUpdateContactOK creates a ContactsUpdateContactOK with default headers values
func NewContactsUpdateContactOK() *ContactsUpdateContactOK {
	return &ContactsUpdateContactOK{}
}

/*
ContactsUpdateContactOK describes a response with status code 200, with default header values.

A successful response.
*/
type ContactsUpdateContactOK struct {
	Payload *models.WebitelContactsContact
}

// IsSuccess returns true when this contacts update contact Ok response has a 2xx status code
func (o *ContactsUpdateContactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contacts update contact Ok response has a 3xx status code
func (o *ContactsUpdateContactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contacts update contact Ok response has a 4xx status code
func (o *ContactsUpdateContactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contacts update contact Ok response has a 5xx status code
func (o *ContactsUpdateContactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contacts update contact Ok response a status code equal to that given
func (o *ContactsUpdateContactOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contacts update contact Ok response
func (o *ContactsUpdateContactOK) Code() int {
	return 200
}

func (o *ContactsUpdateContactOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /contacts/{etag}][%d] contactsUpdateContactOk %s", 200, payload)
}

func (o *ContactsUpdateContactOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /contacts/{etag}][%d] contactsUpdateContactOk %s", 200, payload)
}

func (o *ContactsUpdateContactOK) GetPayload() *models.WebitelContactsContact {
	return o.Payload
}

func (o *ContactsUpdateContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
