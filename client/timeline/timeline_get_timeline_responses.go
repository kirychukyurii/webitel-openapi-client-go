// Code generated by go-swagger; DO NOT EDIT.

package timeline

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

// TimelineGetTimelineReader is a Reader for the TimelineGetTimeline structure.
type TimelineGetTimelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimelineGetTimelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimelineGetTimelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /contacts/{contact_id}/timeline] Timeline_GetTimeline", response, response.Code())
	}
}

// NewTimelineGetTimelineOK creates a TimelineGetTimelineOK with default headers values
func NewTimelineGetTimelineOK() *TimelineGetTimelineOK {
	return &TimelineGetTimelineOK{}
}

/*
TimelineGetTimelineOK describes a response with status code 200, with default header values.

A successful response.
*/
type TimelineGetTimelineOK struct {
	Payload *models.WebitelContactsGetTimelineResponse
}

// IsSuccess returns true when this timeline get timeline Ok response has a 2xx status code
func (o *TimelineGetTimelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this timeline get timeline Ok response has a 3xx status code
func (o *TimelineGetTimelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this timeline get timeline Ok response has a 4xx status code
func (o *TimelineGetTimelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this timeline get timeline Ok response has a 5xx status code
func (o *TimelineGetTimelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this timeline get timeline Ok response a status code equal to that given
func (o *TimelineGetTimelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the timeline get timeline Ok response
func (o *TimelineGetTimelineOK) Code() int {
	return 200
}

func (o *TimelineGetTimelineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/timeline][%d] timelineGetTimelineOk %s", 200, payload)
}

func (o *TimelineGetTimelineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/timeline][%d] timelineGetTimelineOk %s", 200, payload)
}

func (o *TimelineGetTimelineOK) GetPayload() *models.WebitelContactsGetTimelineResponse {
	return o.Payload
}

func (o *TimelineGetTimelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsGetTimelineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
