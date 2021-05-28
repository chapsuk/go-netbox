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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitsCircuitsBulkPartialUpdateReader is a Reader for the CircuitsCircuitsBulkPartialUpdate structure.
type CircuitsCircuitsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitsBulkPartialUpdateOK creates a CircuitsCircuitsBulkPartialUpdateOK with default headers values
func NewCircuitsCircuitsBulkPartialUpdateOK() *CircuitsCircuitsBulkPartialUpdateOK {
	return &CircuitsCircuitsBulkPartialUpdateOK{}
}

/*CircuitsCircuitsBulkPartialUpdateOK handles this case with default header values.

CircuitsCircuitsBulkPartialUpdateOK circuits circuits bulk partial update o k
*/
type CircuitsCircuitsBulkPartialUpdateOK struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuits/][%d] circuitsCircuitsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitsBulkPartialUpdateOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitsBulkPartialUpdateDefault creates a CircuitsCircuitsBulkPartialUpdateDefault with default headers values
func NewCircuitsCircuitsBulkPartialUpdateDefault(code int) *CircuitsCircuitsBulkPartialUpdateDefault {
	return &CircuitsCircuitsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*CircuitsCircuitsBulkPartialUpdateDefault handles this case with default header values.

CircuitsCircuitsBulkPartialUpdateDefault circuits circuits bulk partial update default
*/
type CircuitsCircuitsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuits bulk partial update default response
func (o *CircuitsCircuitsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuits/][%d] circuits_circuits_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
