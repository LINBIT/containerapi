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

// NewLibpodImagesRemoveParams creates a new LibpodImagesRemoveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodImagesRemoveParams() *LibpodImagesRemoveParams {
	return &LibpodImagesRemoveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodImagesRemoveParamsWithTimeout creates a new LibpodImagesRemoveParams object
// with the ability to set a timeout on a request.
func NewLibpodImagesRemoveParamsWithTimeout(timeout time.Duration) *LibpodImagesRemoveParams {
	return &LibpodImagesRemoveParams{
		timeout: timeout,
	}
}

// NewLibpodImagesRemoveParamsWithContext creates a new LibpodImagesRemoveParams object
// with the ability to set a context for a request.
func NewLibpodImagesRemoveParamsWithContext(ctx context.Context) *LibpodImagesRemoveParams {
	return &LibpodImagesRemoveParams{
		Context: ctx,
	}
}

// NewLibpodImagesRemoveParamsWithHTTPClient creates a new LibpodImagesRemoveParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodImagesRemoveParamsWithHTTPClient(client *http.Client) *LibpodImagesRemoveParams {
	return &LibpodImagesRemoveParams{
		HTTPClient: client,
	}
}

/* LibpodImagesRemoveParams contains all the parameters to send to the API endpoint
   for the libpod images remove operation.

   Typically these are written to a http.Request.
*/
type LibpodImagesRemoveParams struct {

	/* All.

	   Remove all images.

	   Default: true
	*/
	All *bool

	/* Force.

	   Force image removal (including containers using the images).
	*/
	Force *bool

	/* Images.

	   Images IDs or names to remove.
	*/
	Images []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod images remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodImagesRemoveParams) WithDefaults() *LibpodImagesRemoveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod images remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodImagesRemoveParams) SetDefaults() {
	var (
		allDefault = bool(true)
	)

	val := LibpodImagesRemoveParams{
		All: &allDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the libpod images remove params
func (o *LibpodImagesRemoveParams) WithTimeout(timeout time.Duration) *LibpodImagesRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod images remove params
func (o *LibpodImagesRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod images remove params
func (o *LibpodImagesRemoveParams) WithContext(ctx context.Context) *LibpodImagesRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod images remove params
func (o *LibpodImagesRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod images remove params
func (o *LibpodImagesRemoveParams) WithHTTPClient(client *http.Client) *LibpodImagesRemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod images remove params
func (o *LibpodImagesRemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the libpod images remove params
func (o *LibpodImagesRemoveParams) WithAll(all *bool) *LibpodImagesRemoveParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the libpod images remove params
func (o *LibpodImagesRemoveParams) SetAll(all *bool) {
	o.All = all
}

// WithForce adds the force to the libpod images remove params
func (o *LibpodImagesRemoveParams) WithForce(force *bool) *LibpodImagesRemoveParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the libpod images remove params
func (o *LibpodImagesRemoveParams) SetForce(force *bool) {
	o.Force = force
}

// WithImages adds the images to the libpod images remove params
func (o *LibpodImagesRemoveParams) WithImages(images []string) *LibpodImagesRemoveParams {
	o.SetImages(images)
	return o
}

// SetImages adds the images to the libpod images remove params
func (o *LibpodImagesRemoveParams) SetImages(images []string) {
	o.Images = images
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodImagesRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool

		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {

			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}
	}

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

	if o.Images != nil {

		// binding items for images
		joinedImages := o.bindParamImages(reg)

		// query array param images
		if err := r.SetQueryParam("images", joinedImages...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLibpodImagesRemove binds the parameter images
func (o *LibpodImagesRemoveParams) bindParamImages(formats strfmt.Registry) []string {
	imagesIR := o.Images

	var imagesIC []string
	for _, imagesIIR := range imagesIR { // explode []string

		imagesIIV := imagesIIR // string as string
		imagesIC = append(imagesIC, imagesIIV)
	}

	// items.CollectionFormat: ""
	imagesIS := swag.JoinByFormat(imagesIC, "")

	return imagesIS
}
