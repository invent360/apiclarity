// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/apiclarity/apiclarity/api/server/models"
)

// GetAPIInventoryAPIIDReconstructedSwaggerJSONOKCode is the HTTP code returned for type GetAPIInventoryAPIIDReconstructedSwaggerJSONOK
const GetAPIInventoryAPIIDReconstructedSwaggerJSONOKCode int = 200

/*GetAPIInventoryAPIIDReconstructedSwaggerJSONOK Success

swagger:response getApiInventoryApiIdReconstructedSwaggerJsonOK
*/
type GetAPIInventoryAPIIDReconstructedSwaggerJSONOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetAPIInventoryAPIIDReconstructedSwaggerJSONOK creates GetAPIInventoryAPIIDReconstructedSwaggerJSONOK with default headers values
func NewGetAPIInventoryAPIIDReconstructedSwaggerJSONOK() *GetAPIInventoryAPIIDReconstructedSwaggerJSONOK {

	return &GetAPIInventoryAPIIDReconstructedSwaggerJSONOK{}
}

// WithPayload adds the payload to the get Api inventory Api Id reconstructed swagger Json o k response
func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSONOK) WithPayload(payload interface{}) *GetAPIInventoryAPIIDReconstructedSwaggerJSONOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api inventory Api Id reconstructed swagger Json o k response
func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSONOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSONOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault unknown error

swagger:response getApiInventoryApiIdReconstructedSwaggerJsonDefault
*/
type GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetAPIInventoryAPIIDReconstructedSwaggerJSONDefault creates GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault with default headers values
func NewGetAPIInventoryAPIIDReconstructedSwaggerJSONDefault(code int) *GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get API inventory API ID reconstructed swagger JSON default response
func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault) WithStatusCode(code int) *GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get API inventory API ID reconstructed swagger JSON default response
func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get API inventory API ID reconstructed swagger JSON default response
func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault) WithPayload(payload *models.APIResponse) *GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get API inventory API ID reconstructed swagger JSON default response
func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSONDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
