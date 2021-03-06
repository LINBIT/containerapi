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
	"github.com/go-openapi/swag"
)

// NewLibpodRemoveImageParams creates a new LibpodRemoveImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodRemoveImageParams() *LibpodRemoveImageParams {
	return &LibpodRemoveImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodRemoveImageParamsWithTimeout creates a new LibpodRemoveImageParams object
// with the ability to set a timeout on a request.
func NewLibpodRemoveImageParamsWithTimeout(timeout time.Duration) *LibpodRemoveImageParams {
	return &LibpodRemoveImageParams{
		timeout: timeout,
	}
}

// NewLibpodRemoveImageParamsWithContext creates a new LibpodRemoveImageParams object
// with the ability to set a context for a request.
func NewLibpodRemoveImageParamsWithContext(ctx context.Context) *LibpodRemoveImageParams {
	return &LibpodRemoveImageParams{
		Context: ctx,
	}
}

// NewLibpodRemoveImageParamsWithHTTPClient creates a new LibpodRemoveImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodRemoveImageParamsWithHTTPClient(client *http.Client) *LibpodRemoveImageParams {
	return &LibpodRemoveImageParams{
		HTTPClient: client,
	}
}

/* LibpodRemoveImageParams contains all the parameters to send to the API endpoint
   for the libpod remove image operation.

   Typically these are written to a http.Request.
*/
type LibpodRemoveImageParams struct {

	/* Force.

	   remove the image even if used by containers or has other tags
	*/
	Force *bool

	/* Name.

	   name or ID of image to remove
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod remove image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodRemoveImageParams) WithDefaults() *LibpodRemoveImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod remove image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodRemoveImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the libpod remove image params
func (o *LibpodRemoveImageParams) WithTimeout(timeout time.Duration) *LibpodRemoveImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod remove image params
func (o *LibpodRemoveImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod remove image params
func (o *LibpodRemoveImageParams) WithContext(ctx context.Context) *LibpodRemoveImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod remove image params
func (o *LibpodRemoveImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod remove image params
func (o *LibpodRemoveImageParams) WithHTTPClient(client *http.Client) *LibpodRemoveImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod remove image params
func (o *LibpodRemoveImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the libpod remove image params
func (o *LibpodRemoveImageParams) WithForce(force *bool) *LibpodRemoveImageParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the libpod remove image params
func (o *LibpodRemoveImageParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the libpod remove image params
func (o *LibpodRemoveImageParams) WithName(name string) *LibpodRemoveImageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod remove image params
func (o *LibpodRemoveImageParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodRemoveImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param name:.*
	if err := r.SetPathParam("name:.*", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
