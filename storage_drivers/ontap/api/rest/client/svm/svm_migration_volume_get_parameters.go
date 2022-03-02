// Code generated by go-swagger; DO NOT EDIT.

package svm

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

// NewSvmMigrationVolumeGetParams creates a new SvmMigrationVolumeGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmMigrationVolumeGetParams() *SvmMigrationVolumeGetParams {
	return &SvmMigrationVolumeGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmMigrationVolumeGetParamsWithTimeout creates a new SvmMigrationVolumeGetParams object
// with the ability to set a timeout on a request.
func NewSvmMigrationVolumeGetParamsWithTimeout(timeout time.Duration) *SvmMigrationVolumeGetParams {
	return &SvmMigrationVolumeGetParams{
		timeout: timeout,
	}
}

// NewSvmMigrationVolumeGetParamsWithContext creates a new SvmMigrationVolumeGetParams object
// with the ability to set a context for a request.
func NewSvmMigrationVolumeGetParamsWithContext(ctx context.Context) *SvmMigrationVolumeGetParams {
	return &SvmMigrationVolumeGetParams{
		Context: ctx,
	}
}

// NewSvmMigrationVolumeGetParamsWithHTTPClient creates a new SvmMigrationVolumeGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmMigrationVolumeGetParamsWithHTTPClient(client *http.Client) *SvmMigrationVolumeGetParams {
	return &SvmMigrationVolumeGetParams{
		HTTPClient: client,
	}
}

/* SvmMigrationVolumeGetParams contains all the parameters to send to the API endpoint
   for the svm migration volume get operation.

   Typically these are written to a http.Request.
*/
type SvmMigrationVolumeGetParams struct {

	/* ErrorsCode.

	   Filter by errors.code
	*/
	ErrorsCodeQueryParameter *int64

	/* ErrorsMessage.

	   Filter by errors.message
	*/
	ErrorsMessageQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Healthy.

	   Filter by healthy
	*/
	HealthyQueryParameter *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* SvmMigrationUUID.

	   Migration UUID
	*/
	SVMMigrationUUIDPathParameter string

	/* TransferState.

	   Filter by transfer_state
	*/
	TransferStateQueryParameter *string

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeNameQueryParameter *string

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm migration volume get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationVolumeGetParams) WithDefaults() *SvmMigrationVolumeGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm migration volume get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationVolumeGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SvmMigrationVolumeGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithTimeout(timeout time.Duration) *SvmMigrationVolumeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithContext(ctx context.Context) *SvmMigrationVolumeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithHTTPClient(client *http.Client) *SvmMigrationVolumeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithErrorsCodeQueryParameter adds the errorsCode to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithErrorsCodeQueryParameter(errorsCode *int64) *SvmMigrationVolumeGetParams {
	o.SetErrorsCodeQueryParameter(errorsCode)
	return o
}

// SetErrorsCodeQueryParameter adds the errorsCode to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetErrorsCodeQueryParameter(errorsCode *int64) {
	o.ErrorsCodeQueryParameter = errorsCode
}

// WithErrorsMessageQueryParameter adds the errorsMessage to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithErrorsMessageQueryParameter(errorsMessage *string) *SvmMigrationVolumeGetParams {
	o.SetErrorsMessageQueryParameter(errorsMessage)
	return o
}

// SetErrorsMessageQueryParameter adds the errorsMessage to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetErrorsMessageQueryParameter(errorsMessage *string) {
	o.ErrorsMessageQueryParameter = errorsMessage
}

// WithFieldsQueryParameter adds the fields to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithFieldsQueryParameter(fields []string) *SvmMigrationVolumeGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithHealthyQueryParameter adds the healthy to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithHealthyQueryParameter(healthy *bool) *SvmMigrationVolumeGetParams {
	o.SetHealthyQueryParameter(healthy)
	return o
}

// SetHealthyQueryParameter adds the healthy to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetHealthyQueryParameter(healthy *bool) {
	o.HealthyQueryParameter = healthy
}

// WithMaxRecordsQueryParameter adds the maxRecords to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SvmMigrationVolumeGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNodeNameQueryParameter adds the nodeName to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithNodeNameQueryParameter(nodeName *string) *SvmMigrationVolumeGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *SvmMigrationVolumeGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderByQueryParameter adds the orderBy to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithOrderByQueryParameter(orderBy []string) *SvmMigrationVolumeGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SvmMigrationVolumeGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SvmMigrationVolumeGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithSVMNameQueryParameter(svmName *string) *SvmMigrationVolumeGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *SvmMigrationVolumeGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithSVMMigrationUUIDPathParameter adds the svmMigrationUUID to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithSVMMigrationUUIDPathParameter(svmMigrationUUID string) *SvmMigrationVolumeGetParams {
	o.SetSVMMigrationUUIDPathParameter(svmMigrationUUID)
	return o
}

// SetSVMMigrationUUIDPathParameter adds the svmMigrationUuid to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetSVMMigrationUUIDPathParameter(svmMigrationUUID string) {
	o.SVMMigrationUUIDPathParameter = svmMigrationUUID
}

// WithTransferStateQueryParameter adds the transferState to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithTransferStateQueryParameter(transferState *string) *SvmMigrationVolumeGetParams {
	o.SetTransferStateQueryParameter(transferState)
	return o
}

// SetTransferStateQueryParameter adds the transferState to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetTransferStateQueryParameter(transferState *string) {
	o.TransferStateQueryParameter = transferState
}

// WithVolumeNameQueryParameter adds the volumeName to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithVolumeNameQueryParameter(volumeName *string) *SvmMigrationVolumeGetParams {
	o.SetVolumeNameQueryParameter(volumeName)
	return o
}

// SetVolumeNameQueryParameter adds the volumeName to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetVolumeNameQueryParameter(volumeName *string) {
	o.VolumeNameQueryParameter = volumeName
}

// WithVolumeUUIDPathParameter adds the volumeUUID to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) WithVolumeUUIDPathParameter(volumeUUID string) *SvmMigrationVolumeGetParams {
	o.SetVolumeUUIDPathParameter(volumeUUID)
	return o
}

// SetVolumeUUIDPathParameter adds the volumeUuid to the svm migration volume get params
func (o *SvmMigrationVolumeGetParams) SetVolumeUUIDPathParameter(volumeUUID string) {
	o.VolumeUUIDPathParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmMigrationVolumeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ErrorsCodeQueryParameter != nil {

		// query param errors.code
		var qrErrorsCode int64

		if o.ErrorsCodeQueryParameter != nil {
			qrErrorsCode = *o.ErrorsCodeQueryParameter
		}
		qErrorsCode := swag.FormatInt64(qrErrorsCode)
		if qErrorsCode != "" {

			if err := r.SetQueryParam("errors.code", qErrorsCode); err != nil {
				return err
			}
		}
	}

	if o.ErrorsMessageQueryParameter != nil {

		// query param errors.message
		var qrErrorsMessage string

		if o.ErrorsMessageQueryParameter != nil {
			qrErrorsMessage = *o.ErrorsMessageQueryParameter
		}
		qErrorsMessage := qrErrorsMessage
		if qErrorsMessage != "" {

			if err := r.SetQueryParam("errors.message", qErrorsMessage); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.HealthyQueryParameter != nil {

		// query param healthy
		var qrHealthy bool

		if o.HealthyQueryParameter != nil {
			qrHealthy = *o.HealthyQueryParameter
		}
		qHealthy := swag.FormatBool(qrHealthy)
		if qHealthy != "" {

			if err := r.SetQueryParam("healthy", qHealthy); err != nil {
				return err
			}
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	// path param svm_migration.uuid
	if err := r.SetPathParam("svm_migration.uuid", o.SVMMigrationUUIDPathParameter); err != nil {
		return err
	}

	if o.TransferStateQueryParameter != nil {

		// query param transfer_state
		var qrTransferState string

		if o.TransferStateQueryParameter != nil {
			qrTransferState = *o.TransferStateQueryParameter
		}
		qTransferState := qrTransferState
		if qTransferState != "" {

			if err := r.SetQueryParam("transfer_state", qTransferState); err != nil {
				return err
			}
		}
	}

	if o.VolumeNameQueryParameter != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeNameQueryParameter != nil {
			qrVolumeName = *o.VolumeNameQueryParameter
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSvmMigrationVolumeGet binds the parameter fields
func (o *SvmMigrationVolumeGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamSvmMigrationVolumeGet binds the parameter order_by
func (o *SvmMigrationVolumeGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}