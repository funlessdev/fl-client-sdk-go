# FunctionCreateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the function | [optional] 
**Code** | Pointer to **os.File** | File with the code of the function | [optional] 
**Events** | Pointer to [**[]FunctionCreateUpdateEventsInner**](FunctionCreateUpdateEventsInner.md) | Events that can trigger the function | [optional] 

## Methods

### NewFunctionCreateUpdate

`func NewFunctionCreateUpdate() *FunctionCreateUpdate`

NewFunctionCreateUpdate instantiates a new FunctionCreateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCreateUpdateWithDefaults

`func NewFunctionCreateUpdateWithDefaults() *FunctionCreateUpdate`

NewFunctionCreateUpdateWithDefaults instantiates a new FunctionCreateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionCreateUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionCreateUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionCreateUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionCreateUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *FunctionCreateUpdate) GetCode() os.File`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FunctionCreateUpdate) GetCodeOk() (*os.File, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FunctionCreateUpdate) SetCode(v os.File)`

SetCode sets Code field to given value.

### HasCode

`func (o *FunctionCreateUpdate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEvents

`func (o *FunctionCreateUpdate) GetEvents() []FunctionCreateUpdateEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FunctionCreateUpdate) GetEventsOk() (*[]FunctionCreateUpdateEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FunctionCreateUpdate) SetEvents(v []FunctionCreateUpdateEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *FunctionCreateUpdate) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


