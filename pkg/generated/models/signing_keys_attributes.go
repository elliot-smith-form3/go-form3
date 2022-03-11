// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SigningKeysAttributes signing keys attributes
// swagger:model SigningKeysAttributes
type SigningKeysAttributes struct {

	// public key
	// Required: true
	PublicKey *string `json:"public_key"`

	// certificate
	Certificate *string `json:"certificate,omitempty"`

	// certificate csr
	CertificateCsr *string `json:"certificate-csr,omitempty"`

	// status
	// Enum: [pending_activation active expired revoked]
	Status *string `json:"status,omitempty"`

	// expiration datetime
	// Format: date-time
	ExpirationDatetime *strfmt.DateTime `json:"expiration_datetime,omitempty"`

	// revocation datetime
	// Format: date-time
	RevocationDatetime *strfmt.DateTime `json:"revocation_datetime,omitempty"`
}

func SigningKeysAttributesWithDefaults(defaults client.Defaults) *SigningKeysAttributes {
	return &SigningKeysAttributes{

		PublicKey: defaults.GetStringPtr("SigningKeysAttributes", "public_key"),

		Certificate: defaults.GetStringPtr("SigningKeysAttributes", "certificate"),

		CertificateCsr: defaults.GetStringPtr("SigningKeysAttributes", "certificate-csr"),

		Status: defaults.GetStringPtr("SigningKeysAttributes", "status"),

		ExpirationDatetime: defaults.GetStrfmtDateTimePtr("SigningKeysAttributes", "expiration_datetime"),

		RevocationDatetime: defaults.GetStrfmtDateTimePtr("SigningKeysAttributes", "revocation_datetime"),
	}
}

func (m *SigningKeysAttributes) WithPublicKey(publicKey string) *SigningKeysAttributes {

	m.PublicKey = &publicKey

	return m
}

func (m *SigningKeysAttributes) WithoutPublicKey() *SigningKeysAttributes {
	m.PublicKey = nil
	return m
}

func (m *SigningKeysAttributes) WithCertificate(certificate string) *SigningKeysAttributes {

	m.Certificate = &certificate

	return m
}

func (m *SigningKeysAttributes) WithoutCertificate() *SigningKeysAttributes {
	m.Certificate = nil
	return m
}

func (m *SigningKeysAttributes) WithCertificateCsr(certificateCsr string) *SigningKeysAttributes {

	m.CertificateCsr = &certificateCsr

	return m
}

func (m *SigningKeysAttributes) WithoutCertificateCsr() *SigningKeysAttributes {
	m.CertificateCsr = nil
	return m
}

func (m *SigningKeysAttributes) WithStatus(status string) *SigningKeysAttributes {

	m.Status = &status

	return m
}

func (m *SigningKeysAttributes) WithoutStatus() *SigningKeysAttributes {
	m.Status = nil
	return m
}

func (m *SigningKeysAttributes) WithExpirationDatetime(expirationDatetime strfmt.DateTime) *SigningKeysAttributes {

	m.ExpirationDatetime = &expirationDatetime

	return m
}

func (m *SigningKeysAttributes) WithoutExpirationDatetime() *SigningKeysAttributes {
	m.ExpirationDatetime = nil
	return m
}

func (m *SigningKeysAttributes) WithRevocationDatetime(revocationDatetime strfmt.DateTime) *SigningKeysAttributes {

	m.RevocationDatetime = &revocationDatetime

	return m
}

func (m *SigningKeysAttributes) WithoutRevocationDatetime() *SigningKeysAttributes {
	m.RevocationDatetime = nil
	return m
}

// Validate validates this signing keys attributes
func (m *SigningKeysAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevocationDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigningKeysAttributes) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("public_key", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

var signingKeysAttributesTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending_activation","active","expired","revoked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		signingKeysAttributesTypeStatusPropEnum = append(signingKeysAttributesTypeStatusPropEnum, v)
	}
}

const (

	// SigningKeysAttributesStatusPendingActivation captures enum value "pending_activation"
	SigningKeysAttributesStatusPendingActivation string = "pending_activation"

	// SigningKeysAttributesStatusActive captures enum value "active"
	SigningKeysAttributesStatusActive string = "active"

	// SigningKeysAttributesStatusExpired captures enum value "expired"
	SigningKeysAttributesStatusExpired string = "expired"

	// SigningKeysAttributesStatusRevoked captures enum value "revoked"
	SigningKeysAttributesStatusRevoked string = "revoked"
)

// prop value enum
func (m *SigningKeysAttributes) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, signingKeysAttributesTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SigningKeysAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *SigningKeysAttributes) validateExpirationDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration_datetime", "body", "date-time", m.ExpirationDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SigningKeysAttributes) validateRevocationDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.RevocationDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("revocation_datetime", "body", "date-time", m.RevocationDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SigningKeysAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigningKeysAttributes) UnmarshalBinary(b []byte) error {
	var res SigningKeysAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SigningKeysAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
