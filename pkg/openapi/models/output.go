// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Output Output
//
// swagger:model Output
type Output struct {

	// bucket id
	// Read Only: true
	BucketID string `json:"bucket_id,omitempty"`

	// created at
	// Read Only: true
	CreatedAt int64 `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Destination, where to forward the webhook, such as 'http://localhost:4000'
	Destination string `json:"destination,omitempty"`

	// ID of the function that will be executed for this output
	FunctionID string `json:"function_id,omitempty"`

	// Headers to override (for example Host)
	Headers Headers `json:"headers,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// When internal output receives webhook, it forwards it to any connected agent (relay CLI, Docker agent or WebSocket connection)
	Internal bool `json:"internal,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Request matching rules, documentation can be found here https://webhookrelay.com/v1/guide/webhook-forwarding.html#Request-matching-rules
	Rules *Rules `json:"rules,omitempty"`

	// Timeout in seconds, value between 0 and 180
	Timeout int64 `json:"timeout,omitempty"`

	// Enforce TLS verification where possible (internal and public destinations)
	TLSVerification bool `json:"tls_verification,omitempty"`

	// updated at
	// Read Only: true
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// Validate validates this output
func (m *Output) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Output) validateHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	if err := m.Headers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("headers")
		}
		return err
	}

	return nil
}

func (m *Output) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	if m.Rules != nil {
		if err := m.Rules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rules")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Output) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Output) UnmarshalBinary(b []byte) error {
	var res Output
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
