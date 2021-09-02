// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/apiclarity/apiclarity/api/server/models"
)

// GetAPIUsageHitCountOKCode is the HTTP code returned for type GetAPIUsageHitCountOK
const GetAPIUsageHitCountOKCode int = 200

/*GetAPIUsageHitCountOK Success

swagger:response getApiUsageHitCountOK
*/
type GetAPIUsageHitCountOK struct {

	/*
	  In: Body
	*/
	Payload []*models.HitCount `json:"body,omitempty"`
}

// NewGetAPIUsageHitCountOK creates GetAPIUsageHitCountOK with default headers values
func NewGetAPIUsageHitCountOK() *GetAPIUsageHitCountOK {

	return &GetAPIUsageHitCountOK{}
}

// WithPayload adds the payload to the get Api usage hit count o k response
func (o *GetAPIUsageHitCountOK) WithPayload(payload []*models.HitCount) *GetAPIUsageHitCountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api usage hit count o k response
func (o *GetAPIUsageHitCountOK) SetPayload(payload []*models.HitCount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIUsageHitCountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.HitCount, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetAPIUsageHitCountDefault unknown error

swagger:response getApiUsageHitCountDefault
*/
type GetAPIUsageHitCountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetAPIUsageHitCountDefault creates GetAPIUsageHitCountDefault with default headers values
func NewGetAPIUsageHitCountDefault(code int) *GetAPIUsageHitCountDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAPIUsageHitCountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get API usage hit count default response
func (o *GetAPIUsageHitCountDefault) WithStatusCode(code int) *GetAPIUsageHitCountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get API usage hit count default response
func (o *GetAPIUsageHitCountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get API usage hit count default response
func (o *GetAPIUsageHitCountDefault) WithPayload(payload *models.APIResponse) *GetAPIUsageHitCountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get API usage hit count default response
func (o *GetAPIUsageHitCountDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIUsageHitCountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
