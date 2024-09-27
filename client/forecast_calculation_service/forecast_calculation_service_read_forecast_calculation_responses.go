// Code generated by go-swagger; DO NOT EDIT.

package forecast_calculation_service

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

// ForecastCalculationServiceReadForecastCalculationReader is a Reader for the ForecastCalculationServiceReadForecastCalculation structure.
type ForecastCalculationServiceReadForecastCalculationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForecastCalculationServiceReadForecastCalculationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForecastCalculationServiceReadForecastCalculationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewForecastCalculationServiceReadForecastCalculationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewForecastCalculationServiceReadForecastCalculationOK creates a ForecastCalculationServiceReadForecastCalculationOK with default headers values
func NewForecastCalculationServiceReadForecastCalculationOK() *ForecastCalculationServiceReadForecastCalculationOK {
	return &ForecastCalculationServiceReadForecastCalculationOK{}
}

/*
ForecastCalculationServiceReadForecastCalculationOK describes a response with status code 200, with default header values.

A successful response.
*/
type ForecastCalculationServiceReadForecastCalculationOK struct {
	Payload *models.WfmReadForecastCalculationResponse
}

// IsSuccess returns true when this forecast calculation service read forecast calculation Ok response has a 2xx status code
func (o *ForecastCalculationServiceReadForecastCalculationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this forecast calculation service read forecast calculation Ok response has a 3xx status code
func (o *ForecastCalculationServiceReadForecastCalculationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this forecast calculation service read forecast calculation Ok response has a 4xx status code
func (o *ForecastCalculationServiceReadForecastCalculationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this forecast calculation service read forecast calculation Ok response has a 5xx status code
func (o *ForecastCalculationServiceReadForecastCalculationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this forecast calculation service read forecast calculation Ok response a status code equal to that given
func (o *ForecastCalculationServiceReadForecastCalculationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the forecast calculation service read forecast calculation Ok response
func (o *ForecastCalculationServiceReadForecastCalculationOK) Code() int {
	return 200
}

func (o *ForecastCalculationServiceReadForecastCalculationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation/{id}][%d] forecastCalculationServiceReadForecastCalculationOk %s", 200, payload)
}

func (o *ForecastCalculationServiceReadForecastCalculationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation/{id}][%d] forecastCalculationServiceReadForecastCalculationOk %s", 200, payload)
}

func (o *ForecastCalculationServiceReadForecastCalculationOK) GetPayload() *models.WfmReadForecastCalculationResponse {
	return o.Payload
}

func (o *ForecastCalculationServiceReadForecastCalculationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmReadForecastCalculationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForecastCalculationServiceReadForecastCalculationDefault creates a ForecastCalculationServiceReadForecastCalculationDefault with default headers values
func NewForecastCalculationServiceReadForecastCalculationDefault(code int) *ForecastCalculationServiceReadForecastCalculationDefault {
	return &ForecastCalculationServiceReadForecastCalculationDefault{
		_statusCode: code,
	}
}

/*
ForecastCalculationServiceReadForecastCalculationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ForecastCalculationServiceReadForecastCalculationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this forecast calculation service read forecast calculation default response has a 2xx status code
func (o *ForecastCalculationServiceReadForecastCalculationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this forecast calculation service read forecast calculation default response has a 3xx status code
func (o *ForecastCalculationServiceReadForecastCalculationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this forecast calculation service read forecast calculation default response has a 4xx status code
func (o *ForecastCalculationServiceReadForecastCalculationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this forecast calculation service read forecast calculation default response has a 5xx status code
func (o *ForecastCalculationServiceReadForecastCalculationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this forecast calculation service read forecast calculation default response a status code equal to that given
func (o *ForecastCalculationServiceReadForecastCalculationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the forecast calculation service read forecast calculation default response
func (o *ForecastCalculationServiceReadForecastCalculationDefault) Code() int {
	return o._statusCode
}

func (o *ForecastCalculationServiceReadForecastCalculationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation/{id}][%d] ForecastCalculationService_ReadForecastCalculation default %s", o._statusCode, payload)
}

func (o *ForecastCalculationServiceReadForecastCalculationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation/{id}][%d] ForecastCalculationService_ReadForecastCalculation default %s", o._statusCode, payload)
}

func (o *ForecastCalculationServiceReadForecastCalculationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ForecastCalculationServiceReadForecastCalculationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
