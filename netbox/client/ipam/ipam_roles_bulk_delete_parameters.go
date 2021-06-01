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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIpamRolesBulkDeleteParams creates a new IpamRolesBulkDeleteParams object
// with the default values initialized.
func NewIpamRolesBulkDeleteParams() *IpamRolesBulkDeleteParams {

	return &IpamRolesBulkDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRolesBulkDeleteParamsWithTimeout creates a new IpamRolesBulkDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamRolesBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamRolesBulkDeleteParams {

	return &IpamRolesBulkDeleteParams{

		timeout: timeout,
	}
}

// NewIpamRolesBulkDeleteParamsWithContext creates a new IpamRolesBulkDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamRolesBulkDeleteParamsWithContext(ctx context.Context) *IpamRolesBulkDeleteParams {

	return &IpamRolesBulkDeleteParams{

		Context: ctx,
	}
}

// NewIpamRolesBulkDeleteParamsWithHTTPClient creates a new IpamRolesBulkDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamRolesBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamRolesBulkDeleteParams {

	return &IpamRolesBulkDeleteParams{
		HTTPClient: client,
	}
}

/*IpamRolesBulkDeleteParams contains all the parameters to send to the API endpoint
for the ipam roles bulk delete operation typically these are written to a http.Request
*/
type IpamRolesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam roles bulk delete params
func (o *IpamRolesBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamRolesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles bulk delete params
func (o *IpamRolesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles bulk delete params
func (o *IpamRolesBulkDeleteParams) WithContext(ctx context.Context) *IpamRolesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles bulk delete params
func (o *IpamRolesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles bulk delete params
func (o *IpamRolesBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamRolesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles bulk delete params
func (o *IpamRolesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRolesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}