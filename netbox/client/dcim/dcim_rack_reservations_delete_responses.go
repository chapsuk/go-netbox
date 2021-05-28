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

// DcimRackReservationsDeleteReader is a Reader for the DcimRackReservationsDelete structure.
type DcimRackReservationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackReservationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackReservationsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackReservationsDeleteNoContent creates a DcimRackReservationsDeleteNoContent with default headers values
func NewDcimRackReservationsDeleteNoContent() *DcimRackReservationsDeleteNoContent {
	return &DcimRackReservationsDeleteNoContent{}
}

/*DcimRackReservationsDeleteNoContent handles this case with default header values.

DcimRackReservationsDeleteNoContent dcim rack reservations delete no content
*/
type DcimRackReservationsDeleteNoContent struct {
}

func (o *DcimRackReservationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-reservations/{id}/][%d] dcimRackReservationsDeleteNoContent ", 204)
}

func (o *DcimRackReservationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRackReservationsDeleteDefault creates a DcimRackReservationsDeleteDefault with default headers values
func NewDcimRackReservationsDeleteDefault(code int) *DcimRackReservationsDeleteDefault {
	return &DcimRackReservationsDeleteDefault{
		_statusCode: code,
	}
}

/*DcimRackReservationsDeleteDefault handles this case with default header values.

DcimRackReservationsDeleteDefault dcim rack reservations delete default
*/
type DcimRackReservationsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack reservations delete default response
func (o *DcimRackReservationsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackReservationsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-reservations/{id}/][%d] dcim_rack-reservations_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackReservationsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackReservationsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
