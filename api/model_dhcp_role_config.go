/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.13.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DhcpRoleConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpRoleConfig{}

// DhcpRoleConfig struct for DhcpRoleConfig
type DhcpRoleConfig struct {
	LeaseNegotiateTimeout *int32 `json:"leaseNegotiateTimeout,omitempty"`
	Port                  *int32 `json:"port,omitempty"`
}

// NewDhcpRoleConfig instantiates a new DhcpRoleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpRoleConfig() *DhcpRoleConfig {
	this := DhcpRoleConfig{}
	return &this
}

// NewDhcpRoleConfigWithDefaults instantiates a new DhcpRoleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpRoleConfigWithDefaults() *DhcpRoleConfig {
	this := DhcpRoleConfig{}
	return &this
}

// GetLeaseNegotiateTimeout returns the LeaseNegotiateTimeout field value if set, zero value otherwise.
func (o *DhcpRoleConfig) GetLeaseNegotiateTimeout() int32 {
	if o == nil || IsNil(o.LeaseNegotiateTimeout) {
		var ret int32
		return ret
	}
	return *o.LeaseNegotiateTimeout
}

// GetLeaseNegotiateTimeoutOk returns a tuple with the LeaseNegotiateTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRoleConfig) GetLeaseNegotiateTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.LeaseNegotiateTimeout) {
		return nil, false
	}
	return o.LeaseNegotiateTimeout, true
}

// HasLeaseNegotiateTimeout returns a boolean if a field has been set.
func (o *DhcpRoleConfig) HasLeaseNegotiateTimeout() bool {
	if o != nil && !IsNil(o.LeaseNegotiateTimeout) {
		return true
	}

	return false
}

// SetLeaseNegotiateTimeout gets a reference to the given int32 and assigns it to the LeaseNegotiateTimeout field.
func (o *DhcpRoleConfig) SetLeaseNegotiateTimeout(v int32) {
	o.LeaseNegotiateTimeout = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DhcpRoleConfig) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRoleConfig) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DhcpRoleConfig) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *DhcpRoleConfig) SetPort(v int32) {
	o.Port = &v
}

func (o DhcpRoleConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpRoleConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LeaseNegotiateTimeout) {
		toSerialize["leaseNegotiateTimeout"] = o.LeaseNegotiateTimeout
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableDhcpRoleConfig struct {
	value *DhcpRoleConfig
	isSet bool
}

func (v NullableDhcpRoleConfig) Get() *DhcpRoleConfig {
	return v.value
}

func (v *NullableDhcpRoleConfig) Set(val *DhcpRoleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpRoleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpRoleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpRoleConfig(val *DhcpRoleConfig) *NullableDhcpRoleConfig {
	return &NullableDhcpRoleConfig{value: val, isSet: true}
}

func (v NullableDhcpRoleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpRoleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
