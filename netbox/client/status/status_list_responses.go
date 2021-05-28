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

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StatusListReader is a Reader for the StatusList structure.
type StatusListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatusListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatusListOK creates a StatusListOK with default headers values
func NewStatusListOK() *StatusListOK {
	return &StatusListOK{}
}

/*StatusListOK handles this case with default header values.

StatusListOK status list o k
*/
type StatusListOK struct {
}

func (o *StatusListOK) Error() string {
	return fmt.Sprintf("[GET /status/][%d] statusListOK ", 200)
}

func (o *StatusListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusListDefault creates a StatusListDefault with default headers values
func NewStatusListDefault(code int) *StatusListDefault {
	return &StatusListDefault{
		_statusCode: code,
	}
}

/*StatusListDefault handles this case with default header values.

StatusListDefault status list default
*/
type StatusListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the status list default response
func (o *StatusListDefault) Code() int {
	return o._statusCode
}

func (o *StatusListDefault) Error() string {
	return fmt.Sprintf("[GET /status/][%d] status_list default  %+v", o._statusCode, o.Payload)
}

func (o *StatusListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *StatusListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
