// Code generated by go-swagger; DO NOT EDIT.

package payments

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// Client.CreatePaymentSubmission creates a new CreatePaymentSubmissionRequest object
// with the default values initialized.
func (c *Client) CreatePaymentSubmission() *CreatePaymentSubmissionRequest {
	var ()
	return &CreatePaymentSubmissionRequest{

		PaymentSubmissionCreation: models.PaymentSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentSubmission", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentSubmissionRequest struct {

	/*SubmissionCreationRequest*/

	*models.PaymentSubmissionCreation

	/*ID      Payment Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentSubmissionRequest) FromJson(j string) *CreatePaymentSubmissionRequest {

	var m models.PaymentSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.PaymentSubmissionCreation = &m

	return o
}

func (o *CreatePaymentSubmissionRequest) WithSubmissionCreationRequest(submissionCreationRequest models.PaymentSubmissionCreation) *CreatePaymentSubmissionRequest {

	o.PaymentSubmissionCreation = &submissionCreationRequest

	return o
}

func (o *CreatePaymentSubmissionRequest) WithoutSubmissionCreationRequest() *CreatePaymentSubmissionRequest {

	o.PaymentSubmissionCreation = &models.PaymentSubmissionCreation{}

	return o
}

func (o *CreatePaymentSubmissionRequest) WithID(id strfmt.UUID) *CreatePaymentSubmissionRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create payment submission Request
func (o *CreatePaymentSubmissionRequest) WithContext(ctx context.Context) *CreatePaymentSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment submission Request
func (o *CreatePaymentSubmissionRequest) WithHTTPClient(client *http.Client) *CreatePaymentSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.PaymentSubmissionCreation != nil {
		if err := r.SetBodyParam(o.PaymentSubmissionCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
