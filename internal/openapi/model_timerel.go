/*
ETSI ISG CIM / NGSI-LD API

This OAS file describes the NGSI-LD API defined by the ETSI ISG CIM group. This Cross-domain Context Information Management API allows to provide, consume and subscribe to context information in multiple scenarios and involving multiple stakeholders

API version: latest
Contact: NGSI-LD@etsi.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Timerel the model 'Timerel'
type Timerel string

// List of timerel
const (
	BEFORE Timerel = "before"
	AFTER Timerel = "after"
	BETWEEN Timerel = "between"
)

// All allowed values of Timerel enum
var AllowedTimerelEnumValues = []Timerel{
	"before",
	"after",
	"between",
}

func (v *Timerel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Timerel(value)
	for _, existing := range AllowedTimerelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Timerel", value)
}

// NewTimerelFromValue returns a pointer to a valid Timerel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimerelFromValue(v string) (*Timerel, error) {
	ev := Timerel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Timerel: valid values are %v", v, AllowedTimerelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Timerel) IsValid() bool {
	for _, existing := range AllowedTimerelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to timerel value
func (v Timerel) Ptr() *Timerel {
	return &v
}

type NullableTimerel struct {
	value *Timerel
	isSet bool
}

func (v NullableTimerel) Get() *Timerel {
	return v.value
}

func (v *NullableTimerel) Set(val *Timerel) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerel) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerel(val *Timerel) *NullableTimerel {
	return &NullableTimerel{value: val, isSet: true}
}

func (v NullableTimerel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
