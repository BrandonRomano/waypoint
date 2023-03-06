// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpWaypointInstanceType Instances are one of a these types.
//
//   - LONG_RUNNING: The "traditional" instance type, a process that is running
//
// constantly for a long period of time.
//   - ON_DEMAND: An instance that was launched in response to a request and will
//
// disappear quickly.
//   - VIRTUAL: An instance that is not actually running any code, but registers
//
// itself as an instance for the purposes of interacting with the
// exec and logs functionality
//
// swagger:model hashicorp.waypoint.Instance.Type
type HashicorpWaypointInstanceType string

func NewHashicorpWaypointInstanceType(value HashicorpWaypointInstanceType) *HashicorpWaypointInstanceType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpWaypointInstanceType.
func (m HashicorpWaypointInstanceType) Pointer() *HashicorpWaypointInstanceType {
	return &m
}

const (

	// HashicorpWaypointInstanceTypeLONGRUNNING captures enum value "LONG_RUNNING"
	HashicorpWaypointInstanceTypeLONGRUNNING HashicorpWaypointInstanceType = "LONG_RUNNING"

	// HashicorpWaypointInstanceTypeONDEMAND captures enum value "ON_DEMAND"
	HashicorpWaypointInstanceTypeONDEMAND HashicorpWaypointInstanceType = "ON_DEMAND"

	// HashicorpWaypointInstanceTypeVIRTUAL captures enum value "VIRTUAL"
	HashicorpWaypointInstanceTypeVIRTUAL HashicorpWaypointInstanceType = "VIRTUAL"
)

// for schema
var hashicorpWaypointInstanceTypeEnum []interface{}

func init() {
	var res []HashicorpWaypointInstanceType
	if err := json.Unmarshal([]byte(`["LONG_RUNNING","ON_DEMAND","VIRTUAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpWaypointInstanceTypeEnum = append(hashicorpWaypointInstanceTypeEnum, v)
	}
}

func (m HashicorpWaypointInstanceType) validateHashicorpWaypointInstanceTypeEnum(path, location string, value HashicorpWaypointInstanceType) error {
	if err := validate.EnumCase(path, location, value, hashicorpWaypointInstanceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp waypoint instance type
func (m HashicorpWaypointInstanceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpWaypointInstanceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp waypoint instance type based on context it is used
func (m HashicorpWaypointInstanceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}