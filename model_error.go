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

// ModelError struct for ModelError
type ModelError struct {
	Errors ListModulesDefaultResponseErrors `json:"errors"`
}

// NewModelError instantiates a new ModelError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelError(errors ListModulesDefaultResponseErrors) *ModelError {
	this := ModelError{}
	this.Errors = errors
	return &this
}

// NewModelErrorWithDefaults instantiates a new ModelError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelErrorWithDefaults() *ModelError {
	this := ModelError{}
	return &this
}

// GetErrors returns the Errors field value
func (o *ModelError) GetErrors() ListModulesDefaultResponseErrors {
	if o == nil {
		var ret ListModulesDefaultResponseErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorsOk() (*ListModulesDefaultResponseErrors, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *ModelError) SetErrors(v ListModulesDefaultResponseErrors) {
	o.Errors = v
}

func (o ModelError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableModelError struct {
	value *ModelError
	isSet bool
}

func (v NullableModelError) Get() *ModelError {
	return v.value
}

func (v *NullableModelError) Set(val *ModelError) {
	v.value = val
	v.isSet = true
}

func (v NullableModelError) IsSet() bool {
	return v.isSet
}

func (v *NullableModelError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelError(val *ModelError) *NullableModelError {
	return &NullableModelError{value: val, isSet: true}
}

func (v NullableModelError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


