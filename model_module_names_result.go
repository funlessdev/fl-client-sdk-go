/*
FunLess Platfom API

The API for the FunLess Platform

API version: 0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ModuleNamesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleNamesResult{}

// ModuleNamesResult struct for ModuleNamesResult
type ModuleNamesResult struct {
	Data []SubjectNameSubject `json:"data,omitempty"`
}

// NewModuleNamesResult instantiates a new ModuleNamesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleNamesResult() *ModuleNamesResult {
	this := ModuleNamesResult{}
	return &this
}

// NewModuleNamesResultWithDefaults instantiates a new ModuleNamesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleNamesResultWithDefaults() *ModuleNamesResult {
	this := ModuleNamesResult{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModuleNamesResult) GetData() []SubjectNameSubject {
	if o == nil || IsNil(o.Data) {
		var ret []SubjectNameSubject
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleNamesResult) GetDataOk() ([]SubjectNameSubject, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModuleNamesResult) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SubjectNameSubject and assigns it to the Data field.
func (o *ModuleNamesResult) SetData(v []SubjectNameSubject) {
	o.Data = v
}

func (o ModuleNamesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleNamesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModuleNamesResult struct {
	value *ModuleNamesResult
	isSet bool
}

func (v NullableModuleNamesResult) Get() *ModuleNamesResult {
	return v.value
}

func (v *NullableModuleNamesResult) Set(val *ModuleNamesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleNamesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleNamesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleNamesResult(val *ModuleNamesResult) *NullableModuleNamesResult {
	return &NullableModuleNamesResult{value: val, isSet: true}
}

func (v NullableModuleNamesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleNamesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


