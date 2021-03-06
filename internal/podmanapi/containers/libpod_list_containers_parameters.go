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

// NewLibpodListContainersParams creates a new LibpodListContainersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLibpodListContainersParams() *LibpodListContainersParams {
	return &LibpodListContainersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodListContainersParamsWithTimeout creates a new LibpodListContainersParams object
// with the ability to set a timeout on a request.
func NewLibpodListContainersParamsWithTimeout(timeout time.Duration) *LibpodListContainersParams {
	return &LibpodListContainersParams{
		timeout: timeout,
	}
}

// NewLibpodListContainersParamsWithContext creates a new LibpodListContainersParams object
// with the ability to set a context for a request.
func NewLibpodListContainersParamsWithContext(ctx context.Context) *LibpodListContainersParams {
	return &LibpodListContainersParams{
		Context: ctx,
	}
}

// NewLibpodListContainersParamsWithHTTPClient creates a new LibpodListContainersParams object
// with the ability to set a custom HTTPClient for a request.
func NewLibpodListContainersParamsWithHTTPClient(client *http.Client) *LibpodListContainersParams {
	return &LibpodListContainersParams{
		HTTPClient: client,
	}
}

/* LibpodListContainersParams contains all the parameters to send to the API endpoint
   for the libpod list containers operation.

   Typically these are written to a http.Request.
*/
type LibpodListContainersParams struct {

	/* All.

	   Return all containers. By default, only running containers are shown
	*/
	All *bool

	/* Filters.

	     A JSON encoded value of the filters (a `map[string][]string`) to process on the containers list. Available filters:
	- `ancestor`=(`<image-name>[:<tag>]`, `<image id>`, or `<image@digest>`)
	- `before`=(`<container id>` or `<container name>`)
	- `expose`=(`<port>[/<proto>]` or `<startport-endport>/[<proto>]`)
	- `exited=<int>` containers with exit code of `<int>`
	- `health`=(`starting`, `healthy`, `unhealthy` or `none`)
	- `id=<ID>` a container's ID
	- `is-task`=(`true` or `false`)
	- `label`=(`key` or `"key=value"`) of an container label
	- `name=<name>` a container's name
	- `network`=(`<network id>` or `<network name>`)
	- `publish`=(`<port>[/<proto>]` or `<startport-endport>/[<proto>]`)
	- `since`=(`<container id>` or `<container name>`)
	- `status`=(`created`, `restarting`, `running`, `removing`, `paused`, `exited` or `dead`)
	- `volume`=(`<volume name>` or `<mount point destination>`)

	*/
	Filters *string

	/* Limit.

	   Return this number of most recently created containers, including non-running ones.
	*/
	Limit *int64

	/* Pod.

	   Ignored. Previously included details on pod name and ID that are currently included by default.
	*/
	Pod *bool

	/* Size.

	   Return the size of container as fields SizeRw and SizeRootFs.
	*/
	Size *bool

	/* Sync.

	   Sync container state with OCI runtime
	*/
	Sync *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the libpod list containers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodListContainersParams) WithDefaults() *LibpodListContainersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the libpod list containers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LibpodListContainersParams) SetDefaults() {
	var (
		allDefault = bool(false)

		podDefault = bool(false)

		sizeDefault = bool(false)

		syncDefault = bool(false)
	)

	val := LibpodListContainersParams{
		All:  &allDefault,
		Pod:  &podDefault,
		Size: &sizeDefault,
		Sync: &syncDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the libpod list containers params
func (o *LibpodListContainersParams) WithTimeout(timeout time.Duration) *LibpodListContainersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod list containers params
func (o *LibpodListContainersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod list containers params
func (o *LibpodListContainersParams) WithContext(ctx context.Context) *LibpodListContainersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod list containers params
func (o *LibpodListContainersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod list containers params
func (o *LibpodListContainersParams) WithHTTPClient(client *http.Client) *LibpodListContainersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod list containers params
func (o *LibpodListContainersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the libpod list containers params
func (o *LibpodListContainersParams) WithAll(all *bool) *LibpodListContainersParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the libpod list containers params
func (o *LibpodListContainersParams) SetAll(all *bool) {
	o.All = all
}

// WithFilters adds the filters to the libpod list containers params
func (o *LibpodListContainersParams) WithFilters(filters *string) *LibpodListContainersParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the libpod list containers params
func (o *LibpodListContainersParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithLimit adds the limit to the libpod list containers params
func (o *LibpodListContainersParams) WithLimit(limit *int64) *LibpodListContainersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the libpod list containers params
func (o *LibpodListContainersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPod adds the pod to the libpod list containers params
func (o *LibpodListContainersParams) WithPod(pod *bool) *LibpodListContainersParams {
	o.SetPod(pod)
	return o
}

// SetPod adds the pod to the libpod list containers params
func (o *LibpodListContainersParams) SetPod(pod *bool) {
	o.Pod = pod
}

// WithSize adds the size to the libpod list containers params
func (o *LibpodListContainersParams) WithSize(size *bool) *LibpodListContainersParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the libpod list containers params
func (o *LibpodListContainersParams) SetSize(size *bool) {
	o.Size = size
}

// WithSync adds the sync to the libpod list containers params
func (o *LibpodListContainersParams) WithSync(sync *bool) *LibpodListContainersParams {
	o.SetSync(sync)
	return o
}

// SetSync adds the sync to the libpod list containers params
func (o *LibpodListContainersParams) SetSync(sync *bool) {
	o.Sync = sync
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodListContainersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Filters != nil {

		// query param filters
		var qrFilters string

		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {

			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Pod != nil {

		// query param pod
		var qrPod bool

		if o.Pod != nil {
			qrPod = *o.Pod
		}
		qPod := swag.FormatBool(qrPod)
		if qPod != "" {

			if err := r.SetQueryParam("pod", qPod); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize bool

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatBool(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Sync != nil {

		// query param sync
		var qrSync bool

		if o.Sync != nil {
			qrSync = *o.Sync
		}
		qSync := swag.FormatBool(qrSync)
		if qSync != "" {

			if err := r.SetQueryParam("sync", qSync); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
