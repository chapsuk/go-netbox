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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// UsersPermissionsBulkUpdateReader is a Reader for the UsersPermissionsBulkUpdate structure.
type UsersPermissionsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersPermissionsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersPermissionsBulkUpdateOK creates a UsersPermissionsBulkUpdateOK with default headers values
func NewUsersPermissionsBulkUpdateOK() *UsersPermissionsBulkUpdateOK {
	return &UsersPermissionsBulkUpdateOK{}
}

/*UsersPermissionsBulkUpdateOK handles this case with default header values.

UsersPermissionsBulkUpdateOK users permissions bulk update o k
*/
type UsersPermissionsBulkUpdateOK struct {
	Payload *models.ObjectPermission
}

func (o *UsersPermissionsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/permissions/][%d] usersPermissionsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersPermissionsBulkUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPermissionsBulkUpdateDefault creates a UsersPermissionsBulkUpdateDefault with default headers values
func NewUsersPermissionsBulkUpdateDefault(code int) *UsersPermissionsBulkUpdateDefault {
	return &UsersPermissionsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*UsersPermissionsBulkUpdateDefault handles this case with default header values.

UsersPermissionsBulkUpdateDefault users permissions bulk update default
*/
type UsersPermissionsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users permissions bulk update default response
func (o *UsersPermissionsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UsersPermissionsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /users/permissions/][%d] users_permissions_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPermissionsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
