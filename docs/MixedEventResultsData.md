# MixedEventResultsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The name of the function | [optional] 
**Events** | Pointer to **interface{}** | The results of event connection, both successful and failed | [optional] 
**Metadata** | Pointer to [**MixedEventResultsDataMetadata**](MixedEventResultsDataMetadata.md) |  | [optional] 

## Methods

### NewMixedEventResultsData

`func NewMixedEventResultsData() *MixedEventResultsData`

NewMixedEventResultsData instantiates a new MixedEventResultsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixedEventResultsDataWithDefaults

`func NewMixedEventResultsDataWithDefaults() *MixedEventResultsData`

NewMixedEventResultsDataWithDefaults instantiates a new MixedEventResultsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MixedEventResultsData) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MixedEventResultsData) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MixedEventResultsData) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *MixedEventResultsData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MixedEventResultsData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MixedEventResultsData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEvents

`func (o *MixedEventResultsData) GetEvents() interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MixedEventResultsData) GetEventsOk() (*interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MixedEventResultsData) SetEvents(v interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *MixedEventResultsData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *MixedEventResultsData) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *MixedEventResultsData) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetMetadata

`func (o *MixedEventResultsData) GetMetadata() MixedEventResultsDataMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MixedEventResultsData) GetMetadataOk() (*MixedEventResultsDataMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MixedEventResultsData) SetMetadata(v MixedEventResultsDataMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MixedEventResultsData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


