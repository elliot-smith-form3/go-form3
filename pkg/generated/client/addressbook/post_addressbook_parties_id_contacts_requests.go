// Code generated by go-swagger; DO NOT EDIT.

package addressbook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// Client.PostAddressbookPartiesIDContacts creates a new PostAddressbookPartiesIDContactsRequest object
// with the default values initialized.
func (c *Client) PostAddressbookPartiesIDContacts() *PostAddressbookPartiesIDContactsRequest {
	var ()
	return &PostAddressbookPartiesIDContactsRequest{

		ContactCreation: models.ContactCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("PostAddressbookPartiesIDContacts", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostAddressbookPartiesIDContactsRequest struct {

	/*CreationRequest*/

	*models.ContactCreation

	/*ID      Id of party      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostAddressbookPartiesIDContactsRequest) FromJson(j string) (*PostAddressbookPartiesIDContactsRequest, error) {

	var m models.ContactCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.ContactCreation = &m

	return o, nil
}

func (o *PostAddressbookPartiesIDContactsRequest) WithCreationRequest(creationRequest models.ContactCreation) *PostAddressbookPartiesIDContactsRequest {

	o.ContactCreation = &creationRequest

	return o
}

func (o *PostAddressbookPartiesIDContactsRequest) WithoutCreationRequest() *PostAddressbookPartiesIDContactsRequest {

	o.ContactCreation = &models.ContactCreation{}

	return o
}

func (o *PostAddressbookPartiesIDContactsRequest) WithID(id strfmt.UUID) *PostAddressbookPartiesIDContactsRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the post addressbook parties ID contacts Request
func (o *PostAddressbookPartiesIDContactsRequest) WithContext(ctx context.Context) *PostAddressbookPartiesIDContactsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post addressbook parties ID contacts Request
func (o *PostAddressbookPartiesIDContactsRequest) WithHTTPClient(client *http.Client) *PostAddressbookPartiesIDContactsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostAddressbookPartiesIDContactsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ContactCreation != nil {
		if err := r.SetBodyParam(o.ContactCreation); err != nil {
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
