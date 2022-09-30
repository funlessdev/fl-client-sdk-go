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

// FunctionDeletionSuccess struct for FunctionDeletionSuccess
type FunctionDeletionSuccess struct {
	// The name of the function
	Result *string `json:"result,omitempty"`
}

// NewFunctionDeletionSuccess instantiates a new FunctionDeletionSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionDeletionSuccess() *FunctionDeletionSuccess {
	this := FunctionDeletionSuccess{}
	return &this
}

// NewFunctionDeletionSuccessWithDefaults instantiates a new FunctionDeletionSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionDeletionSuccessWithDefaults() *FunctionDeletionSuccess {
	this := FunctionDeletionSuccess{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FunctionDeletionSuccess) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionDeletionSuccess) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FunctionDeletionSuccess) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *FunctionDeletionSuccess) SetResult(v string) {
	o.Result = &v
}

func (o FunctionDeletionSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionDeletionSuccess struct {
	value *FunctionDeletionSuccess
	isSet bool
}

func (v NullableFunctionDeletionSuccess) Get() *FunctionDeletionSuccess {
	return v.value
}

func (v *NullableFunctionDeletionSuccess) Set(val *FunctionDeletionSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionDeletionSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionDeletionSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionDeletionSuccess(val *FunctionDeletionSuccess) *NullableFunctionDeletionSuccess {
	return &NullableFunctionDeletionSuccess{value: val, isSet: true}
}

func (v NullableFunctionDeletionSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionDeletionSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


