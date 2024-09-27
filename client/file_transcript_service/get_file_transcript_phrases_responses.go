// Code generated by go-swagger; DO NOT EDIT.

package file_transcript_service

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

// GetFileTranscriptPhrasesReader is a Reader for the GetFileTranscriptPhrases structure.
type GetFileTranscriptPhrasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileTranscriptPhrasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileTranscriptPhrasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /storage/transcript_file/{id}/phrases] GetFileTranscriptPhrases", response, response.Code())
	}
}

// NewGetFileTranscriptPhrasesOK creates a GetFileTranscriptPhrasesOK with default headers values
func NewGetFileTranscriptPhrasesOK() *GetFileTranscriptPhrasesOK {
	return &GetFileTranscriptPhrasesOK{}
}

/*
GetFileTranscriptPhrasesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetFileTranscriptPhrasesOK struct {
	Payload *models.StorageListPhrases
}

// IsSuccess returns true when this get file transcript phrases Ok response has a 2xx status code
func (o *GetFileTranscriptPhrasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file transcript phrases Ok response has a 3xx status code
func (o *GetFileTranscriptPhrasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file transcript phrases Ok response has a 4xx status code
func (o *GetFileTranscriptPhrasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file transcript phrases Ok response has a 5xx status code
func (o *GetFileTranscriptPhrasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file transcript phrases Ok response a status code equal to that given
func (o *GetFileTranscriptPhrasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file transcript phrases Ok response
func (o *GetFileTranscriptPhrasesOK) Code() int {
	return 200
}

func (o *GetFileTranscriptPhrasesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/transcript_file/{id}/phrases][%d] getFileTranscriptPhrasesOk %s", 200, payload)
}

func (o *GetFileTranscriptPhrasesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/transcript_file/{id}/phrases][%d] getFileTranscriptPhrasesOk %s", 200, payload)
}

func (o *GetFileTranscriptPhrasesOK) GetPayload() *models.StorageListPhrases {
	return o.Payload
}

func (o *GetFileTranscriptPhrasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageListPhrases)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
