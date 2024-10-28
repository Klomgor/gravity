/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.13.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ApiAPIToolTracerouteInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPIToolTracerouteInput{}

// ApiAPIToolTracerouteInput struct for ApiAPIToolTracerouteInput
type ApiAPIToolTracerouteInput struct {
	Host *string `json:"host,omitempty"`
}

// NewApiAPIToolTracerouteInput instantiates a new ApiAPIToolTracerouteInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIToolTracerouteInput() *ApiAPIToolTracerouteInput {
	this := ApiAPIToolTracerouteInput{}
	return &this
}

// NewApiAPIToolTracerouteInputWithDefaults instantiates a new ApiAPIToolTracerouteInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIToolTracerouteInputWithDefaults() *ApiAPIToolTracerouteInput {
	this := ApiAPIToolTracerouteInput{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ApiAPIToolTracerouteInput) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPIToolTracerouteInput) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ApiAPIToolTracerouteInput) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ApiAPIToolTracerouteInput) SetHost(v string) {
	o.Host = &v
}

func (o ApiAPIToolTracerouteInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPIToolTracerouteInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	return toSerialize, nil
}

type NullableApiAPIToolTracerouteInput struct {
	value *ApiAPIToolTracerouteInput
	isSet bool
}

func (v NullableApiAPIToolTracerouteInput) Get() *ApiAPIToolTracerouteInput {
	return v.value
}

func (v *NullableApiAPIToolTracerouteInput) Set(val *ApiAPIToolTracerouteInput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIToolTracerouteInput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIToolTracerouteInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIToolTracerouteInput(val *ApiAPIToolTracerouteInput) *NullableApiAPIToolTracerouteInput {
	return &NullableApiAPIToolTracerouteInput{value: val, isSet: true}
}

func (v NullableApiAPIToolTracerouteInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIToolTracerouteInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
