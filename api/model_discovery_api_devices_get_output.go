/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.18.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DiscoveryAPIDevicesGetOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryAPIDevicesGetOutput{}

// DiscoveryAPIDevicesGetOutput struct for DiscoveryAPIDevicesGetOutput
type DiscoveryAPIDevicesGetOutput struct {
	Devices []DiscoveryAPIDevice `json:"devices,omitempty"`
}

// NewDiscoveryAPIDevicesGetOutput instantiates a new DiscoveryAPIDevicesGetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryAPIDevicesGetOutput() *DiscoveryAPIDevicesGetOutput {
	this := DiscoveryAPIDevicesGetOutput{}
	return &this
}

// NewDiscoveryAPIDevicesGetOutputWithDefaults instantiates a new DiscoveryAPIDevicesGetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryAPIDevicesGetOutputWithDefaults() *DiscoveryAPIDevicesGetOutput {
	this := DiscoveryAPIDevicesGetOutput{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryAPIDevicesGetOutput) GetDevices() []DiscoveryAPIDevice {
	if o == nil {
		var ret []DiscoveryAPIDevice
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryAPIDevicesGetOutput) GetDevicesOk() ([]DiscoveryAPIDevice, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *DiscoveryAPIDevicesGetOutput) HasDevices() bool {
	if o != nil && IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []DiscoveryAPIDevice and assigns it to the Devices field.
func (o *DiscoveryAPIDevicesGetOutput) SetDevices(v []DiscoveryAPIDevice) {
	o.Devices = v
}

func (o DiscoveryAPIDevicesGetOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryAPIDevicesGetOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Devices != nil {
		toSerialize["devices"] = o.Devices
	}
	return toSerialize, nil
}

type NullableDiscoveryAPIDevicesGetOutput struct {
	value *DiscoveryAPIDevicesGetOutput
	isSet bool
}

func (v NullableDiscoveryAPIDevicesGetOutput) Get() *DiscoveryAPIDevicesGetOutput {
	return v.value
}

func (v *NullableDiscoveryAPIDevicesGetOutput) Set(val *DiscoveryAPIDevicesGetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryAPIDevicesGetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryAPIDevicesGetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryAPIDevicesGetOutput(val *DiscoveryAPIDevicesGetOutput) *NullableDiscoveryAPIDevicesGetOutput {
	return &NullableDiscoveryAPIDevicesGetOutput{value: val, isSet: true}
}

func (v NullableDiscoveryAPIDevicesGetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryAPIDevicesGetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
