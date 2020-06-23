// Code generated by go-swagger; DO NOT EDIT.

package functions

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

	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

// NewPostV1FunctionsParams creates a new PostV1FunctionsParams object
// with the default values initialized.
func NewPostV1FunctionsParams() *PostV1FunctionsParams {
	var ()
	return &PostV1FunctionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1FunctionsParamsWithTimeout creates a new PostV1FunctionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostV1FunctionsParamsWithTimeout(timeout time.Duration) *PostV1FunctionsParams {
	var ()
	return &PostV1FunctionsParams{

		timeout: timeout,
	}
}

// NewPostV1FunctionsParamsWithContext creates a new PostV1FunctionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostV1FunctionsParamsWithContext(ctx context.Context) *PostV1FunctionsParams {
	var ()
	return &PostV1FunctionsParams{

		Context: ctx,
	}
}

// NewPostV1FunctionsParamsWithHTTPClient creates a new PostV1FunctionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostV1FunctionsParamsWithHTTPClient(client *http.Client) *PostV1FunctionsParams {
	var ()
	return &PostV1FunctionsParams{
		HTTPClient: client,
	}
}

/*PostV1FunctionsParams contains all the parameters to send to the API endpoint
for the post v1 functions operation typically these are written to a http.Request
*/
type PostV1FunctionsParams struct {

	/*Body*/
	Body *models.FunctionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post v1 functions params
func (o *PostV1FunctionsParams) WithTimeout(timeout time.Duration) *PostV1FunctionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 functions params
func (o *PostV1FunctionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 functions params
func (o *PostV1FunctionsParams) WithContext(ctx context.Context) *PostV1FunctionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 functions params
func (o *PostV1FunctionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 functions params
func (o *PostV1FunctionsParams) WithHTTPClient(client *http.Client) *PostV1FunctionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 functions params
func (o *PostV1FunctionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post v1 functions params
func (o *PostV1FunctionsParams) WithBody(body *models.FunctionRequest) *PostV1FunctionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 functions params
func (o *PostV1FunctionsParams) SetBody(body *models.FunctionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1FunctionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
