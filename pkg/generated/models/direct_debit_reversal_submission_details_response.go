// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DirectDebitReversalSubmissionDetailsResponse direct debit reversal submission details response
// swagger:model DirectDebitReversalSubmissionDetailsResponse
type DirectDebitReversalSubmissionDetailsResponse struct {

	// data
	Data *DirectDebitReversalSubmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func DirectDebitReversalSubmissionDetailsResponseWithDefaults(defaults client.Defaults) *DirectDebitReversalSubmissionDetailsResponse {
	return &DirectDebitReversalSubmissionDetailsResponse{

		Data: DirectDebitReversalSubmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *DirectDebitReversalSubmissionDetailsResponse) WithData(data DirectDebitReversalSubmission) *DirectDebitReversalSubmissionDetailsResponse {

	m.Data = &data

	return m
}

func (m *DirectDebitReversalSubmissionDetailsResponse) WithoutData() *DirectDebitReversalSubmissionDetailsResponse {
	m.Data = nil
	return m
}

func (m *DirectDebitReversalSubmissionDetailsResponse) WithLinks(links Links) *DirectDebitReversalSubmissionDetailsResponse {

	m.Links = &links

	return m
}

func (m *DirectDebitReversalSubmissionDetailsResponse) WithoutLinks() *DirectDebitReversalSubmissionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this direct debit reversal submission details response
func (m *DirectDebitReversalSubmissionDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalSubmissionDetailsResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReversalSubmissionDetailsResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalSubmissionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalSubmissionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
