// Code generated by go-swagger; DO NOT EDIT.

package team_hook_service

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

// PatchTeamHookReader is a Reader for the PatchTeamHook structure.
type PatchTeamHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTeamHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchTeamHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchTeamHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchTeamHookOK creates a PatchTeamHookOK with default headers values
func NewPatchTeamHookOK() *PatchTeamHookOK {
	return &PatchTeamHookOK{}
}

/*
PatchTeamHookOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchTeamHookOK struct {
	Payload *models.EngineTeamHook
}

// IsSuccess returns true when this patch team hook Ok response has a 2xx status code
func (o *PatchTeamHookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch team hook Ok response has a 3xx status code
func (o *PatchTeamHookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch team hook Ok response has a 4xx status code
func (o *PatchTeamHookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch team hook Ok response has a 5xx status code
func (o *PatchTeamHookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch team hook Ok response a status code equal to that given
func (o *PatchTeamHookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch team hook Ok response
func (o *PatchTeamHookOK) Code() int {
	return 200
}

func (o *PatchTeamHookOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/teams/{team_id}/hooks/{id}][%d] patchTeamHookOk %s", 200, payload)
}

func (o *PatchTeamHookOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/teams/{team_id}/hooks/{id}][%d] patchTeamHookOk %s", 200, payload)
}

func (o *PatchTeamHookOK) GetPayload() *models.EngineTeamHook {
	return o.Payload
}

func (o *PatchTeamHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineTeamHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamHookDefault creates a PatchTeamHookDefault with default headers values
func NewPatchTeamHookDefault(code int) *PatchTeamHookDefault {
	return &PatchTeamHookDefault{
		_statusCode: code,
	}
}

/*
PatchTeamHookDefault describes a response with status code -1, with default header values.

Server error
*/
type PatchTeamHookDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this patch team hook default response has a 2xx status code
func (o *PatchTeamHookDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch team hook default response has a 3xx status code
func (o *PatchTeamHookDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch team hook default response has a 4xx status code
func (o *PatchTeamHookDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch team hook default response has a 5xx status code
func (o *PatchTeamHookDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch team hook default response a status code equal to that given
func (o *PatchTeamHookDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch team hook default response
func (o *PatchTeamHookDefault) Code() int {
	return o._statusCode
}

func (o *PatchTeamHookDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/teams/{team_id}/hooks/{id}][%d] PatchTeamHook default %s", o._statusCode, payload)
}

func (o *PatchTeamHookDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/teams/{team_id}/hooks/{id}][%d] PatchTeamHook default %s", o._statusCode, payload)
}

func (o *PatchTeamHookDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *PatchTeamHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
