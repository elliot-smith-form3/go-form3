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
)

// AuditEntryListResponse audit entry list response
// swagger:model AuditEntryListResponse
type AuditEntryListResponse struct {

	// data
	Data []*AuditEntry `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func AuditEntryListResponseWithDefaults(defaults client.Defaults) *AuditEntryListResponse {
	return &AuditEntryListResponse{

		Data: make([]*AuditEntry, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *AuditEntryListResponse) WithData(data []*AuditEntry) *AuditEntryListResponse {

	m.Data = data

	return m
}

func (m *AuditEntryListResponse) WithLinks(links Links) *AuditEntryListResponse {

	m.Links = &links

	return m
}

func (m *AuditEntryListResponse) WithoutLinks() *AuditEntryListResponse {
	m.Links = nil
	return m
}

// Validate validates this audit entry list response
func (m *AuditEntryListResponse) Validate(formats strfmt.Registry) error {
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

func (m *AuditEntryListResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
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

func (m *AuditEntryListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *AuditEntryListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditEntryListResponse) UnmarshalBinary(b []byte) error {
	var res AuditEntryListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AuditEntryListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
