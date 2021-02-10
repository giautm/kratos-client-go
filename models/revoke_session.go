// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession RevokeSession revoke session
//
// swagger:model revokeSession
type RevokeSession struct {

	// The Session Token
	//
	// Invalidate this session token.
	// Required: true
	SessionToken *string `json:"session_token"`
}

// Validate validates this revoke session
func (m *RevokeSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSessionToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RevokeSession) validateSessionToken(formats strfmt.Registry) error {

	if err := validate.Required("session_token", "body", m.SessionToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this revoke session based on context it is used
func (m *RevokeSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RevokeSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RevokeSession) UnmarshalBinary(b []byte) error {
	var res RevokeSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
