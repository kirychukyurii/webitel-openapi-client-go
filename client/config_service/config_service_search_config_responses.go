// Code generated by go-swagger; DO NOT EDIT.

package config_service

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

// ConfigServiceSearchConfigReader is a Reader for the ConfigServiceSearchConfig structure.
type ConfigServiceSearchConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigServiceSearchConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigServiceSearchConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigServiceSearchConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigServiceSearchConfigOK creates a ConfigServiceSearchConfigOK with default headers values
func NewConfigServiceSearchConfigOK() *ConfigServiceSearchConfigOK {
	return &ConfigServiceSearchConfigOK{}
}

/*
ConfigServiceSearchConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConfigServiceSearchConfigOK struct {
	Payload *models.LoggerConfigs
}

// IsSuccess returns true when this config service search config Ok response has a 2xx status code
func (o *ConfigServiceSearchConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this config service search config Ok response has a 3xx status code
func (o *ConfigServiceSearchConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config service search config Ok response has a 4xx status code
func (o *ConfigServiceSearchConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this config service search config Ok response has a 5xx status code
func (o *ConfigServiceSearchConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this config service search config Ok response a status code equal to that given
func (o *ConfigServiceSearchConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the config service search config Ok response
func (o *ConfigServiceSearchConfigOK) Code() int {
	return 200
}

func (o *ConfigServiceSearchConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /logger/config][%d] configServiceSearchConfigOk %s", 200, payload)
}

func (o *ConfigServiceSearchConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /logger/config][%d] configServiceSearchConfigOk %s", 200, payload)
}

func (o *ConfigServiceSearchConfigOK) GetPayload() *models.LoggerConfigs {
	return o.Payload
}

func (o *ConfigServiceSearchConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoggerConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigServiceSearchConfigDefault creates a ConfigServiceSearchConfigDefault with default headers values
func NewConfigServiceSearchConfigDefault(code int) *ConfigServiceSearchConfigDefault {
	return &ConfigServiceSearchConfigDefault{
		_statusCode: code,
	}
}

/*
ConfigServiceSearchConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ConfigServiceSearchConfigDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this config service search config default response has a 2xx status code
func (o *ConfigServiceSearchConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this config service search config default response has a 3xx status code
func (o *ConfigServiceSearchConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this config service search config default response has a 4xx status code
func (o *ConfigServiceSearchConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this config service search config default response has a 5xx status code
func (o *ConfigServiceSearchConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this config service search config default response a status code equal to that given
func (o *ConfigServiceSearchConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the config service search config default response
func (o *ConfigServiceSearchConfigDefault) Code() int {
	return o._statusCode
}

func (o *ConfigServiceSearchConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /logger/config][%d] ConfigService_SearchConfig default %s", o._statusCode, payload)
}

func (o *ConfigServiceSearchConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /logger/config][%d] ConfigService_SearchConfig default %s", o._statusCode, payload)
}

func (o *ConfigServiceSearchConfigDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ConfigServiceSearchConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
