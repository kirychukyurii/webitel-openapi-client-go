// Code generated by go-swagger; DO NOT EDIT.

package call_service

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

// SearchHistoryCallReader is a Reader for the SearchHistoryCall structure.
type SearchHistoryCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchHistoryCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchHistoryCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchHistoryCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchHistoryCallOK creates a SearchHistoryCallOK with default headers values
func NewSearchHistoryCallOK() *SearchHistoryCallOK {
	return &SearchHistoryCallOK{}
}

/*
SearchHistoryCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchHistoryCallOK struct {
	Payload *models.EngineListHistoryCall
}

// IsSuccess returns true when this search history call Ok response has a 2xx status code
func (o *SearchHistoryCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search history call Ok response has a 3xx status code
func (o *SearchHistoryCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search history call Ok response has a 4xx status code
func (o *SearchHistoryCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search history call Ok response has a 5xx status code
func (o *SearchHistoryCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search history call Ok response a status code equal to that given
func (o *SearchHistoryCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search history call Ok response
func (o *SearchHistoryCallOK) Code() int {
	return 200
}

func (o *SearchHistoryCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /calls/history][%d] searchHistoryCallOk %s", 200, payload)
}

func (o *SearchHistoryCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /calls/history][%d] searchHistoryCallOk %s", 200, payload)
}

func (o *SearchHistoryCallOK) GetPayload() *models.EngineListHistoryCall {
	return o.Payload
}

func (o *SearchHistoryCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListHistoryCall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchHistoryCallDefault creates a SearchHistoryCallDefault with default headers values
func NewSearchHistoryCallDefault(code int) *SearchHistoryCallDefault {
	return &SearchHistoryCallDefault{
		_statusCode: code,
	}
}

/*
SearchHistoryCallDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchHistoryCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search history call default response has a 2xx status code
func (o *SearchHistoryCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search history call default response has a 3xx status code
func (o *SearchHistoryCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search history call default response has a 4xx status code
func (o *SearchHistoryCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search history call default response has a 5xx status code
func (o *SearchHistoryCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search history call default response a status code equal to that given
func (o *SearchHistoryCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search history call default response
func (o *SearchHistoryCallDefault) Code() int {
	return o._statusCode
}

func (o *SearchHistoryCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /calls/history][%d] SearchHistoryCall default %s", o._statusCode, payload)
}

func (o *SearchHistoryCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /calls/history][%d] SearchHistoryCall default %s", o._statusCode, payload)
}

func (o *SearchHistoryCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchHistoryCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
