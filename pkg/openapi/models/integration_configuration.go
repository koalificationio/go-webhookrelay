// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntegrationConfiguration Integration Configuration
//
// swagger:model IntegrationConfiguration
type IntegrationConfiguration struct {

	// buckets
	Buckets []*Bucket `json:"buckets"`

	// created
	// Read Only: true
	Created int64 `json:"created,omitempty"`

	// Free form description field
	Description string `json:"description,omitempty"`

	// Disable/enable notification configuration. Consistently failing configurations will be disabled
	Disabled bool `json:"disabled,omitempty"`

	// Event types, currently supported are: 'forwarding.function.error' for failed functions and 'forwarding.delivery.error' for failed webhook delivery. Learn more here https://webhookrelay.com/v1/guide/integrations
	Events []string `json:"events"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Status description
	Message string `json:"message,omitempty"`

	// Plugin to use, more info about webhook plugin can be found in the docs here: https://webhookrelay.com/v1/guide/integrations#Webhook
	// Enum: [slack webhook]
	Plugin string `json:"plugin,omitempty"`

	// Plugin status
	// Read Only: true
	Status string `json:"status,omitempty"`

	// updated
	// Read Only: true
	Updated int64 `json:"updated,omitempty"`
}

// Validate validates this integration configuration
func (m *IntegrationConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuckets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationConfiguration) validateBuckets(formats strfmt.Registry) error {

	if swag.IsZero(m.Buckets) { // not required
		return nil
	}

	for i := 0; i < len(m.Buckets); i++ {
		if swag.IsZero(m.Buckets[i]) { // not required
			continue
		}

		if m.Buckets[i] != nil {
			if err := m.Buckets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var integrationConfigurationEventsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["forwarding.function.error","forwarding.delivery.error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationConfigurationEventsItemsEnum = append(integrationConfigurationEventsItemsEnum, v)
	}
}

func (m *IntegrationConfiguration) validateEventsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, integrationConfigurationEventsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *IntegrationConfiguration) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {

		// value enum
		if err := m.validateEventsItemsEnum("events"+"."+strconv.Itoa(i), "body", m.Events[i]); err != nil {
			return err
		}

	}

	return nil
}

var integrationConfigurationTypePluginPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["slack","webhook"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationConfigurationTypePluginPropEnum = append(integrationConfigurationTypePluginPropEnum, v)
	}
}

const (

	// IntegrationConfigurationPluginSlack captures enum value "slack"
	IntegrationConfigurationPluginSlack string = "slack"

	// IntegrationConfigurationPluginWebhook captures enum value "webhook"
	IntegrationConfigurationPluginWebhook string = "webhook"
)

// prop value enum
func (m *IntegrationConfiguration) validatePluginEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, integrationConfigurationTypePluginPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IntegrationConfiguration) validatePlugin(formats strfmt.Registry) error {

	if swag.IsZero(m.Plugin) { // not required
		return nil
	}

	// value enum
	if err := m.validatePluginEnum("plugin", "body", m.Plugin); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationConfiguration) UnmarshalBinary(b []byte) error {
	var res IntegrationConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
