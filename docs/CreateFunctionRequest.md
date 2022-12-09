# CreateFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [***os.File**](*os.File.md) |  | [optional] 

## Methods

### NewCreateFunctionRequest

`func NewCreateFunctionRequest() *CreateFunctionRequest`

NewCreateFunctionRequest instantiates a new CreateFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFunctionRequestWithDefaults

`func NewCreateFunctionRequestWithDefaults() *CreateFunctionRequest`

NewCreateFunctionRequestWithDefaults instantiates a new CreateFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateFunctionRequest) GetCode() *os.File`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateFunctionRequest) GetCodeOk() (**os.File, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateFunctionRequest) SetCode(v *os.File)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateFunctionRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


