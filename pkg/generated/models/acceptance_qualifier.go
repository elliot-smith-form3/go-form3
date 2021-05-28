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

// AcceptanceQualifier All accepted payments will receive the matching qualifier code
// swagger:model AcceptanceQualifier
type AcceptanceQualifier string

const (

	// AcceptanceQualifierSomeOtherTime captures enum value "some_other_time"
	AcceptanceQualifierSomeOtherTime AcceptanceQualifier = "some_other_time"

	// AcceptanceQualifierSameDay captures enum value "same_day"
	AcceptanceQualifierSameDay AcceptanceQualifier = "same_day"

	// AcceptanceQualifierNextCalendarDay captures enum value "next_calendar_day"
	AcceptanceQualifierNextCalendarDay AcceptanceQualifier = "next_calendar_day"

	// AcceptanceQualifierNextWorkingDay captures enum value "next_working_day"
	AcceptanceQualifierNextWorkingDay AcceptanceQualifier = "next_working_day"

	// AcceptanceQualifierAfterNextWorkingDay captures enum value "after_next_working_day"
	AcceptanceQualifierAfterNextWorkingDay AcceptanceQualifier = "after_next_working_day"

	// AcceptanceQualifierNone captures enum value "none"
	AcceptanceQualifierNone AcceptanceQualifier = "none"
)

// for schema
var acceptanceQualifierEnum []interface{}

func init() {
	var res []AcceptanceQualifier
	if err := json.Unmarshal([]byte(`["some_other_time","same_day","next_calendar_day","next_working_day","after_next_working_day","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		acceptanceQualifierEnum = append(acceptanceQualifierEnum, v)
	}
}

func (m AcceptanceQualifier) validateAcceptanceQualifierEnum(path, location string, value AcceptanceQualifier) error {
	if err := validate.Enum(path, location, value, acceptanceQualifierEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this acceptance qualifier
func (m AcceptanceQualifier) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAcceptanceQualifierEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcceptanceQualifier) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
