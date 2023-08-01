# AppCreateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the APP script | [optional] 
**File** | Pointer to ***os.File** | File with the APP script | [optional] 

## Methods

### NewAppCreateUpdate

`func NewAppCreateUpdate() *AppCreateUpdate`

NewAppCreateUpdate instantiates a new AppCreateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCreateUpdateWithDefaults

`func NewAppCreateUpdateWithDefaults() *AppCreateUpdate`

NewAppCreateUpdateWithDefaults instantiates a new AppCreateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppCreateUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppCreateUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppCreateUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppCreateUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFile

`func (o *AppCreateUpdate) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AppCreateUpdate) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AppCreateUpdate) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *AppCreateUpdate) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


