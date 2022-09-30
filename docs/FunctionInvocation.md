# FunctionInvocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Function** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFunctionInvocation

`func NewFunctionInvocation() *FunctionInvocation`

NewFunctionInvocation instantiates a new FunctionInvocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionInvocationWithDefaults

`func NewFunctionInvocationWithDefaults() *FunctionInvocation`

NewFunctionInvocationWithDefaults instantiates a new FunctionInvocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *FunctionInvocation) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FunctionInvocation) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FunctionInvocation) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FunctionInvocation) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetFunction

`func (o *FunctionInvocation) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *FunctionInvocation) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *FunctionInvocation) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *FunctionInvocation) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetArgs

`func (o *FunctionInvocation) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *FunctionInvocation) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *FunctionInvocation) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *FunctionInvocation) HasArgs() bool`

HasArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


