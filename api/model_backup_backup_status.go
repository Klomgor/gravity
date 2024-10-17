/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the BackupBackupStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupBackupStatus{}

// BackupBackupStatus struct for BackupBackupStatus
type BackupBackupStatus struct {
	Duration int32     `json:"duration"`
	Error    string    `json:"error"`
	Filename string    `json:"filename"`
	Size     int32     `json:"size"`
	Status   string    `json:"status"`
	Time     time.Time `json:"time"`
}

// NewBackupBackupStatus instantiates a new BackupBackupStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupBackupStatus(duration int32, error_ string, filename string, size int32, status string, time time.Time) *BackupBackupStatus {
	this := BackupBackupStatus{}
	this.Duration = duration
	this.Error = error_
	this.Filename = filename
	this.Size = size
	this.Status = status
	this.Time = time
	return &this
}

// NewBackupBackupStatusWithDefaults instantiates a new BackupBackupStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupBackupStatusWithDefaults() *BackupBackupStatus {
	this := BackupBackupStatus{}
	return &this
}

// GetDuration returns the Duration field value
func (o *BackupBackupStatus) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *BackupBackupStatus) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *BackupBackupStatus) SetDuration(v int32) {
	o.Duration = v
}

// GetError returns the Error field value
func (o *BackupBackupStatus) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *BackupBackupStatus) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *BackupBackupStatus) SetError(v string) {
	o.Error = v
}

// GetFilename returns the Filename field value
func (o *BackupBackupStatus) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *BackupBackupStatus) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *BackupBackupStatus) SetFilename(v string) {
	o.Filename = v
}

// GetSize returns the Size field value
func (o *BackupBackupStatus) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *BackupBackupStatus) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *BackupBackupStatus) SetSize(v int32) {
	o.Size = v
}

// GetStatus returns the Status field value
func (o *BackupBackupStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BackupBackupStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BackupBackupStatus) SetStatus(v string) {
	o.Status = v
}

// GetTime returns the Time field value
func (o *BackupBackupStatus) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *BackupBackupStatus) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *BackupBackupStatus) SetTime(v time.Time) {
	o.Time = v
}

func (o BackupBackupStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupBackupStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["error"] = o.Error
	toSerialize["filename"] = o.Filename
	toSerialize["size"] = o.Size
	toSerialize["status"] = o.Status
	toSerialize["time"] = o.Time
	return toSerialize, nil
}

type NullableBackupBackupStatus struct {
	value *BackupBackupStatus
	isSet bool
}

func (v NullableBackupBackupStatus) Get() *BackupBackupStatus {
	return v.value
}

func (v *NullableBackupBackupStatus) Set(val *BackupBackupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupBackupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupBackupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupBackupStatus(val *BackupBackupStatus) *NullableBackupBackupStatus {
	return &NullableBackupBackupStatus{value: val, isSet: true}
}

func (v NullableBackupBackupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupBackupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
