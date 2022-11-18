/*
Funless Platfom API

The API for the Funless Platform

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FunctionCreationSuccess struct for FunctionCreationSuccess
type FunctionCreationSuccess struct {
	// The name of the function
	Result *string `json:"result,omitempty"`
}

// NewFunctionCreationSuccess instantiates a new FunctionCreationSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionCreationSuccess() *FunctionCreationSuccess {
	this := FunctionCreationSuccess{}
	return &this
}

// NewFunctionCreationSuccessWithDefaults instantiates a new FunctionCreationSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionCreationSuccessWithDefaults() *FunctionCreationSuccess {
	this := FunctionCreationSuccess{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FunctionCreationSuccess) GetResult() string {
	if o == nil || isNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCreationSuccess) GetResultOk() (*string, bool) {
	if o == nil || isNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FunctionCreationSuccess) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *FunctionCreationSuccess) SetResult(v string) {
	o.Result = &v
}

func (o FunctionCreationSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionCreationSuccess struct {
	value *FunctionCreationSuccess
	isSet bool
}

func (v NullableFunctionCreationSuccess) Get() *FunctionCreationSuccess {
	return v.value
}

func (v *NullableFunctionCreationSuccess) Set(val *FunctionCreationSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionCreationSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionCreationSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionCreationSuccess(val *FunctionCreationSuccess) *NullableFunctionCreationSuccess {
	return &NullableFunctionCreationSuccess{value: val, isSet: true}
}

func (v NullableFunctionCreationSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionCreationSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


