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
	"github.com/go-openapi/swag"
)

// NewLibpodPlayKubeParams creates a new LibpodPlayKubeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodPlayKubeParams() *LibpodPlayKubeParams {
	return &LibpodPlayKubeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodPlayKubeParamsWithTimeout creates a new LibpodPlayKubeParams object
// with the ability to set a timeout on a request.
func NewLibpodPlayKubeParamsWithTimeout(timeout time.Duration) *LibpodPlayKubeParams {
	return &LibpodPlayKubeParams{
		timeout: timeout,
	}
}

// NewLibpodPlayKubeParamsWithContext creates a new LibpodPlayKubeParams object
// with the ability to set a context for a request.
func NewLibpodPlayKubeParamsWithContext(ctx context.Context) *LibpodPlayKubeParams {
	return &LibpodPlayKubeParams{
		Context: ctx,
	}
}

// NewLibpodPlayKubeParamsWithHTTPClient creates a new LibpodPlayKubeParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodPlayKubeParamsWithHTTPClient(client *http.Client) *LibpodPlayKubeParams {
	return &LibpodPlayKubeParams{
		HTTPClient: client,
	}
}

/* LibpodPlayKubeParams contains all the parameters to send to the API endpoint
   for the libpod play kube operation.

   Typically these are written to a http.Request.
*/
type LibpodPlayKubeParams struct {

	/* Network.

	   Connect the pod to this network.
	*/
	Network *string

	/* Request.

	   Kubernetes YAML file.
	*/
	Request string

	/* TLSVerify.

	   Require HTTPS and verify signatures when contating registries.

	   Default: true
	*/
	TLSVerify *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod play kube params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodPlayKubeParams) WithDefaults() *LibpodPlayKubeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod play kube params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodPlayKubeParams) SetDefaults() {
	var (
		tLSVerifyDefault = bool(true)
	)

	val := LibpodPlayKubeParams{
		TLSVerify: &tLSVerifyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the libpod play kube params
func (o *LibpodPlayKubeParams) WithTimeout(timeout time.Duration) *LibpodPlayKubeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod play kube params
func (o *LibpodPlayKubeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod play kube params
func (o *LibpodPlayKubeParams) WithContext(ctx context.Context) *LibpodPlayKubeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod play kube params
func (o *LibpodPlayKubeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod play kube params
func (o *LibpodPlayKubeParams) WithHTTPClient(client *http.Client) *LibpodPlayKubeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod play kube params
func (o *LibpodPlayKubeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetwork adds the network to the libpod play kube params
func (o *LibpodPlayKubeParams) WithNetwork(network *string) *LibpodPlayKubeParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the libpod play kube params
func (o *LibpodPlayKubeParams) SetNetwork(network *string) {
	o.Network = network
}

// WithRequest adds the request to the libpod play kube params
func (o *LibpodPlayKubeParams) WithRequest(request string) *LibpodPlayKubeParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the libpod play kube params
func (o *LibpodPlayKubeParams) SetRequest(request string) {
	o.Request = request
}

// WithTLSVerify adds the tLSVerify to the libpod play kube params
func (o *LibpodPlayKubeParams) WithTLSVerify(tLSVerify *bool) *LibpodPlayKubeParams {
	o.SetTLSVerify(tLSVerify)
	return o
}

// SetTLSVerify adds the tlsVerify to the libpod play kube params
func (o *LibpodPlayKubeParams) SetTLSVerify(tLSVerify *bool) {
	o.TLSVerify = tLSVerify
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodPlayKubeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Network != nil {

		// query param network
		var qrNetwork string

		if o.Network != nil {
			qrNetwork = *o.Network
		}
		qNetwork := qrNetwork
		if qNetwork != "" {

			if err := r.SetQueryParam("network", qNetwork); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if o.TLSVerify != nil {

		// query param tlsVerify
		var qrTLSVerify bool

		if o.TLSVerify != nil {
			qrTLSVerify = *o.TLSVerify
		}
		qTLSVerify := swag.FormatBool(qrTLSVerify)
		if qTLSVerify != "" {

			if err := r.SetQueryParam("tlsVerify", qTLSVerify); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
