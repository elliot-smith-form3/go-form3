// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// BranchValidationType optional validation to apply to the branch
// swagger:model BranchValidationType
type BranchValidationType string

const (

	// BranchValidationTypeCard captures enum value "card"
	BranchValidationTypeCard BranchValidationType = "card"

	// BranchValidationTypeAccountModCheck captures enum value "account_mod_check"
	BranchValidationTypeAccountModCheck BranchValidationType = "account_mod_check"

	// BranchValidationTypeNone captures enum value "none"
	BranchValidationTypeNone BranchValidationType = "none"
)

// for schema
var branchValidationTypeEnum []interface{}

func init() {
	var res []BranchValidationType
	if err := json.Unmarshal([]byte(`["card","account_mod_check","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		branchValidationTypeEnum = append(branchValidationTypeEnum, v)
	}
}

func (m BranchValidationType) validateBranchValidationTypeEnum(path, location string, value BranchValidationType) error {
	if err := validate.Enum(path, location, value, branchValidationTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this branch validation type
func (m BranchValidationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBranchValidationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchValidationType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
