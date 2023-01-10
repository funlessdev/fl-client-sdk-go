# MixedEventResultsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the function | [optional] 
**Events** | Pointer to [**[]MixedEventResultsDataEventsInner**](MixedEventResultsDataEventsInner.md) | The results of event connection, both successful and failed | [optional] 
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

`func (o *MixedEventResultsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MixedEventResultsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MixedEventResultsData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MixedEventResultsData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEvents

`func (o *MixedEventResultsData) GetEvents() []MixedEventResultsDataEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MixedEventResultsData) GetEventsOk() (*[]MixedEventResultsDataEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MixedEventResultsData) SetEvents(v []MixedEventResultsDataEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *MixedEventResultsData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

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


