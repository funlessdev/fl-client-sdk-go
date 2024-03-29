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

// checks if the MixedResultsDataEventsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MixedResultsDataEventsMetadata{}

// MixedResultsDataEventsMetadata struct for MixedResultsDataEventsMetadata
type MixedResultsDataEventsMetadata struct {
	// The amount of events that was successfully connected
	Successful *int32 `json:"successful,omitempty"`
	// The amount of events that wasn't successfully connected
	Failed *int32 `json:"failed,omitempty"`
	// The total amount of events that was passed
	Total *int32 `json:"total,omitempty"`
}

// NewMixedResultsDataEventsMetadata instantiates a new MixedResultsDataEventsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMixedResultsDataEventsMetadata() *MixedResultsDataEventsMetadata {
	this := MixedResultsDataEventsMetadata{}
	return &this
}

// NewMixedResultsDataEventsMetadataWithDefaults instantiates a new MixedResultsDataEventsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMixedResultsDataEventsMetadataWithDefaults() *MixedResultsDataEventsMetadata {
	this := MixedResultsDataEventsMetadata{}
	return &this
}

// GetSuccessful returns the Successful field value if set, zero value otherwise.
func (o *MixedResultsDataEventsMetadata) GetSuccessful() int32 {
	if o == nil || IsNil(o.Successful) {
		var ret int32
		return ret
	}
	return *o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixedResultsDataEventsMetadata) GetSuccessfulOk() (*int32, bool) {
	if o == nil || IsNil(o.Successful) {
		return nil, false
	}
	return o.Successful, true
}

// HasSuccessful returns a boolean if a field has been set.
func (o *MixedResultsDataEventsMetadata) HasSuccessful() bool {
	if o != nil && !IsNil(o.Successful) {
		return true
	}

	return false
}

// SetSuccessful gets a reference to the given int32 and assigns it to the Successful field.
func (o *MixedResultsDataEventsMetadata) SetSuccessful(v int32) {
	o.Successful = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *MixedResultsDataEventsMetadata) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixedResultsDataEventsMetadata) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *MixedResultsDataEventsMetadata) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *MixedResultsDataEventsMetadata) SetFailed(v int32) {
	o.Failed = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MixedResultsDataEventsMetadata) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixedResultsDataEventsMetadata) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MixedResultsDataEventsMetadata) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MixedResultsDataEventsMetadata) SetTotal(v int32) {
	o.Total = &v
}

func (o MixedResultsDataEventsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MixedResultsDataEventsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Successful) {
		toSerialize["successful"] = o.Successful
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMixedResultsDataEventsMetadata struct {
	value *MixedResultsDataEventsMetadata
	isSet bool
}

func (v NullableMixedResultsDataEventsMetadata) Get() *MixedResultsDataEventsMetadata {
	return v.value
}

func (v *NullableMixedResultsDataEventsMetadata) Set(val *MixedResultsDataEventsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMixedResultsDataEventsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMixedResultsDataEventsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMixedResultsDataEventsMetadata(val *MixedResultsDataEventsMetadata) *NullableMixedResultsDataEventsMetadata {
	return &NullableMixedResultsDataEventsMetadata{value: val, isSet: true}
}

func (v NullableMixedResultsDataEventsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMixedResultsDataEventsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


