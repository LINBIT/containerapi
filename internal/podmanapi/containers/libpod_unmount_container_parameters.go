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

// NewLibpodUnmountContainerParams creates a new LibpodUnmountContainerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodUnmountContainerParams() *LibpodUnmountContainerParams {
	return &LibpodUnmountContainerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodUnmountContainerParamsWithTimeout creates a new LibpodUnmountContainerParams object
// with the ability to set a timeout on a request.
func NewLibpodUnmountContainerParamsWithTimeout(timeout time.Duration) *LibpodUnmountContainerParams {
	return &LibpodUnmountContainerParams{
		timeout: timeout,
	}
}

// NewLibpodUnmountContainerParamsWithContext creates a new LibpodUnmountContainerParams object
// with the ability to set a context for a request.
func NewLibpodUnmountContainerParamsWithContext(ctx context.Context) *LibpodUnmountContainerParams {
	return &LibpodUnmountContainerParams{
		Context: ctx,
	}
}

// NewLibpodUnmountContainerParamsWithHTTPClient creates a new LibpodUnmountContainerParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodUnmountContainerParamsWithHTTPClient(client *http.Client) *LibpodUnmountContainerParams {
	return &LibpodUnmountContainerParams{
		HTTPClient: client,
	}
}

/* LibpodUnmountContainerParams contains all the parameters to send to the API endpoint
   for the libpod unmount container operation.

   Typically these are written to a http.Request.
*/
type LibpodUnmountContainerParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod unmount container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodUnmountContainerParams) WithDefaults() *LibpodUnmountContainerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod unmount container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodUnmountContainerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the libpod unmount container params
func (o *LibpodUnmountContainerParams) WithTimeout(timeout time.Duration) *LibpodUnmountContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod unmount container params
func (o *LibpodUnmountContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod unmount container params
func (o *LibpodUnmountContainerParams) WithContext(ctx context.Context) *LibpodUnmountContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod unmount container params
func (o *LibpodUnmountContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod unmount container params
func (o *LibpodUnmountContainerParams) WithHTTPClient(client *http.Client) *LibpodUnmountContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod unmount container params
func (o *LibpodUnmountContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the libpod unmount container params
func (o *LibpodUnmountContainerParams) WithName(name string) *LibpodUnmountContainerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod unmount container params
func (o *LibpodUnmountContainerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodUnmountContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
