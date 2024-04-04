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

// StorageProxyMetricsCasWriteUnavailablesGetReader is a Reader for the StorageProxyMetricsCasWriteUnavailablesGet structure.
type StorageProxyMetricsCasWriteUnavailablesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsCasWriteUnavailablesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsCasWriteUnavailablesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsCasWriteUnavailablesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsCasWriteUnavailablesGetOK creates a StorageProxyMetricsCasWriteUnavailablesGetOK with default headers values
func NewStorageProxyMetricsCasWriteUnavailablesGetOK() *StorageProxyMetricsCasWriteUnavailablesGetOK {
	return &StorageProxyMetricsCasWriteUnavailablesGetOK{}
}

/*
StorageProxyMetricsCasWriteUnavailablesGetOK handles this case with default header values.

Success
*/
type StorageProxyMetricsCasWriteUnavailablesGetOK struct {
	Payload interface{}
}

func (o *StorageProxyMetricsCasWriteUnavailablesGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyMetricsCasWriteUnavailablesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsCasWriteUnavailablesGetDefault creates a StorageProxyMetricsCasWriteUnavailablesGetDefault with default headers values
func NewStorageProxyMetricsCasWriteUnavailablesGetDefault(code int) *StorageProxyMetricsCasWriteUnavailablesGetDefault {
	return &StorageProxyMetricsCasWriteUnavailablesGetDefault{
		_statusCode: code,
	}
}

/*
StorageProxyMetricsCasWriteUnavailablesGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsCasWriteUnavailablesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics cas write unavailables get default response
func (o *StorageProxyMetricsCasWriteUnavailablesGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsCasWriteUnavailablesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsCasWriteUnavailablesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsCasWriteUnavailablesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
