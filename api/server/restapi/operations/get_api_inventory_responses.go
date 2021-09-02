// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/apiclarity/apiclarity/api/server/models"
)

// GetAPIInventoryOKCode is the HTTP code returned for type GetAPIInventoryOK
const GetAPIInventoryOKCode int = 200

/*GetAPIInventoryOK Success

swagger:response getApiInventoryOK
*/
type GetAPIInventoryOK struct {

	/*
	  In: Body
	*/
	Payload *GetAPIInventoryOKBody `json:"body,omitempty"`
}

// NewGetAPIInventoryOK creates GetAPIInventoryOK with default headers values
func NewGetAPIInventoryOK() *GetAPIInventoryOK {

	return &GetAPIInventoryOK{}
}

// WithPayload adds the payload to the get Api inventory o k response
func (o *GetAPIInventoryOK) WithPayload(payload *GetAPIInventoryOKBody) *GetAPIInventoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api inventory o k response
func (o *GetAPIInventoryOK) SetPayload(payload *GetAPIInventoryOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAPIInventoryDefault unknown error

swagger:response getApiInventoryDefault
*/
type GetAPIInventoryDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetAPIInventoryDefault creates GetAPIInventoryDefault with default headers values
func NewGetAPIInventoryDefault(code int) *GetAPIInventoryDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAPIInventoryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get API inventory default response
func (o *GetAPIInventoryDefault) WithStatusCode(code int) *GetAPIInventoryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get API inventory default response
func (o *GetAPIInventoryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get API inventory default response
func (o *GetAPIInventoryDefault) WithPayload(payload *models.APIResponse) *GetAPIInventoryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get API inventory default response
func (o *GetAPIInventoryDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
