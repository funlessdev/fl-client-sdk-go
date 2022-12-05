# \DefaultApi

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FnCreatePost**](DefaultApi.md#V1FnCreatePost) | **Post** /v1/fn/create | Create a new function
[**V1FnDeleteDelete**](DefaultApi.md#V1FnDeleteDelete) | **Delete** /v1/fn/delete | Delete a function
[**V1FnInvokePost**](DefaultApi.md#V1FnInvokePost) | **Post** /v1/fn/invoke | Invoke a function
[**V1FnListFnModuleGet**](DefaultApi.md#V1FnListFnModuleGet) | **Get** /v1/fn/list/{fnModule} | List functions



## V1FnCreatePost

> FunctionCreationSuccess V1FnCreatePost(ctx).Name(name).Module(module).Code(code).Execute()

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
    name := "name_example" // string |  (optional)
    module := "module_example" // string |  (optional)
    code := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1FnCreatePost(context.Background()).Name(name).Module(module).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1FnCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1FnCreatePost`: FunctionCreationSuccess
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1FnCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FnCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **module** | **string** |  | 
 **code** | ***os.File** |  | 

### Return type

[**FunctionCreationSuccess**](FunctionCreationSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1FnDeleteDelete

> FunctionDeletionSuccess V1FnDeleteDelete(ctx).FunctionDeletion(functionDeletion).Execute()

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
    functionDeletion := *openapiclient.NewFunctionDeletion() // FunctionDeletion | Object containing the function's name and module to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1FnDeleteDelete(context.Background()).FunctionDeletion(functionDeletion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1FnDeleteDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1FnDeleteDelete`: FunctionDeletionSuccess
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1FnDeleteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FnDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionDeletion** | [**FunctionDeletion**](FunctionDeletion.md) | Object containing the function&#39;s name and module to delete | 

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


## V1FnInvokePost

> FunctionInvocationSuccess V1FnInvokePost(ctx).FunctionInvocation(functionInvocation).Execute()

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
    functionInvocation := *openapiclient.NewFunctionInvocation() // FunctionInvocation | Object containing the function to invoke, the module and optional parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1FnInvokePost(context.Background()).FunctionInvocation(functionInvocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1FnInvokePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1FnInvokePost`: FunctionInvocationSuccess
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1FnInvokePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FnInvokePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionInvocation** | [**FunctionInvocation**](FunctionInvocation.md) | Object containing the function to invoke, the module and optional parameters | 

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


## V1FnListFnModuleGet

> FunctionListSuccess V1FnListFnModuleGet(ctx, fnModule).Execute()

List functions



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
    fnModule := "fnModule_example" // string | Module of the listed functions

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.V1FnListFnModuleGet(context.Background(), fnModule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1FnListFnModuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1FnListFnModuleGet`: FunctionListSuccess
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1FnListFnModuleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fnModule** | **string** | Module of the listed functions | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FnListFnModuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionListSuccess**](FunctionListSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

