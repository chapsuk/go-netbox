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

// DcimDeviceBayTemplatesDeleteReader is a Reader for the DcimDeviceBayTemplatesDelete structure.
type DcimDeviceBayTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceBayTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBayTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesDeleteNoContent creates a DcimDeviceBayTemplatesDeleteNoContent with default headers values
func NewDcimDeviceBayTemplatesDeleteNoContent() *DcimDeviceBayTemplatesDeleteNoContent {
	return &DcimDeviceBayTemplatesDeleteNoContent{}
}

/*DcimDeviceBayTemplatesDeleteNoContent handles this case with default header values.

DcimDeviceBayTemplatesDeleteNoContent dcim device bay templates delete no content
*/
type DcimDeviceBayTemplatesDeleteNoContent struct {
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDeviceBayTemplatesDeleteDefault creates a DcimDeviceBayTemplatesDeleteDefault with default headers values
func NewDcimDeviceBayTemplatesDeleteDefault(code int) *DcimDeviceBayTemplatesDeleteDefault {
	return &DcimDeviceBayTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/*DcimDeviceBayTemplatesDeleteDefault handles this case with default header values.

DcimDeviceBayTemplatesDeleteDefault dcim device bay templates delete default
*/
type DcimDeviceBayTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bay templates delete default response
func (o *DcimDeviceBayTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceBayTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/{id}/][%d] dcim_device-bay-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
