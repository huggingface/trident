// Code generated by go-swagger; DO NOT EDIT.

package cluster

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
	"github.com/go-openapi/swag"
)

// NewSoftwarePackageGetParams creates a new SoftwarePackageGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSoftwarePackageGetParams() *SoftwarePackageGetParams {
	return &SoftwarePackageGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwarePackageGetParamsWithTimeout creates a new SoftwarePackageGetParams object
// with the ability to set a timeout on a request.
func NewSoftwarePackageGetParamsWithTimeout(timeout time.Duration) *SoftwarePackageGetParams {
	return &SoftwarePackageGetParams{
		timeout: timeout,
	}
}

// NewSoftwarePackageGetParamsWithContext creates a new SoftwarePackageGetParams object
// with the ability to set a context for a request.
func NewSoftwarePackageGetParamsWithContext(ctx context.Context) *SoftwarePackageGetParams {
	return &SoftwarePackageGetParams{
		Context: ctx,
	}
}

// NewSoftwarePackageGetParamsWithHTTPClient creates a new SoftwarePackageGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSoftwarePackageGetParamsWithHTTPClient(client *http.Client) *SoftwarePackageGetParams {
	return &SoftwarePackageGetParams{
		HTTPClient: client,
	}
}

/* SoftwarePackageGetParams contains all the parameters to send to the API endpoint
   for the software package get operation.

   Typically these are written to a http.Request.
*/
type SoftwarePackageGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	// Version.
	VersionPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the software package get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwarePackageGetParams) WithDefaults() *SoftwarePackageGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the software package get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwarePackageGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the software package get params
func (o *SoftwarePackageGetParams) WithTimeout(timeout time.Duration) *SoftwarePackageGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software package get params
func (o *SoftwarePackageGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software package get params
func (o *SoftwarePackageGetParams) WithContext(ctx context.Context) *SoftwarePackageGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software package get params
func (o *SoftwarePackageGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software package get params
func (o *SoftwarePackageGetParams) WithHTTPClient(client *http.Client) *SoftwarePackageGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software package get params
func (o *SoftwarePackageGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the software package get params
func (o *SoftwarePackageGetParams) WithFields(fields []string) *SoftwarePackageGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the software package get params
func (o *SoftwarePackageGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithVersionPathParameter adds the version to the software package get params
func (o *SoftwarePackageGetParams) WithVersionPathParameter(version string) *SoftwarePackageGetParams {
	o.SetVersionPathParameter(version)
	return o
}

// SetVersionPathParameter adds the version to the software package get params
func (o *SoftwarePackageGetParams) SetVersionPathParameter(version string) {
	o.VersionPathParameter = version
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwarePackageGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param version
	if err := r.SetPathParam("version", o.VersionPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSoftwarePackageGet binds the parameter fields
func (o *SoftwarePackageGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}