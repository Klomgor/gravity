/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.27.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ApiAPIToolPortmapOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPIToolPortmapOutput{}

// ApiAPIToolPortmapOutput struct for ApiAPIToolPortmapOutput
type ApiAPIToolPortmapOutput struct {
	Ports []ApiAPIToolPortmapOutputPort `json:"ports,omitempty"`
}

// NewApiAPIToolPortmapOutput instantiates a new ApiAPIToolPortmapOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIToolPortmapOutput() *ApiAPIToolPortmapOutput {
	this := ApiAPIToolPortmapOutput{}
	return &this
}

// NewApiAPIToolPortmapOutputWithDefaults instantiates a new ApiAPIToolPortmapOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIToolPortmapOutputWithDefaults() *ApiAPIToolPortmapOutput {
	this := ApiAPIToolPortmapOutput{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAPIToolPortmapOutput) GetPorts() []ApiAPIToolPortmapOutputPort {
	if o == nil {
		var ret []ApiAPIToolPortmapOutputPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAPIToolPortmapOutput) GetPortsOk() ([]ApiAPIToolPortmapOutputPort, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApiAPIToolPortmapOutput) HasPorts() bool {
	if o != nil && IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ApiAPIToolPortmapOutputPort and assigns it to the Ports field.
func (o *ApiAPIToolPortmapOutput) SetPorts(v []ApiAPIToolPortmapOutputPort) {
	o.Ports = v
}

func (o ApiAPIToolPortmapOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPIToolPortmapOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableApiAPIToolPortmapOutput struct {
	value *ApiAPIToolPortmapOutput
	isSet bool
}

func (v NullableApiAPIToolPortmapOutput) Get() *ApiAPIToolPortmapOutput {
	return v.value
}

func (v *NullableApiAPIToolPortmapOutput) Set(val *ApiAPIToolPortmapOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIToolPortmapOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIToolPortmapOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIToolPortmapOutput(val *ApiAPIToolPortmapOutput) *NullableApiAPIToolPortmapOutput {
	return &NullableApiAPIToolPortmapOutput{value: val, isSet: true}
}

func (v NullableApiAPIToolPortmapOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIToolPortmapOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
