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

// ForecastCalculationServiceExecuteForecastCalculationReader is a Reader for the ForecastCalculationServiceExecuteForecastCalculation structure.
type ForecastCalculationServiceExecuteForecastCalculationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForecastCalculationServiceExecuteForecastCalculationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForecastCalculationServiceExecuteForecastCalculationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewForecastCalculationServiceExecuteForecastCalculationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewForecastCalculationServiceExecuteForecastCalculationOK creates a ForecastCalculationServiceExecuteForecastCalculationOK with default headers values
func NewForecastCalculationServiceExecuteForecastCalculationOK() *ForecastCalculationServiceExecuteForecastCalculationOK {
	return &ForecastCalculationServiceExecuteForecastCalculationOK{}
}

/*
ForecastCalculationServiceExecuteForecastCalculationOK describes a response with status code 200, with default header values.

A successful response.
*/
type ForecastCalculationServiceExecuteForecastCalculationOK struct {
	Payload *models.WfmExecuteForecastCalculationResponse
}

// IsSuccess returns true when this forecast calculation service execute forecast calculation Ok response has a 2xx status code
func (o *ForecastCalculationServiceExecuteForecastCalculationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this forecast calculation service execute forecast calculation Ok response has a 3xx status code
func (o *ForecastCalculationServiceExecuteForecastCalculationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this forecast calculation service execute forecast calculation Ok response has a 4xx status code
func (o *ForecastCalculationServiceExecuteForecastCalculationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this forecast calculation service execute forecast calculation Ok response has a 5xx status code
func (o *ForecastCalculationServiceExecuteForecastCalculationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this forecast calculation service execute forecast calculation Ok response a status code equal to that given
func (o *ForecastCalculationServiceExecuteForecastCalculationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the forecast calculation service execute forecast calculation Ok response
func (o *ForecastCalculationServiceExecuteForecastCalculationOK) Code() int {
	return 200
}

func (o *ForecastCalculationServiceExecuteForecastCalculationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation/{id_1}][%d] forecastCalculationServiceExecuteForecastCalculationOk %s", 200, payload)
}

func (o *ForecastCalculationServiceExecuteForecastCalculationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation/{id_1}][%d] forecastCalculationServiceExecuteForecastCalculationOk %s", 200, payload)
}

func (o *ForecastCalculationServiceExecuteForecastCalculationOK) GetPayload() *models.WfmExecuteForecastCalculationResponse {
	return o.Payload
}

func (o *ForecastCalculationServiceExecuteForecastCalculationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmExecuteForecastCalculationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForecastCalculationServiceExecuteForecastCalculationDefault creates a ForecastCalculationServiceExecuteForecastCalculationDefault with default headers values
func NewForecastCalculationServiceExecuteForecastCalculationDefault(code int) *ForecastCalculationServiceExecuteForecastCalculationDefault {
	return &ForecastCalculationServiceExecuteForecastCalculationDefault{
		_statusCode: code,
	}
}

/*
ForecastCalculationServiceExecuteForecastCalculationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ForecastCalculationServiceExecuteForecastCalculationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this forecast calculation service execute forecast calculation default response has a 2xx status code
func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this forecast calculation service execute forecast calculation default response has a 3xx status code
func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this forecast calculation service execute forecast calculation default response has a 4xx status code
func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this forecast calculation service execute forecast calculation default response has a 5xx status code
func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this forecast calculation service execute forecast calculation default response a status code equal to that given
func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the forecast calculation service execute forecast calculation default response
func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) Code() int {
	return o._statusCode
}

func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation/{id_1}][%d] ForecastCalculationService_ExecuteForecastCalculation default %s", o._statusCode, payload)
}

func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation/{id_1}][%d] ForecastCalculationService_ExecuteForecastCalculation default %s", o._statusCode, payload)
}

func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ForecastCalculationServiceExecuteForecastCalculationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
