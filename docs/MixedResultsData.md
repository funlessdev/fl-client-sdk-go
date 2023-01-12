# MixedResultsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the function | [optional] 
**Events** | Pointer to [**[]MixedResultsDataEventsInner**](MixedResultsDataEventsInner.md) | The results of event connection, both successful and failed | [optional] 
**Sinks** | Pointer to [**[]MixedResultsDataSinksInner**](MixedResultsDataSinksInner.md) | The results of sink connection, both successful and failed | [optional] 
**SinksMetadata** | Pointer to [**MixedResultsDataSinksMetadata**](MixedResultsDataSinksMetadata.md) |  | [optional] 
**EventsMetadata** | Pointer to [**MixedResultsDataEventsMetadata**](MixedResultsDataEventsMetadata.md) |  | [optional] 

## Methods

### NewMixedResultsData

`func NewMixedResultsData() *MixedResultsData`

NewMixedResultsData instantiates a new MixedResultsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixedResultsDataWithDefaults

`func NewMixedResultsDataWithDefaults() *MixedResultsData`

NewMixedResultsDataWithDefaults instantiates a new MixedResultsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MixedResultsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MixedResultsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MixedResultsData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MixedResultsData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEvents

`func (o *MixedResultsData) GetEvents() []MixedResultsDataEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MixedResultsData) GetEventsOk() (*[]MixedResultsDataEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MixedResultsData) SetEvents(v []MixedResultsDataEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *MixedResultsData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetSinks

`func (o *MixedResultsData) GetSinks() []MixedResultsDataSinksInner`

GetSinks returns the Sinks field if non-nil, zero value otherwise.

### GetSinksOk

`func (o *MixedResultsData) GetSinksOk() (*[]MixedResultsDataSinksInner, bool)`

GetSinksOk returns a tuple with the Sinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinks

`func (o *MixedResultsData) SetSinks(v []MixedResultsDataSinksInner)`

SetSinks sets Sinks field to given value.

### HasSinks

`func (o *MixedResultsData) HasSinks() bool`

HasSinks returns a boolean if a field has been set.

### GetSinksMetadata

`func (o *MixedResultsData) GetSinksMetadata() MixedResultsDataSinksMetadata`

GetSinksMetadata returns the SinksMetadata field if non-nil, zero value otherwise.

### GetSinksMetadataOk

`func (o *MixedResultsData) GetSinksMetadataOk() (*MixedResultsDataSinksMetadata, bool)`

GetSinksMetadataOk returns a tuple with the SinksMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinksMetadata

`func (o *MixedResultsData) SetSinksMetadata(v MixedResultsDataSinksMetadata)`

SetSinksMetadata sets SinksMetadata field to given value.

### HasSinksMetadata

`func (o *MixedResultsData) HasSinksMetadata() bool`

HasSinksMetadata returns a boolean if a field has been set.

### GetEventsMetadata

`func (o *MixedResultsData) GetEventsMetadata() MixedResultsDataEventsMetadata`

GetEventsMetadata returns the EventsMetadata field if non-nil, zero value otherwise.

### GetEventsMetadataOk

`func (o *MixedResultsData) GetEventsMetadataOk() (*MixedResultsDataEventsMetadata, bool)`

GetEventsMetadataOk returns a tuple with the EventsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsMetadata

`func (o *MixedResultsData) SetEventsMetadata(v MixedResultsDataEventsMetadata)`

SetEventsMetadata sets EventsMetadata field to given value.

### HasEventsMetadata

`func (o *MixedResultsData) HasEventsMetadata() bool`

HasEventsMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


