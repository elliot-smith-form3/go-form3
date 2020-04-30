// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package fx_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fx api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for fx api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create deal API
*/
func (a *CreateDealRequest) Do() (*CreateDealCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDeal",
		Method:             "POST",
		PathPattern:        "/fx/deals",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateDealReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDealCreated), nil

}

func (a *CreateDealRequest) MustDo() *CreateDealCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create deal submission API
*/
func (a *CreateDealSubmissionRequest) Do() (*CreateDealSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDealSubmission",
		Method:             "POST",
		PathPattern:        "/fx/deals/{fx_deal_id}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateDealSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDealSubmissionCreated), nil

}

func (a *CreateDealSubmissionRequest) MustDo() *CreateDealSubmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get deal API
*/
func (a *GetDealRequest) Do() (*GetDealOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDeal",
		Method:             "GET",
		PathPattern:        "/fx/deals/{fx_deal_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDealReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDealOK), nil

}

func (a *GetDealRequest) MustDo() *GetDealOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get deal submission API
*/
func (a *GetDealSubmissionRequest) Do() (*GetDealSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDealSubmission",
		Method:             "GET",
		PathPattern:        "/fx/deals/{fx_deal_id}/submissions/{fx_deal_submission_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDealSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDealSubmissionOK), nil

}

func (a *GetDealSubmissionRequest) MustDo() *GetDealSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list deals API
*/
func (a *ListDealsRequest) Do() (*ListDealsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListDeals",
		Method:             "GET",
		PathPattern:        "/fx/deals",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListDealsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDealsOK), nil

}

func (a *ListDealsRequest) MustDo() *ListDealsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list rates API
*/
func (a *ListRatesRequest) Do() (*ListRatesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListRates",
		Method:             "GET",
		PathPattern:        "/fx/rates",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListRatesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRatesOK), nil

}

func (a *ListRatesRequest) MustDo() *ListRatesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
modify deal submission API
*/
func (a *ModifyDealSubmissionRequest) Do() (*ModifyDealSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyDealSubmission",
		Method:             "PATCH",
		PathPattern:        "/fx/deals/{fx_deal_id}/submissions/{fx_deal_submission_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ModifyDealSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyDealSubmissionOK), nil

}

func (a *ModifyDealSubmissionRequest) MustDo() *ModifyDealSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}