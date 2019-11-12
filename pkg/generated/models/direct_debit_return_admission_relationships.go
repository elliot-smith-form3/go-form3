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
)

// DirectDebitReturnAdmissionRelationships direct debit return admission relationships
// swagger:model DirectDebitReturnAdmissionRelationships
type DirectDebitReturnAdmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitReturnAdmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit return
	DirectDebitReturn *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn `json:"direct_debit_return,omitempty"`
}

func DirectDebitReturnAdmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitReturnAdmissionRelationships {
	return &DirectDebitReturnAdmissionRelationships{

		DirectDebit: DirectDebitReturnAdmissionRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitReturn: DirectDebitReturnAdmissionRelationshipsDirectDebitReturnWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnAdmissionRelationships) WithDirectDebit(directDebit DirectDebitReturnAdmissionRelationshipsDirectDebit) *DirectDebitReturnAdmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitReturnAdmissionRelationships) WithoutDirectDebit() *DirectDebitReturnAdmissionRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitReturnAdmissionRelationships) WithDirectDebitReturn(directDebitReturn DirectDebitReturnAdmissionRelationshipsDirectDebitReturn) *DirectDebitReturnAdmissionRelationships {

	m.DirectDebitReturn = &directDebitReturn

	return m
}

func (m *DirectDebitReturnAdmissionRelationships) WithoutDirectDebitReturn() *DirectDebitReturnAdmissionRelationships {
	m.DirectDebitReturn = nil
	return m
}

// Validate validates this direct debit return admission relationships
func (m *DirectDebitReturnAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnAdmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnAdmissionRelationships) validateDirectDebitReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturn) { // not required
		return nil
	}

	if m.DirectDebitReturn != nil {
		if err := m.DirectDebitReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit_return")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnAdmissionRelationshipsDirectDebit direct debit return admission relationships direct debit
// swagger:model DirectDebitReturnAdmissionRelationshipsDirectDebit
type DirectDebitReturnAdmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitReturnAdmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitReturnAdmissionRelationshipsDirectDebit {
	return &DirectDebitReturnAdmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitReturnAdmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitReturnAdmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit return admission relationships direct debit
func (m *DirectDebitReturnAdmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnAdmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnAdmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnAdmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnAdmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnAdmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnAdmissionRelationshipsDirectDebitReturn direct debit return admission relationships direct debit return
// swagger:model DirectDebitReturnAdmissionRelationshipsDirectDebitReturn
type DirectDebitReturnAdmissionRelationshipsDirectDebitReturn struct {

	// data
	Data []*DirectDebitReturn `json:"data"`
}

func DirectDebitReturnAdmissionRelationshipsDirectDebitReturnWithDefaults(defaults client.Defaults) *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn {
	return &DirectDebitReturnAdmissionRelationshipsDirectDebitReturn{

		Data: make([]*DirectDebitReturn, 0),
	}
}

func (m *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn) WithData(data []*DirectDebitReturn) *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn {

	m.Data = data

	return m
}

// Validate validates this direct debit return admission relationships direct debit return
func (m *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnAdmissionRelationshipsDirectDebitReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnAdmissionRelationshipsDirectDebitReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
