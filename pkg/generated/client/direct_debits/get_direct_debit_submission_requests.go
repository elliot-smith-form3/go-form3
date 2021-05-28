// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetDirectDebitSubmission creates a new GetDirectDebitSubmissionRequest object
// with the default values initialized.
func (c *Client) GetDirectDebitSubmission() *GetDirectDebitSubmissionRequest {
	var ()
	return &GetDirectDebitSubmissionRequest{

		ID: c.Defaults.GetStrfmtUUID("GetDirectDebitSubmission", "id"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetDirectDebitSubmission", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectDebitSubmissionRequest struct {

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*SubmissionID      Direct Debit decision submission id      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetDirectDebitSubmissionRequest) FromJson(j string) (*GetDirectDebitSubmissionRequest, error) {

	return o, nil
}

func (o *GetDirectDebitSubmissionRequest) WithID(id strfmt.UUID) *GetDirectDebitSubmissionRequest {

	o.ID = id

	return o
}

func (o *GetDirectDebitSubmissionRequest) WithSubmissionID(submissionID strfmt.UUID) *GetDirectDebitSubmissionRequest {

	o.SubmissionID = submissionID

	return o
}

//////////////////
// WithContext adds the context to the get direct debit submission Request
func (o *GetDirectDebitSubmissionRequest) WithContext(ctx context.Context) *GetDirectDebitSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get direct debit submission Request
func (o *GetDirectDebitSubmissionRequest) WithHTTPClient(client *http.Client) *GetDirectDebitSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectDebitSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
