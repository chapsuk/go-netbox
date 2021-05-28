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

// DcimRearPortTemplatesBulkDeleteReader is a Reader for the DcimRearPortTemplatesBulkDelete structure.
type DcimRearPortTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRearPortTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortTemplatesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortTemplatesBulkDeleteNoContent creates a DcimRearPortTemplatesBulkDeleteNoContent with default headers values
func NewDcimRearPortTemplatesBulkDeleteNoContent() *DcimRearPortTemplatesBulkDeleteNoContent {
	return &DcimRearPortTemplatesBulkDeleteNoContent{}
}

/*DcimRearPortTemplatesBulkDeleteNoContent handles this case with default header values.

DcimRearPortTemplatesBulkDeleteNoContent dcim rear port templates bulk delete no content
*/
type DcimRearPortTemplatesBulkDeleteNoContent struct {
}

func (o *DcimRearPortTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-port-templates/][%d] dcimRearPortTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimRearPortTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRearPortTemplatesBulkDeleteDefault creates a DcimRearPortTemplatesBulkDeleteDefault with default headers values
func NewDcimRearPortTemplatesBulkDeleteDefault(code int) *DcimRearPortTemplatesBulkDeleteDefault {
	return &DcimRearPortTemplatesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*DcimRearPortTemplatesBulkDeleteDefault handles this case with default header values.

DcimRearPortTemplatesBulkDeleteDefault dcim rear port templates bulk delete default
*/
type DcimRearPortTemplatesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear port templates bulk delete default response
func (o *DcimRearPortTemplatesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortTemplatesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-port-templates/][%d] dcim_rear-port-templates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortTemplatesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
