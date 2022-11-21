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

// checks if the FunctionCreationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionCreationError{}

// FunctionCreationError struct for FunctionCreationError
type FunctionCreationError struct {
	// The error message
	Error *string `json:"error,omitempty"`
}

// NewFunctionCreationError instantiates a new FunctionCreationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionCreationError() *FunctionCreationError {
	this := FunctionCreationError{}
	return &this
}

// NewFunctionCreationErrorWithDefaults instantiates a new FunctionCreationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionCreationErrorWithDefaults() *FunctionCreationError {
	this := FunctionCreationError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FunctionCreationError) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCreationError) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FunctionCreationError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *FunctionCreationError) SetError(v string) {
	o.Error = &v
}

func (o FunctionCreationError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionCreationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableFunctionCreationError struct {
	value *FunctionCreationError
	isSet bool
}

func (v NullableFunctionCreationError) Get() *FunctionCreationError {
	return v.value
}

func (v *NullableFunctionCreationError) Set(val *FunctionCreationError) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionCreationError) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionCreationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionCreationError(val *FunctionCreationError) *NullableFunctionCreationError {
	return &NullableFunctionCreationError{value: val, isSet: true}
}

func (v NullableFunctionCreationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionCreationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


