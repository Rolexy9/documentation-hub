// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v5/utils"
)

// DryRunAccount dry run account
// swagger:model DryRunAccount
type DryRunAccount struct {

	// amount
	// Required: true
	Amount utils.BigInt `json:"amount"`

	// pub key
	// Required: true
	PubKey *string `json:"pub_key"`
}

// Validate validates this dry run account
func (m *DryRunAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DryRunAccount) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *DryRunAccount) validatePubKey(formats strfmt.Registry) error {

	if err := validate.Required("pub_key", "body", m.PubKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DryRunAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DryRunAccount) UnmarshalBinary(b []byte) error {
	var res DryRunAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}