/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.26.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DhcpAPIScopesImportOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpAPIScopesImportOutput{}

// DhcpAPIScopesImportOutput struct for DhcpAPIScopesImportOutput
type DhcpAPIScopesImportOutput struct {
	Successful *bool `json:"successful,omitempty"`
}

// NewDhcpAPIScopesImportOutput instantiates a new DhcpAPIScopesImportOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpAPIScopesImportOutput() *DhcpAPIScopesImportOutput {
	this := DhcpAPIScopesImportOutput{}
	return &this
}

// NewDhcpAPIScopesImportOutputWithDefaults instantiates a new DhcpAPIScopesImportOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpAPIScopesImportOutputWithDefaults() *DhcpAPIScopesImportOutput {
	this := DhcpAPIScopesImportOutput{}
	return &this
}

// GetSuccessful returns the Successful field value if set, zero value otherwise.
func (o *DhcpAPIScopesImportOutput) GetSuccessful() bool {
	if o == nil || IsNil(o.Successful) {
		var ret bool
		return ret
	}
	return *o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpAPIScopesImportOutput) GetSuccessfulOk() (*bool, bool) {
	if o == nil || IsNil(o.Successful) {
		return nil, false
	}
	return o.Successful, true
}

// HasSuccessful returns a boolean if a field has been set.
func (o *DhcpAPIScopesImportOutput) HasSuccessful() bool {
	if o != nil && !IsNil(o.Successful) {
		return true
	}

	return false
}

// SetSuccessful gets a reference to the given bool and assigns it to the Successful field.
func (o *DhcpAPIScopesImportOutput) SetSuccessful(v bool) {
	o.Successful = &v
}

func (o DhcpAPIScopesImportOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpAPIScopesImportOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Successful) {
		toSerialize["successful"] = o.Successful
	}
	return toSerialize, nil
}

type NullableDhcpAPIScopesImportOutput struct {
	value *DhcpAPIScopesImportOutput
	isSet bool
}

func (v NullableDhcpAPIScopesImportOutput) Get() *DhcpAPIScopesImportOutput {
	return v.value
}

func (v *NullableDhcpAPIScopesImportOutput) Set(val *DhcpAPIScopesImportOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpAPIScopesImportOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpAPIScopesImportOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpAPIScopesImportOutput(val *DhcpAPIScopesImportOutput) *NullableDhcpAPIScopesImportOutput {
	return &NullableDhcpAPIScopesImportOutput{value: val, isSet: true}
}

func (v NullableDhcpAPIScopesImportOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpAPIScopesImportOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
