# SingleModuleResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Functions** | Pointer to [**[]SubjectNameSubject**](SubjectNameSubject.md) |  | [optional] 

## Methods

### NewSingleModuleResultData

`func NewSingleModuleResultData() *SingleModuleResultData`

NewSingleModuleResultData instantiates a new SingleModuleResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleModuleResultDataWithDefaults

`func NewSingleModuleResultDataWithDefaults() *SingleModuleResultData`

NewSingleModuleResultDataWithDefaults instantiates a new SingleModuleResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SingleModuleResultData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SingleModuleResultData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SingleModuleResultData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SingleModuleResultData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFunctions

`func (o *SingleModuleResultData) GetFunctions() []SubjectNameSubject`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *SingleModuleResultData) GetFunctionsOk() (*[]SubjectNameSubject, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *SingleModuleResultData) SetFunctions(v []SubjectNameSubject)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *SingleModuleResultData) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


