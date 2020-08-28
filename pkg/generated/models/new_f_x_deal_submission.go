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
	"github.com/go-openapi/validate"
)

// NewFXDealSubmission new f x deal submission
// swagger:model NewFXDealSubmission
type NewFXDealSubmission struct {

	// attributes
	// Required: true
	Attributes *NewFXDealSubmissionAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// type
	// Required: true
	Type FXDealSubmissionResourceType `json:"type"`
}

func NewFXDealSubmissionWithDefaults(defaults client.Defaults) *NewFXDealSubmission {
	return &NewFXDealSubmission{

		Attributes: NewFXDealSubmissionAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("NewFXDealSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewFXDealSubmission", "organisation_id"),

		// TODO Type: FXDealSubmissionResourceType,

	}
}

func (m *NewFXDealSubmission) WithAttributes(attributes NewFXDealSubmissionAttributes) *NewFXDealSubmission {

	m.Attributes = &attributes

	return m
}

func (m *NewFXDealSubmission) WithoutAttributes() *NewFXDealSubmission {
	m.Attributes = nil
	return m
}

func (m *NewFXDealSubmission) WithID(id strfmt.UUID) *NewFXDealSubmission {

	m.ID = &id

	return m
}

func (m *NewFXDealSubmission) WithoutID() *NewFXDealSubmission {
	m.ID = nil
	return m
}

func (m *NewFXDealSubmission) WithOrganisationID(organisationID strfmt.UUID) *NewFXDealSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewFXDealSubmission) WithoutOrganisationID() *NewFXDealSubmission {
	m.OrganisationID = nil
	return m
}

func (m *NewFXDealSubmission) WithType(typeVar FXDealSubmissionResourceType) *NewFXDealSubmission {

	m.Type = typeVar

	return m
}

// Validate validates this new f x deal submission
func (m *NewFXDealSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewFXDealSubmission) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *NewFXDealSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewFXDealSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewFXDealSubmission) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewFXDealSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewFXDealSubmission) UnmarshalBinary(b []byte) error {
	var res NewFXDealSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewFXDealSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
