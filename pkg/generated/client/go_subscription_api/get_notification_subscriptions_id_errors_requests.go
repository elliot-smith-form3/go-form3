// Code generated by go-swagger; DO NOT EDIT.

package go_subscription_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetNotificationSubscriptionsIDErrors creates a new GetNotificationSubscriptionsIDErrorsRequest object
// with the default values initialized.
func (c *Client) GetNotificationSubscriptionsIDErrors() *GetNotificationSubscriptionsIDErrorsRequest {
	var ()
	return &GetNotificationSubscriptionsIDErrorsRequest{

		ID: c.Defaults.GetStrfmtUUID("GetNotificationSubscriptionsIDErrors", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetNotificationSubscriptionsIDErrorsRequest struct {

	/*ID*/

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetNotificationSubscriptionsIDErrorsRequest) FromJson(j string) (*GetNotificationSubscriptionsIDErrorsRequest, error) {

	return o, nil
}

func (o *GetNotificationSubscriptionsIDErrorsRequest) WithID(id strfmt.UUID) *GetNotificationSubscriptionsIDErrorsRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get notification subscriptions ID errors Request
func (o *GetNotificationSubscriptionsIDErrorsRequest) WithContext(ctx context.Context) *GetNotificationSubscriptionsIDErrorsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get notification subscriptions ID errors Request
func (o *GetNotificationSubscriptionsIDErrorsRequest) WithHTTPClient(client *http.Client) *GetNotificationSubscriptionsIDErrorsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetNotificationSubscriptionsIDErrorsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
