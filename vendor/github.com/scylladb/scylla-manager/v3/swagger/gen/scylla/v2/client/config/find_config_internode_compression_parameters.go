// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigInternodeCompressionParams creates a new FindConfigInternodeCompressionParams object
// with the default values initialized.
func NewFindConfigInternodeCompressionParams() *FindConfigInternodeCompressionParams {

	return &FindConfigInternodeCompressionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigInternodeCompressionParamsWithTimeout creates a new FindConfigInternodeCompressionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigInternodeCompressionParamsWithTimeout(timeout time.Duration) *FindConfigInternodeCompressionParams {

	return &FindConfigInternodeCompressionParams{

		timeout: timeout,
	}
}

// NewFindConfigInternodeCompressionParamsWithContext creates a new FindConfigInternodeCompressionParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigInternodeCompressionParamsWithContext(ctx context.Context) *FindConfigInternodeCompressionParams {

	return &FindConfigInternodeCompressionParams{

		Context: ctx,
	}
}

// NewFindConfigInternodeCompressionParamsWithHTTPClient creates a new FindConfigInternodeCompressionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigInternodeCompressionParamsWithHTTPClient(client *http.Client) *FindConfigInternodeCompressionParams {

	return &FindConfigInternodeCompressionParams{
		HTTPClient: client,
	}
}

/*
FindConfigInternodeCompressionParams contains all the parameters to send to the API endpoint
for the find config internode compression operation typically these are written to a http.Request
*/
type FindConfigInternodeCompressionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config internode compression params
func (o *FindConfigInternodeCompressionParams) WithTimeout(timeout time.Duration) *FindConfigInternodeCompressionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config internode compression params
func (o *FindConfigInternodeCompressionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config internode compression params
func (o *FindConfigInternodeCompressionParams) WithContext(ctx context.Context) *FindConfigInternodeCompressionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config internode compression params
func (o *FindConfigInternodeCompressionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config internode compression params
func (o *FindConfigInternodeCompressionParams) WithHTTPClient(client *http.Client) *FindConfigInternodeCompressionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config internode compression params
func (o *FindConfigInternodeCompressionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigInternodeCompressionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}