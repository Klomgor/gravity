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

// checks if the ApiAPITransportEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAPITransportEntry{}

// ApiAPITransportEntry struct for ApiAPITransportEntry
type ApiAPITransportEntry struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewApiAPITransportEntry instantiates a new ApiAPITransportEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPITransportEntry() *ApiAPITransportEntry {
	this := ApiAPITransportEntry{}
	return &this
}

// NewApiAPITransportEntryWithDefaults instantiates a new ApiAPITransportEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPITransportEntryWithDefaults() *ApiAPITransportEntry {
	this := ApiAPITransportEntry{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApiAPITransportEntry) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPITransportEntry) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApiAPITransportEntry) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApiAPITransportEntry) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiAPITransportEntry) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAPITransportEntry) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiAPITransportEntry) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiAPITransportEntry) SetValue(v string) {
	o.Value = &v
}

func (o ApiAPITransportEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAPITransportEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiAPITransportEntry struct {
	value *ApiAPITransportEntry
	isSet bool
}

func (v NullableApiAPITransportEntry) Get() *ApiAPITransportEntry {
	return v.value
}

func (v *NullableApiAPITransportEntry) Set(val *ApiAPITransportEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPITransportEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPITransportEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPITransportEntry(val *ApiAPITransportEntry) *NullableApiAPITransportEntry {
	return &NullableApiAPITransportEntry{value: val, isSet: true}
}

func (v NullableApiAPITransportEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPITransportEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
