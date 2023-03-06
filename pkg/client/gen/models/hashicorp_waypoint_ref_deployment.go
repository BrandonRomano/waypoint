// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointRefDeployment hashicorp waypoint ref deployment
//
// swagger:model hashicorp.waypoint.Ref.Deployment
type HashicorpWaypointRefDeployment struct {

	// latest
	Latest bool `json:"latest,omitempty"`

	// sequence
	Sequence string `json:"sequence,omitempty"`
}

// Validate validates this hashicorp waypoint ref deployment
func (m *HashicorpWaypointRefDeployment) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint ref deployment based on context it is used
func (m *HashicorpWaypointRefDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointRefDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointRefDeployment) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointRefDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}