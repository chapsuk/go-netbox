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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDevicesBulkDeleteReader is a Reader for the DcimDevicesBulkDelete structure.
type DcimDevicesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDevicesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesBulkDeleteNoContent creates a DcimDevicesBulkDeleteNoContent with default headers values
func NewDcimDevicesBulkDeleteNoContent() *DcimDevicesBulkDeleteNoContent {
	return &DcimDevicesBulkDeleteNoContent{}
}

/*DcimDevicesBulkDeleteNoContent handles this case with default header values.

DcimDevicesBulkDeleteNoContent dcim devices bulk delete no content
*/
type DcimDevicesBulkDeleteNoContent struct {
}

func (o *DcimDevicesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcimDevicesBulkDeleteNoContent ", 204)
}

func (o *DcimDevicesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDevicesBulkDeleteDefault creates a DcimDevicesBulkDeleteDefault with default headers values
func NewDcimDevicesBulkDeleteDefault(code int) *DcimDevicesBulkDeleteDefault {
	return &DcimDevicesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*DcimDevicesBulkDeleteDefault handles this case with default header values.

DcimDevicesBulkDeleteDefault dcim devices bulk delete default
*/
type DcimDevicesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim devices bulk delete default response
func (o *DcimDevicesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimDevicesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcim_devices_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
