// Code generated by go-swagger; DO NOT EDIT.

package payments

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetPaymentReversal creates a new GetPaymentReversalRequest object
// with the default values initialized.
func (c *Client) GetPaymentReversal() *GetPaymentReversalRequest {
	var ()
	return &GetPaymentReversalRequest{

		ID: c.Defaults.GetStrfmtUUID("GetPaymentReversal", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetPaymentReversal", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentReversalRequest struct {

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*ReversalID      Reversal Id      */

	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentReversalRequest) FromJson(j string) (*GetPaymentReversalRequest, error) {

	return o, nil
}

func (o *GetPaymentReversalRequest) WithID(id strfmt.UUID) *GetPaymentReversalRequest {

	o.ID = id

	return o
}

func (o *GetPaymentReversalRequest) WithReversalID(reversalID strfmt.UUID) *GetPaymentReversalRequest {

	o.ReversalID = reversalID

	return o
}

//////////////////
// WithContext adds the context to the get payment reversal Request
func (o *GetPaymentReversalRequest) WithContext(ctx context.Context) *GetPaymentReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment reversal Request
func (o *GetPaymentReversalRequest) WithHTTPClient(client *http.Client) *GetPaymentReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
