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

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SecretsSecretRolesDeleteReader is a Reader for the SecretsSecretRolesDelete structure.
type SecretsSecretRolesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretRolesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSecretsSecretRolesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecretsSecretRolesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsSecretRolesDeleteNoContent creates a SecretsSecretRolesDeleteNoContent with default headers values
func NewSecretsSecretRolesDeleteNoContent() *SecretsSecretRolesDeleteNoContent {
	return &SecretsSecretRolesDeleteNoContent{}
}

/*SecretsSecretRolesDeleteNoContent handles this case with default header values.

SecretsSecretRolesDeleteNoContent secrets secret roles delete no content
*/
type SecretsSecretRolesDeleteNoContent struct {
}

func (o *SecretsSecretRolesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /secrets/secret-roles/{id}/][%d] secretsSecretRolesDeleteNoContent ", 204)
}

func (o *SecretsSecretRolesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecretsSecretRolesDeleteDefault creates a SecretsSecretRolesDeleteDefault with default headers values
func NewSecretsSecretRolesDeleteDefault(code int) *SecretsSecretRolesDeleteDefault {
	return &SecretsSecretRolesDeleteDefault{
		_statusCode: code,
	}
}

/*SecretsSecretRolesDeleteDefault handles this case with default header values.

SecretsSecretRolesDeleteDefault secrets secret roles delete default
*/
type SecretsSecretRolesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the secrets secret roles delete default response
func (o *SecretsSecretRolesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SecretsSecretRolesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /secrets/secret-roles/{id}/][%d] secrets_secret-roles_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsSecretRolesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *SecretsSecretRolesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
