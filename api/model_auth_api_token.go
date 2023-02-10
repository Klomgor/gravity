/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthAPIToken struct for AuthAPIToken
type AuthAPIToken struct {
	Key      string `json:"key"`
	Username string `json:"username"`
}

// NewAuthAPIToken instantiates a new AuthAPIToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPIToken(key string, username string) *AuthAPIToken {
	this := AuthAPIToken{}
	this.Key = key
	this.Username = username
	return &this
}

// NewAuthAPITokenWithDefaults instantiates a new AuthAPIToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAPITokenWithDefaults() *AuthAPIToken {
	this := AuthAPIToken{}
	return &this
}

// GetKey returns the Key field value
func (o *AuthAPIToken) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AuthAPIToken) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AuthAPIToken) SetKey(v string) {
	o.Key = v
}

// GetUsername returns the Username field value
func (o *AuthAPIToken) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthAPIToken) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthAPIToken) SetUsername(v string) {
	o.Username = v
}

func (o AuthAPIToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAPIToken struct {
	value *AuthAPIToken
	isSet bool
}

func (v NullableAuthAPIToken) Get() *AuthAPIToken {
	return v.value
}

func (v *NullableAuthAPIToken) Set(val *AuthAPIToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAPIToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAPIToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAPIToken(val *AuthAPIToken) *NullableAuthAPIToken {
	return &NullableAuthAPIToken{value: val, isSet: true}
}

func (v NullableAuthAPIToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAPIToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
