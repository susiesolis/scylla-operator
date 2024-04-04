// Code generated by go-swagger; DO NOT EDIT.

package operations

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
)

// NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetParams creates a new StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams object
// with the default values initialized.
func NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetParams() *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	var ()
	return &StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetParamsWithTimeout creates a new StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetParamsWithTimeout(timeout time.Duration) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	var ()
	return &StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetParamsWithContext creates a new StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetParamsWithContext(ctx context.Context) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	var ()
	return &StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams{

		Context: ctx,
	}
}

// NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetParamsWithHTTPClient creates a new StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceViewBuildStatusesByKeyspaceAndViewGetParamsWithHTTPClient(client *http.Client) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	var ()
	return &StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams contains all the parameters to send to the API endpoint
for the storage service view build statuses by keyspace and view get operation typically these are written to a http.Request
*/
type StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams struct {

	/*Keyspace
	  The keyspace

	*/
	Keyspace string
	/*View
	  View name

	*/
	View string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) WithTimeout(timeout time.Duration) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) WithContext(ctx context.Context) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) WithHTTPClient(client *http.Client) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyspace adds the keyspace to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) WithKeyspace(keyspace string) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) SetKeyspace(keyspace string) {
	o.Keyspace = keyspace
}

// WithView adds the view to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) WithView(view string) *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams {
	o.SetView(view)
	return o
}

// SetView adds the view to the storage service view build statuses by keyspace and view get params
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) SetView(view string) {
	o.View = view
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceViewBuildStatusesByKeyspaceAndViewGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyspace
	if err := r.SetPathParam("keyspace", o.Keyspace); err != nil {
		return err
	}

	// path param view
	if err := r.SetPathParam("view", o.View); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
