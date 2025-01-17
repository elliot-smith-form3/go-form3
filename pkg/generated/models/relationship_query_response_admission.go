// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RelationshipQueryResponseAdmission relationship query response admission
// swagger:model RelationshipQueryResponseAdmission
type RelationshipQueryResponseAdmission struct {

	// data
	Data []*QueryResponseAdmission `json:"data"`
}

func RelationshipQueryResponseAdmissionWithDefaults(defaults client.Defaults) *RelationshipQueryResponseAdmission {
	return &RelationshipQueryResponseAdmission{

		Data: make([]*QueryResponseAdmission, 0),
	}
}

func (m *RelationshipQueryResponseAdmission) WithData(data []*QueryResponseAdmission) *RelationshipQueryResponseAdmission {

	m.Data = data

	return m
}

// Validate validates this relationship query response admission
func (m *RelationshipQueryResponseAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationshipQueryResponseAdmission) validateData(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *RelationshipQueryResponseAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipQueryResponseAdmission) UnmarshalBinary(b []byte) error {
	var res RelationshipQueryResponseAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipQueryResponseAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
