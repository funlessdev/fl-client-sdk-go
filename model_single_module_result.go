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

// checks if the SingleModuleResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleModuleResult{}

// SingleModuleResult struct for SingleModuleResult
type SingleModuleResult struct {
	Data *SingleModuleResultData `json:"data,omitempty"`
}

// NewSingleModuleResult instantiates a new SingleModuleResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleModuleResult() *SingleModuleResult {
	this := SingleModuleResult{}
	return &this
}

// NewSingleModuleResultWithDefaults instantiates a new SingleModuleResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleModuleResultWithDefaults() *SingleModuleResult {
	this := SingleModuleResult{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SingleModuleResult) GetData() SingleModuleResultData {
	if o == nil || isNil(o.Data) {
		var ret SingleModuleResultData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleModuleResult) GetDataOk() (*SingleModuleResultData, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SingleModuleResult) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SingleModuleResultData and assigns it to the Data field.
func (o *SingleModuleResult) SetData(v SingleModuleResultData) {
	o.Data = &v
}

func (o SingleModuleResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleModuleResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSingleModuleResult struct {
	value *SingleModuleResult
	isSet bool
}

func (v NullableSingleModuleResult) Get() *SingleModuleResult {
	return v.value
}

func (v *NullableSingleModuleResult) Set(val *SingleModuleResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleModuleResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleModuleResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleModuleResult(val *SingleModuleResult) *NullableSingleModuleResult {
	return &NullableSingleModuleResult{value: val, isSet: true}
}

func (v NullableSingleModuleResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleModuleResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


