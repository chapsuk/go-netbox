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

package virtualization

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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewVirtualizationClusterGroupsBulkPartialUpdateParams creates a new VirtualizationClusterGroupsBulkPartialUpdateParams object
// with the default values initialized.
func NewVirtualizationClusterGroupsBulkPartialUpdateParams() *VirtualizationClusterGroupsBulkPartialUpdateParams {
	var ()
	return &VirtualizationClusterGroupsBulkPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterGroupsBulkPartialUpdateParamsWithTimeout creates a new VirtualizationClusterGroupsBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClusterGroupsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClusterGroupsBulkPartialUpdateParams {
	var ()
	return &VirtualizationClusterGroupsBulkPartialUpdateParams{

		timeout: timeout,
	}
}

// NewVirtualizationClusterGroupsBulkPartialUpdateParamsWithContext creates a new VirtualizationClusterGroupsBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClusterGroupsBulkPartialUpdateParamsWithContext(ctx context.Context) *VirtualizationClusterGroupsBulkPartialUpdateParams {
	var ()
	return &VirtualizationClusterGroupsBulkPartialUpdateParams{

		Context: ctx,
	}
}

// NewVirtualizationClusterGroupsBulkPartialUpdateParamsWithHTTPClient creates a new VirtualizationClusterGroupsBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClusterGroupsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClusterGroupsBulkPartialUpdateParams {
	var ()
	return &VirtualizationClusterGroupsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*VirtualizationClusterGroupsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
for the virtualization cluster groups bulk partial update operation typically these are written to a http.Request
*/
type VirtualizationClusterGroupsBulkPartialUpdateParams struct {

	/*Data*/
	Data *models.ClusterGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization cluster groups bulk partial update params
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClusterGroupsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster groups bulk partial update params
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster groups bulk partial update params
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) WithContext(ctx context.Context) *VirtualizationClusterGroupsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster groups bulk partial update params
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster groups bulk partial update params
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClusterGroupsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster groups bulk partial update params
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization cluster groups bulk partial update params
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) WithData(data *models.ClusterGroup) *VirtualizationClusterGroupsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization cluster groups bulk partial update params
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) SetData(data *models.ClusterGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterGroupsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
