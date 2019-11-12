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

// DirectDebitResourceType direct debit resource type
// swagger:model DirectDebitResourceType
type DirectDebitResourceType string

const (

	// DirectDebitResourceTypeDirectDebits captures enum value "direct_debits"
	DirectDebitResourceTypeDirectDebits DirectDebitResourceType = "direct_debits"

	// DirectDebitResourceTypeDirectDebitAdmissions captures enum value "direct_debit_admissions"
	DirectDebitResourceTypeDirectDebitAdmissions DirectDebitResourceType = "direct_debit_admissions"

	// DirectDebitResourceTypeDirectDebitSubmissions captures enum value "direct_debit_submissions"
	DirectDebitResourceTypeDirectDebitSubmissions DirectDebitResourceType = "direct_debit_submissions"

	// DirectDebitResourceTypeDirectDebitSubmissionValidations captures enum value "direct_debit_submission_validations"
	DirectDebitResourceTypeDirectDebitSubmissionValidations DirectDebitResourceType = "direct_debit_submission_validations"

	// DirectDebitResourceTypeDirectDebitReversals captures enum value "direct_debit_reversals"
	DirectDebitResourceTypeDirectDebitReversals DirectDebitResourceType = "direct_debit_reversals"

	// DirectDebitResourceTypeDirectDebitReversalAdmissions captures enum value "direct_debit_reversal_admissions"
	DirectDebitResourceTypeDirectDebitReversalAdmissions DirectDebitResourceType = "direct_debit_reversal_admissions"

	// DirectDebitResourceTypeDirectDebitReturns captures enum value "direct_debit_returns"
	DirectDebitResourceTypeDirectDebitReturns DirectDebitResourceType = "direct_debit_returns"

	// DirectDebitResourceTypeDirectDebitReturnAdmissions captures enum value "direct_debit_return_admissions"
	DirectDebitResourceTypeDirectDebitReturnAdmissions DirectDebitResourceType = "direct_debit_return_admissions"

	// DirectDebitResourceTypeDirectDebitReturnSubmissions captures enum value "direct_debit_return_submissions"
	DirectDebitResourceTypeDirectDebitReturnSubmissions DirectDebitResourceType = "direct_debit_return_submissions"

	// DirectDebitResourceTypeDirectDebitReturnSubmissionValidations captures enum value "direct_debit_return_submission_validations"
	DirectDebitResourceTypeDirectDebitReturnSubmissionValidations DirectDebitResourceType = "direct_debit_return_submission_validations"

	// DirectDebitResourceTypeDirectDebitReturnReversals captures enum value "direct_debit_return_reversals"
	DirectDebitResourceTypeDirectDebitReturnReversals DirectDebitResourceType = "direct_debit_return_reversals"

	// DirectDebitResourceTypeDirectDebitReturnReversalAdmissions captures enum value "direct_debit_return_reversal_admissions"
	DirectDebitResourceTypeDirectDebitReturnReversalAdmissions DirectDebitResourceType = "direct_debit_return_reversal_admissions"

	// DirectDebitResourceTypeDirectDebitAutomaticReturns captures enum value "direct_debit_automatic_returns"
	DirectDebitResourceTypeDirectDebitAutomaticReturns DirectDebitResourceType = "direct_debit_automatic_returns"
)

// for schema
var directDebitResourceTypeEnum []interface{}

func init() {
	var res []DirectDebitResourceType
	if err := json.Unmarshal([]byte(`["direct_debits","direct_debit_admissions","direct_debit_submissions","direct_debit_submission_validations","direct_debit_reversals","direct_debit_reversal_admissions","direct_debit_returns","direct_debit_return_admissions","direct_debit_return_submissions","direct_debit_return_submission_validations","direct_debit_return_reversals","direct_debit_return_reversal_admissions","direct_debit_automatic_returns"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitResourceTypeEnum = append(directDebitResourceTypeEnum, v)
	}
}

func (m DirectDebitResourceType) validateDirectDebitResourceTypeEnum(path, location string, value DirectDebitResourceType) error {
	if err := validate.Enum(path, location, value, directDebitResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this direct debit resource type
func (m DirectDebitResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectDebitResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
