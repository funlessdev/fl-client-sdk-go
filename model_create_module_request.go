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

// checks if the CreateModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateModuleRequest{}

// CreateModuleRequest struct for CreateModuleRequest
type CreateModuleRequest struct {
	Name *string `json:"name,omitempty"`
}

// NewCreateModuleRequest instantiates a new CreateModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateModuleRequest() *CreateModuleRequest {
	this := CreateModuleRequest{}
	return &this
}

// NewCreateModuleRequestWithDefaults instantiates a new CreateModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateModuleRequestWithDefaults() *CreateModuleRequest {
	this := CreateModuleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateModuleRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateModuleRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateModuleRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateModuleRequest) SetName(v string) {
	o.Name = &v
}

func (o CreateModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateModuleRequest struct {
	value *CreateModuleRequest
	isSet bool
}

func (v NullableCreateModuleRequest) Get() *CreateModuleRequest {
	return v.value
}

func (v *NullableCreateModuleRequest) Set(val *CreateModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateModuleRequest(val *CreateModuleRequest) *NullableCreateModuleRequest {
	return &NullableCreateModuleRequest{value: val, isSet: true}
}

func (v NullableCreateModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


