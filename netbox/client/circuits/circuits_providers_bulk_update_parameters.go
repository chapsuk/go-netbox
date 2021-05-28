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

package circuits

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

// NewCircuitsProvidersBulkUpdateParams creates a new CircuitsProvidersBulkUpdateParams object
// with the default values initialized.
func NewCircuitsProvidersBulkUpdateParams() *CircuitsProvidersBulkUpdateParams {
	var ()
	return &CircuitsProvidersBulkUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersBulkUpdateParamsWithTimeout creates a new CircuitsProvidersBulkUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsProvidersBulkUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersBulkUpdateParams {
	var ()
	return &CircuitsProvidersBulkUpdateParams{

		timeout: timeout,
	}
}

// NewCircuitsProvidersBulkUpdateParamsWithContext creates a new CircuitsProvidersBulkUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsProvidersBulkUpdateParamsWithContext(ctx context.Context) *CircuitsProvidersBulkUpdateParams {
	var ()
	return &CircuitsProvidersBulkUpdateParams{

		Context: ctx,
	}
}

// NewCircuitsProvidersBulkUpdateParamsWithHTTPClient creates a new CircuitsProvidersBulkUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsProvidersBulkUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersBulkUpdateParams {
	var ()
	return &CircuitsProvidersBulkUpdateParams{
		HTTPClient: client,
	}
}

/*CircuitsProvidersBulkUpdateParams contains all the parameters to send to the API endpoint
for the circuits providers bulk update operation typically these are written to a http.Request
*/
type CircuitsProvidersBulkUpdateParams struct {

	/*Data*/
	Data *models.Provider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) WithContext(ctx context.Context) *CircuitsProvidersBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) WithData(data *models.Provider) *CircuitsProvidersBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
