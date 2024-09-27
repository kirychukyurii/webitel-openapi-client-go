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

// ForecastCalculationServiceSearchForecastCalculationReader is a Reader for the ForecastCalculationServiceSearchForecastCalculation structure.
type ForecastCalculationServiceSearchForecastCalculationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForecastCalculationServiceSearchForecastCalculationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForecastCalculationServiceSearchForecastCalculationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewForecastCalculationServiceSearchForecastCalculationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewForecastCalculationServiceSearchForecastCalculationOK creates a ForecastCalculationServiceSearchForecastCalculationOK with default headers values
func NewForecastCalculationServiceSearchForecastCalculationOK() *ForecastCalculationServiceSearchForecastCalculationOK {
	return &ForecastCalculationServiceSearchForecastCalculationOK{}
}

/*
ForecastCalculationServiceSearchForecastCalculationOK describes a response with status code 200, with default header values.

A successful response.
*/
type ForecastCalculationServiceSearchForecastCalculationOK struct {
	Payload *models.WfmSearchForecastCalculationResponse
}

// IsSuccess returns true when this forecast calculation service search forecast calculation Ok response has a 2xx status code
func (o *ForecastCalculationServiceSearchForecastCalculationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this forecast calculation service search forecast calculation Ok response has a 3xx status code
func (o *ForecastCalculationServiceSearchForecastCalculationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this forecast calculation service search forecast calculation Ok response has a 4xx status code
func (o *ForecastCalculationServiceSearchForecastCalculationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this forecast calculation service search forecast calculation Ok response has a 5xx status code
func (o *ForecastCalculationServiceSearchForecastCalculationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this forecast calculation service search forecast calculation Ok response a status code equal to that given
func (o *ForecastCalculationServiceSearchForecastCalculationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the forecast calculation service search forecast calculation Ok response
func (o *ForecastCalculationServiceSearchForecastCalculationOK) Code() int {
	return 200
}

func (o *ForecastCalculationServiceSearchForecastCalculationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation][%d] forecastCalculationServiceSearchForecastCalculationOk %s", 200, payload)
}

func (o *ForecastCalculationServiceSearchForecastCalculationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation][%d] forecastCalculationServiceSearchForecastCalculationOk %s", 200, payload)
}

func (o *ForecastCalculationServiceSearchForecastCalculationOK) GetPayload() *models.WfmSearchForecastCalculationResponse {
	return o.Payload
}

func (o *ForecastCalculationServiceSearchForecastCalculationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmSearchForecastCalculationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForecastCalculationServiceSearchForecastCalculationDefault creates a ForecastCalculationServiceSearchForecastCalculationDefault with default headers values
func NewForecastCalculationServiceSearchForecastCalculationDefault(code int) *ForecastCalculationServiceSearchForecastCalculationDefault {
	return &ForecastCalculationServiceSearchForecastCalculationDefault{
		_statusCode: code,
	}
}

/*
ForecastCalculationServiceSearchForecastCalculationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ForecastCalculationServiceSearchForecastCalculationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this forecast calculation service search forecast calculation default response has a 2xx status code
func (o *ForecastCalculationServiceSearchForecastCalculationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this forecast calculation service search forecast calculation default response has a 3xx status code
func (o *ForecastCalculationServiceSearchForecastCalculationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this forecast calculation service search forecast calculation default response has a 4xx status code
func (o *ForecastCalculationServiceSearchForecastCalculationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this forecast calculation service search forecast calculation default response has a 5xx status code
func (o *ForecastCalculationServiceSearchForecastCalculationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this forecast calculation service search forecast calculation default response a status code equal to that given
func (o *ForecastCalculationServiceSearchForecastCalculationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the forecast calculation service search forecast calculation default response
func (o *ForecastCalculationServiceSearchForecastCalculationDefault) Code() int {
	return o._statusCode
}

func (o *ForecastCalculationServiceSearchForecastCalculationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation][%d] ForecastCalculationService_SearchForecastCalculation default %s", o._statusCode, payload)
}

func (o *ForecastCalculationServiceSearchForecastCalculationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/forecast_calculation][%d] ForecastCalculationService_SearchForecastCalculation default %s", o._statusCode, payload)
}

func (o *ForecastCalculationServiceSearchForecastCalculationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ForecastCalculationServiceSearchForecastCalculationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
