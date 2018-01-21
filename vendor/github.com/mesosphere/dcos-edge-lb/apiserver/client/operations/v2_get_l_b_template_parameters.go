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

// NewV2GetLBTemplateParams creates a new V2GetLBTemplateParams object
// with the default values initialized.
func NewV2GetLBTemplateParams() *V2GetLBTemplateParams {
	var ()
	return &V2GetLBTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetLBTemplateParamsWithTimeout creates a new V2GetLBTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2GetLBTemplateParamsWithTimeout(timeout time.Duration) *V2GetLBTemplateParams {
	var ()
	return &V2GetLBTemplateParams{

		timeout: timeout,
	}
}

// NewV2GetLBTemplateParamsWithContext creates a new V2GetLBTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2GetLBTemplateParamsWithContext(ctx context.Context) *V2GetLBTemplateParams {
	var ()
	return &V2GetLBTemplateParams{

		Context: ctx,
	}
}

// NewV2GetLBTemplateParamsWithHTTPClient creates a new V2GetLBTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2GetLBTemplateParamsWithHTTPClient(client *http.Client) *V2GetLBTemplateParams {
	var ()
	return &V2GetLBTemplateParams{
		HTTPClient: client,
	}
}

/*V2GetLBTemplateParams contains all the parameters to send to the API endpoint
for the v2 get l b template operation typically these are written to a http.Request
*/
type V2GetLBTemplateParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 get l b template params
func (o *V2GetLBTemplateParams) WithTimeout(timeout time.Duration) *V2GetLBTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get l b template params
func (o *V2GetLBTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get l b template params
func (o *V2GetLBTemplateParams) WithContext(ctx context.Context) *V2GetLBTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get l b template params
func (o *V2GetLBTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get l b template params
func (o *V2GetLBTemplateParams) WithHTTPClient(client *http.Client) *V2GetLBTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get l b template params
func (o *V2GetLBTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 get l b template params
func (o *V2GetLBTemplateParams) WithName(name string) *V2GetLBTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 get l b template params
func (o *V2GetLBTemplateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetLBTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
