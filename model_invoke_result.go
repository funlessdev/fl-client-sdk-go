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

// checks if the InvokeResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvokeResult{}

// InvokeResult struct for InvokeResult
type InvokeResult struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewInvokeResult instantiates a new InvokeResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvokeResult() *InvokeResult {
	this := InvokeResult{}
	return &this
}

// NewInvokeResultWithDefaults instantiates a new InvokeResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvokeResultWithDefaults() *InvokeResult {
	this := InvokeResult{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InvokeResult) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeResult) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InvokeResult) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *InvokeResult) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o InvokeResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvokeResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableInvokeResult struct {
	value *InvokeResult
	isSet bool
}

func (v NullableInvokeResult) Get() *InvokeResult {
	return v.value
}

func (v *NullableInvokeResult) Set(val *InvokeResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInvokeResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInvokeResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvokeResult(val *InvokeResult) *NullableInvokeResult {
	return &NullableInvokeResult{value: val, isSet: true}
}

func (v NullableInvokeResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvokeResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


