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
)

// IpamRouteTargetsDeleteReader is a Reader for the IpamRouteTargetsDelete structure.
type IpamRouteTargetsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRouteTargetsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRouteTargetsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsDeleteNoContent creates a IpamRouteTargetsDeleteNoContent with default headers values
func NewIpamRouteTargetsDeleteNoContent() *IpamRouteTargetsDeleteNoContent {
	return &IpamRouteTargetsDeleteNoContent{}
}

/*IpamRouteTargetsDeleteNoContent handles this case with default header values.

IpamRouteTargetsDeleteNoContent ipam route targets delete no content
*/
type IpamRouteTargetsDeleteNoContent struct {
}

func (o *IpamRouteTargetsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/{id}/][%d] ipamRouteTargetsDeleteNoContent ", 204)
}

func (o *IpamRouteTargetsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamRouteTargetsDeleteDefault creates a IpamRouteTargetsDeleteDefault with default headers values
func NewIpamRouteTargetsDeleteDefault(code int) *IpamRouteTargetsDeleteDefault {
	return &IpamRouteTargetsDeleteDefault{
		_statusCode: code,
	}
}

/*IpamRouteTargetsDeleteDefault handles this case with default header values.

IpamRouteTargetsDeleteDefault ipam route targets delete default
*/
type IpamRouteTargetsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam route targets delete default response
func (o *IpamRouteTargetsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamRouteTargetsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/{id}/][%d] ipam_route-targets_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
