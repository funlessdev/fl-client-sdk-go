# \DefaultApi

All URIs are relative to *http://localhost:4001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePost**](DefaultApi.md#CreatePost) | **Post** /create | Create a new function
[**DeletePost**](DefaultApi.md#DeletePost) | **Post** /delete | Delete a function
[**InvokePost**](DefaultApi.md#InvokePost) | **Post** /invoke | Invoke a function



## CreatePost

> FunctionCreationSuccess CreatePost(ctx).FunctionCreation(functionCreation).Execute()

Create a new function



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
    functionCreation := *openapiclient.NewFunctionCreation() // FunctionCreation | Object containing the function to create, with name, optional namespace, code and runtime image identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreatePost(context.Background()).FunctionCreation(functionCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePost`: FunctionCreationSuccess
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionCreation** | [**FunctionCreation**](FunctionCreation.md) | Object containing the function to create, with name, optional namespace, code and runtime image identifier | 

### Return type

[**FunctionCreationSuccess**](FunctionCreationSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePost

> FunctionDeletionSuccess DeletePost(ctx).FunctionDeletion(functionDeletion).Execute()

Delete a function



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
    functionDeletion := *openapiclient.NewFunctionDeletion() // FunctionDeletion | Object containing the name and namespace of the function to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeletePost(context.Background()).FunctionDeletion(functionDeletion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePost`: FunctionDeletionSuccess
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionDeletion** | [**FunctionDeletion**](FunctionDeletion.md) | Object containing the name and namespace of the function to delete | 

### Return type

[**FunctionDeletionSuccess**](FunctionDeletionSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokePost

> FunctionInvocationSuccess InvokePost(ctx).FunctionInvocation(functionInvocation).Execute()

Invoke a function



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
    functionInvocation := *openapiclient.NewFunctionInvocation() // FunctionInvocation | Object containing the function to invoke, the namespace and optional parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InvokePost(context.Background()).FunctionInvocation(functionInvocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InvokePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokePost`: FunctionInvocationSuccess
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InvokePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvokePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionInvocation** | [**FunctionInvocation**](FunctionInvocation.md) | Object containing the function to invoke, the namespace and optional parameters | 

### Return type

[**FunctionInvocationSuccess**](FunctionInvocationSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

