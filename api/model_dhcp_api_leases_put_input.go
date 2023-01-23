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

// DhcpAPILeasesPutInput struct for DhcpAPILeasesPutInput
type DhcpAPILeasesPutInput struct {
	Address          string  `json:"address"`
	AddressLeaseTime string  `json:"addressLeaseTime"`
	DnsZone          *string `json:"dnsZone,omitempty"`
	Hostname         string  `json:"hostname"`
}

// NewDhcpAPILeasesPutInput instantiates a new DhcpAPILeasesPutInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpAPILeasesPutInput(address string, addressLeaseTime string, hostname string) *DhcpAPILeasesPutInput {
	this := DhcpAPILeasesPutInput{}
	this.Address = address
	this.AddressLeaseTime = addressLeaseTime
	this.Hostname = hostname
	return &this
}

// NewDhcpAPILeasesPutInputWithDefaults instantiates a new DhcpAPILeasesPutInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpAPILeasesPutInputWithDefaults() *DhcpAPILeasesPutInput {
	this := DhcpAPILeasesPutInput{}
	return &this
}

// GetAddress returns the Address field value
func (o *DhcpAPILeasesPutInput) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DhcpAPILeasesPutInput) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DhcpAPILeasesPutInput) SetAddress(v string) {
	o.Address = v
}

// GetAddressLeaseTime returns the AddressLeaseTime field value
func (o *DhcpAPILeasesPutInput) GetAddressLeaseTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLeaseTime
}

// GetAddressLeaseTimeOk returns a tuple with the AddressLeaseTime field value
// and a boolean to check if the value has been set.
func (o *DhcpAPILeasesPutInput) GetAddressLeaseTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLeaseTime, true
}

// SetAddressLeaseTime sets field value
func (o *DhcpAPILeasesPutInput) SetAddressLeaseTime(v string) {
	o.AddressLeaseTime = v
}

// GetDnsZone returns the DnsZone field value if set, zero value otherwise.
func (o *DhcpAPILeasesPutInput) GetDnsZone() string {
	if o == nil || o.DnsZone == nil {
		var ret string
		return ret
	}
	return *o.DnsZone
}

// GetDnsZoneOk returns a tuple with the DnsZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpAPILeasesPutInput) GetDnsZoneOk() (*string, bool) {
	if o == nil || o.DnsZone == nil {
		return nil, false
	}
	return o.DnsZone, true
}

// HasDnsZone returns a boolean if a field has been set.
func (o *DhcpAPILeasesPutInput) HasDnsZone() bool {
	if o != nil && o.DnsZone != nil {
		return true
	}

	return false
}

// SetDnsZone gets a reference to the given string and assigns it to the DnsZone field.
func (o *DhcpAPILeasesPutInput) SetDnsZone(v string) {
	o.DnsZone = &v
}

// GetHostname returns the Hostname field value
func (o *DhcpAPILeasesPutInput) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *DhcpAPILeasesPutInput) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *DhcpAPILeasesPutInput) SetHostname(v string) {
	o.Hostname = v
}

func (o DhcpAPILeasesPutInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["addressLeaseTime"] = o.AddressLeaseTime
	}
	if o.DnsZone != nil {
		toSerialize["dnsZone"] = o.DnsZone
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpAPILeasesPutInput struct {
	value *DhcpAPILeasesPutInput
	isSet bool
}

func (v NullableDhcpAPILeasesPutInput) Get() *DhcpAPILeasesPutInput {
	return v.value
}

func (v *NullableDhcpAPILeasesPutInput) Set(val *DhcpAPILeasesPutInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpAPILeasesPutInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpAPILeasesPutInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpAPILeasesPutInput(val *DhcpAPILeasesPutInput) *NullableDhcpAPILeasesPutInput {
	return &NullableDhcpAPILeasesPutInput{value: val, isSet: true}
}

func (v NullableDhcpAPILeasesPutInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpAPILeasesPutInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
