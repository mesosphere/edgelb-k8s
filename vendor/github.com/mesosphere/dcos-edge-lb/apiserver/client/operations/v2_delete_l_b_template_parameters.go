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

// NewV2DeleteLBTemplateParams creates a new V2DeleteLBTemplateParams object
// with the default values initialized.
func NewV2DeleteLBTemplateParams() *V2DeleteLBTemplateParams {
	var ()
	return &V2DeleteLBTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2DeleteLBTemplateParamsWithTimeout creates a new V2DeleteLBTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2DeleteLBTemplateParamsWithTimeout(timeout time.Duration) *V2DeleteLBTemplateParams {
	var ()
	return &V2DeleteLBTemplateParams{

		timeout: timeout,
	}
}

// NewV2DeleteLBTemplateParamsWithContext creates a new V2DeleteLBTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2DeleteLBTemplateParamsWithContext(ctx context.Context) *V2DeleteLBTemplateParams {
	var ()
	return &V2DeleteLBTemplateParams{

		Context: ctx,
	}
}

// NewV2DeleteLBTemplateParamsWithHTTPClient creates a new V2DeleteLBTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2DeleteLBTemplateParamsWithHTTPClient(client *http.Client) *V2DeleteLBTemplateParams {
	var ()
	return &V2DeleteLBTemplateParams{
		HTTPClient: client,
	}
}

/*V2DeleteLBTemplateParams contains all the parameters to send to the API endpoint
for the v2 delete l b template operation typically these are written to a http.Request
*/
type V2DeleteLBTemplateParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 delete l b template params
func (o *V2DeleteLBTemplateParams) WithTimeout(timeout time.Duration) *V2DeleteLBTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 delete l b template params
func (o *V2DeleteLBTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 delete l b template params
func (o *V2DeleteLBTemplateParams) WithContext(ctx context.Context) *V2DeleteLBTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 delete l b template params
func (o *V2DeleteLBTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 delete l b template params
func (o *V2DeleteLBTemplateParams) WithHTTPClient(client *http.Client) *V2DeleteLBTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 delete l b template params
func (o *V2DeleteLBTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 delete l b template params
func (o *V2DeleteLBTemplateParams) WithName(name string) *V2DeleteLBTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 delete l b template params
func (o *V2DeleteLBTemplateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V2DeleteLBTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
