// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FileAccessFilter ONTAP allows creation of filters for file access tracing for both CIFS and NFS. These filters have protocols, path, username  and client IP based on which file access operations are logged.
//
// swagger:model file_access_filter
type FileAccessFilter struct {

	// Specifies the IP address from which the client accesses the file or directory.
	// Example: 10.140.68.143
	ClientIP string `json:"client_ip,omitempty"`

	// Specifies whether to enable or disable the filter. Filters are enabled by default and are deleted after 60 mins.
	Enabled *bool `json:"enabled,omitempty"`

	// Position of the file access tracing filter.
	// Example: 1
	// Read Only: true
	Index int64 `json:"index,omitempty"`

	// Specifies the path for which permission tracing can be applied. The value can be complete path from root of CIFS share or root of volume for NFS.
	// Example: /dir1/dir2
	Path string `json:"path,omitempty"`

	// Specifies the protocol for which permission trace is required.
	// Enum: [cifs nfs]
	Protocol *string `json:"protocol,omitempty"`

	// svm
	Svm *SvmReference `json:"svm,omitempty"`

	// Specifies if the filter can trace file access denied and allowed events. The value of trace-allow is false by default, and it traces access denied events. The value is set to true for tracing access allowed events.
	TraceAllowedOps *bool `json:"trace_allowed_ops,omitempty"`

	// Specifies the UNIX username whose access requests you want to trace. The filter would match only if the request is received with this user.
	// Example: root
	UnixUser string `json:"unix_user,omitempty"`

	// Specifies the Windows username whose access requests you want to trace. The filter would match only if the request is received with this user.
	// Example: cifs1/administrator
	WindowsUser string `json:"windows_user,omitempty"`
}

// Validate validates this file access filter
func (m *FileAccessFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fileAccessFilterTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cifs","nfs"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileAccessFilterTypeProtocolPropEnum = append(fileAccessFilterTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// file_access_filter
	// FileAccessFilter
	// protocol
	// Protocol
	// cifs
	// END DEBUGGING
	// FileAccessFilterProtocolCifs captures enum value "cifs"
	FileAccessFilterProtocolCifs string = "cifs"

	// BEGIN DEBUGGING
	// file_access_filter
	// FileAccessFilter
	// protocol
	// Protocol
	// nfs
	// END DEBUGGING
	// FileAccessFilterProtocolNfs captures enum value "nfs"
	FileAccessFilterProtocolNfs string = "nfs"
)

// prop value enum
func (m *FileAccessFilter) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileAccessFilterTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileAccessFilter) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *FileAccessFilter) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file access filter based on the context it is used
func (m *FileAccessFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileAccessFilter) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", int64(m.Index)); err != nil {
		return err
	}

	return nil
}

func (m *FileAccessFilter) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileAccessFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileAccessFilter) UnmarshalBinary(b []byte) error {
	var res FileAccessFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}