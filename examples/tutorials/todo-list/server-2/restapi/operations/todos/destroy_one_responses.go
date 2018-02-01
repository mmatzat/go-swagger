// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mmatzat/go-swagger/examples/tutorials/todo-list/server-2/models"
)

// DestroyOneNoContentCode is the HTTP code returned for type DestroyOneNoContent
const DestroyOneNoContentCode int = 204

/*DestroyOneNoContent Deleted

swagger:response destroyOneNoContent
*/
type DestroyOneNoContent struct {
}

// NewDestroyOneNoContent creates DestroyOneNoContent with default headers values
func NewDestroyOneNoContent() *DestroyOneNoContent {
	return &DestroyOneNoContent{}
}

// WriteResponse to the client
func (o *DestroyOneNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DestroyOneDefault error

swagger:response destroyOneDefault
*/
type DestroyOneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDestroyOneDefault creates DestroyOneDefault with default headers values
func NewDestroyOneDefault(code int) *DestroyOneDefault {
	if code <= 0 {
		code = 500
	}

	return &DestroyOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the destroy one default response
func (o *DestroyOneDefault) WithStatusCode(code int) *DestroyOneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the destroy one default response
func (o *DestroyOneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the destroy one default response
func (o *DestroyOneDefault) WithPayload(payload *models.Error) *DestroyOneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the destroy one default response
func (o *DestroyOneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DestroyOneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
