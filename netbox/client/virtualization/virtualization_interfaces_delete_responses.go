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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationInterfacesDeleteReader is a Reader for the VirtualizationInterfacesDelete structure.
type VirtualizationInterfacesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationInterfacesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesDeleteNoContent creates a VirtualizationInterfacesDeleteNoContent with default headers values
func NewVirtualizationInterfacesDeleteNoContent() *VirtualizationInterfacesDeleteNoContent {
	return &VirtualizationInterfacesDeleteNoContent{}
}

/*VirtualizationInterfacesDeleteNoContent handles this case with default header values.

VirtualizationInterfacesDeleteNoContent virtualization interfaces delete no content
*/
type VirtualizationInterfacesDeleteNoContent struct {
}

func (o *VirtualizationInterfacesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/interfaces/{id}/][%d] virtualizationInterfacesDeleteNoContent ", 204)
}

func (o *VirtualizationInterfacesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationInterfacesDeleteDefault creates a VirtualizationInterfacesDeleteDefault with default headers values
func NewVirtualizationInterfacesDeleteDefault(code int) *VirtualizationInterfacesDeleteDefault {
	return &VirtualizationInterfacesDeleteDefault{
		_statusCode: code,
	}
}

/*VirtualizationInterfacesDeleteDefault handles this case with default header values.

VirtualizationInterfacesDeleteDefault virtualization interfaces delete default
*/
type VirtualizationInterfacesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization interfaces delete default response
func (o *VirtualizationInterfacesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationInterfacesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/interfaces/{id}/][%d] virtualization_interfaces_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationInterfacesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
