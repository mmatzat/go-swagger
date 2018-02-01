// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/mmatzat/go-swagger/examples/todo-list/models"
)

// NewAddOneParams creates a new AddOneParams object
// with the default values initialized.
func NewAddOneParams() *AddOneParams {
	var ()
	return &AddOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddOneParamsWithTimeout creates a new AddOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddOneParamsWithTimeout(timeout time.Duration) *AddOneParams {
	var ()
	return &AddOneParams{

		timeout: timeout,
	}
}

// NewAddOneParamsWithContext creates a new AddOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddOneParamsWithContext(ctx context.Context) *AddOneParams {
	var ()
	return &AddOneParams{

		Context: ctx,
	}
}

// NewAddOneParamsWithHTTPClient creates a new AddOneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddOneParamsWithHTTPClient(client *http.Client) *AddOneParams {
	var ()
	return &AddOneParams{
		HTTPClient: client,
	}
}

/*AddOneParams contains all the parameters to send to the API endpoint
for the add one operation typically these are written to a http.Request
*/
type AddOneParams struct {

	/*Body*/
	Body *models.Item

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add one params
func (o *AddOneParams) WithTimeout(timeout time.Duration) *AddOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add one params
func (o *AddOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add one params
func (o *AddOneParams) WithContext(ctx context.Context) *AddOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add one params
func (o *AddOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add one params
func (o *AddOneParams) WithHTTPClient(client *http.Client) *AddOneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add one params
func (o *AddOneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add one params
func (o *AddOneParams) WithBody(body *models.Item) *AddOneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add one params
func (o *AddOneParams) SetBody(body *models.Item) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.Item)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
