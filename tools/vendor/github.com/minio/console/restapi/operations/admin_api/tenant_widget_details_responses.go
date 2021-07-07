// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// TenantWidgetDetailsOKCode is the HTTP code returned for type TenantWidgetDetailsOK
const TenantWidgetDetailsOKCode int = 200

/*TenantWidgetDetailsOK A successful response.

swagger:response tenantWidgetDetailsOK
*/
type TenantWidgetDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.WidgetDetails `json:"body,omitempty"`
}

// NewTenantWidgetDetailsOK creates TenantWidgetDetailsOK with default headers values
func NewTenantWidgetDetailsOK() *TenantWidgetDetailsOK {

	return &TenantWidgetDetailsOK{}
}

// WithPayload adds the payload to the tenant widget details o k response
func (o *TenantWidgetDetailsOK) WithPayload(payload *models.WidgetDetails) *TenantWidgetDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tenant widget details o k response
func (o *TenantWidgetDetailsOK) SetPayload(payload *models.WidgetDetails) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TenantWidgetDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TenantWidgetDetailsDefault Generic error response.

swagger:response tenantWidgetDetailsDefault
*/
type TenantWidgetDetailsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTenantWidgetDetailsDefault creates TenantWidgetDetailsDefault with default headers values
func NewTenantWidgetDetailsDefault(code int) *TenantWidgetDetailsDefault {
	if code <= 0 {
		code = 500
	}

	return &TenantWidgetDetailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the tenant widget details default response
func (o *TenantWidgetDetailsDefault) WithStatusCode(code int) *TenantWidgetDetailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the tenant widget details default response
func (o *TenantWidgetDetailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the tenant widget details default response
func (o *TenantWidgetDetailsDefault) WithPayload(payload *models.Error) *TenantWidgetDetailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tenant widget details default response
func (o *TenantWidgetDetailsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TenantWidgetDetailsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}