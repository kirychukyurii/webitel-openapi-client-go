// Code generated by go-swagger; DO NOT EDIT.

package phones

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

// PhonesUpdatePhoneReader is a Reader for the PhonesUpdatePhone structure.
type PhonesUpdatePhoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhonesUpdatePhoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhonesUpdatePhoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /contacts/{contact_id}/phones/{etag}] Phones_UpdatePhone", response, response.Code())
	}
}

// NewPhonesUpdatePhoneOK creates a PhonesUpdatePhoneOK with default headers values
func NewPhonesUpdatePhoneOK() *PhonesUpdatePhoneOK {
	return &PhonesUpdatePhoneOK{}
}

/*
PhonesUpdatePhoneOK describes a response with status code 200, with default header values.

A successful response.
*/
type PhonesUpdatePhoneOK struct {
	Payload *models.WebitelContactsPhoneList
}

// IsSuccess returns true when this phones update phone Ok response has a 2xx status code
func (o *PhonesUpdatePhoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this phones update phone Ok response has a 3xx status code
func (o *PhonesUpdatePhoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this phones update phone Ok response has a 4xx status code
func (o *PhonesUpdatePhoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this phones update phone Ok response has a 5xx status code
func (o *PhonesUpdatePhoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this phones update phone Ok response a status code equal to that given
func (o *PhonesUpdatePhoneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the phones update phone Ok response
func (o *PhonesUpdatePhoneOK) Code() int {
	return 200
}

func (o *PhonesUpdatePhoneOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contacts/{contact_id}/phones/{etag}][%d] phonesUpdatePhoneOk %s", 200, payload)
}

func (o *PhonesUpdatePhoneOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contacts/{contact_id}/phones/{etag}][%d] phonesUpdatePhoneOk %s", 200, payload)
}

func (o *PhonesUpdatePhoneOK) GetPayload() *models.WebitelContactsPhoneList {
	return o.Payload
}

func (o *PhonesUpdatePhoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsPhoneList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
