// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodAttachContainerReader is a Reader for the LibpodAttachContainer structure.
type LibpodAttachContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodAttachContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 101:
		result := NewLibpodAttachContainerSwitchingProtocols()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewLibpodAttachContainerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLibpodAttachContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodAttachContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodAttachContainerSwitchingProtocols creates a LibpodAttachContainerSwitchingProtocols with default headers values
func NewLibpodAttachContainerSwitchingProtocols() *LibpodAttachContainerSwitchingProtocols {
	return &LibpodAttachContainerSwitchingProtocols{}
}

/* LibpodAttachContainerSwitchingProtocols describes a response with status code 101, with default header values.

No error, connection has been hijacked for transporting streams.
*/
type LibpodAttachContainerSwitchingProtocols struct {
}

func (o *LibpodAttachContainerSwitchingProtocols) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] libpodAttachContainerSwitchingProtocols ", 101)
}

func (o *LibpodAttachContainerSwitchingProtocols) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodAttachContainerBadRequest creates a LibpodAttachContainerBadRequest with default headers values
func NewLibpodAttachContainerBadRequest() *LibpodAttachContainerBadRequest {
	return &LibpodAttachContainerBadRequest{}
}

/* LibpodAttachContainerBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type LibpodAttachContainerBadRequest struct {
	Payload *LibpodAttachContainerBadRequestBody
}

func (o *LibpodAttachContainerBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] libpodAttachContainerBadRequest  %+v", 400, o.Payload)
}
func (o *LibpodAttachContainerBadRequest) GetPayload() *LibpodAttachContainerBadRequestBody {
	return o.Payload
}

func (o *LibpodAttachContainerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodAttachContainerBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodAttachContainerNotFound creates a LibpodAttachContainerNotFound with default headers values
func NewLibpodAttachContainerNotFound() *LibpodAttachContainerNotFound {
	return &LibpodAttachContainerNotFound{}
}

/* LibpodAttachContainerNotFound describes a response with status code 404, with default header values.

No such container
*/
type LibpodAttachContainerNotFound struct {
	Payload *LibpodAttachContainerNotFoundBody
}

func (o *LibpodAttachContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] libpodAttachContainerNotFound  %+v", 404, o.Payload)
}
func (o *LibpodAttachContainerNotFound) GetPayload() *LibpodAttachContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodAttachContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodAttachContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodAttachContainerInternalServerError creates a LibpodAttachContainerInternalServerError with default headers values
func NewLibpodAttachContainerInternalServerError() *LibpodAttachContainerInternalServerError {
	return &LibpodAttachContainerInternalServerError{}
}

/* LibpodAttachContainerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodAttachContainerInternalServerError struct {
	Payload *LibpodAttachContainerInternalServerErrorBody
}

func (o *LibpodAttachContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] libpodAttachContainerInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodAttachContainerInternalServerError) GetPayload() *LibpodAttachContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodAttachContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodAttachContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodAttachContainerBadRequestBody libpod attach container bad request body
swagger:model LibpodAttachContainerBadRequestBody
*/
type LibpodAttachContainerBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod attach container bad request body
func (o *LibpodAttachContainerBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod attach container bad request body based on context it is used
func (o *LibpodAttachContainerBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodAttachContainerBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodAttachContainerBadRequestBody) UnmarshalBinary(b []byte) error {
	var res LibpodAttachContainerBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodAttachContainerInternalServerErrorBody libpod attach container internal server error body
swagger:model LibpodAttachContainerInternalServerErrorBody
*/
type LibpodAttachContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod attach container internal server error body
func (o *LibpodAttachContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod attach container internal server error body based on context it is used
func (o *LibpodAttachContainerInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodAttachContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodAttachContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodAttachContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodAttachContainerNotFoundBody libpod attach container not found body
swagger:model LibpodAttachContainerNotFoundBody
*/
type LibpodAttachContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod attach container not found body
func (o *LibpodAttachContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod attach container not found body based on context it is used
func (o *LibpodAttachContainerNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodAttachContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodAttachContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodAttachContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
