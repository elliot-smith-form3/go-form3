// Code generated by go-swagger; DO NOT EDIT.

package addressbook

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

// Client.GetAddressbookPartiesIDAccountsPartyAccountID creates a new GetAddressbookPartiesIDAccountsPartyAccountIDRequest object
// with the default values initialized.
func (c *Client) GetAddressbookPartiesIDAccountsPartyAccountID() *GetAddressbookPartiesIDAccountsPartyAccountIDRequest {
	var ()
	return &GetAddressbookPartiesIDAccountsPartyAccountIDRequest{

		ID: c.Defaults.GetStrfmtUUID("GetAddressbookPartiesIDAccountsPartyAccountID", "id"),

		PartyAccountID: c.Defaults.GetStrfmtUUID("GetAddressbookPartiesIDAccountsPartyAccountID", "party_account_id"),

		Version: c.Defaults.GetInt64Ptr("GetAddressbookPartiesIDAccountsPartyAccountID", "version"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAddressbookPartiesIDAccountsPartyAccountIDRequest struct {

	/*ID      Id of party      */

	ID strfmt.UUID

	/*PartyAccountID      Id of party account      */

	PartyAccountID strfmt.UUID

	/*Version*/

	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) FromJson(j string) (*GetAddressbookPartiesIDAccountsPartyAccountIDRequest, error) {

	return o, nil
}

func (o *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) WithID(id strfmt.UUID) *GetAddressbookPartiesIDAccountsPartyAccountIDRequest {

	o.ID = id

	return o
}

func (o *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) WithPartyAccountID(partyAccountID strfmt.UUID) *GetAddressbookPartiesIDAccountsPartyAccountIDRequest {

	o.PartyAccountID = partyAccountID

	return o
}

func (o *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) WithVersion(version int64) *GetAddressbookPartiesIDAccountsPartyAccountIDRequest {

	o.Version = &version

	return o
}

func (o *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) WithoutVersion() *GetAddressbookPartiesIDAccountsPartyAccountIDRequest {

	o.Version = nil

	return o
}

//////////////////
// WithContext adds the context to the get addressbook parties ID accounts party account ID Request
func (o *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) WithContext(ctx context.Context) *GetAddressbookPartiesIDAccountsPartyAccountIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get addressbook parties ID accounts party account ID Request
func (o *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) WithHTTPClient(client *http.Client) *GetAddressbookPartiesIDAccountsPartyAccountIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAddressbookPartiesIDAccountsPartyAccountIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param party_account_id
	if err := r.SetPathParam("party_account_id", o.PartyAccountID.String()); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
