// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LibpodListImagesReader is a Reader for the LibpodListImages structure.
type LibpodListImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodListImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodListImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLibpodListImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLibpodListImagesOK creates a LibpodListImagesOK with default headers values
func NewLibpodListImagesOK() *LibpodListImagesOK {
	return &LibpodListImagesOK{}
}

/* LibpodListImagesOK describes a response with status code 200, with default header values.

Image summary
*/
type LibpodListImagesOK struct {
	Payload []*LibpodListImagesOKBodyItems0
}

func (o *LibpodListImagesOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/json][%d] libpodListImagesOK  %+v", 200, o.Payload)
}
func (o *LibpodListImagesOK) GetPayload() []*LibpodListImagesOKBodyItems0 {
	return o.Payload
}

func (o *LibpodListImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodListImagesInternalServerError creates a LibpodListImagesInternalServerError with default headers values
func NewLibpodListImagesInternalServerError() *LibpodListImagesInternalServerError {
	return &LibpodListImagesInternalServerError{}
}

/* LibpodListImagesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LibpodListImagesInternalServerError struct {
	Payload *LibpodListImagesInternalServerErrorBody
}

func (o *LibpodListImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/json][%d] libpodListImagesInternalServerError  %+v", 500, o.Payload)
}
func (o *LibpodListImagesInternalServerError) GetPayload() *LibpodListImagesInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodListImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodListImagesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodListImagesInternalServerErrorBody libpod list images internal server error body
swagger:model LibpodListImagesInternalServerErrorBody
*/
type LibpodListImagesInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod list images internal server error body
func (o *LibpodListImagesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod list images internal server error body based on context it is used
func (o *LibpodListImagesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodListImagesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodListImagesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodListImagesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodListImagesOKBodyItems0 ImageSummary image summary
swagger:model LibpodListImagesOKBodyItems0
*/
type LibpodListImagesOKBodyItems0 struct {

	// containers
	// Required: true
	Containers *int64 `json:"Containers"`

	// created
	// Required: true
	Created *int64 `json:"Created"`

	// Id
	// Required: true
	ID *string `json:"Id"`

	// labels
	// Required: true
	Labels map[string]string `json:"Labels"`

	// parent Id
	// Required: true
	ParentID *string `json:"ParentId"`

	// repo digests
	// Required: true
	RepoDigests []string `json:"RepoDigests"`

	// repo tags
	// Required: true
	RepoTags []string `json:"RepoTags"`

	// shared size
	// Required: true
	SharedSize *int64 `json:"SharedSize"`

	// size
	// Required: true
	Size *int64 `json:"Size"`

	// virtual size
	// Required: true
	VirtualSize *int64 `json:"VirtualSize"`
}

// Validate validates this libpod list images o k body items0
func (o *LibpodListImagesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRepoDigests(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRepoTags(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSharedSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVirtualSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("Containers", "body", o.Containers); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("Created", "body", o.Created); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("Id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("Labels", "body", o.Labels); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("ParentId", "body", o.ParentID); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateRepoDigests(formats strfmt.Registry) error {

	if err := validate.Required("RepoDigests", "body", o.RepoDigests); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateRepoTags(formats strfmt.Registry) error {

	if err := validate.Required("RepoTags", "body", o.RepoTags); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateSharedSize(formats strfmt.Registry) error {

	if err := validate.Required("SharedSize", "body", o.SharedSize); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

func (o *LibpodListImagesOKBodyItems0) validateVirtualSize(formats strfmt.Registry) error {

	if err := validate.Required("VirtualSize", "body", o.VirtualSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this libpod list images o k body items0 based on context it is used
func (o *LibpodListImagesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodListImagesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodListImagesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res LibpodListImagesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}