// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CapacityPoolResponse Information on capacity pool licenses associated with the cluster.
//
// swagger:model capacity_pool_response
type CapacityPoolResponse struct {

	// links
	Links *CollectionLinks `json:"_links,omitempty"`

	// Number of records
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*CapacityPoolResponseRecordsItems0 `json:"records,omitempty"`
}

// Validate validates this capacity pool response
func (m *CapacityPoolResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolResponse) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this capacity pool response based on the context it is used
func (m *CapacityPoolResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityPoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityPoolResponse) UnmarshalBinary(b []byte) error {
	var res CapacityPoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapacityPoolResponseRecordsItems0 Information on a capacity pool license and how it is associated with the cluster.
//
// swagger:model CapacityPoolResponseRecordsItems0
type CapacityPoolResponseRecordsItems0 struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// license manager
	LicenseManager *CapacityPoolResponseRecordsItems0LicenseManager `json:"license_manager,omitempty"`

	// Nodes in the cluster associated with this capacity pool.
	Nodes []*CapacityPoolResponseRecordsItems0NodesItems0 `json:"nodes,omitempty"`

	// Serial number of the capacity pool license.
	// Example: 390000100
	SerialNumber string `json:"serial_number,omitempty"`
}

// Validate validates this capacity pool response records items0
func (m *CapacityPoolResponseRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolResponseRecordsItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolResponseRecordsItems0) validateLicenseManager(formats strfmt.Registry) error {
	if swag.IsZero(m.LicenseManager) { // not required
		return nil
	}

	if m.LicenseManager != nil {
		if err := m.LicenseManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license_manager")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolResponseRecordsItems0) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this capacity pool response records items0 based on the context it is used
func (m *CapacityPoolResponseRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLicenseManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolResponseRecordsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolResponseRecordsItems0) contextValidateLicenseManager(ctx context.Context, formats strfmt.Registry) error {

	if m.LicenseManager != nil {
		if err := m.LicenseManager.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license_manager")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolResponseRecordsItems0) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityPoolResponseRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityPoolResponseRecordsItems0) UnmarshalBinary(b []byte) error {
	var res CapacityPoolResponseRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapacityPoolResponseRecordsItems0LicenseManager License manager instance where this capacity pool license in installed.
//
// swagger:model CapacityPoolResponseRecordsItems0LicenseManager
type CapacityPoolResponseRecordsItems0LicenseManager struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// uuid
	// Example: 4ea7a442-86d1-11e0-ae1c-112233445566
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this capacity pool response records items0 license manager
func (m *CapacityPoolResponseRecordsItems0LicenseManager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolResponseRecordsItems0LicenseManager) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license_manager" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolResponseRecordsItems0LicenseManager) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("license_manager"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this capacity pool response records items0 license manager based on the context it is used
func (m *CapacityPoolResponseRecordsItems0LicenseManager) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolResponseRecordsItems0LicenseManager) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license_manager" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityPoolResponseRecordsItems0LicenseManager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityPoolResponseRecordsItems0LicenseManager) UnmarshalBinary(b []byte) error {
	var res CapacityPoolResponseRecordsItems0LicenseManager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CapacityPoolResponseRecordsItems0NodesItems0 Information on a node from the capacity licensing perspective.
//
// swagger:model CapacityPoolResponseRecordsItems0NodesItems0
type CapacityPoolResponseRecordsItems0NodesItems0 struct {

	// node
	Node *NodeReference `json:"node,omitempty"`

	// Capacity, in bytes, that is currently used by the node.
	// Read Only: true
	UsedSize int64 `json:"used_size,omitempty"`
}

// Validate validates this capacity pool response records items0 nodes items0
func (m *CapacityPoolResponseRecordsItems0NodesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolResponseRecordsItems0NodesItems0) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this capacity pool response records items0 nodes items0 based on the context it is used
func (m *CapacityPoolResponseRecordsItems0NodesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsedSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPoolResponseRecordsItems0NodesItems0) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPoolResponseRecordsItems0NodesItems0) contextValidateUsedSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "used_size", "body", int64(m.UsedSize)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityPoolResponseRecordsItems0NodesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityPoolResponseRecordsItems0NodesItems0) UnmarshalBinary(b []byte) error {
	var res CapacityPoolResponseRecordsItems0NodesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}