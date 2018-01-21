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

// NewV1GetLoadBalancerPoolsParams creates a new V1GetLoadBalancerPoolsParams object
// with the default values initialized.
func NewV1GetLoadBalancerPoolsParams() *V1GetLoadBalancerPoolsParams {

	return &V1GetLoadBalancerPoolsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1GetLoadBalancerPoolsParamsWithTimeout creates a new V1GetLoadBalancerPoolsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1GetLoadBalancerPoolsParamsWithTimeout(timeout time.Duration) *V1GetLoadBalancerPoolsParams {

	return &V1GetLoadBalancerPoolsParams{

		timeout: timeout,
	}
}

// NewV1GetLoadBalancerPoolsParamsWithContext creates a new V1GetLoadBalancerPoolsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1GetLoadBalancerPoolsParamsWithContext(ctx context.Context) *V1GetLoadBalancerPoolsParams {

	return &V1GetLoadBalancerPoolsParams{

		Context: ctx,
	}
}

// NewV1GetLoadBalancerPoolsParamsWithHTTPClient creates a new V1GetLoadBalancerPoolsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1GetLoadBalancerPoolsParamsWithHTTPClient(client *http.Client) *V1GetLoadBalancerPoolsParams {

	return &V1GetLoadBalancerPoolsParams{
		HTTPClient: client,
	}
}

/*V1GetLoadBalancerPoolsParams contains all the parameters to send to the API endpoint
for the v1 get load balancer pools operation typically these are written to a http.Request
*/
type V1GetLoadBalancerPoolsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 get load balancer pools params
func (o *V1GetLoadBalancerPoolsParams) WithTimeout(timeout time.Duration) *V1GetLoadBalancerPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 get load balancer pools params
func (o *V1GetLoadBalancerPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 get load balancer pools params
func (o *V1GetLoadBalancerPoolsParams) WithContext(ctx context.Context) *V1GetLoadBalancerPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 get load balancer pools params
func (o *V1GetLoadBalancerPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 get load balancer pools params
func (o *V1GetLoadBalancerPoolsParams) WithHTTPClient(client *http.Client) *V1GetLoadBalancerPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 get load balancer pools params
func (o *V1GetLoadBalancerPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1GetLoadBalancerPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
