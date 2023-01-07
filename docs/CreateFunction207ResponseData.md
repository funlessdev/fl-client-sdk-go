# CreateFunction207ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the function | [optional] 
**Events** | Pointer to [**[]CreateFunction207ResponseDataEventsInner**](CreateFunction207ResponseDataEventsInner.md) | The results of event connection, both successful and failed | [optional] 
**Metadata** | Pointer to [**CreateFunction207ResponseDataMetadata**](CreateFunction207ResponseDataMetadata.md) |  | [optional] 

## Methods

### NewCreateFunction207ResponseData

`func NewCreateFunction207ResponseData() *CreateFunction207ResponseData`

NewCreateFunction207ResponseData instantiates a new CreateFunction207ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFunction207ResponseDataWithDefaults

`func NewCreateFunction207ResponseDataWithDefaults() *CreateFunction207ResponseData`

NewCreateFunction207ResponseDataWithDefaults instantiates a new CreateFunction207ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFunction207ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFunction207ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFunction207ResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateFunction207ResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEvents

`func (o *CreateFunction207ResponseData) GetEvents() []CreateFunction207ResponseDataEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateFunction207ResponseData) GetEventsOk() (*[]CreateFunction207ResponseDataEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateFunction207ResponseData) SetEvents(v []CreateFunction207ResponseDataEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CreateFunction207ResponseData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateFunction207ResponseData) GetMetadata() CreateFunction207ResponseDataMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateFunction207ResponseData) GetMetadataOk() (*CreateFunction207ResponseDataMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateFunction207ResponseData) SetMetadata(v CreateFunction207ResponseDataMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateFunction207ResponseData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


