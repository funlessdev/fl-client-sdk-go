# MixedEventResultsDataMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | Pointer to **int32** | The amount of events that was successfully connected | [optional] 
**Failed** | Pointer to **int32** | The amount of events that wasn&#39;t successfully connected | [optional] 
**Total** | Pointer to **int32** | The total amount of events that was passed | [optional] 

## Methods

### NewMixedEventResultsDataMetadata

`func NewMixedEventResultsDataMetadata() *MixedEventResultsDataMetadata`

NewMixedEventResultsDataMetadata instantiates a new MixedEventResultsDataMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixedEventResultsDataMetadataWithDefaults

`func NewMixedEventResultsDataMetadataWithDefaults() *MixedEventResultsDataMetadata`

NewMixedEventResultsDataMetadataWithDefaults instantiates a new MixedEventResultsDataMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessful

`func (o *MixedEventResultsDataMetadata) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *MixedEventResultsDataMetadata) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *MixedEventResultsDataMetadata) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *MixedEventResultsDataMetadata) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetFailed

`func (o *MixedEventResultsDataMetadata) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *MixedEventResultsDataMetadata) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *MixedEventResultsDataMetadata) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *MixedEventResultsDataMetadata) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetTotal

`func (o *MixedEventResultsDataMetadata) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MixedEventResultsDataMetadata) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MixedEventResultsDataMetadata) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MixedEventResultsDataMetadata) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


