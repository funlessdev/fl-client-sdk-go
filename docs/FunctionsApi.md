# \FunctionsApi

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](FunctionsApi.md#CreateFunction) | **Post** /v1/fn/{module_name} | Create new function
[**DeleteFunction**](FunctionsApi.md#DeleteFunction) | **Delete** /v1/fn/{module_name}/{function_name} | Delete function
[**InvokeFunction**](FunctionsApi.md#InvokeFunction) | **Post** /v1/fn/{module_name}/{function_name} | Invoke function
[**ShowFunctionByName**](FunctionsApi.md#ShowFunctionByName) | **Get** /v1/fn/{module_name}/{function_name} | Show function info
[**UpdateFunction**](FunctionsApi.md#UpdateFunction) | **Put** /v1/fn/{module_name}/{function_name} | Update function code



## CreateFunction

> CreateFunction(ctx).CreateFunctionRequest(createFunctionRequest).Execute()

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
    createFunctionRequest := *openapiclient.NewCreateFunctionRequest() // CreateFunctionRequest | Function code to upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.CreateFunction(context.Background()).CreateFunctionRequest(createFunctionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.CreateFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFunctionRequest** | [**CreateFunctionRequest**](CreateFunctionRequest.md) | Function code to upload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeFunction

> InvokeFunction200Response InvokeFunction(ctx, moduleName, functionName).InvokeFunctionRequest(invokeFunctionRequest).Execute()

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
    invokeFunctionRequest := *openapiclient.NewInvokeFunctionRequest() // InvokeFunctionRequest | Function input

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.InvokeFunction(context.Background(), moduleName, functionName).InvokeFunctionRequest(invokeFunctionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.InvokeFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeFunction`: InvokeFunction200Response
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


 **invokeFunctionRequest** | [**InvokeFunctionRequest**](InvokeFunctionRequest.md) | Function input | 

### Return type

[**InvokeFunction200Response**](InvokeFunction200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFunctionByName

> ShowFunctionByName200Response ShowFunctionByName(ctx, functionName).Execute()

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
    functionName := "functionName_example" // string | The name of the function

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.ShowFunctionByName(context.Background(), functionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.ShowFunctionByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowFunctionByName`: ShowFunctionByName200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.ShowFunctionByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionName** | **string** | The name of the function | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowFunctionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShowFunctionByName200Response**](ShowFunctionByName200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> UpdateFunction(ctx, moduleName, functionName).Body(body).Execute()

Update function code



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
    body := Object(987) // Object | New function code to use

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsApi.UpdateFunction(context.Background(), moduleName, functionName).Body(body).Execute()
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


 **body** | **Object** | New function code to use | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

