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

// DcimPlatformsBulkDeleteReader is a Reader for the DcimPlatformsBulkDelete structure.
type DcimPlatformsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPlatformsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPlatformsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsBulkDeleteNoContent creates a DcimPlatformsBulkDeleteNoContent with default headers values
func NewDcimPlatformsBulkDeleteNoContent() *DcimPlatformsBulkDeleteNoContent {
	return &DcimPlatformsBulkDeleteNoContent{}
}

/*DcimPlatformsBulkDeleteNoContent handles this case with default header values.

DcimPlatformsBulkDeleteNoContent dcim platforms bulk delete no content
*/
type DcimPlatformsBulkDeleteNoContent struct {
}

func (o *DcimPlatformsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/][%d] dcimPlatformsBulkDeleteNoContent ", 204)
}

func (o *DcimPlatformsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPlatformsBulkDeleteDefault creates a DcimPlatformsBulkDeleteDefault with default headers values
func NewDcimPlatformsBulkDeleteDefault(code int) *DcimPlatformsBulkDeleteDefault {
	return &DcimPlatformsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*DcimPlatformsBulkDeleteDefault handles this case with default header values.

DcimPlatformsBulkDeleteDefault dcim platforms bulk delete default
*/
type DcimPlatformsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms bulk delete default response
func (o *DcimPlatformsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimPlatformsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/][%d] dcim_platforms_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
