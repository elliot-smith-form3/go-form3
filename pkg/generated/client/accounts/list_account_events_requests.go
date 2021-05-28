// Code generated by go-swagger; DO NOT EDIT.

package accounts

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.ListAccountEvents creates a new ListAccountEventsRequest object
// with the default values initialized.
func (c *Client) ListAccountEvents() *ListAccountEventsRequest {
	var ()
	return &ListAccountEventsRequest{

		ID: c.Defaults.GetStrfmtUUID("ListAccountEvents", "id"),

		PageNumber: c.Defaults.GetStringPtr("ListAccountEvents", "page[number]"),

		PageSize: c.Defaults.GetInt64Ptr("ListAccountEvents", "page[size]"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListAccountEventsRequest struct {

	/*ID      Account Id      */

	ID strfmt.UUID

	/*PageNumber      Which page to select      */

	PageNumber *string

	/*PageSize      Number of items to select      */

	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ListAccountEventsRequest) FromJson(j string) (*ListAccountEventsRequest, error) {

	return o, nil
}

func (o *ListAccountEventsRequest) WithID(id strfmt.UUID) *ListAccountEventsRequest {

	o.ID = id

	return o
}

func (o *ListAccountEventsRequest) WithPageNumber(pageNumber string) *ListAccountEventsRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListAccountEventsRequest) WithoutPageNumber() *ListAccountEventsRequest {

	o.PageNumber = nil

	return o
}

func (o *ListAccountEventsRequest) WithPageSize(pageSize int64) *ListAccountEventsRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListAccountEventsRequest) WithoutPageSize() *ListAccountEventsRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list account events Request
func (o *ListAccountEventsRequest) WithContext(ctx context.Context) *ListAccountEventsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list account events Request
func (o *ListAccountEventsRequest) WithHTTPClient(client *http.Client) *ListAccountEventsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListAccountEventsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
