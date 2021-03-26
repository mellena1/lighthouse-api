// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostScanOKCode is the HTTP code returned for type PostScanOK
const PostScanOKCode int = 200

/*PostScanOK Rescan successfully completed

swagger:response postScanOK
*/
type PostScanOK struct {

	/*
	  In: Body
	*/
	Payload *PostScanOKBody `json:"body,omitempty"`
}

// NewPostScanOK creates PostScanOK with default headers values
func NewPostScanOK() *PostScanOK {

	return &PostScanOK{}
}

// WithPayload adds the payload to the post scan o k response
func (o *PostScanOK) WithPayload(payload *PostScanOKBody) *PostScanOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post scan o k response
func (o *PostScanOK) SetPayload(payload *PostScanOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostScanOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostScanBadRequestCode is the HTTP code returned for type PostScanBadRequest
const PostScanBadRequestCode int = 400

/*PostScanBadRequest Malformed request

swagger:response postScanBadRequest
*/
type PostScanBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostScanBadRequestBody `json:"body,omitempty"`
}

// NewPostScanBadRequest creates PostScanBadRequest with default headers values
func NewPostScanBadRequest() *PostScanBadRequest {

	return &PostScanBadRequest{}
}

// WithPayload adds the payload to the post scan bad request response
func (o *PostScanBadRequest) WithPayload(payload *PostScanBadRequestBody) *PostScanBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post scan bad request response
func (o *PostScanBadRequest) SetPayload(payload *PostScanBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostScanBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostScanTooManyRequestsCode is the HTTP code returned for type PostScanTooManyRequests
const PostScanTooManyRequestsCode int = 429

/*PostScanTooManyRequests Rate limited (can only rescan so often)

swagger:response postScanTooManyRequests
*/
type PostScanTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *PostScanTooManyRequestsBody `json:"body,omitempty"`
}

// NewPostScanTooManyRequests creates PostScanTooManyRequests with default headers values
func NewPostScanTooManyRequests() *PostScanTooManyRequests {

	return &PostScanTooManyRequests{}
}

// WithPayload adds the payload to the post scan too many requests response
func (o *PostScanTooManyRequests) WithPayload(payload *PostScanTooManyRequestsBody) *PostScanTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post scan too many requests response
func (o *PostScanTooManyRequests) SetPayload(payload *PostScanTooManyRequestsBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostScanTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}