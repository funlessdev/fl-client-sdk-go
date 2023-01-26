# \FunctionsApi

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](FunctionsApi.md#CreateFunction) | **Post** /v1/fn/{module_name} | Create new function
[**DeleteFunction**](FunctionsApi.md#DeleteFunction) | **Delete** /v1/fn/{module_name}/{function_name} | Delete function
[**InvokeFunction**](FunctionsApi.md#InvokeFunction) | **Post** /v1/fn/{module_name}/{function_name} | Invoke function
[**ShowFunctionByName**](FunctionsApi.md#ShowFunctionByName) | **Get** /v1/fn/{module_name}/{function_name} | Show function info
[**UpdateFunction**](FunctionsApi.md#UpdateFunction) | **Put** /v1/fn/{module_name}/{function_name} | Update function



## CreateFunction

> CreateFunction(ctx, moduleName).Name(name).Code(code).Events(events).Sinks(sinks).Execute()

Create new function



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moduleName := "moduleName_example" // string | The name of the module to retrieve
    name := "name_example" // string | Name of the function (optional)
    code := os.NewFile(1234, "some_file") // *os.File | File with the code of the function (optional)
    events := []openapiclient.FunctionCreateUpdateEventsInner{*openapiclient.NewFunctionCreateUpdateEventsInner()} // []FunctionCreateUpdateEventsInner | Events that can trigger the function (optional)
    sinks := []openapiclient.FunctionCreateUpdateSinksInner{*openapiclient.NewFunctionCreateUpdateSinksInner()} // []FunctionCreateUpdateSinksInner | Data sinks that receive invocation's results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.CreateFunction(context.Background(), moduleName).Name(name).Code(code).Events(events).Sinks(sinks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.CreateFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** | The name of the module to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the function | 
 **code** | ***os.File** | File with the code of the function | 
 **events** | [**[]FunctionCreateUpdateEventsInner**](FunctionCreateUpdateEventsInner.md) | Events that can trigger the function | 
 **sinks** | [**[]FunctionCreateUpdateSinksInner**](FunctionCreateUpdateSinksInner.md) | Data sinks that receive invocation&#39;s results | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunction

> DeleteFunction(ctx, moduleName, functionName).Execute()

Delete function



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moduleName := "moduleName_example" // string | The name of the module
    functionName := "functionName_example" // string | The name of the function

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.DeleteFunction(context.Background(), moduleName, functionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.DeleteFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** | The name of the module | 
**functionName** | **string** | The name of the function | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeFunction

> InvokeResult InvokeFunction(ctx, moduleName, functionName).InvokeInput(invokeInput).Execute()

Invoke function



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moduleName := "moduleName_example" // string | The name of the module
    functionName := "functionName_example" // string | The name of the function
    invokeInput := *openapiclient.NewInvokeInput() // InvokeInput | Function input

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.InvokeFunction(context.Background(), moduleName, functionName).InvokeInput(invokeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.InvokeFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeFunction`: InvokeResult
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.InvokeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** | The name of the module | 
**functionName** | **string** | The name of the function | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **invokeInput** | [**InvokeInput**](InvokeInput.md) | Function input | 

### Return type

[**InvokeResult**](InvokeResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFunctionByName

> SingleFunctionResult ShowFunctionByName(ctx, moduleName, functionName).Execute()

Show function info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moduleName := "moduleName_example" // string | The name of the module to retrieve
    functionName := "functionName_example" // string | The name of the function

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.ShowFunctionByName(context.Background(), moduleName, functionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.ShowFunctionByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowFunctionByName`: SingleFunctionResult
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.ShowFunctionByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** | The name of the module to retrieve | 
**functionName** | **string** | The name of the function | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowFunctionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SingleFunctionResult**](SingleFunctionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> UpdateFunction(ctx, moduleName, functionName).Name(name).Code(code).Events(events).Sinks(sinks).Execute()

Update function



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moduleName := "moduleName_example" // string | The name of the module
    functionName := "functionName_example" // string | The name of the function
    name := "name_example" // string | Name of the function (optional)
    code := os.NewFile(1234, "some_file") // *os.File | File with the code of the function (optional)
    events := []openapiclient.FunctionCreateUpdateEventsInner{*openapiclient.NewFunctionCreateUpdateEventsInner()} // []FunctionCreateUpdateEventsInner | Events that can trigger the function (optional)
    sinks := []openapiclient.FunctionCreateUpdateSinksInner{*openapiclient.NewFunctionCreateUpdateSinksInner()} // []FunctionCreateUpdateSinksInner | Data sinks that receive invocation's results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.UpdateFunction(context.Background(), moduleName, functionName).Name(name).Code(code).Events(events).Sinks(sinks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.UpdateFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** | The name of the module | 
**functionName** | **string** | The name of the function | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | Name of the function | 
 **code** | ***os.File** | File with the code of the function | 
 **events** | [**[]FunctionCreateUpdateEventsInner**](FunctionCreateUpdateEventsInner.md) | Events that can trigger the function | 
 **sinks** | [**[]FunctionCreateUpdateSinksInner**](FunctionCreateUpdateSinksInner.md) | Data sinks that receive invocation&#39;s results | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

