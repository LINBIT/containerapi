// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewLibpodImageExistsParams creates a new LibpodImageExistsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodImageExistsParams() *LibpodImageExistsParams {
	return &LibpodImageExistsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodImageExistsParamsWithTimeout creates a new LibpodImageExistsParams object
// with the ability to set a timeout on a request.
func NewLibpodImageExistsParamsWithTimeout(timeout time.Duration) *LibpodImageExistsParams {
	return &LibpodImageExistsParams{
		timeout: timeout,
	}
}

// NewLibpodImageExistsParamsWithContext creates a new LibpodImageExistsParams object
// with the ability to set a context for a request.
func NewLibpodImageExistsParamsWithContext(ctx context.Context) *LibpodImageExistsParams {
	return &LibpodImageExistsParams{
		Context: ctx,
	}
}

// NewLibpodImageExistsParamsWithHTTPClient creates a new LibpodImageExistsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodImageExistsParamsWithHTTPClient(client *http.Client) *LibpodImageExistsParams {
	return &LibpodImageExistsParams{
		HTTPClient: client,
	}
}

/* LibpodImageExistsParams contains all the parameters to send to the API endpoint
   for the libpod image exists operation.

   Typically these are written to a http.Request.
*/
type LibpodImageExistsParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod image exists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodImageExistsParams) WithDefaults() *LibpodImageExistsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod image exists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodImageExistsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the libpod image exists params
func (o *LibpodImageExistsParams) WithTimeout(timeout time.Duration) *LibpodImageExistsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod image exists params
func (o *LibpodImageExistsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod image exists params
func (o *LibpodImageExistsParams) WithContext(ctx context.Context) *LibpodImageExistsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod image exists params
func (o *LibpodImageExistsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod image exists params
func (o *LibpodImageExistsParams) WithHTTPClient(client *http.Client) *LibpodImageExistsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod image exists params
func (o *LibpodImageExistsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the libpod image exists params
func (o *LibpodImageExistsParams) WithName(name string) *LibpodImageExistsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod image exists params
func (o *LibpodImageExistsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodImageExistsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name:.*
	if err := r.SetPathParam("name:.*", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}