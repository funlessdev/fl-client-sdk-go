# FunctionCreateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the function | [optional] 
**Code** | Pointer to ***os.File** | File with the code of the function | [optional] 
**WaitForWorkers** | Pointer to **bool** | Whether to wait for all workers to receive the code of the function. If false, the request returns as soon as the creation request terminates. | [optional] [default to true]
**Events** | Pointer to [**[]FunctionCreateUpdateEventsInner**](FunctionCreateUpdateEventsInner.md) | Events that can trigger the function | [optional] 
**Sinks** | Pointer to [**[]FunctionCreateUpdateSinksInner**](FunctionCreateUpdateSinksInner.md) | Data sinks that receive invocation&#39;s results | [optional] 

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

`func (o *FunctionCreateUpdate) GetCode() *os.File`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FunctionCreateUpdate) GetCodeOk() (**os.File, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FunctionCreateUpdate) SetCode(v *os.File)`

SetCode sets Code field to given value.

### HasCode

`func (o *FunctionCreateUpdate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetWaitForWorkers

`func (o *FunctionCreateUpdate) GetWaitForWorkers() bool`

GetWaitForWorkers returns the WaitForWorkers field if non-nil, zero value otherwise.

### GetWaitForWorkersOk

`func (o *FunctionCreateUpdate) GetWaitForWorkersOk() (*bool, bool)`

GetWaitForWorkersOk returns a tuple with the WaitForWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForWorkers

`func (o *FunctionCreateUpdate) SetWaitForWorkers(v bool)`

SetWaitForWorkers sets WaitForWorkers field to given value.

### HasWaitForWorkers

`func (o *FunctionCreateUpdate) HasWaitForWorkers() bool`

HasWaitForWorkers returns a boolean if a field has been set.

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

### GetSinks

`func (o *FunctionCreateUpdate) GetSinks() []FunctionCreateUpdateSinksInner`

GetSinks returns the Sinks field if non-nil, zero value otherwise.

### GetSinksOk

`func (o *FunctionCreateUpdate) GetSinksOk() (*[]FunctionCreateUpdateSinksInner, bool)`

GetSinksOk returns a tuple with the Sinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinks

`func (o *FunctionCreateUpdate) SetSinks(v []FunctionCreateUpdateSinksInner)`

SetSinks sets Sinks field to given value.

### HasSinks

`func (o *FunctionCreateUpdate) HasSinks() bool`

HasSinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


