// Code generated by go-swagger; DO NOT EDIT.

package domains

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

// DomainsSearchDomainsReader is a Reader for the DomainsSearchDomains structure.
type DomainsSearchDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainsSearchDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainsSearchDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /domains] Domains_SearchDomains", response, response.Code())
	}
}

// NewDomainsSearchDomainsOK creates a DomainsSearchDomainsOK with default headers values
func NewDomainsSearchDomainsOK() *DomainsSearchDomainsOK {
	return &DomainsSearchDomainsOK{}
}

/*
DomainsSearchDomainsOK describes a response with status code 200, with default header values.

A successful response.
*/
type DomainsSearchDomainsOK struct {
	Payload *models.APISearchDomainsResponse
}

// IsSuccess returns true when this domains search domains Ok response has a 2xx status code
func (o *DomainsSearchDomainsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domains search domains Ok response has a 3xx status code
func (o *DomainsSearchDomainsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domains search domains Ok response has a 4xx status code
func (o *DomainsSearchDomainsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domains search domains Ok response has a 5xx status code
func (o *DomainsSearchDomainsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domains search domains Ok response a status code equal to that given
func (o *DomainsSearchDomainsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domains search domains Ok response
func (o *DomainsSearchDomainsOK) Code() int {
	return 200
}

func (o *DomainsSearchDomainsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains][%d] domainsSearchDomainsOk %s", 200, payload)
}

func (o *DomainsSearchDomainsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains][%d] domainsSearchDomainsOk %s", 200, payload)
}

func (o *DomainsSearchDomainsOK) GetPayload() *models.APISearchDomainsResponse {
	return o.Payload
}

func (o *DomainsSearchDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APISearchDomainsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
