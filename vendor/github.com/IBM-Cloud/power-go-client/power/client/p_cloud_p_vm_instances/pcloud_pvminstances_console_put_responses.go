// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesConsolePutReader is a Reader for the PcloudPvminstancesConsolePut structure.
type PcloudPvminstancesConsolePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesConsolePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesConsolePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesConsolePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesConsolePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesConsolePutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesConsolePutOK creates a PcloudPvminstancesConsolePutOK with default headers values
func NewPcloudPvminstancesConsolePutOK() *PcloudPvminstancesConsolePutOK {
	return &PcloudPvminstancesConsolePutOK{}
}

/* PcloudPvminstancesConsolePutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesConsolePutOK struct {
	Payload *models.ConsoleLanguage
}

func (o *PcloudPvminstancesConsolePutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutOK  %+v", 200, o.Payload)
}
func (o *PcloudPvminstancesConsolePutOK) GetPayload() *models.ConsoleLanguage {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleLanguage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePutBadRequest creates a PcloudPvminstancesConsolePutBadRequest with default headers values
func NewPcloudPvminstancesConsolePutBadRequest() *PcloudPvminstancesConsolePutBadRequest {
	return &PcloudPvminstancesConsolePutBadRequest{}
}

/* PcloudPvminstancesConsolePutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesConsolePutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsolePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudPvminstancesConsolePutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePutUnauthorized creates a PcloudPvminstancesConsolePutUnauthorized with default headers values
func NewPcloudPvminstancesConsolePutUnauthorized() *PcloudPvminstancesConsolePutUnauthorized {
	return &PcloudPvminstancesConsolePutUnauthorized{}
}

/* PcloudPvminstancesConsolePutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesConsolePutUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsolePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudPvminstancesConsolePutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePutInternalServerError creates a PcloudPvminstancesConsolePutInternalServerError with default headers values
func NewPcloudPvminstancesConsolePutInternalServerError() *PcloudPvminstancesConsolePutInternalServerError {
	return &PcloudPvminstancesConsolePutInternalServerError{}
}

/* PcloudPvminstancesConsolePutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesConsolePutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsolePutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudPvminstancesConsolePutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}