# FunctionCreateUpdateSinksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the data sink | [optional] 
**Params** | Pointer to **map[string]interface{}** | Additional parameters for the data sink (usually connection params) | [optional] 

## Methods

### NewFunctionCreateUpdateSinksInner

`func NewFunctionCreateUpdateSinksInner() *FunctionCreateUpdateSinksInner`

NewFunctionCreateUpdateSinksInner instantiates a new FunctionCreateUpdateSinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCreateUpdateSinksInnerWithDefaults

`func NewFunctionCreateUpdateSinksInnerWithDefaults() *FunctionCreateUpdateSinksInner`

NewFunctionCreateUpdateSinksInnerWithDefaults instantiates a new FunctionCreateUpdateSinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionCreateUpdateSinksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionCreateUpdateSinksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionCreateUpdateSinksInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FunctionCreateUpdateSinksInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParams

`func (o *FunctionCreateUpdateSinksInner) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *FunctionCreateUpdateSinksInner) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *FunctionCreateUpdateSinksInner) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *FunctionCreateUpdateSinksInner) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


