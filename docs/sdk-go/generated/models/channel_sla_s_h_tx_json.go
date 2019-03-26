// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChannelSLASHTxJSON channel slash tx JSON
// swagger:model ChannelSlashTxJSON
type ChannelSLASHTxJSON struct {
	versionField *uint64

	ChannelSLASHTx
}

// Type gets the type of this subtype
func (m *ChannelSLASHTxJSON) Type() string {
	return "ChannelSlashTxJSON"
}

// SetType sets the type of this subtype
func (m *ChannelSLASHTxJSON) SetType(val string) {

}

// Version gets the version of this subtype
func (m *ChannelSLASHTxJSON) Version() *uint64 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *ChannelSLASHTxJSON) SetVersion(val *uint64) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ChannelSLASHTxJSON) UnmarshalJSON(raw []byte) error {
	var data struct {
		ChannelSLASHTx
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`

		Version *uint64 `json:"version"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ChannelSLASHTxJSON

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.versionField = base.Version

	result.ChannelSLASHTx = data.ChannelSLASHTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ChannelSLASHTxJSON) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		ChannelSLASHTx
	}{

		ChannelSLASHTx: m.ChannelSLASHTx,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`

		Version *uint64 `json:"version"`
	}{

		Type: m.Type(),

		Version: m.Version(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this channel slash tx JSON
func (m *ChannelSLASHTxJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with ChannelSLASHTx
	if err := m.ChannelSLASHTx.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelSLASHTxJSON) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelSLASHTxJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelSLASHTxJSON) UnmarshalBinary(b []byte) error {
	var res ChannelSLASHTxJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
