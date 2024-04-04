// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// StorageServiceForceRemoveCompletionPostReader is a Reader for the StorageServiceForceRemoveCompletionPost structure.
type StorageServiceForceRemoveCompletionPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceForceRemoveCompletionPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceForceRemoveCompletionPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceForceRemoveCompletionPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceForceRemoveCompletionPostOK creates a StorageServiceForceRemoveCompletionPostOK with default headers values
func NewStorageServiceForceRemoveCompletionPostOK() *StorageServiceForceRemoveCompletionPostOK {
	return &StorageServiceForceRemoveCompletionPostOK{}
}

/*
StorageServiceForceRemoveCompletionPostOK handles this case with default header values.

Success
*/
type StorageServiceForceRemoveCompletionPostOK struct {
}

func (o *StorageServiceForceRemoveCompletionPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceForceRemoveCompletionPostDefault creates a StorageServiceForceRemoveCompletionPostDefault with default headers values
func NewStorageServiceForceRemoveCompletionPostDefault(code int) *StorageServiceForceRemoveCompletionPostDefault {
	return &StorageServiceForceRemoveCompletionPostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceForceRemoveCompletionPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceForceRemoveCompletionPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service force remove completion post default response
func (o *StorageServiceForceRemoveCompletionPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceForceRemoveCompletionPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceForceRemoveCompletionPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceForceRemoveCompletionPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}