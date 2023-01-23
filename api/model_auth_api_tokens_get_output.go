/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthAPITokensGetOutput struct for AuthAPITokensGetOutput
type AuthAPITokensGetOutput struct {
	Tokens []AuthAPIToken `json:"tokens"`
}

// NewAuthAPITokensGetOutput instantiates a new AuthAPITokensGetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPITokensGetOutput(tokens []AuthAPIToken) *AuthAPITokensGetOutput {
	this := AuthAPITokensGetOutput{}
	this.Tokens = tokens
	return &this
}

// NewAuthAPITokensGetOutputWithDefaults instantiates a new AuthAPITokensGetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAPITokensGetOutputWithDefaults() *AuthAPITokensGetOutput {
	this := AuthAPITokensGetOutput{}
	return &this
}

// GetTokens returns the Tokens field value
// If the value is explicit nil, the zero value for []AuthAPIToken will be returned
func (o *AuthAPITokensGetOutput) GetTokens() []AuthAPIToken {
	if o == nil {
		var ret []AuthAPIToken
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthAPITokensGetOutput) GetTokensOk() ([]AuthAPIToken, bool) {
	if o == nil || o.Tokens == nil {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *AuthAPITokensGetOutput) SetTokens(v []AuthAPIToken) {
	o.Tokens = v
}

func (o AuthAPITokensGetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tokens != nil {
		toSerialize["tokens"] = o.Tokens
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAPITokensGetOutput struct {
	value *AuthAPITokensGetOutput
	isSet bool
}

func (v NullableAuthAPITokensGetOutput) Get() *AuthAPITokensGetOutput {
	return v.value
}

func (v *NullableAuthAPITokensGetOutput) Set(val *AuthAPITokensGetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAPITokensGetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAPITokensGetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAPITokensGetOutput(val *AuthAPITokensGetOutput) *NullableAuthAPITokensGetOutput {
	return &NullableAuthAPITokensGetOutput{value: val, isSet: true}
}

func (v NullableAuthAPITokensGetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAPITokensGetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
