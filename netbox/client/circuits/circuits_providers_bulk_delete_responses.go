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
)

// CircuitsProvidersBulkDeleteReader is a Reader for the CircuitsProvidersBulkDelete structure.
type CircuitsProvidersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsProvidersBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProvidersBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersBulkDeleteNoContent creates a CircuitsProvidersBulkDeleteNoContent with default headers values
func NewCircuitsProvidersBulkDeleteNoContent() *CircuitsProvidersBulkDeleteNoContent {
	return &CircuitsProvidersBulkDeleteNoContent{}
}

/*CircuitsProvidersBulkDeleteNoContent handles this case with default header values.

CircuitsProvidersBulkDeleteNoContent circuits providers bulk delete no content
*/
type CircuitsProvidersBulkDeleteNoContent struct {
}

func (o *CircuitsProvidersBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuitsProvidersBulkDeleteNoContent ", 204)
}

func (o *CircuitsProvidersBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCircuitsProvidersBulkDeleteDefault creates a CircuitsProvidersBulkDeleteDefault with default headers values
func NewCircuitsProvidersBulkDeleteDefault(code int) *CircuitsProvidersBulkDeleteDefault {
	return &CircuitsProvidersBulkDeleteDefault{
		_statusCode: code,
	}
}

/*CircuitsProvidersBulkDeleteDefault handles this case with default header values.

CircuitsProvidersBulkDeleteDefault circuits providers bulk delete default
*/
type CircuitsProvidersBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers bulk delete default response
func (o *CircuitsProvidersBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsProvidersBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuits_providers_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
