// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/apiclarity/apiclarity/api/server/models"
)

// GetAPIInventoryAPIIDProvidedSwaggerJSONOKCode is the HTTP code returned for type GetAPIInventoryAPIIDProvidedSwaggerJSONOK
const GetAPIInventoryAPIIDProvidedSwaggerJSONOKCode int = 200

/*GetAPIInventoryAPIIDProvidedSwaggerJSONOK Success

swagger:response getApiInventoryApiIdProvidedSwaggerJsonOK
*/
type GetAPIInventoryAPIIDProvidedSwaggerJSONOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetAPIInventoryAPIIDProvidedSwaggerJSONOK creates GetAPIInventoryAPIIDProvidedSwaggerJSONOK with default headers values
func NewGetAPIInventoryAPIIDProvidedSwaggerJSONOK() *GetAPIInventoryAPIIDProvidedSwaggerJSONOK {

	return &GetAPIInventoryAPIIDProvidedSwaggerJSONOK{}
}

// WithPayload adds the payload to the get Api inventory Api Id provided swagger Json o k response
func (o *GetAPIInventoryAPIIDProvidedSwaggerJSONOK) WithPayload(payload interface{}) *GetAPIInventoryAPIIDProvidedSwaggerJSONOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api inventory Api Id provided swagger Json o k response
func (o *GetAPIInventoryAPIIDProvidedSwaggerJSONOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryAPIIDProvidedSwaggerJSONOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetAPIInventoryAPIIDProvidedSwaggerJSONDefault unknown error

swagger:response getApiInventoryApiIdProvidedSwaggerJsonDefault
*/
type GetAPIInventoryAPIIDProvidedSwaggerJSONDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetAPIInventoryAPIIDProvidedSwaggerJSONDefault creates GetAPIInventoryAPIIDProvidedSwaggerJSONDefault with default headers values
func NewGetAPIInventoryAPIIDProvidedSwaggerJSONDefault(code int) *GetAPIInventoryAPIIDProvidedSwaggerJSONDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAPIInventoryAPIIDProvidedSwaggerJSONDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get API inventory API ID provided swagger JSON default response
func (o *GetAPIInventoryAPIIDProvidedSwaggerJSONDefault) WithStatusCode(code int) *GetAPIInventoryAPIIDProvidedSwaggerJSONDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get API inventory API ID provided swagger JSON default response
func (o *GetAPIInventoryAPIIDProvidedSwaggerJSONDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get API inventory API ID provided swagger JSON default response
func (o *GetAPIInventoryAPIIDProvidedSwaggerJSONDefault) WithPayload(payload *models.APIResponse) *GetAPIInventoryAPIIDProvidedSwaggerJSONDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get API inventory API ID provided swagger JSON default response
func (o *GetAPIInventoryAPIIDProvidedSwaggerJSONDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryAPIIDProvidedSwaggerJSONDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
