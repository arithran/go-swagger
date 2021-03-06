// Code generated by go-swagger; DO NOT EDIT.

package store

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
	"github.com/go-openapi/swag"
)

// NewOrderGetParams creates a new OrderGetParams object
// with the default values initialized.
func NewOrderGetParams() *OrderGetParams {
	var ()
	return &OrderGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderGetParamsWithTimeout creates a new OrderGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderGetParamsWithTimeout(timeout time.Duration) *OrderGetParams {
	var ()
	return &OrderGetParams{

		timeout: timeout,
	}
}

// NewOrderGetParamsWithContext creates a new OrderGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderGetParamsWithContext(ctx context.Context) *OrderGetParams {
	var ()
	return &OrderGetParams{

		Context: ctx,
	}
}

// NewOrderGetParamsWithHTTPClient creates a new OrderGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderGetParamsWithHTTPClient(client *http.Client) *OrderGetParams {
	var ()
	return &OrderGetParams{
		HTTPClient: client,
	}
}

/*OrderGetParams contains all the parameters to send to the API endpoint
for the order get operation typically these are written to a http.Request
*/
type OrderGetParams struct {

	/*OrderID
	  ID of pet that needs to be fetched

	*/
	OrderID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order get params
func (o *OrderGetParams) WithTimeout(timeout time.Duration) *OrderGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order get params
func (o *OrderGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order get params
func (o *OrderGetParams) WithContext(ctx context.Context) *OrderGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order get params
func (o *OrderGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order get params
func (o *OrderGetParams) WithHTTPClient(client *http.Client) *OrderGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order get params
func (o *OrderGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the order get params
func (o *OrderGetParams) WithOrderID(orderID int64) *OrderGetParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the order get params
func (o *OrderGetParams) SetOrderID(orderID int64) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *OrderGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderId
	if err := r.SetPathParam("orderId", swag.FormatInt64(o.OrderID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
