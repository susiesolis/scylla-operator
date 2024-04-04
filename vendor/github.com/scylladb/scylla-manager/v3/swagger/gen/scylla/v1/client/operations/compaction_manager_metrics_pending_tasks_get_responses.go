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

// CompactionManagerMetricsPendingTasksGetReader is a Reader for the CompactionManagerMetricsPendingTasksGet structure.
type CompactionManagerMetricsPendingTasksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompactionManagerMetricsPendingTasksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompactionManagerMetricsPendingTasksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCompactionManagerMetricsPendingTasksGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompactionManagerMetricsPendingTasksGetOK creates a CompactionManagerMetricsPendingTasksGetOK with default headers values
func NewCompactionManagerMetricsPendingTasksGetOK() *CompactionManagerMetricsPendingTasksGetOK {
	return &CompactionManagerMetricsPendingTasksGetOK{}
}

/*
CompactionManagerMetricsPendingTasksGetOK handles this case with default header values.

Success
*/
type CompactionManagerMetricsPendingTasksGetOK struct {
	Payload int32
}

func (o *CompactionManagerMetricsPendingTasksGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *CompactionManagerMetricsPendingTasksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompactionManagerMetricsPendingTasksGetDefault creates a CompactionManagerMetricsPendingTasksGetDefault with default headers values
func NewCompactionManagerMetricsPendingTasksGetDefault(code int) *CompactionManagerMetricsPendingTasksGetDefault {
	return &CompactionManagerMetricsPendingTasksGetDefault{
		_statusCode: code,
	}
}

/*
CompactionManagerMetricsPendingTasksGetDefault handles this case with default header values.

internal server error
*/
type CompactionManagerMetricsPendingTasksGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the compaction manager metrics pending tasks get default response
func (o *CompactionManagerMetricsPendingTasksGetDefault) Code() int {
	return o._statusCode
}

func (o *CompactionManagerMetricsPendingTasksGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CompactionManagerMetricsPendingTasksGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CompactionManagerMetricsPendingTasksGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
