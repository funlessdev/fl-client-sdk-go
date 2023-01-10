# MixedEventResultsDataMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | Pointer to **interface{}** | The amount of events that was successfully connected | [optional] 
**Failed** | Pointer to **interface{}** | The amount of events that wasn&#39;t successfully connected | [optional] 
**Total** | Pointer to **interface{}** | The total amount of events that was passed | [optional] 

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

`func (o *MixedEventResultsDataMetadata) GetSuccessful() interface{}`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *MixedEventResultsDataMetadata) GetSuccessfulOk() (*interface{}, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *MixedEventResultsDataMetadata) SetSuccessful(v interface{})`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *MixedEventResultsDataMetadata) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### SetSuccessfulNil

`func (o *MixedEventResultsDataMetadata) SetSuccessfulNil(b bool)`

 SetSuccessfulNil sets the value for Successful to be an explicit nil

### UnsetSuccessful
`func (o *MixedEventResultsDataMetadata) UnsetSuccessful()`

UnsetSuccessful ensures that no value is present for Successful, not even an explicit nil
### GetFailed

`func (o *MixedEventResultsDataMetadata) GetFailed() interface{}`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *MixedEventResultsDataMetadata) GetFailedOk() (*interface{}, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *MixedEventResultsDataMetadata) SetFailed(v interface{})`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *MixedEventResultsDataMetadata) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *MixedEventResultsDataMetadata) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *MixedEventResultsDataMetadata) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetTotal

`func (o *MixedEventResultsDataMetadata) GetTotal() interface{}`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MixedEventResultsDataMetadata) GetTotalOk() (*interface{}, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MixedEventResultsDataMetadata) SetTotal(v interface{})`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MixedEventResultsDataMetadata) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *MixedEventResultsDataMetadata) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *MixedEventResultsDataMetadata) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


