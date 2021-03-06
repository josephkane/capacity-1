// Code generated by go-swagger; DO NOT EDIT.

package workers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewCreateWorkerParams creates a new CreateWorkerParams object
// with the default values initialized.
func NewCreateWorkerParams() *CreateWorkerParams {

	return &CreateWorkerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWorkerParamsWithTimeout creates a new CreateWorkerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWorkerParamsWithTimeout(timeout time.Duration) *CreateWorkerParams {

	return &CreateWorkerParams{

		timeout: timeout,
	}
}

// NewCreateWorkerParamsWithContext creates a new CreateWorkerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWorkerParamsWithContext(ctx context.Context) *CreateWorkerParams {

	return &CreateWorkerParams{

		Context: ctx,
	}
}

// NewCreateWorkerParamsWithHTTPClient creates a new CreateWorkerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWorkerParamsWithHTTPClient(client *http.Client) *CreateWorkerParams {

	return &CreateWorkerParams{
		HTTPClient: client,
	}
}

/*CreateWorkerParams contains all the parameters to send to the API endpoint
for the create worker operation typically these are written to a http.Request
*/
type CreateWorkerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create worker params
func (o *CreateWorkerParams) WithTimeout(timeout time.Duration) *CreateWorkerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create worker params
func (o *CreateWorkerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create worker params
func (o *CreateWorkerParams) WithContext(ctx context.Context) *CreateWorkerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create worker params
func (o *CreateWorkerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create worker params
func (o *CreateWorkerParams) WithHTTPClient(client *http.Client) *CreateWorkerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create worker params
func (o *CreateWorkerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWorkerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
