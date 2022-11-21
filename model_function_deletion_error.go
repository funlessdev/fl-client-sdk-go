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

// checks if the FunctionDeletionError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionDeletionError{}

// FunctionDeletionError struct for FunctionDeletionError
type FunctionDeletionError struct {
	// The error message
	Error *string `json:"error,omitempty"`
}

// NewFunctionDeletionError instantiates a new FunctionDeletionError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionDeletionError() *FunctionDeletionError {
	this := FunctionDeletionError{}
	return &this
}

// NewFunctionDeletionErrorWithDefaults instantiates a new FunctionDeletionError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionDeletionErrorWithDefaults() *FunctionDeletionError {
	this := FunctionDeletionError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FunctionDeletionError) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionDeletionError) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FunctionDeletionError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *FunctionDeletionError) SetError(v string) {
	o.Error = &v
}

func (o FunctionDeletionError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionDeletionError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableFunctionDeletionError struct {
	value *FunctionDeletionError
	isSet bool
}

func (v NullableFunctionDeletionError) Get() *FunctionDeletionError {
	return v.value
}

func (v *NullableFunctionDeletionError) Set(val *FunctionDeletionError) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionDeletionError) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionDeletionError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionDeletionError(val *FunctionDeletionError) *NullableFunctionDeletionError {
	return &NullableFunctionDeletionError{value: val, isSet: true}
}

func (v NullableFunctionDeletionError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionDeletionError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


