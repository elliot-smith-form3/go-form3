// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserCredentialListResponse user credential list response
// swagger:model UserCredentialListResponse
type UserCredentialListResponse struct {

	// data
	// Required: true
	Data []*Credential `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func UserCredentialListResponseWithDefaults(defaults client.Defaults) *UserCredentialListResponse {
	return &UserCredentialListResponse{

		Data: make([]*Credential, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *UserCredentialListResponse) WithData(data []*Credential) *UserCredentialListResponse {

	m.Data = data

	return m
}

func (m *UserCredentialListResponse) WithLinks(links Links) *UserCredentialListResponse {

	m.Links = &links

	return m
}

func (m *UserCredentialListResponse) WithoutLinks() *UserCredentialListResponse {
	m.Links = nil
	return m
}

// Validate validates this user credential list response
func (m *UserCredentialListResponse) Validate(formats strfmt.Registry) error {
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

func (m *UserCredentialListResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserCredentialListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *UserCredentialListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCredentialListResponse) UnmarshalBinary(b []byte) error {
	var res UserCredentialListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *UserCredentialListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
