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

// InvokeInput struct for InvokeInput
type InvokeInput struct {
	Args map[string]interface{} `json:"args,omitempty"`
}

// NewInvokeInput instantiates a new InvokeInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvokeInput() *InvokeInput {
	this := InvokeInput{}
	return &this
}

// NewInvokeInputWithDefaults instantiates a new InvokeInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvokeInputWithDefaults() *InvokeInput {
	this := InvokeInput{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *InvokeInput) GetArgs() map[string]interface{} {
	if o == nil || isNil(o.Args) {
		var ret map[string]interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeInput) GetArgsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Args) {
    return map[string]interface{}{}, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *InvokeInput) HasArgs() bool {
	if o != nil && !isNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given map[string]interface{} and assigns it to the Args field.
func (o *InvokeInput) SetArgs(v map[string]interface{}) {
	o.Args = v
}

func (o InvokeInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	return json.Marshal(toSerialize)
}

type NullableInvokeInput struct {
	value *InvokeInput
	isSet bool
}

func (v NullableInvokeInput) Get() *InvokeInput {
	return v.value
}

func (v *NullableInvokeInput) Set(val *InvokeInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInvokeInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInvokeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvokeInput(val *InvokeInput) *NullableInvokeInput {
	return &NullableInvokeInput{value: val, isSet: true}
}

func (v NullableInvokeInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvokeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


