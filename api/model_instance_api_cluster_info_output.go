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

// checks if the InstanceAPIClusterInfoOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceAPIClusterInfoOutput{}

// InstanceAPIClusterInfoOutput struct for InstanceAPIClusterInfoOutput
type InstanceAPIClusterInfoOutput struct {
	ClusterVersion      string                 `json:"clusterVersion"`
	ClusterVersionShort string                 `json:"clusterVersionShort"`
	Instances           []InstanceInstanceInfo `json:"instances"`
}

// NewInstanceAPIClusterInfoOutput instantiates a new InstanceAPIClusterInfoOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceAPIClusterInfoOutput(clusterVersion string, clusterVersionShort string, instances []InstanceInstanceInfo) *InstanceAPIClusterInfoOutput {
	this := InstanceAPIClusterInfoOutput{}
	this.ClusterVersion = clusterVersion
	this.ClusterVersionShort = clusterVersionShort
	this.Instances = instances
	return &this
}

// NewInstanceAPIClusterInfoOutputWithDefaults instantiates a new InstanceAPIClusterInfoOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceAPIClusterInfoOutputWithDefaults() *InstanceAPIClusterInfoOutput {
	this := InstanceAPIClusterInfoOutput{}
	return &this
}

// GetClusterVersion returns the ClusterVersion field value
func (o *InstanceAPIClusterInfoOutput) GetClusterVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterVersion
}

// GetClusterVersionOk returns a tuple with the ClusterVersion field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIClusterInfoOutput) GetClusterVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterVersion, true
}

// SetClusterVersion sets field value
func (o *InstanceAPIClusterInfoOutput) SetClusterVersion(v string) {
	o.ClusterVersion = v
}

// GetClusterVersionShort returns the ClusterVersionShort field value
func (o *InstanceAPIClusterInfoOutput) GetClusterVersionShort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterVersionShort
}

// GetClusterVersionShortOk returns a tuple with the ClusterVersionShort field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIClusterInfoOutput) GetClusterVersionShortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterVersionShort, true
}

// SetClusterVersionShort sets field value
func (o *InstanceAPIClusterInfoOutput) SetClusterVersionShort(v string) {
	o.ClusterVersionShort = v
}

// GetInstances returns the Instances field value
// If the value is explicit nil, the zero value for []InstanceInstanceInfo will be returned
func (o *InstanceAPIClusterInfoOutput) GetInstances() []InstanceInstanceInfo {
	if o == nil {
		var ret []InstanceInstanceInfo
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAPIClusterInfoOutput) GetInstancesOk() ([]InstanceInstanceInfo, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *InstanceAPIClusterInfoOutput) SetInstances(v []InstanceInstanceInfo) {
	o.Instances = v
}

func (o InstanceAPIClusterInfoOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceAPIClusterInfoOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterVersion"] = o.ClusterVersion
	toSerialize["clusterVersionShort"] = o.ClusterVersionShort
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	return toSerialize, nil
}

type NullableInstanceAPIClusterInfoOutput struct {
	value *InstanceAPIClusterInfoOutput
	isSet bool
}

func (v NullableInstanceAPIClusterInfoOutput) Get() *InstanceAPIClusterInfoOutput {
	return v.value
}

func (v *NullableInstanceAPIClusterInfoOutput) Set(val *InstanceAPIClusterInfoOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceAPIClusterInfoOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceAPIClusterInfoOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceAPIClusterInfoOutput(val *InstanceAPIClusterInfoOutput) *NullableInstanceAPIClusterInfoOutput {
	return &NullableInstanceAPIClusterInfoOutput{value: val, isSet: true}
}

func (v NullableInstanceAPIClusterInfoOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceAPIClusterInfoOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
