// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// QueryResponseCompensation query response compensation
// swagger:model QueryResponseCompensation
type QueryResponseCompensation struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode string `json:"account_number_code,omitempty"`

	// amount
	Amount string `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`
}

func QueryResponseCompensationWithDefaults(defaults client.Defaults) *QueryResponseCompensation {
	return &QueryResponseCompensation{

		AccountNumber: defaults.GetString("QueryResponseCompensation", "account_number"),

		AccountNumberCode: defaults.GetString("QueryResponseCompensation", "account_number_code"),

		Amount: defaults.GetString("QueryResponseCompensation", "amount"),

		Currency: defaults.GetString("QueryResponseCompensation", "currency"),
	}
}

func (m *QueryResponseCompensation) WithAccountNumber(accountNumber string) *QueryResponseCompensation {

	m.AccountNumber = accountNumber

	return m
}

func (m *QueryResponseCompensation) WithAccountNumberCode(accountNumberCode string) *QueryResponseCompensation {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *QueryResponseCompensation) WithAmount(amount string) *QueryResponseCompensation {

	m.Amount = amount

	return m
}

func (m *QueryResponseCompensation) WithCurrency(currency string) *QueryResponseCompensation {

	m.Currency = currency

	return m
}

// Validate validates this query response compensation
func (m *QueryResponseCompensation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseCompensation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseCompensation) UnmarshalBinary(b []byte) error {
	var res QueryResponseCompensation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseCompensation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
