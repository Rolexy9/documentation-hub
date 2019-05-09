// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PoI po i
// swagger:model PoI
type PoI struct {

	// Proof of inclusion
	// Required: true
	Poi *string `json:"poi"`
}

// Validate validates this po i
func (m *PoI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePoi(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoI) validatePoi(formats strfmt.Registry) error {

	if err := validate.Required("poi", "body", m.Poi); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoI) UnmarshalBinary(b []byte) error {
	var res PoI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}