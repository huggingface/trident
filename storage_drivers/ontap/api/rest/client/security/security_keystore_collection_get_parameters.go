// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityKeystoreCollectionGetParams creates a new SecurityKeystoreCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeystoreCollectionGetParams() *SecurityKeystoreCollectionGetParams {
	return &SecurityKeystoreCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeystoreCollectionGetParamsWithTimeout creates a new SecurityKeystoreCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSecurityKeystoreCollectionGetParamsWithTimeout(timeout time.Duration) *SecurityKeystoreCollectionGetParams {
	return &SecurityKeystoreCollectionGetParams{
		timeout: timeout,
	}
}

// NewSecurityKeystoreCollectionGetParamsWithContext creates a new SecurityKeystoreCollectionGetParams object
// with the ability to set a context for a request.
func NewSecurityKeystoreCollectionGetParamsWithContext(ctx context.Context) *SecurityKeystoreCollectionGetParams {
	return &SecurityKeystoreCollectionGetParams{
		Context: ctx,
	}
}

// NewSecurityKeystoreCollectionGetParamsWithHTTPClient creates a new SecurityKeystoreCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeystoreCollectionGetParamsWithHTTPClient(client *http.Client) *SecurityKeystoreCollectionGetParams {
	return &SecurityKeystoreCollectionGetParams{
		HTTPClient: client,
	}
}

/* SecurityKeystoreCollectionGetParams contains all the parameters to send to the API endpoint
   for the security keystore collection get operation.

   Typically these are written to a http.Request.
*/
type SecurityKeystoreCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Location.

	   Filter by location
	*/
	LocationQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

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

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security keystore collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeystoreCollectionGetParams) WithDefaults() *SecurityKeystoreCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security keystore collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeystoreCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SecurityKeystoreCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithTimeout(timeout time.Duration) *SecurityKeystoreCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithContext(ctx context.Context) *SecurityKeystoreCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithHTTPClient(client *http.Client) *SecurityKeystoreCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithFieldsQueryParameter(fields []string) *SecurityKeystoreCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithLocationQueryParameter adds the location to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithLocationQueryParameter(location *string) *SecurityKeystoreCollectionGetParams {
	o.SetLocationQueryParameter(location)
	return o
}

// SetLocationQueryParameter adds the location to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetLocationQueryParameter(location *string) {
	o.LocationQueryParameter = location
}

// WithMaxRecordsQueryParameter adds the maxRecords to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SecurityKeystoreCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *SecurityKeystoreCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SecurityKeystoreCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SecurityKeystoreCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *SecurityKeystoreCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *SecurityKeystoreCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTypeQueryParameter adds the typeVar to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithTypeQueryParameter(typeVar *string) *SecurityKeystoreCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithUUIDQueryParameter adds the uuid to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) WithUUIDQueryParameter(uuid *string) *SecurityKeystoreCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the security keystore collection get params
func (o *SecurityKeystoreCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeystoreCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.LocationQueryParameter != nil {

		// query param location
		var qrLocation string

		if o.LocationQueryParameter != nil {
			qrLocation = *o.LocationQueryParameter
		}
		qLocation := qrLocation
		if qLocation != "" {

			if err := r.SetQueryParam("location", qLocation); err != nil {
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

	if o.TypeQueryParameter != nil {

		// query param type
		var qrType string

		if o.TypeQueryParameter != nil {
			qrType = *o.TypeQueryParameter
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityKeystoreCollectionGet binds the parameter fields
func (o *SecurityKeystoreCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSecurityKeystoreCollectionGet binds the parameter order_by
func (o *SecurityKeystoreCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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