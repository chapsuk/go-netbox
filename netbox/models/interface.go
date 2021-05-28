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

// Interface interface
//
// swagger:model Interface
type Interface struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable peer
	//
	//
	// Return the appropriate serializer for the cable termination model.
	//
	// Read Only: true
	CablePeer map[string]*string `json:"cable_peer,omitempty"`

	// Cable peer type
	// Read Only: true
	CablePeerType string `json:"cable_peer_type,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]*string `json:"connected_endpoint,omitempty"`

	// Connected endpoint reachable
	// Read Only: true
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Count ipaddresses
	// Read Only: true
	CountIpaddresses int64 `json:"count_ipaddresses,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// lag
	Lag *NestedInterface `json:"lag,omitempty"`

	// MAC Address
	MacAddress *string `json:"mac_address,omitempty"`

	// Management only
	//
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// mode
	Mode *InterfaceMode `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []*NestedVLAN `json:"tagged_vlans"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	// Required: true
	Type *InterfaceType `json:"type"`

	// untagged vlan
	UntaggedVlan *NestedVLAN `json:"untagged_vlan,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this interface
func (m *Interface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUntaggedVlan(formats); err != nil {
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

func (m *Interface) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", string(m.Label), 64); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateLag(formats strfmt.Registry) error {

	if swag.IsZero(m.Lag) { // not required
		return nil
	}

	if m.Lag != nil {
		if err := m.Lag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lag")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(*m.Mtu), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(*m.Mtu), 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *Interface) validateTaggedVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	for i := 0; i < len(m.TaggedVlans); i++ {
		if swag.IsZero(m.TaggedVlans[i]) { // not required
			continue
		}

		if m.TaggedVlans[i] != nil {
			if err := m.TaggedVlans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagged_vlans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Interface) validateTags(formats strfmt.Registry) error {

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

func (m *Interface) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateUntaggedVlan(formats strfmt.Registry) error {

	if swag.IsZero(m.UntaggedVlan) { // not required
		return nil
	}

	if m.UntaggedVlan != nil {
		if err := m.UntaggedVlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("untagged_vlan")
			}
			return err
		}
	}

	return nil
}

func (m *Interface) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Interface) UnmarshalBinary(b []byte) error {
	var res Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceMode Mode
//
// swagger:model InterfaceMode
type InterfaceMode struct {

	// label
	// Required: true
	// Enum: [Access Tagged Tagged (All)]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [access tagged tagged-all]
	Value *string `json:"value"`
}

// Validate validates this interface mode
func (m *InterfaceMode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var interfaceModeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Access","Tagged","Tagged (All)"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceModeTypeLabelPropEnum = append(interfaceModeTypeLabelPropEnum, v)
	}
}

const (

	// InterfaceModeLabelAccess captures enum value "Access"
	InterfaceModeLabelAccess string = "Access"

	// InterfaceModeLabelTagged captures enum value "Tagged"
	InterfaceModeLabelTagged string = "Tagged"

	// InterfaceModeLabelTaggedAll captures enum value "Tagged (All)"
	InterfaceModeLabelTaggedAll string = "Tagged (All)"
)

// prop value enum
func (m *InterfaceMode) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceModeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceMode) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("mode"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var interfaceModeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","tagged","tagged-all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceModeTypeValuePropEnum = append(interfaceModeTypeValuePropEnum, v)
	}
}

const (

	// InterfaceModeValueAccess captures enum value "access"
	InterfaceModeValueAccess string = "access"

	// InterfaceModeValueTagged captures enum value "tagged"
	InterfaceModeValueTagged string = "tagged"

	// InterfaceModeValueTaggedAll captures enum value "tagged-all"
	InterfaceModeValueTaggedAll string = "tagged-all"
)

// prop value enum
func (m *InterfaceMode) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceModeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceMode) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("mode"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceMode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceMode) UnmarshalBinary(b []byte) error {
	var res InterfaceMode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceType Type
//
// swagger:model InterfaceType
type InterfaceType struct {

	// label
	// Required: true
	// Enum: [Virtual Link Aggregation Group (LAG) 100BASE-TX (10/100ME) 1000BASE-T (1GE) 2.5GBASE-T (2.5GE) 5GBASE-T (5GE) 10GBASE-T (10GE) 10GBASE-CX4 (10GE) GBIC (1GE) SFP (1GE) SFP+ (10GE) XFP (10GE) XENPAK (10GE) X2 (10GE) SFP28 (25GE) SFP56 (50GE) QSFP+ (40GE) QSFP28 (50GE) CFP (100GE) CFP2 (100GE) CFP2 (200GE) CFP4 (100GE) Cisco CPAK (100GE) QSFP28 (100GE) QSFP56 (200GE) QSFP-DD (400GE) OSFP (400GE) IEEE 802.11a IEEE 802.11b/g IEEE 802.11n IEEE 802.11ac IEEE 802.11ad IEEE 802.11ax GSM CDMA LTE OC-3/STM-1 OC-12/STM-4 OC-48/STM-16 OC-192/STM-64 OC-768/STM-256 OC-1920/STM-640 OC-3840/STM-1234 SFP (1GFC) SFP (2GFC) SFP (4GFC) SFP+ (8GFC) SFP+ (16GFC) SFP28 (32GFC) QSFP+ (64GFC) QSFP28 (128GFC) SDR (2 Gbps) DDR (4 Gbps) QDR (8 Gbps) FDR10 (10 Gbps) FDR (13.5 Gbps) EDR (25 Gbps) HDR (50 Gbps) NDR (100 Gbps) XDR (250 Gbps) T1 (1.544 Mbps) E1 (2.048 Mbps) T3 (45 Mbps) E3 (34 Mbps) Cisco StackWise Cisco StackWise Plus Cisco FlexStack Cisco FlexStack Plus Juniper VCP Extreme SummitStack Extreme SummitStack-128 Extreme SummitStack-256 Extreme SummitStack-512 Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [virtual lag 100base-tx 1000base-t 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 1000base-x-gbic 1000base-x-sfp 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 50gbase-x-sfp56 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 200gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cpak 100gbase-x-qsfp28 200gbase-x-qsfp56 400gbase-x-qsfpdd 400gbase-x-osfp ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 64gfc-qsfpp 128gfc-sfp28 infiniband-sdr infiniband-ddr infiniband-qdr infiniband-fdr10 infiniband-fdr infiniband-edr infiniband-hdr infiniband-ndr infiniband-xdr t1 e1 t3 e3 cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]
	Value *string `json:"value"`
}

// Validate validates this interface type
func (m *InterfaceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var interfaceTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Virtual","Link Aggregation Group (LAG)","100BASE-TX (10/100ME)","1000BASE-T (1GE)","2.5GBASE-T (2.5GE)","5GBASE-T (5GE)","10GBASE-T (10GE)","10GBASE-CX4 (10GE)","GBIC (1GE)","SFP (1GE)","SFP+ (10GE)","XFP (10GE)","XENPAK (10GE)","X2 (10GE)","SFP28 (25GE)","SFP56 (50GE)","QSFP+ (40GE)","QSFP28 (50GE)","CFP (100GE)","CFP2 (100GE)","CFP2 (200GE)","CFP4 (100GE)","Cisco CPAK (100GE)","QSFP28 (100GE)","QSFP56 (200GE)","QSFP-DD (400GE)","OSFP (400GE)","IEEE 802.11a","IEEE 802.11b/g","IEEE 802.11n","IEEE 802.11ac","IEEE 802.11ad","IEEE 802.11ax","GSM","CDMA","LTE","OC-3/STM-1","OC-12/STM-4","OC-48/STM-16","OC-192/STM-64","OC-768/STM-256","OC-1920/STM-640","OC-3840/STM-1234","SFP (1GFC)","SFP (2GFC)","SFP (4GFC)","SFP+ (8GFC)","SFP+ (16GFC)","SFP28 (32GFC)","QSFP+ (64GFC)","QSFP28 (128GFC)","SDR (2 Gbps)","DDR (4 Gbps)","QDR (8 Gbps)","FDR10 (10 Gbps)","FDR (13.5 Gbps)","EDR (25 Gbps)","HDR (50 Gbps)","NDR (100 Gbps)","XDR (250 Gbps)","T1 (1.544 Mbps)","E1 (2.048 Mbps)","T3 (45 Mbps)","E3 (34 Mbps)","Cisco StackWise","Cisco StackWise Plus","Cisco FlexStack","Cisco FlexStack Plus","Juniper VCP","Extreme SummitStack","Extreme SummitStack-128","Extreme SummitStack-256","Extreme SummitStack-512","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceTypeTypeLabelPropEnum = append(interfaceTypeTypeLabelPropEnum, v)
	}
}

const (

	// InterfaceTypeLabelVirtual captures enum value "Virtual"
	InterfaceTypeLabelVirtual string = "Virtual"

	// InterfaceTypeLabelLinkAggregationGroupLAG captures enum value "Link Aggregation Group (LAG)"
	InterfaceTypeLabelLinkAggregationGroupLAG string = "Link Aggregation Group (LAG)"

	// InterfaceTypeLabelNr100BASETX10100ME captures enum value "100BASE-TX (10/100ME)"
	InterfaceTypeLabelNr100BASETX10100ME string = "100BASE-TX (10/100ME)"

	// InterfaceTypeLabelNr1000BASET1GE captures enum value "1000BASE-T (1GE)"
	InterfaceTypeLabelNr1000BASET1GE string = "1000BASE-T (1GE)"

	// InterfaceTypeLabelNr25GBASET25GE captures enum value "2.5GBASE-T (2.5GE)"
	InterfaceTypeLabelNr25GBASET25GE string = "2.5GBASE-T (2.5GE)"

	// InterfaceTypeLabelNr5GBASET5GE captures enum value "5GBASE-T (5GE)"
	InterfaceTypeLabelNr5GBASET5GE string = "5GBASE-T (5GE)"

	// InterfaceTypeLabelNr10GBASET10GE captures enum value "10GBASE-T (10GE)"
	InterfaceTypeLabelNr10GBASET10GE string = "10GBASE-T (10GE)"

	// InterfaceTypeLabelNr10GBASECX410GE captures enum value "10GBASE-CX4 (10GE)"
	InterfaceTypeLabelNr10GBASECX410GE string = "10GBASE-CX4 (10GE)"

	// InterfaceTypeLabelGBIC1GE captures enum value "GBIC (1GE)"
	InterfaceTypeLabelGBIC1GE string = "GBIC (1GE)"

	// InterfaceTypeLabelSFP1GE captures enum value "SFP (1GE)"
	InterfaceTypeLabelSFP1GE string = "SFP (1GE)"

	// InterfaceTypeLabelSFP10GE captures enum value "SFP+ (10GE)"
	InterfaceTypeLabelSFP10GE string = "SFP+ (10GE)"

	// InterfaceTypeLabelXFP10GE captures enum value "XFP (10GE)"
	InterfaceTypeLabelXFP10GE string = "XFP (10GE)"

	// InterfaceTypeLabelXENPAK10GE captures enum value "XENPAK (10GE)"
	InterfaceTypeLabelXENPAK10GE string = "XENPAK (10GE)"

	// InterfaceTypeLabelX210GE captures enum value "X2 (10GE)"
	InterfaceTypeLabelX210GE string = "X2 (10GE)"

	// InterfaceTypeLabelSFP2825GE captures enum value "SFP28 (25GE)"
	InterfaceTypeLabelSFP2825GE string = "SFP28 (25GE)"

	// InterfaceTypeLabelSFP5650GE captures enum value "SFP56 (50GE)"
	InterfaceTypeLabelSFP5650GE string = "SFP56 (50GE)"

	// InterfaceTypeLabelQSFP40GE captures enum value "QSFP+ (40GE)"
	InterfaceTypeLabelQSFP40GE string = "QSFP+ (40GE)"

	// InterfaceTypeLabelQSFP2850GE captures enum value "QSFP28 (50GE)"
	InterfaceTypeLabelQSFP2850GE string = "QSFP28 (50GE)"

	// InterfaceTypeLabelCFP100GE captures enum value "CFP (100GE)"
	InterfaceTypeLabelCFP100GE string = "CFP (100GE)"

	// InterfaceTypeLabelCFP2100GE captures enum value "CFP2 (100GE)"
	InterfaceTypeLabelCFP2100GE string = "CFP2 (100GE)"

	// InterfaceTypeLabelCFP2200GE captures enum value "CFP2 (200GE)"
	InterfaceTypeLabelCFP2200GE string = "CFP2 (200GE)"

	// InterfaceTypeLabelCFP4100GE captures enum value "CFP4 (100GE)"
	InterfaceTypeLabelCFP4100GE string = "CFP4 (100GE)"

	// InterfaceTypeLabelCiscoCPAK100GE captures enum value "Cisco CPAK (100GE)"
	InterfaceTypeLabelCiscoCPAK100GE string = "Cisco CPAK (100GE)"

	// InterfaceTypeLabelQSFP28100GE captures enum value "QSFP28 (100GE)"
	InterfaceTypeLabelQSFP28100GE string = "QSFP28 (100GE)"

	// InterfaceTypeLabelQSFP56200GE captures enum value "QSFP56 (200GE)"
	InterfaceTypeLabelQSFP56200GE string = "QSFP56 (200GE)"

	// InterfaceTypeLabelQSFPDD400GE captures enum value "QSFP-DD (400GE)"
	InterfaceTypeLabelQSFPDD400GE string = "QSFP-DD (400GE)"

	// InterfaceTypeLabelOSFP400GE captures enum value "OSFP (400GE)"
	InterfaceTypeLabelOSFP400GE string = "OSFP (400GE)"

	// InterfaceTypeLabelIEEE80211a captures enum value "IEEE 802.11a"
	InterfaceTypeLabelIEEE80211a string = "IEEE 802.11a"

	// InterfaceTypeLabelIEEE80211bg captures enum value "IEEE 802.11b/g"
	InterfaceTypeLabelIEEE80211bg string = "IEEE 802.11b/g"

	// InterfaceTypeLabelIEEE80211n captures enum value "IEEE 802.11n"
	InterfaceTypeLabelIEEE80211n string = "IEEE 802.11n"

	// InterfaceTypeLabelIEEE80211ac captures enum value "IEEE 802.11ac"
	InterfaceTypeLabelIEEE80211ac string = "IEEE 802.11ac"

	// InterfaceTypeLabelIEEE80211ad captures enum value "IEEE 802.11ad"
	InterfaceTypeLabelIEEE80211ad string = "IEEE 802.11ad"

	// InterfaceTypeLabelIEEE80211ax captures enum value "IEEE 802.11ax"
	InterfaceTypeLabelIEEE80211ax string = "IEEE 802.11ax"

	// InterfaceTypeLabelGSM captures enum value "GSM"
	InterfaceTypeLabelGSM string = "GSM"

	// InterfaceTypeLabelCDMA captures enum value "CDMA"
	InterfaceTypeLabelCDMA string = "CDMA"

	// InterfaceTypeLabelLTE captures enum value "LTE"
	InterfaceTypeLabelLTE string = "LTE"

	// InterfaceTypeLabelOC3STM1 captures enum value "OC-3/STM-1"
	InterfaceTypeLabelOC3STM1 string = "OC-3/STM-1"

	// InterfaceTypeLabelOC12STM4 captures enum value "OC-12/STM-4"
	InterfaceTypeLabelOC12STM4 string = "OC-12/STM-4"

	// InterfaceTypeLabelOC48STM16 captures enum value "OC-48/STM-16"
	InterfaceTypeLabelOC48STM16 string = "OC-48/STM-16"

	// InterfaceTypeLabelOC192STM64 captures enum value "OC-192/STM-64"
	InterfaceTypeLabelOC192STM64 string = "OC-192/STM-64"

	// InterfaceTypeLabelOC768STM256 captures enum value "OC-768/STM-256"
	InterfaceTypeLabelOC768STM256 string = "OC-768/STM-256"

	// InterfaceTypeLabelOC1920STM640 captures enum value "OC-1920/STM-640"
	InterfaceTypeLabelOC1920STM640 string = "OC-1920/STM-640"

	// InterfaceTypeLabelOC3840STM1234 captures enum value "OC-3840/STM-1234"
	InterfaceTypeLabelOC3840STM1234 string = "OC-3840/STM-1234"

	// InterfaceTypeLabelSFP1GFC captures enum value "SFP (1GFC)"
	InterfaceTypeLabelSFP1GFC string = "SFP (1GFC)"

	// InterfaceTypeLabelSFP2GFC captures enum value "SFP (2GFC)"
	InterfaceTypeLabelSFP2GFC string = "SFP (2GFC)"

	// InterfaceTypeLabelSFP4GFC captures enum value "SFP (4GFC)"
	InterfaceTypeLabelSFP4GFC string = "SFP (4GFC)"

	// InterfaceTypeLabelSFP8GFC captures enum value "SFP+ (8GFC)"
	InterfaceTypeLabelSFP8GFC string = "SFP+ (8GFC)"

	// InterfaceTypeLabelSFP16GFC captures enum value "SFP+ (16GFC)"
	InterfaceTypeLabelSFP16GFC string = "SFP+ (16GFC)"

	// InterfaceTypeLabelSFP2832GFC captures enum value "SFP28 (32GFC)"
	InterfaceTypeLabelSFP2832GFC string = "SFP28 (32GFC)"

	// InterfaceTypeLabelQSFP64GFC captures enum value "QSFP+ (64GFC)"
	InterfaceTypeLabelQSFP64GFC string = "QSFP+ (64GFC)"

	// InterfaceTypeLabelQSFP28128GFC captures enum value "QSFP28 (128GFC)"
	InterfaceTypeLabelQSFP28128GFC string = "QSFP28 (128GFC)"

	// InterfaceTypeLabelSDR2Gbps captures enum value "SDR (2 Gbps)"
	InterfaceTypeLabelSDR2Gbps string = "SDR (2 Gbps)"

	// InterfaceTypeLabelDDR4Gbps captures enum value "DDR (4 Gbps)"
	InterfaceTypeLabelDDR4Gbps string = "DDR (4 Gbps)"

	// InterfaceTypeLabelQDR8Gbps captures enum value "QDR (8 Gbps)"
	InterfaceTypeLabelQDR8Gbps string = "QDR (8 Gbps)"

	// InterfaceTypeLabelFDR1010Gbps captures enum value "FDR10 (10 Gbps)"
	InterfaceTypeLabelFDR1010Gbps string = "FDR10 (10 Gbps)"

	// InterfaceTypeLabelFDR135Gbps captures enum value "FDR (13.5 Gbps)"
	InterfaceTypeLabelFDR135Gbps string = "FDR (13.5 Gbps)"

	// InterfaceTypeLabelEDR25Gbps captures enum value "EDR (25 Gbps)"
	InterfaceTypeLabelEDR25Gbps string = "EDR (25 Gbps)"

	// InterfaceTypeLabelHDR50Gbps captures enum value "HDR (50 Gbps)"
	InterfaceTypeLabelHDR50Gbps string = "HDR (50 Gbps)"

	// InterfaceTypeLabelNDR100Gbps captures enum value "NDR (100 Gbps)"
	InterfaceTypeLabelNDR100Gbps string = "NDR (100 Gbps)"

	// InterfaceTypeLabelXDR250Gbps captures enum value "XDR (250 Gbps)"
	InterfaceTypeLabelXDR250Gbps string = "XDR (250 Gbps)"

	// InterfaceTypeLabelT11544Mbps captures enum value "T1 (1.544 Mbps)"
	InterfaceTypeLabelT11544Mbps string = "T1 (1.544 Mbps)"

	// InterfaceTypeLabelE12048Mbps captures enum value "E1 (2.048 Mbps)"
	InterfaceTypeLabelE12048Mbps string = "E1 (2.048 Mbps)"

	// InterfaceTypeLabelT345Mbps captures enum value "T3 (45 Mbps)"
	InterfaceTypeLabelT345Mbps string = "T3 (45 Mbps)"

	// InterfaceTypeLabelE334Mbps captures enum value "E3 (34 Mbps)"
	InterfaceTypeLabelE334Mbps string = "E3 (34 Mbps)"

	// InterfaceTypeLabelCiscoStackWise captures enum value "Cisco StackWise"
	InterfaceTypeLabelCiscoStackWise string = "Cisco StackWise"

	// InterfaceTypeLabelCiscoStackWisePlus captures enum value "Cisco StackWise Plus"
	InterfaceTypeLabelCiscoStackWisePlus string = "Cisco StackWise Plus"

	// InterfaceTypeLabelCiscoFlexStack captures enum value "Cisco FlexStack"
	InterfaceTypeLabelCiscoFlexStack string = "Cisco FlexStack"

	// InterfaceTypeLabelCiscoFlexStackPlus captures enum value "Cisco FlexStack Plus"
	InterfaceTypeLabelCiscoFlexStackPlus string = "Cisco FlexStack Plus"

	// InterfaceTypeLabelJuniperVCP captures enum value "Juniper VCP"
	InterfaceTypeLabelJuniperVCP string = "Juniper VCP"

	// InterfaceTypeLabelExtremeSummitStack captures enum value "Extreme SummitStack"
	InterfaceTypeLabelExtremeSummitStack string = "Extreme SummitStack"

	// InterfaceTypeLabelExtremeSummitStack128 captures enum value "Extreme SummitStack-128"
	InterfaceTypeLabelExtremeSummitStack128 string = "Extreme SummitStack-128"

	// InterfaceTypeLabelExtremeSummitStack256 captures enum value "Extreme SummitStack-256"
	InterfaceTypeLabelExtremeSummitStack256 string = "Extreme SummitStack-256"

	// InterfaceTypeLabelExtremeSummitStack512 captures enum value "Extreme SummitStack-512"
	InterfaceTypeLabelExtremeSummitStack512 string = "Extreme SummitStack-512"

	// InterfaceTypeLabelOther captures enum value "Other"
	InterfaceTypeLabelOther string = "Other"
)

// prop value enum
func (m *InterfaceType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var interfaceTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","lag","100base-tx","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","50gbase-x-sfp56","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","64gfc-qsfpp","128gfc-sfp28","infiniband-sdr","infiniband-ddr","infiniband-qdr","infiniband-fdr10","infiniband-fdr","infiniband-edr","infiniband-hdr","infiniband-ndr","infiniband-xdr","t1","e1","t3","e3","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceTypeTypeValuePropEnum = append(interfaceTypeTypeValuePropEnum, v)
	}
}

const (

	// InterfaceTypeValueVirtual captures enum value "virtual"
	InterfaceTypeValueVirtual string = "virtual"

	// InterfaceTypeValueLag captures enum value "lag"
	InterfaceTypeValueLag string = "lag"

	// InterfaceTypeValueNr100baseTx captures enum value "100base-tx"
	InterfaceTypeValueNr100baseTx string = "100base-tx"

	// InterfaceTypeValueNr1000baset captures enum value "1000base-t"
	InterfaceTypeValueNr1000baset string = "1000base-t"

	// InterfaceTypeValueNr25gbaset captures enum value "2.5gbase-t"
	InterfaceTypeValueNr25gbaset string = "2.5gbase-t"

	// InterfaceTypeValueNr5gbaset captures enum value "5gbase-t"
	InterfaceTypeValueNr5gbaset string = "5gbase-t"

	// InterfaceTypeValueNr10gbaset captures enum value "10gbase-t"
	InterfaceTypeValueNr10gbaset string = "10gbase-t"

	// InterfaceTypeValueNr10gbaseCx4 captures enum value "10gbase-cx4"
	InterfaceTypeValueNr10gbaseCx4 string = "10gbase-cx4"

	// InterfaceTypeValueNr1000basexGbic captures enum value "1000base-x-gbic"
	InterfaceTypeValueNr1000basexGbic string = "1000base-x-gbic"

	// InterfaceTypeValueNr1000basexSfp captures enum value "1000base-x-sfp"
	InterfaceTypeValueNr1000basexSfp string = "1000base-x-sfp"

	// InterfaceTypeValueNr10gbasexSfpp captures enum value "10gbase-x-sfpp"
	InterfaceTypeValueNr10gbasexSfpp string = "10gbase-x-sfpp"

	// InterfaceTypeValueNr10gbasexXfp captures enum value "10gbase-x-xfp"
	InterfaceTypeValueNr10gbasexXfp string = "10gbase-x-xfp"

	// InterfaceTypeValueNr10gbasexXenpak captures enum value "10gbase-x-xenpak"
	InterfaceTypeValueNr10gbasexXenpak string = "10gbase-x-xenpak"

	// InterfaceTypeValueNr10gbasexX2 captures enum value "10gbase-x-x2"
	InterfaceTypeValueNr10gbasexX2 string = "10gbase-x-x2"

	// InterfaceTypeValueNr25gbasexSfp28 captures enum value "25gbase-x-sfp28"
	InterfaceTypeValueNr25gbasexSfp28 string = "25gbase-x-sfp28"

	// InterfaceTypeValueNr50gbasexSfp56 captures enum value "50gbase-x-sfp56"
	InterfaceTypeValueNr50gbasexSfp56 string = "50gbase-x-sfp56"

	// InterfaceTypeValueNr40gbasexQsfpp captures enum value "40gbase-x-qsfpp"
	InterfaceTypeValueNr40gbasexQsfpp string = "40gbase-x-qsfpp"

	// InterfaceTypeValueNr50gbasexSfp28 captures enum value "50gbase-x-sfp28"
	InterfaceTypeValueNr50gbasexSfp28 string = "50gbase-x-sfp28"

	// InterfaceTypeValueNr100gbasexCfp captures enum value "100gbase-x-cfp"
	InterfaceTypeValueNr100gbasexCfp string = "100gbase-x-cfp"

	// InterfaceTypeValueNr100gbasexCfp2 captures enum value "100gbase-x-cfp2"
	InterfaceTypeValueNr100gbasexCfp2 string = "100gbase-x-cfp2"

	// InterfaceTypeValueNr200gbasexCfp2 captures enum value "200gbase-x-cfp2"
	InterfaceTypeValueNr200gbasexCfp2 string = "200gbase-x-cfp2"

	// InterfaceTypeValueNr100gbasexCfp4 captures enum value "100gbase-x-cfp4"
	InterfaceTypeValueNr100gbasexCfp4 string = "100gbase-x-cfp4"

	// InterfaceTypeValueNr100gbasexCpak captures enum value "100gbase-x-cpak"
	InterfaceTypeValueNr100gbasexCpak string = "100gbase-x-cpak"

	// InterfaceTypeValueNr100gbasexQsfp28 captures enum value "100gbase-x-qsfp28"
	InterfaceTypeValueNr100gbasexQsfp28 string = "100gbase-x-qsfp28"

	// InterfaceTypeValueNr200gbasexQsfp56 captures enum value "200gbase-x-qsfp56"
	InterfaceTypeValueNr200gbasexQsfp56 string = "200gbase-x-qsfp56"

	// InterfaceTypeValueNr400gbasexQsfpdd captures enum value "400gbase-x-qsfpdd"
	InterfaceTypeValueNr400gbasexQsfpdd string = "400gbase-x-qsfpdd"

	// InterfaceTypeValueNr400gbasexOsfp captures enum value "400gbase-x-osfp"
	InterfaceTypeValueNr400gbasexOsfp string = "400gbase-x-osfp"

	// InterfaceTypeValueIeee80211a captures enum value "ieee802.11a"
	InterfaceTypeValueIeee80211a string = "ieee802.11a"

	// InterfaceTypeValueIeee80211g captures enum value "ieee802.11g"
	InterfaceTypeValueIeee80211g string = "ieee802.11g"

	// InterfaceTypeValueIeee80211n captures enum value "ieee802.11n"
	InterfaceTypeValueIeee80211n string = "ieee802.11n"

	// InterfaceTypeValueIeee80211ac captures enum value "ieee802.11ac"
	InterfaceTypeValueIeee80211ac string = "ieee802.11ac"

	// InterfaceTypeValueIeee80211ad captures enum value "ieee802.11ad"
	InterfaceTypeValueIeee80211ad string = "ieee802.11ad"

	// InterfaceTypeValueIeee80211ax captures enum value "ieee802.11ax"
	InterfaceTypeValueIeee80211ax string = "ieee802.11ax"

	// InterfaceTypeValueGsm captures enum value "gsm"
	InterfaceTypeValueGsm string = "gsm"

	// InterfaceTypeValueCdma captures enum value "cdma"
	InterfaceTypeValueCdma string = "cdma"

	// InterfaceTypeValueLte captures enum value "lte"
	InterfaceTypeValueLte string = "lte"

	// InterfaceTypeValueSonetOc3 captures enum value "sonet-oc3"
	InterfaceTypeValueSonetOc3 string = "sonet-oc3"

	// InterfaceTypeValueSonetOc12 captures enum value "sonet-oc12"
	InterfaceTypeValueSonetOc12 string = "sonet-oc12"

	// InterfaceTypeValueSonetOc48 captures enum value "sonet-oc48"
	InterfaceTypeValueSonetOc48 string = "sonet-oc48"

	// InterfaceTypeValueSonetOc192 captures enum value "sonet-oc192"
	InterfaceTypeValueSonetOc192 string = "sonet-oc192"

	// InterfaceTypeValueSonetOc768 captures enum value "sonet-oc768"
	InterfaceTypeValueSonetOc768 string = "sonet-oc768"

	// InterfaceTypeValueSonetOc1920 captures enum value "sonet-oc1920"
	InterfaceTypeValueSonetOc1920 string = "sonet-oc1920"

	// InterfaceTypeValueSonetOc3840 captures enum value "sonet-oc3840"
	InterfaceTypeValueSonetOc3840 string = "sonet-oc3840"

	// InterfaceTypeValueNr1gfcSfp captures enum value "1gfc-sfp"
	InterfaceTypeValueNr1gfcSfp string = "1gfc-sfp"

	// InterfaceTypeValueNr2gfcSfp captures enum value "2gfc-sfp"
	InterfaceTypeValueNr2gfcSfp string = "2gfc-sfp"

	// InterfaceTypeValueNr4gfcSfp captures enum value "4gfc-sfp"
	InterfaceTypeValueNr4gfcSfp string = "4gfc-sfp"

	// InterfaceTypeValueNr8gfcSfpp captures enum value "8gfc-sfpp"
	InterfaceTypeValueNr8gfcSfpp string = "8gfc-sfpp"

	// InterfaceTypeValueNr16gfcSfpp captures enum value "16gfc-sfpp"
	InterfaceTypeValueNr16gfcSfpp string = "16gfc-sfpp"

	// InterfaceTypeValueNr32gfcSfp28 captures enum value "32gfc-sfp28"
	InterfaceTypeValueNr32gfcSfp28 string = "32gfc-sfp28"

	// InterfaceTypeValueNr64gfcQsfpp captures enum value "64gfc-qsfpp"
	InterfaceTypeValueNr64gfcQsfpp string = "64gfc-qsfpp"

	// InterfaceTypeValueNr128gfcSfp28 captures enum value "128gfc-sfp28"
	InterfaceTypeValueNr128gfcSfp28 string = "128gfc-sfp28"

	// InterfaceTypeValueInfinibandSdr captures enum value "infiniband-sdr"
	InterfaceTypeValueInfinibandSdr string = "infiniband-sdr"

	// InterfaceTypeValueInfinibandDdr captures enum value "infiniband-ddr"
	InterfaceTypeValueInfinibandDdr string = "infiniband-ddr"

	// InterfaceTypeValueInfinibandQdr captures enum value "infiniband-qdr"
	InterfaceTypeValueInfinibandQdr string = "infiniband-qdr"

	// InterfaceTypeValueInfinibandFdr10 captures enum value "infiniband-fdr10"
	InterfaceTypeValueInfinibandFdr10 string = "infiniband-fdr10"

	// InterfaceTypeValueInfinibandFdr captures enum value "infiniband-fdr"
	InterfaceTypeValueInfinibandFdr string = "infiniband-fdr"

	// InterfaceTypeValueInfinibandEdr captures enum value "infiniband-edr"
	InterfaceTypeValueInfinibandEdr string = "infiniband-edr"

	// InterfaceTypeValueInfinibandHdr captures enum value "infiniband-hdr"
	InterfaceTypeValueInfinibandHdr string = "infiniband-hdr"

	// InterfaceTypeValueInfinibandNdr captures enum value "infiniband-ndr"
	InterfaceTypeValueInfinibandNdr string = "infiniband-ndr"

	// InterfaceTypeValueInfinibandXdr captures enum value "infiniband-xdr"
	InterfaceTypeValueInfinibandXdr string = "infiniband-xdr"

	// InterfaceTypeValueT1 captures enum value "t1"
	InterfaceTypeValueT1 string = "t1"

	// InterfaceTypeValueE1 captures enum value "e1"
	InterfaceTypeValueE1 string = "e1"

	// InterfaceTypeValueT3 captures enum value "t3"
	InterfaceTypeValueT3 string = "t3"

	// InterfaceTypeValueE3 captures enum value "e3"
	InterfaceTypeValueE3 string = "e3"

	// InterfaceTypeValueCiscoStackwise captures enum value "cisco-stackwise"
	InterfaceTypeValueCiscoStackwise string = "cisco-stackwise"

	// InterfaceTypeValueCiscoStackwisePlus captures enum value "cisco-stackwise-plus"
	InterfaceTypeValueCiscoStackwisePlus string = "cisco-stackwise-plus"

	// InterfaceTypeValueCiscoFlexstack captures enum value "cisco-flexstack"
	InterfaceTypeValueCiscoFlexstack string = "cisco-flexstack"

	// InterfaceTypeValueCiscoFlexstackPlus captures enum value "cisco-flexstack-plus"
	InterfaceTypeValueCiscoFlexstackPlus string = "cisco-flexstack-plus"

	// InterfaceTypeValueJuniperVcp captures enum value "juniper-vcp"
	InterfaceTypeValueJuniperVcp string = "juniper-vcp"

	// InterfaceTypeValueExtremeSummitstack captures enum value "extreme-summitstack"
	InterfaceTypeValueExtremeSummitstack string = "extreme-summitstack"

	// InterfaceTypeValueExtremeSummitstack128 captures enum value "extreme-summitstack-128"
	InterfaceTypeValueExtremeSummitstack128 string = "extreme-summitstack-128"

	// InterfaceTypeValueExtremeSummitstack256 captures enum value "extreme-summitstack-256"
	InterfaceTypeValueExtremeSummitstack256 string = "extreme-summitstack-256"

	// InterfaceTypeValueExtremeSummitstack512 captures enum value "extreme-summitstack-512"
	InterfaceTypeValueExtremeSummitstack512 string = "extreme-summitstack-512"

	// InterfaceTypeValueOther captures enum value "other"
	InterfaceTypeValueOther string = "other"
)

// prop value enum
func (m *InterfaceType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceType) UnmarshalBinary(b []byte) error {
	var res InterfaceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
