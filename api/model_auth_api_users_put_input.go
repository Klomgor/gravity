/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthAPIUsersPutInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthAPIUsersPutInput{}

// AuthAPIUsersPutInput struct for AuthAPIUsersPutInput
type AuthAPIUsersPutInput struct {
	Password    string           `json:"password"`
	Permissions []AuthPermission `json:"permissions"`
}

// NewAuthAPIUsersPutInput instantiates a new AuthAPIUsersPutInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAPIUsersPutInput(password string, permissions []AuthPermission) *AuthAPIUsersPutInput {
	this := AuthAPIUsersPutInput{}
	this.Password = password
	this.Permissions = permissions
	return &this
}

// NewAuthAPIUsersPutInputWithDefaults instantiates a new AuthAPIUsersPutInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAPIUsersPutInputWithDefaults() *AuthAPIUsersPutInput {
	this := AuthAPIUsersPutInput{}
	return &this
}

// GetPassword returns the Password field value
func (o *AuthAPIUsersPutInput) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthAPIUsersPutInput) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthAPIUsersPutInput) SetPassword(v string) {
	o.Password = v
}

// GetPermissions returns the Permissions field value
// If the value is explicit nil, the zero value for []AuthPermission will be returned
func (o *AuthAPIUsersPutInput) GetPermissions() []AuthPermission {
	if o == nil {
		var ret []AuthPermission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthAPIUsersPutInput) GetPermissionsOk() ([]AuthPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *AuthAPIUsersPutInput) SetPermissions(v []AuthPermission) {
	o.Permissions = v
}

func (o AuthAPIUsersPutInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthAPIUsersPutInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableAuthAPIUsersPutInput struct {
	value *AuthAPIUsersPutInput
	isSet bool
}

func (v NullableAuthAPIUsersPutInput) Get() *AuthAPIUsersPutInput {
	return v.value
}

func (v *NullableAuthAPIUsersPutInput) Set(val *AuthAPIUsersPutInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAPIUsersPutInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAPIUsersPutInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAPIUsersPutInput(val *AuthAPIUsersPutInput) *NullableAuthAPIUsersPutInput {
	return &NullableAuthAPIUsersPutInput{value: val, isSet: true}
}

func (v NullableAuthAPIUsersPutInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAPIUsersPutInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
