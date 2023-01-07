# \ModulesApi

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModule**](ModulesApi.md#CreateModule) | **Post** /v1/fn | Create
[**DeleteModule**](ModulesApi.md#DeleteModule) | **Delete** /v1/fn/{module_name} | Delete module
[**ListModules**](ModulesApi.md#ListModules) | **Get** /v1/fn | List modules
[**ShowModuleByName**](ModulesApi.md#ShowModuleByName) | **Get** /v1/fn/{module_name} | Show module info
[**UpdateModule**](ModulesApi.md#UpdateModule) | **Put** /v1/fn/{module_name} | Update module name



## CreateModule

> CreateModule(ctx).CreateModuleRequest(createModuleRequest).Execute()

Create



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
    createModuleRequest := *openapiclient.NewCreateModuleRequest() // CreateModuleRequest | Module to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModulesApi.CreateModule(context.Background()).CreateModuleRequest(createModuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesApi.CreateModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createModuleRequest** | [**CreateModuleRequest**](CreateModuleRequest.md) | Module to create | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteModule

> DeleteModule(ctx, moduleName).Execute()

Delete module



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModulesApi.DeleteModule(context.Background(), moduleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesApi.DeleteModule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteModuleRequest struct via the builder pattern


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


## ListModules

> ListModules200Response ListModules(ctx).Execute()

List modules



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModulesApi.ListModules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesApi.ListModules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListModules`: ListModules200Response
    fmt.Fprintf(os.Stdout, "Response from `ModulesApi.ListModules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListModulesRequest struct via the builder pattern


### Return type

[**ListModules200Response**](ListModules200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowModuleByName

> ShowModuleByName200Response ShowModuleByName(ctx, moduleName).Execute()

Show module info



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModulesApi.ShowModuleByName(context.Background(), moduleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesApi.ShowModuleByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowModuleByName`: ShowModuleByName200Response
    fmt.Fprintf(os.Stdout, "Response from `ModulesApi.ShowModuleByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** | The name of the module to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowModuleByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShowModuleByName200Response**](ShowModuleByName200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModule

> UpdateModule(ctx, moduleName).CreateModuleRequest(createModuleRequest).Execute()

Update module name



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
    createModuleRequest := *openapiclient.NewCreateModuleRequest() // CreateModuleRequest | New module name to use

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModulesApi.UpdateModule(context.Background(), moduleName).CreateModuleRequest(createModuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesApi.UpdateModule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createModuleRequest** | [**CreateModuleRequest**](CreateModuleRequest.md) | New module name to use | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

