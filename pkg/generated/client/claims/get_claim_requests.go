// Code generated by go-swagger; DO NOT EDIT.

package claims

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetClaim creates a new GetClaimRequest object
// with the default values initialized.
func (c *Client) GetClaim() *GetClaimRequest {
	var ()
	return &GetClaimRequest{

		ID: c.Defaults.GetStrfmtUUID("GetClaim", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetClaimRequest struct {

	/*ID      Claim Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetClaimRequest) FromJson(j string) (*GetClaimRequest, error) {

	return o, nil
}

func (o *GetClaimRequest) WithID(id strfmt.UUID) *GetClaimRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get claim Request
func (o *GetClaimRequest) WithContext(ctx context.Context) *GetClaimRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get claim Request
func (o *GetClaimRequest) WithHTTPClient(client *http.Client) *GetClaimRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetClaimRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
