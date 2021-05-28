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
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	StatusList(params *StatusListParams, authInfo runtime.ClientAuthInfoWriter) (*StatusListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StatusList A lightweight read-only endpoint for conveying NetBox's current operational status.
*/
func (a *Client) StatusList(params *StatusListParams, authInfo runtime.ClientAuthInfoWriter) (*StatusListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "status_list",
		Method:             "GET",
		PathPattern:        "/status/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StatusListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatusListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StatusListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
