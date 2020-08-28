// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package fundingrequest_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fundingrequest api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for fundingrequest api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get transaction fundingrequests API
*/
func (a *GetTransactionFundingrequestsRequest) Do() (*GetTransactionFundingrequestsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionFundingrequests",
		Method:             "GET",
		PathPattern:        "/transaction/fundingrequests",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionFundingrequestsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionFundingrequestsOK), nil

}

func (a *GetTransactionFundingrequestsRequest) MustDo() *GetTransactionFundingrequestsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction fundingrequests funding request ID API
*/
func (a *GetTransactionFundingrequestsFundingRequestIDRequest) Do() (*GetTransactionFundingrequestsFundingRequestIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionFundingrequestsFundingRequestID",
		Method:             "GET",
		PathPattern:        "/transaction/fundingrequests/{funding_request_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionFundingrequestsFundingRequestIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionFundingrequestsFundingRequestIDOK), nil

}

func (a *GetTransactionFundingrequestsFundingRequestIDRequest) MustDo() *GetTransactionFundingrequestsFundingRequestIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction fundingrequests funding request ID admissions funding request admission ID API
*/
func (a *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDRequest) Do() (*GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionID",
		Method:             "GET",
		PathPattern:        "/transaction/fundingrequests/{funding_request_id}/admissions/{funding_request_admission_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK), nil

}

func (a *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDRequest) MustDo() *GetTransactionFundingrequestsFundingRequestIDAdmissionsFundingRequestAdmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction fundingrequests API
*/
func (a *PostTransactionFundingrequestsRequest) Do() (*PostTransactionFundingrequestsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionFundingrequests",
		Method:             "POST",
		PathPattern:        "/transaction/fundingrequests",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionFundingrequestsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionFundingrequestsCreated), nil

}

func (a *PostTransactionFundingrequestsRequest) MustDo() *PostTransactionFundingrequestsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction fundingrequests funding request ID admissions API
*/
func (a *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) Do() (*PostTransactionFundingrequestsFundingRequestIDAdmissionsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionFundingrequestsFundingRequestIDAdmissions",
		Method:             "POST",
		PathPattern:        "/transaction/fundingrequests/{funding_request_id}/admissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionFundingrequestsFundingRequestIDAdmissionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionFundingrequestsFundingRequestIDAdmissionsCreated), nil

}

func (a *PostTransactionFundingrequestsFundingRequestIDAdmissionsRequest) MustDo() *PostTransactionFundingrequestsFundingRequestIDAdmissionsCreated {
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
