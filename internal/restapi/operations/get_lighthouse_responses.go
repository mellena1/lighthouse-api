// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetLighthouseOKCode is the HTTP code returned for type GetLighthouseOK
const GetLighthouseOKCode int = 200

/*GetLighthouseOK All known lighthouses

swagger:response getLighthouseOK
*/
type GetLighthouseOK struct {

	/*
	  In: Body
	*/
	Payload *GetLighthouseOKBody `json:"body,omitempty"`
}

// NewGetLighthouseOK creates GetLighthouseOK with default headers values
func NewGetLighthouseOK() *GetLighthouseOK {

	return &GetLighthouseOK{}
}

// WithPayload adds the payload to the get lighthouse o k response
func (o *GetLighthouseOK) WithPayload(payload *GetLighthouseOKBody) *GetLighthouseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lighthouse o k response
func (o *GetLighthouseOK) SetPayload(payload *GetLighthouseOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLighthouseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}