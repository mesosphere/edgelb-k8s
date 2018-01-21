// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewV1GetLoadBalancerPoolParams creates a new V1GetLoadBalancerPoolParams object
// with the default values initialized.
func NewV1GetLoadBalancerPoolParams() *V1GetLoadBalancerPoolParams {
	var ()
	return &V1GetLoadBalancerPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1GetLoadBalancerPoolParamsWithTimeout creates a new V1GetLoadBalancerPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1GetLoadBalancerPoolParamsWithTimeout(timeout time.Duration) *V1GetLoadBalancerPoolParams {
	var ()
	return &V1GetLoadBalancerPoolParams{

		timeout: timeout,
	}
}

// NewV1GetLoadBalancerPoolParamsWithContext creates a new V1GetLoadBalancerPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1GetLoadBalancerPoolParamsWithContext(ctx context.Context) *V1GetLoadBalancerPoolParams {
	var ()
	return &V1GetLoadBalancerPoolParams{

		Context: ctx,
	}
}

// NewV1GetLoadBalancerPoolParamsWithHTTPClient creates a new V1GetLoadBalancerPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1GetLoadBalancerPoolParamsWithHTTPClient(client *http.Client) *V1GetLoadBalancerPoolParams {
	var ()
	return &V1GetLoadBalancerPoolParams{
		HTTPClient: client,
	}
}

/*V1GetLoadBalancerPoolParams contains all the parameters to send to the API endpoint
for the v1 get load balancer pool operation typically these are written to a http.Request
*/
type V1GetLoadBalancerPoolParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 get load balancer pool params
func (o *V1GetLoadBalancerPoolParams) WithTimeout(timeout time.Duration) *V1GetLoadBalancerPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 get load balancer pool params
func (o *V1GetLoadBalancerPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 get load balancer pool params
func (o *V1GetLoadBalancerPoolParams) WithContext(ctx context.Context) *V1GetLoadBalancerPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 get load balancer pool params
func (o *V1GetLoadBalancerPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 get load balancer pool params
func (o *V1GetLoadBalancerPoolParams) WithHTTPClient(client *http.Client) *V1GetLoadBalancerPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 get load balancer pool params
func (o *V1GetLoadBalancerPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v1 get load balancer pool params
func (o *V1GetLoadBalancerPoolParams) WithName(name string) *V1GetLoadBalancerPoolParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v1 get load balancer pool params
func (o *V1GetLoadBalancerPoolParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V1GetLoadBalancerPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
