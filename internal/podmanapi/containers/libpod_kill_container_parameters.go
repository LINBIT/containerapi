// Code generated by go-swagger; DO NOT EDIT.

package containers

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

// NewLibpodKillContainerParams creates a new LibpodKillContainerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodKillContainerParams() *LibpodKillContainerParams {
	return &LibpodKillContainerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodKillContainerParamsWithTimeout creates a new LibpodKillContainerParams object
// with the ability to set a timeout on a request.
func NewLibpodKillContainerParamsWithTimeout(timeout time.Duration) *LibpodKillContainerParams {
	return &LibpodKillContainerParams{
		timeout: timeout,
	}
}

// NewLibpodKillContainerParamsWithContext creates a new LibpodKillContainerParams object
// with the ability to set a context for a request.
func NewLibpodKillContainerParamsWithContext(ctx context.Context) *LibpodKillContainerParams {
	return &LibpodKillContainerParams{
		Context: ctx,
	}
}

// NewLibpodKillContainerParamsWithHTTPClient creates a new LibpodKillContainerParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodKillContainerParamsWithHTTPClient(client *http.Client) *LibpodKillContainerParams {
	return &LibpodKillContainerParams{
		HTTPClient: client,
	}
}

/* LibpodKillContainerParams contains all the parameters to send to the API endpoint
   for the libpod kill container operation.

   Typically these are written to a http.Request.
*/
type LibpodKillContainerParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	/* Signal.

	   signal to be sent to container, either by integer or SIG_ name

	   Default: "TERM"
	*/
	Signal *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod kill container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodKillContainerParams) WithDefaults() *LibpodKillContainerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod kill container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodKillContainerParams) SetDefaults() {
	var (
		signalDefault = string("TERM")
	)

	val := LibpodKillContainerParams{
		Signal: &signalDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the libpod kill container params
func (o *LibpodKillContainerParams) WithTimeout(timeout time.Duration) *LibpodKillContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod kill container params
func (o *LibpodKillContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod kill container params
func (o *LibpodKillContainerParams) WithContext(ctx context.Context) *LibpodKillContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod kill container params
func (o *LibpodKillContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod kill container params
func (o *LibpodKillContainerParams) WithHTTPClient(client *http.Client) *LibpodKillContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod kill container params
func (o *LibpodKillContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the libpod kill container params
func (o *LibpodKillContainerParams) WithName(name string) *LibpodKillContainerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod kill container params
func (o *LibpodKillContainerParams) SetName(name string) {
	o.Name = name
}

// WithSignal adds the signal to the libpod kill container params
func (o *LibpodKillContainerParams) WithSignal(signal *string) *LibpodKillContainerParams {
	o.SetSignal(signal)
	return o
}

// SetSignal adds the signal to the libpod kill container params
func (o *LibpodKillContainerParams) SetSignal(signal *string) {
	o.Signal = signal
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodKillContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Signal != nil {

		// query param signal
		var qrSignal string

		if o.Signal != nil {
			qrSignal = *o.Signal
		}
		qSignal := qrSignal
		if qSignal != "" {

			if err := r.SetQueryParam("signal", qSignal); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
