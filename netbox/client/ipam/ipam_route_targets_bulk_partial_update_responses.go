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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamRouteTargetsBulkPartialUpdateReader is a Reader for the IpamRouteTargetsBulkPartialUpdate structure.
type IpamRouteTargetsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRouteTargetsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRouteTargetsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsBulkPartialUpdateOK creates a IpamRouteTargetsBulkPartialUpdateOK with default headers values
func NewIpamRouteTargetsBulkPartialUpdateOK() *IpamRouteTargetsBulkPartialUpdateOK {
	return &IpamRouteTargetsBulkPartialUpdateOK{}
}

/*IpamRouteTargetsBulkPartialUpdateOK handles this case with default header values.

IpamRouteTargetsBulkPartialUpdateOK ipam route targets bulk partial update o k
*/
type IpamRouteTargetsBulkPartialUpdateOK struct {
	Payload *models.RouteTarget
}

func (o *IpamRouteTargetsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/route-targets/][%d] ipamRouteTargetsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRouteTargetsBulkPartialUpdateOK) GetPayload() *models.RouteTarget {
	return o.Payload
}

func (o *IpamRouteTargetsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsBulkPartialUpdateDefault creates a IpamRouteTargetsBulkPartialUpdateDefault with default headers values
func NewIpamRouteTargetsBulkPartialUpdateDefault(code int) *IpamRouteTargetsBulkPartialUpdateDefault {
	return &IpamRouteTargetsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*IpamRouteTargetsBulkPartialUpdateDefault handles this case with default header values.

IpamRouteTargetsBulkPartialUpdateDefault ipam route targets bulk partial update default
*/
type IpamRouteTargetsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam route targets bulk partial update default response
func (o *IpamRouteTargetsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamRouteTargetsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/route-targets/][%d] ipam_route-targets_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
