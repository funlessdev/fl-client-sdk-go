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

// checks if the SingleModuleResultData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleModuleResultData{}

// SingleModuleResultData struct for SingleModuleResultData
type SingleModuleResultData struct {
	Name *string `json:"name,omitempty"`
	Functions []ModuleNameModule `json:"functions,omitempty"`
}

// NewSingleModuleResultData instantiates a new SingleModuleResultData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleModuleResultData() *SingleModuleResultData {
	this := SingleModuleResultData{}
	return &this
}

// NewSingleModuleResultDataWithDefaults instantiates a new SingleModuleResultData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleModuleResultDataWithDefaults() *SingleModuleResultData {
	this := SingleModuleResultData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SingleModuleResultData) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleModuleResultData) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SingleModuleResultData) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SingleModuleResultData) SetName(v string) {
	o.Name = &v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
func (o *SingleModuleResultData) GetFunctions() []ModuleNameModule {
	if o == nil || isNil(o.Functions) {
		var ret []ModuleNameModule
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleModuleResultData) GetFunctionsOk() ([]ModuleNameModule, bool) {
	if o == nil || isNil(o.Functions) {
		return nil, false
	}
	return o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *SingleModuleResultData) HasFunctions() bool {
	if o != nil && !isNil(o.Functions) {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given []ModuleNameModule and assigns it to the Functions field.
func (o *SingleModuleResultData) SetFunctions(v []ModuleNameModule) {
	o.Functions = v
}

func (o SingleModuleResultData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleModuleResultData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Functions) {
		toSerialize["functions"] = o.Functions
	}
	return toSerialize, nil
}

type NullableSingleModuleResultData struct {
	value *SingleModuleResultData
	isSet bool
}

func (v NullableSingleModuleResultData) Get() *SingleModuleResultData {
	return v.value
}

func (v *NullableSingleModuleResultData) Set(val *SingleModuleResultData) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleModuleResultData) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleModuleResultData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleModuleResultData(val *SingleModuleResultData) *NullableSingleModuleResultData {
	return &NullableSingleModuleResultData{value: val, isSet: true}
}

func (v NullableSingleModuleResultData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleModuleResultData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


