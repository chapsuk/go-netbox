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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasCustomFieldsBulkUpdateReader is a Reader for the ExtrasCustomFieldsBulkUpdate structure.
type ExtrasCustomFieldsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldsBulkUpdateOK creates a ExtrasCustomFieldsBulkUpdateOK with default headers values
func NewExtrasCustomFieldsBulkUpdateOK() *ExtrasCustomFieldsBulkUpdateOK {
	return &ExtrasCustomFieldsBulkUpdateOK{}
}

/*ExtrasCustomFieldsBulkUpdateOK handles this case with default header values.

ExtrasCustomFieldsBulkUpdateOK extras custom fields bulk update o k
*/
type ExtrasCustomFieldsBulkUpdateOK struct {
	Payload *models.CustomField
}

func (o *ExtrasCustomFieldsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/][%d] extrasCustomFieldsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsBulkUpdateOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldsBulkUpdateDefault creates a ExtrasCustomFieldsBulkUpdateDefault with default headers values
func NewExtrasCustomFieldsBulkUpdateDefault(code int) *ExtrasCustomFieldsBulkUpdateDefault {
	return &ExtrasCustomFieldsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*ExtrasCustomFieldsBulkUpdateDefault handles this case with default header values.

ExtrasCustomFieldsBulkUpdateDefault extras custom fields bulk update default
*/
type ExtrasCustomFieldsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom fields bulk update default response
func (o *ExtrasCustomFieldsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomFieldsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/][%d] extras_custom-fields_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
