/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.13.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// TypesAPIMetricsRole the model 'TypesAPIMetricsRole'
type TypesAPIMetricsRole string

// List of TypesAPIMetricsRole
const (
	TYPESAPIMETRICSROLE_SYSTEM TypesAPIMetricsRole = "system"
	TYPESAPIMETRICSROLE_DNS    TypesAPIMetricsRole = "dns"
)

// All allowed values of TypesAPIMetricsRole enum
var AllowedTypesAPIMetricsRoleEnumValues = []TypesAPIMetricsRole{
	"system",
	"dns",
}

func (v *TypesAPIMetricsRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesAPIMetricsRole(value)
	for _, existing := range AllowedTypesAPIMetricsRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesAPIMetricsRole", value)
}

// NewTypesAPIMetricsRoleFromValue returns a pointer to a valid TypesAPIMetricsRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesAPIMetricsRoleFromValue(v string) (*TypesAPIMetricsRole, error) {
	ev := TypesAPIMetricsRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesAPIMetricsRole: valid values are %v", v, AllowedTypesAPIMetricsRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesAPIMetricsRole) IsValid() bool {
	for _, existing := range AllowedTypesAPIMetricsRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TypesAPIMetricsRole value
func (v TypesAPIMetricsRole) Ptr() *TypesAPIMetricsRole {
	return &v
}

type NullableTypesAPIMetricsRole struct {
	value *TypesAPIMetricsRole
	isSet bool
}

func (v NullableTypesAPIMetricsRole) Get() *TypesAPIMetricsRole {
	return v.value
}

func (v *NullableTypesAPIMetricsRole) Set(val *TypesAPIMetricsRole) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesAPIMetricsRole) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesAPIMetricsRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesAPIMetricsRole(val *TypesAPIMetricsRole) *NullableTypesAPIMetricsRole {
	return &NullableTypesAPIMetricsRole{value: val, isSet: true}
}

func (v NullableTypesAPIMetricsRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesAPIMetricsRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
