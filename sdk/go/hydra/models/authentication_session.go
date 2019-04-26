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

// AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession AuthenticationSession authentication session
// swagger:model AuthenticationSession
type AuthenticationSession struct {

	// authenticated at
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	AuthenticatedAt strfmt.DateTime `json:"AuthenticatedAt,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// subject
	Subject string `json:"Subject,omitempty"`
}

// Validate validates this authentication session
func (m *AuthenticationSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticationSession) validateAuthenticatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("AuthenticatedAt", "body", "date-time", m.AuthenticatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationSession) UnmarshalBinary(b []byte) error {
	var res AuthenticationSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
