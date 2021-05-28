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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimRacksElevationReader is a Reader for the DcimRacksElevation structure.
type DcimRacksElevationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksElevationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksElevationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRacksElevationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksElevationOK creates a DcimRacksElevationOK with default headers values
func NewDcimRacksElevationOK() *DcimRacksElevationOK {
	return &DcimRacksElevationOK{}
}

/*DcimRacksElevationOK handles this case with default header values.

DcimRacksElevationOK dcim racks elevation o k
*/
type DcimRacksElevationOK struct {
	Payload []*models.RackUnit
}

func (o *DcimRacksElevationOK) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/elevation/][%d] dcimRacksElevationOK  %+v", 200, o.Payload)
}

func (o *DcimRacksElevationOK) GetPayload() []*models.RackUnit {
	return o.Payload
}

func (o *DcimRacksElevationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksElevationDefault creates a DcimRacksElevationDefault with default headers values
func NewDcimRacksElevationDefault(code int) *DcimRacksElevationDefault {
	return &DcimRacksElevationDefault{
		_statusCode: code,
	}
}

/*DcimRacksElevationDefault handles this case with default header values.

DcimRacksElevationDefault dcim racks elevation default
*/
type DcimRacksElevationDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim racks elevation default response
func (o *DcimRacksElevationDefault) Code() int {
	return o._statusCode
}

func (o *DcimRacksElevationDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/elevation/][%d] dcim_racks_elevation default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksElevationDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksElevationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
