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

// DcimConsoleServerPortTemplatesBulkPartialUpdateReader is a Reader for the DcimConsoleServerPortTemplatesBulkPartialUpdate structure.
type DcimConsoleServerPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesBulkPartialUpdateOK creates a DcimConsoleServerPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesBulkPartialUpdateOK() *DcimConsoleServerPortTemplatesBulkPartialUpdateOK {
	return &DcimConsoleServerPortTemplatesBulkPartialUpdateOK{}
}

/* DcimConsoleServerPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesBulkPartialUpdateOK dcim console server port templates bulk partial update o k
*/
type DcimConsoleServerPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesBulkPartialUpdateDefault creates a DcimConsoleServerPortTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimConsoleServerPortTemplatesBulkPartialUpdateDefault(code int) *DcimConsoleServerPortTemplatesBulkPartialUpdateDefault {
	return &DcimConsoleServerPortTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimConsoleServerPortTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortTemplatesBulkPartialUpdateDefault dcim console server port templates bulk partial update default
*/
type DcimConsoleServerPortTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates bulk partial update default response
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/][%d] dcim_console-server-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
