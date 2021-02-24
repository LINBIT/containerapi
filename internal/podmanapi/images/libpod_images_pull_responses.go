// Code generated by go-swagger; DO NOT EDIT.

package images

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

// LibpodImagesPullReader is a Reader for the LibpodImagesPull structure.
type LibpodImagesPullReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *LibpodImagesPullReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodImagesPullOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLibpodImagesPullBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodImagesPullInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodImagesPullOK creates a LibpodImagesPullOK with default headers values
func NewLibpodImagesPullOK(writer io.Writer) *LibpodImagesPullOK {
	return &LibpodImagesPullOK{

		Payload: writer,
	}
}

/* LibpodImagesPullOK describes a response with status code 200, with default header values.

JSON lines holding progress information on the pull process
*/
type LibpodImagesPullOK struct {
	Payload io.Writer
}

func (o *LibpodImagesPullOK) Error() string {
	return fmt.Sprintf("[POST /libpod/images/pull][%d] libpodImagesPullOK  %+v", 200, o.Payload)
}
func (o *LibpodImagesPullOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *LibpodImagesPullOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodImagesPullBadRequest creates a LibpodImagesPullBadRequest with default headers values
func NewLibpodImagesPullBadRequest() *LibpodImagesPullBadRequest {
	return &LibpodImagesPullBadRequest{}
}

/* LibpodImagesPullBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type LibpodImagesPullBadRequest struct {
	Payload *LibpodImagesPullBadRequestBody
}

func (o *LibpodImagesPullBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/images/pull][%d] libpodImagesPullBadRequest  %+v", 400, o.Payload)
}
func (o *LibpodImagesPullBadRequest) GetPayload() *LibpodImagesPullBadRequestBody {
	return o.Payload
}

func (o *LibpodImagesPullBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImagesPullBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodImagesPullInternalServerError creates a LibpodImagesPullInternalServerError with default headers values
func NewLibpodImagesPullInternalServerError() *LibpodImagesPullInternalServerError {
	return &LibpodImagesPullInternalServerError{}
}

/* LibpodImagesPullInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodImagesPullInternalServerError struct {
	Payload *LibpodImagesPullInternalServerErrorBody
}

func (o *LibpodImagesPullInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/pull][%d] libpodImagesPullInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodImagesPullInternalServerError) GetPayload() *LibpodImagesPullInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodImagesPullInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodImagesPullInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodImagesPullBadRequestBody libpod images pull bad request body
swagger:model LibpodImagesPullBadRequestBody
*/
type LibpodImagesPullBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod images pull bad request body
func (o *LibpodImagesPullBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod images pull bad request body based on context it is used
func (o *LibpodImagesPullBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImagesPullBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImagesPullBadRequestBody) UnmarshalBinary(b []byte) error {
	var res LibpodImagesPullBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodImagesPullInternalServerErrorBody libpod images pull internal server error body
swagger:model LibpodImagesPullInternalServerErrorBody
*/
type LibpodImagesPullInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod images pull internal server error body
func (o *LibpodImagesPullInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod images pull internal server error body based on context it is used
func (o *LibpodImagesPullInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodImagesPullInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodImagesPullInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodImagesPullInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}