// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableDeviceType writable device type
//
// swagger:model WritableDeviceType
type WritableDeviceType struct {

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// Front image
	// Read Only: true
	// Format: uri
	FrontImage strfmt.URI `json:"front_image,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Is full depth
	//
	// Device consumes both front and rear rack faces
	IsFullDepth bool `json:"is_full_depth,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Manufacturer
	// Required: true
	Manufacturer *int64 `json:"manufacturer"`

	// Model
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Model *string `json:"model"`

	// Part number
	//
	// Discrete part number (optional)
	// Max Length: 50
	PartNumber string `json:"part_number,omitempty"`

	// Rear image
	// Read Only: true
	// Format: uri
	RearImage strfmt.URI `json:"rear_image,omitempty"`

	// Slug
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// Parent/child status
	//
	// Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child.
	// Enum: [parent child]
	SubdeviceRole string `json:"subdevice_role,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Height (U)
	// Maximum: 32767
	// Minimum: 0
	UHeight *int64 `json:"u_height,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable device type
func (m *WritableDeviceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubdeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableDeviceType) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validateFrontImage(formats strfmt.Registry) error {

	if swag.IsZero(m.FrontImage) { // not required
		return nil
	}

	if err := validate.FormatOf("front_image", "body", "uri", m.FrontImage.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if err := validate.MinLength("model", "body", string(*m.Model), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("model", "body", string(*m.Model), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validatePartNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.PartNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("part_number", "body", string(m.PartNumber), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validateRearImage(formats strfmt.Registry) error {

	if swag.IsZero(m.RearImage) { // not required
		return nil
	}

	if err := validate.FormatOf("rear_image", "body", "uri", m.RearImage.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", string(*m.Slug), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", string(*m.Slug), 100); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", string(*m.Slug), `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

var writableDeviceTypeTypeSubdeviceRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["parent","child"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceTypeTypeSubdeviceRolePropEnum = append(writableDeviceTypeTypeSubdeviceRolePropEnum, v)
	}
}

const (

	// WritableDeviceTypeSubdeviceRoleParent captures enum value "parent"
	WritableDeviceTypeSubdeviceRoleParent string = "parent"

	// WritableDeviceTypeSubdeviceRoleChild captures enum value "child"
	WritableDeviceTypeSubdeviceRoleChild string = "child"
)

// prop value enum
func (m *WritableDeviceType) validateSubdeviceRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableDeviceTypeTypeSubdeviceRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceType) validateSubdeviceRole(formats strfmt.Registry) error {

	if swag.IsZero(m.SubdeviceRole) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubdeviceRoleEnum("subdevice_role", "body", m.SubdeviceRole); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableDeviceType) validateUHeight(formats strfmt.Registry) error {

	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", int64(*m.UHeight), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", int64(*m.UHeight), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceType) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableDeviceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableDeviceType) UnmarshalBinary(b []byte) error {
	var res WritableDeviceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
