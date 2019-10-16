// Code generated by go-swagger; DO NOT EDIT.

package inputs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteV1BucketsBucketIDInputsInputIDParams creates a new DeleteV1BucketsBucketIDInputsInputIDParams object
// with the default values initialized.
func NewDeleteV1BucketsBucketIDInputsInputIDParams() *DeleteV1BucketsBucketIDInputsInputIDParams {
	var ()
	return &DeleteV1BucketsBucketIDInputsInputIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1BucketsBucketIDInputsInputIDParamsWithTimeout creates a new DeleteV1BucketsBucketIDInputsInputIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1BucketsBucketIDInputsInputIDParamsWithTimeout(timeout time.Duration) *DeleteV1BucketsBucketIDInputsInputIDParams {
	var ()
	return &DeleteV1BucketsBucketIDInputsInputIDParams{

		timeout: timeout,
	}
}

// NewDeleteV1BucketsBucketIDInputsInputIDParamsWithContext creates a new DeleteV1BucketsBucketIDInputsInputIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1BucketsBucketIDInputsInputIDParamsWithContext(ctx context.Context) *DeleteV1BucketsBucketIDInputsInputIDParams {
	var ()
	return &DeleteV1BucketsBucketIDInputsInputIDParams{

		Context: ctx,
	}
}

// NewDeleteV1BucketsBucketIDInputsInputIDParamsWithHTTPClient creates a new DeleteV1BucketsBucketIDInputsInputIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV1BucketsBucketIDInputsInputIDParamsWithHTTPClient(client *http.Client) *DeleteV1BucketsBucketIDInputsInputIDParams {
	var ()
	return &DeleteV1BucketsBucketIDInputsInputIDParams{
		HTTPClient: client,
	}
}

/*DeleteV1BucketsBucketIDInputsInputIDParams contains all the parameters to send to the API endpoint
for the delete v1 buckets bucket ID inputs input ID operation typically these are written to a http.Request
*/
type DeleteV1BucketsBucketIDInputsInputIDParams struct {

	/*BucketID
	  ID of a bucket to create input in

	*/
	BucketID string
	/*InputID
	  ID of an input to delete

	*/
	InputID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) WithTimeout(timeout time.Duration) *DeleteV1BucketsBucketIDInputsInputIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) WithContext(ctx context.Context) *DeleteV1BucketsBucketIDInputsInputIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) WithHTTPClient(client *http.Client) *DeleteV1BucketsBucketIDInputsInputIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) WithBucketID(bucketID string) *DeleteV1BucketsBucketIDInputsInputIDParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WithInputID adds the inputID to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) WithInputID(inputID string) *DeleteV1BucketsBucketIDInputsInputIDParams {
	o.SetInputID(inputID)
	return o
}

// SetInputID adds the inputId to the delete v1 buckets bucket ID inputs input ID params
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) SetInputID(inputID string) {
	o.InputID = inputID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1BucketsBucketIDInputsInputIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucketID
	if err := r.SetPathParam("bucketID", o.BucketID); err != nil {
		return err
	}

	// path param inputID
	if err := r.SetPathParam("inputID", o.InputID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
