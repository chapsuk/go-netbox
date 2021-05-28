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

// DcimConsoleServerPortTemplatesReadReader is a Reader for the DcimConsoleServerPortTemplatesRead structure.
type DcimConsoleServerPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesReadOK creates a DcimConsoleServerPortTemplatesReadOK with default headers values
func NewDcimConsoleServerPortTemplatesReadOK() *DcimConsoleServerPortTemplatesReadOK {
	return &DcimConsoleServerPortTemplatesReadOK{}
}

/*DcimConsoleServerPortTemplatesReadOK handles this case with default header values.

DcimConsoleServerPortTemplatesReadOK dcim console server port templates read o k
*/
type DcimConsoleServerPortTemplatesReadOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesReadOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesReadDefault creates a DcimConsoleServerPortTemplatesReadDefault with default headers values
func NewDcimConsoleServerPortTemplatesReadDefault(code int) *DcimConsoleServerPortTemplatesReadDefault {
	return &DcimConsoleServerPortTemplatesReadDefault{
		_statusCode: code,
	}
}

/*DcimConsoleServerPortTemplatesReadDefault handles this case with default header values.

DcimConsoleServerPortTemplatesReadDefault dcim console server port templates read default
*/
type DcimConsoleServerPortTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates read default response
func (o *DcimConsoleServerPortTemplatesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
