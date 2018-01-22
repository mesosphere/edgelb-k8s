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

	"github.com/mesosphere/dcos-edge-lb/apiserver/models"
)

// NewV1UpdateConfigParams creates a new V1UpdateConfigParams object
// with the default values initialized.
func NewV1UpdateConfigParams() *V1UpdateConfigParams {
	var ()
	return &V1UpdateConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UpdateConfigParamsWithTimeout creates a new V1UpdateConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UpdateConfigParamsWithTimeout(timeout time.Duration) *V1UpdateConfigParams {
	var ()
	return &V1UpdateConfigParams{

		timeout: timeout,
	}
}

// NewV1UpdateConfigParamsWithContext creates a new V1UpdateConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UpdateConfigParamsWithContext(ctx context.Context) *V1UpdateConfigParams {
	var ()
	return &V1UpdateConfigParams{

		Context: ctx,
	}
}

// NewV1UpdateConfigParamsWithHTTPClient creates a new V1UpdateConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UpdateConfigParamsWithHTTPClient(client *http.Client) *V1UpdateConfigParams {
	var ()
	return &V1UpdateConfigParams{
		HTTPClient: client,
	}
}

/*V1UpdateConfigParams contains all the parameters to send to the API endpoint
for the v1 update config operation typically these are written to a http.Request
*/
type V1UpdateConfigParams struct {

	/*Config*/
	Config *models.V1Config
	/*Token
	  DCOS Auth Token

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 update config params
func (o *V1UpdateConfigParams) WithTimeout(timeout time.Duration) *V1UpdateConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 update config params
func (o *V1UpdateConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 update config params
func (o *V1UpdateConfigParams) WithContext(ctx context.Context) *V1UpdateConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 update config params
func (o *V1UpdateConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 update config params
func (o *V1UpdateConfigParams) WithHTTPClient(client *http.Client) *V1UpdateConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 update config params
func (o *V1UpdateConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the v1 update config params
func (o *V1UpdateConfigParams) WithConfig(config *models.V1Config) *V1UpdateConfigParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the v1 update config params
func (o *V1UpdateConfigParams) SetConfig(config *models.V1Config) {
	o.Config = config
}

// WithToken adds the token to the v1 update config params
func (o *V1UpdateConfigParams) WithToken(token *string) *V1UpdateConfigParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the v1 update config params
func (o *V1UpdateConfigParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *V1UpdateConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config == nil {
		o.Config = new(models.V1Config)
	}

	if err := r.SetBodyParam(o.Config); err != nil {
		return err
	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}