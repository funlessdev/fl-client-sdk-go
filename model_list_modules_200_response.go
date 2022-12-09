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

// ListModules200Response struct for ListModules200Response
type ListModules200Response struct {
	Data []string `json:"data,omitempty"`
}

// NewListModules200Response instantiates a new ListModules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListModules200Response() *ListModules200Response {
	this := ListModules200Response{}
	return &this
}

// NewListModules200ResponseWithDefaults instantiates a new ListModules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListModules200ResponseWithDefaults() *ListModules200Response {
	this := ListModules200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListModules200Response) GetData() []string {
	if o == nil || isNil(o.Data) {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListModules200Response) GetDataOk() ([]string, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListModules200Response) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ListModules200Response) SetData(v []string) {
	o.Data = v
}

func (o ListModules200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListModules200Response struct {
	value *ListModules200Response
	isSet bool
}

func (v NullableListModules200Response) Get() *ListModules200Response {
	return v.value
}

func (v *NullableListModules200Response) Set(val *ListModules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListModules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListModules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListModules200Response(val *ListModules200Response) *NullableListModules200Response {
	return &NullableListModules200Response{value: val, isSet: true}
}

func (v NullableListModules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListModules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


