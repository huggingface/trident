// Code generated by go-swagger; DO NOT EDIT.

package storage

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewVolumeModifyParams creates a new VolumeModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeModifyParams() *VolumeModifyParams {
	return &VolumeModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeModifyParamsWithTimeout creates a new VolumeModifyParams object
// with the ability to set a timeout on a request.
func NewVolumeModifyParamsWithTimeout(timeout time.Duration) *VolumeModifyParams {
	return &VolumeModifyParams{
		timeout: timeout,
	}
}

// NewVolumeModifyParamsWithContext creates a new VolumeModifyParams object
// with the ability to set a context for a request.
func NewVolumeModifyParamsWithContext(ctx context.Context) *VolumeModifyParams {
	return &VolumeModifyParams{
		Context: ctx,
	}
}

// NewVolumeModifyParamsWithHTTPClient creates a new VolumeModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeModifyParamsWithHTTPClient(client *http.Client) *VolumeModifyParams {
	return &VolumeModifyParams{
		HTTPClient: client,
	}
}

/* VolumeModifyParams contains all the parameters to send to the API endpoint
   for the volume modify operation.

   Typically these are written to a http.Request.
*/
type VolumeModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Volume

	/* RestoreToSnapshotName.

	   Name of the Snapshot copy to restore volume to the point in time the Snapshot copy was taken.
	*/
	RestoreToSnapshotNameQueryParameter *string

	/* RestoreToSnapshotUUID.

	   UUID of the Snapshot copy to restore volume to the point in time the Snapshot copy was taken.
	*/
	RestoreToSnapshotUUIDQueryParameter *string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* SizingMethod.

	     Represents the method to modify the size of a Flexgroup. The following methods are supported:
	* use_existing_resources - Increases or decreases the size of the FlexGroup by increasing or decreasing the size of the current FlexGroup resources
	* add_new_resources - Increases the size of the FlexGroup by adding new resources


	     Default: "use_existing_resources"
	*/
	SizingMethodQueryParameter *string

	/* UUID.

	   Unique identifier of the volume.
	*/
	UUIDPathParameter string

	/* ValidateOnly.

	   Validate the operation and its parameters, without actually performing the operation.
	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeModifyParams) WithDefaults() *VolumeModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)

		sizingMethodQueryParameterDefault = string("use_existing_resources")
	)

	val := VolumeModifyParams{
		ReturnTimeout:              &returnTimeoutDefault,
		SizingMethodQueryParameter: &sizingMethodQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the volume modify params
func (o *VolumeModifyParams) WithTimeout(timeout time.Duration) *VolumeModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume modify params
func (o *VolumeModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume modify params
func (o *VolumeModifyParams) WithContext(ctx context.Context) *VolumeModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume modify params
func (o *VolumeModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume modify params
func (o *VolumeModifyParams) WithHTTPClient(client *http.Client) *VolumeModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume modify params
func (o *VolumeModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the volume modify params
func (o *VolumeModifyParams) WithInfo(info *models.Volume) *VolumeModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the volume modify params
func (o *VolumeModifyParams) SetInfo(info *models.Volume) {
	o.Info = info
}

// WithRestoreToSnapshotNameQueryParameter adds the restoreToSnapshotName to the volume modify params
func (o *VolumeModifyParams) WithRestoreToSnapshotNameQueryParameter(restoreToSnapshotName *string) *VolumeModifyParams {
	o.SetRestoreToSnapshotNameQueryParameter(restoreToSnapshotName)
	return o
}

// SetRestoreToSnapshotNameQueryParameter adds the restoreToSnapshotName to the volume modify params
func (o *VolumeModifyParams) SetRestoreToSnapshotNameQueryParameter(restoreToSnapshotName *string) {
	o.RestoreToSnapshotNameQueryParameter = restoreToSnapshotName
}

// WithRestoreToSnapshotUUIDQueryParameter adds the restoreToSnapshotUUID to the volume modify params
func (o *VolumeModifyParams) WithRestoreToSnapshotUUIDQueryParameter(restoreToSnapshotUUID *string) *VolumeModifyParams {
	o.SetRestoreToSnapshotUUIDQueryParameter(restoreToSnapshotUUID)
	return o
}

// SetRestoreToSnapshotUUIDQueryParameter adds the restoreToSnapshotUuid to the volume modify params
func (o *VolumeModifyParams) SetRestoreToSnapshotUUIDQueryParameter(restoreToSnapshotUUID *string) {
	o.RestoreToSnapshotUUIDQueryParameter = restoreToSnapshotUUID
}

// WithReturnTimeout adds the returnTimeout to the volume modify params
func (o *VolumeModifyParams) WithReturnTimeout(returnTimeout *int64) *VolumeModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the volume modify params
func (o *VolumeModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSizingMethodQueryParameter adds the sizingMethod to the volume modify params
func (o *VolumeModifyParams) WithSizingMethodQueryParameter(sizingMethod *string) *VolumeModifyParams {
	o.SetSizingMethodQueryParameter(sizingMethod)
	return o
}

// SetSizingMethodQueryParameter adds the sizingMethod to the volume modify params
func (o *VolumeModifyParams) SetSizingMethodQueryParameter(sizingMethod *string) {
	o.SizingMethodQueryParameter = sizingMethod
}

// WithUUIDPathParameter adds the uuid to the volume modify params
func (o *VolumeModifyParams) WithUUIDPathParameter(uuid string) *VolumeModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the volume modify params
func (o *VolumeModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WithValidateOnly adds the validateOnly to the volume modify params
func (o *VolumeModifyParams) WithValidateOnly(validateOnly *bool) *VolumeModifyParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the volume modify params
func (o *VolumeModifyParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.RestoreToSnapshotNameQueryParameter != nil {

		// query param restore_to.snapshot.name
		var qrRestoreToSnapshotName string

		if o.RestoreToSnapshotNameQueryParameter != nil {
			qrRestoreToSnapshotName = *o.RestoreToSnapshotNameQueryParameter
		}
		qRestoreToSnapshotName := qrRestoreToSnapshotName
		if qRestoreToSnapshotName != "" {

			if err := r.SetQueryParam("restore_to.snapshot.name", qRestoreToSnapshotName); err != nil {
				return err
			}
		}
	}

	if o.RestoreToSnapshotUUIDQueryParameter != nil {

		// query param restore_to.snapshot.uuid
		var qrRestoreToSnapshotUUID string

		if o.RestoreToSnapshotUUIDQueryParameter != nil {
			qrRestoreToSnapshotUUID = *o.RestoreToSnapshotUUIDQueryParameter
		}
		qRestoreToSnapshotUUID := qrRestoreToSnapshotUUID
		if qRestoreToSnapshotUUID != "" {

			if err := r.SetQueryParam("restore_to.snapshot.uuid", qRestoreToSnapshotUUID); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SizingMethodQueryParameter != nil {

		// query param sizing_method
		var qrSizingMethod string

		if o.SizingMethodQueryParameter != nil {
			qrSizingMethod = *o.SizingMethodQueryParameter
		}
		qSizingMethod := qrSizingMethod
		if qSizingMethod != "" {

			if err := r.SetQueryParam("sizing_method", qSizingMethod); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool

		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {

			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}