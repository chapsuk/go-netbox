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

// ExtrasContentTypesReadReader is a Reader for the ExtrasContentTypesRead structure.
type ExtrasContentTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasContentTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasContentTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasContentTypesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasContentTypesReadOK creates a ExtrasContentTypesReadOK with default headers values
func NewExtrasContentTypesReadOK() *ExtrasContentTypesReadOK {
	return &ExtrasContentTypesReadOK{}
}

/* ExtrasContentTypesReadOK describes a response with status code 200, with default header values.

ExtrasContentTypesReadOK extras content types read o k
*/
type ExtrasContentTypesReadOK struct {
	Payload *models.ContentType
}

func (o *ExtrasContentTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extrasContentTypesReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasContentTypesReadOK) GetPayload() *models.ContentType {
	return o.Payload
}

func (o *ExtrasContentTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasContentTypesReadDefault creates a ExtrasContentTypesReadDefault with default headers values
func NewExtrasContentTypesReadDefault(code int) *ExtrasContentTypesReadDefault {
	return &ExtrasContentTypesReadDefault{
		_statusCode: code,
	}
}

/* ExtrasContentTypesReadDefault describes a response with status code -1, with default header values.

ExtrasContentTypesReadDefault extras content types read default
*/
type ExtrasContentTypesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras content types read default response
func (o *ExtrasContentTypesReadDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasContentTypesReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extras_content-types_read default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasContentTypesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasContentTypesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
