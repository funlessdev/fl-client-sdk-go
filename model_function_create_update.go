/*
FunLess Platfom API

The API for the FunLess Platform

API version: 0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"os"
)

// checks if the FunctionCreateUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionCreateUpdate{}

// FunctionCreateUpdate struct for FunctionCreateUpdate
type FunctionCreateUpdate struct {
	// Name of the function
	Name *string `json:"name,omitempty"`
	// File with the code of the function
	Code **os.File `json:"code,omitempty"`
	// Events that can trigger the function
	Events []FunctionCreateUpdateEventsInner `json:"events,omitempty"`
	// Data sinks that receive invocation's results
	Sinks []FunctionCreateUpdateSinksInner `json:"sinks,omitempty"`
}

// NewFunctionCreateUpdate instantiates a new FunctionCreateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionCreateUpdate() *FunctionCreateUpdate {
	this := FunctionCreateUpdate{}
	return &this
}

// NewFunctionCreateUpdateWithDefaults instantiates a new FunctionCreateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionCreateUpdateWithDefaults() *FunctionCreateUpdate {
	this := FunctionCreateUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FunctionCreateUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCreateUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FunctionCreateUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FunctionCreateUpdate) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FunctionCreateUpdate) GetCode() *os.File {
	if o == nil || isNil(o.Code) {
		var ret *os.File
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCreateUpdate) GetCodeOk() (**os.File, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FunctionCreateUpdate) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given *os.File and assigns it to the Code field.
func (o *FunctionCreateUpdate) SetCode(v *os.File) {
	o.Code = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *FunctionCreateUpdate) GetEvents() []FunctionCreateUpdateEventsInner {
	if o == nil || isNil(o.Events) {
		var ret []FunctionCreateUpdateEventsInner
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCreateUpdate) GetEventsOk() ([]FunctionCreateUpdateEventsInner, bool) {
	if o == nil || isNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *FunctionCreateUpdate) HasEvents() bool {
	if o != nil && !isNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []FunctionCreateUpdateEventsInner and assigns it to the Events field.
func (o *FunctionCreateUpdate) SetEvents(v []FunctionCreateUpdateEventsInner) {
	o.Events = v
}

// GetSinks returns the Sinks field value if set, zero value otherwise.
func (o *FunctionCreateUpdate) GetSinks() []FunctionCreateUpdateSinksInner {
	if o == nil || isNil(o.Sinks) {
		var ret []FunctionCreateUpdateSinksInner
		return ret
	}
	return o.Sinks
}

// GetSinksOk returns a tuple with the Sinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCreateUpdate) GetSinksOk() ([]FunctionCreateUpdateSinksInner, bool) {
	if o == nil || isNil(o.Sinks) {
		return nil, false
	}
	return o.Sinks, true
}

// HasSinks returns a boolean if a field has been set.
func (o *FunctionCreateUpdate) HasSinks() bool {
	if o != nil && !isNil(o.Sinks) {
		return true
	}

	return false
}

// SetSinks gets a reference to the given []FunctionCreateUpdateSinksInner and assigns it to the Sinks field.
func (o *FunctionCreateUpdate) SetSinks(v []FunctionCreateUpdateSinksInner) {
	o.Sinks = v
}

func (o FunctionCreateUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionCreateUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !isNil(o.Sinks) {
		toSerialize["sinks"] = o.Sinks
	}
	return toSerialize, nil
}

type NullableFunctionCreateUpdate struct {
	value *FunctionCreateUpdate
	isSet bool
}

func (v NullableFunctionCreateUpdate) Get() *FunctionCreateUpdate {
	return v.value
}

func (v *NullableFunctionCreateUpdate) Set(val *FunctionCreateUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionCreateUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionCreateUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionCreateUpdate(val *FunctionCreateUpdate) *NullableFunctionCreateUpdate {
	return &NullableFunctionCreateUpdate{value: val, isSet: true}
}

func (v NullableFunctionCreateUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionCreateUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


