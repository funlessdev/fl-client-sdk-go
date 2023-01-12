# MixedResultsDataEventsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | Pointer to **int32** | The amount of events that was successfully connected | [optional] 
**Failed** | Pointer to **int32** | The amount of events that wasn&#39;t successfully connected | [optional] 
**Total** | Pointer to **int32** | The total amount of events that was passed | [optional] 

## Methods

### NewMixedResultsDataEventsMetadata

`func NewMixedResultsDataEventsMetadata() *MixedResultsDataEventsMetadata`

NewMixedResultsDataEventsMetadata instantiates a new MixedResultsDataEventsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixedResultsDataEventsMetadataWithDefaults

`func NewMixedResultsDataEventsMetadataWithDefaults() *MixedResultsDataEventsMetadata`

NewMixedResultsDataEventsMetadataWithDefaults instantiates a new MixedResultsDataEventsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessful

`func (o *MixedResultsDataEventsMetadata) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *MixedResultsDataEventsMetadata) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *MixedResultsDataEventsMetadata) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *MixedResultsDataEventsMetadata) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetFailed

`func (o *MixedResultsDataEventsMetadata) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *MixedResultsDataEventsMetadata) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *MixedResultsDataEventsMetadata) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *MixedResultsDataEventsMetadata) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetTotal

`func (o *MixedResultsDataEventsMetadata) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MixedResultsDataEventsMetadata) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MixedResultsDataEventsMetadata) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MixedResultsDataEventsMetadata) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


