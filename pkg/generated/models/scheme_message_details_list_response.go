// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SchemeMessageDetailsListResponse scheme message details list response
// swagger:model SchemeMessageDetailsListResponse
type SchemeMessageDetailsListResponse struct {

	// data
	// Required: true
	Data []*SchemeMessage `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func SchemeMessageDetailsListResponseWithDefaults(defaults client.Defaults) *SchemeMessageDetailsListResponse {
	return &SchemeMessageDetailsListResponse{

		Data: make([]*SchemeMessage, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *SchemeMessageDetailsListResponse) WithData(data []*SchemeMessage) *SchemeMessageDetailsListResponse {

	m.Data = data

	return m
}

func (m *SchemeMessageDetailsListResponse) WithLinks(links Links) *SchemeMessageDetailsListResponse {

	m.Links = &links

	return m
}

func (m *SchemeMessageDetailsListResponse) WithoutLinks() *SchemeMessageDetailsListResponse {
	m.Links = nil
	return m
}

// Validate validates this scheme message details list response
func (m *SchemeMessageDetailsListResponse) Validate(formats strfmt.Registry) error {
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

func (m *SchemeMessageDetailsListResponse) validateData(formats strfmt.Registry) error {

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

func (m *SchemeMessageDetailsListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *SchemeMessageDetailsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessageDetailsListResponse) UnmarshalBinary(b []byte) error {
	var res SchemeMessageDetailsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessageDetailsListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
