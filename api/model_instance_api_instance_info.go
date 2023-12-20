/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.8.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InstanceAPIInstanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceAPIInstanceInfo{}

// InstanceAPIInstanceInfo struct for InstanceAPIInstanceInfo
type InstanceAPIInstanceInfo struct {
	BuildHash                 string                 `json:"buildHash"`
	CurrentInstanceIP         string                 `json:"currentInstanceIP"`
	CurrentInstanceIdentifier string                 `json:"currentInstanceIdentifier"`
	Dirs                      ExtconfigExtConfigDirs `json:"dirs"`
	Version                   string                 `json:"version"`
}

// NewInstanceAPIInstanceInfo instantiates a new InstanceAPIInstanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceAPIInstanceInfo(buildHash string, currentInstanceIP string, currentInstanceIdentifier string, dirs ExtconfigExtConfigDirs, version string) *InstanceAPIInstanceInfo {
	this := InstanceAPIInstanceInfo{}
	this.BuildHash = buildHash
	this.CurrentInstanceIP = currentInstanceIP
	this.CurrentInstanceIdentifier = currentInstanceIdentifier
	this.Dirs = dirs
	this.Version = version
	return &this
}

// NewInstanceAPIInstanceInfoWithDefaults instantiates a new InstanceAPIInstanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceAPIInstanceInfoWithDefaults() *InstanceAPIInstanceInfo {
	this := InstanceAPIInstanceInfo{}
	return &this
}

// GetBuildHash returns the BuildHash field value
func (o *InstanceAPIInstanceInfo) GetBuildHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildHash
}

// GetBuildHashOk returns a tuple with the BuildHash field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetBuildHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildHash, true
}

// SetBuildHash sets field value
func (o *InstanceAPIInstanceInfo) SetBuildHash(v string) {
	o.BuildHash = v
}

// GetCurrentInstanceIP returns the CurrentInstanceIP field value
func (o *InstanceAPIInstanceInfo) GetCurrentInstanceIP() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentInstanceIP
}

// GetCurrentInstanceIPOk returns a tuple with the CurrentInstanceIP field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetCurrentInstanceIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentInstanceIP, true
}

// SetCurrentInstanceIP sets field value
func (o *InstanceAPIInstanceInfo) SetCurrentInstanceIP(v string) {
	o.CurrentInstanceIP = v
}

// GetCurrentInstanceIdentifier returns the CurrentInstanceIdentifier field value
func (o *InstanceAPIInstanceInfo) GetCurrentInstanceIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentInstanceIdentifier
}

// GetCurrentInstanceIdentifierOk returns a tuple with the CurrentInstanceIdentifier field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetCurrentInstanceIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentInstanceIdentifier, true
}

// SetCurrentInstanceIdentifier sets field value
func (o *InstanceAPIInstanceInfo) SetCurrentInstanceIdentifier(v string) {
	o.CurrentInstanceIdentifier = v
}

// GetDirs returns the Dirs field value
func (o *InstanceAPIInstanceInfo) GetDirs() ExtconfigExtConfigDirs {
	if o == nil {
		var ret ExtconfigExtConfigDirs
		return ret
	}

	return o.Dirs
}

// GetDirsOk returns a tuple with the Dirs field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetDirsOk() (*ExtconfigExtConfigDirs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dirs, true
}

// SetDirs sets field value
func (o *InstanceAPIInstanceInfo) SetDirs(v ExtconfigExtConfigDirs) {
	o.Dirs = v
}

// GetVersion returns the Version field value
func (o *InstanceAPIInstanceInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *InstanceAPIInstanceInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *InstanceAPIInstanceInfo) SetVersion(v string) {
	o.Version = v
}

func (o InstanceAPIInstanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceAPIInstanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buildHash"] = o.BuildHash
	toSerialize["currentInstanceIP"] = o.CurrentInstanceIP
	toSerialize["currentInstanceIdentifier"] = o.CurrentInstanceIdentifier
	toSerialize["dirs"] = o.Dirs
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableInstanceAPIInstanceInfo struct {
	value *InstanceAPIInstanceInfo
	isSet bool
}

func (v NullableInstanceAPIInstanceInfo) Get() *InstanceAPIInstanceInfo {
	return v.value
}

func (v *NullableInstanceAPIInstanceInfo) Set(val *InstanceAPIInstanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceAPIInstanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceAPIInstanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceAPIInstanceInfo(val *InstanceAPIInstanceInfo) *NullableInstanceAPIInstanceInfo {
	return &NullableInstanceAPIInstanceInfo{value: val, isSet: true}
}

func (v NullableInstanceAPIInstanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceAPIInstanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
