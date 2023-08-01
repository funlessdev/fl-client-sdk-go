# \ModulesAPI

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModule**](ModulesAPI.md#CreateModule) | **Post** /v1/fn | Create
[**DeleteModule**](ModulesAPI.md#DeleteModule) | **Delete** /v1/fn/{module_name} | Delete module
[**ListModules**](ModulesAPI.md#ListModules) | **Get** /v1/fn | List modules
[**ShowModuleByName**](ModulesAPI.md#ShowModuleByName) | **Get** /v1/fn/{module_name} | Show module info
[**UpdateModule**](ModulesAPI.md#UpdateModule) | **Put** /v1/fn/{module_name} | Update module name



## CreateModule

> CreateModule(ctx).ModuleName(moduleName).Execute()

Create



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/funlessdev/fl-client-sdk-go"
)

func main() {
    moduleName := *openapiclient.NewModuleName() // ModuleName | Module to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ModulesAPI.CreateModule(context.Background()).ModuleName(moduleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesAPI.CreateModule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moduleName** | [**ModuleName**](ModuleName.md) | Module to create | 

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
    openapiclient "github.com/funlessdev/fl-client-sdk-go"
)

func main() {
    moduleName := "moduleName_example" // string | The name of the module to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ModulesAPI.DeleteModule(context.Background(), moduleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesAPI.DeleteModule``: %v\n", err)
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

> ModuleNamesResult ListModules(ctx).Execute()

List modules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/funlessdev/fl-client-sdk-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModulesAPI.ListModules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesAPI.ListModules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListModules`: ModuleNamesResult
    fmt.Fprintf(os.Stdout, "Response from `ModulesAPI.ListModules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListModulesRequest struct via the builder pattern


### Return type

[**ModuleNamesResult**](ModuleNamesResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowModuleByName

> SingleModuleResult ShowModuleByName(ctx, moduleName).Execute()

Show module info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/funlessdev/fl-client-sdk-go"
)

func main() {
    moduleName := "moduleName_example" // string | The name of the module to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModulesAPI.ShowModuleByName(context.Background(), moduleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesAPI.ShowModuleByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowModuleByName`: SingleModuleResult
    fmt.Fprintf(os.Stdout, "Response from `ModulesAPI.ShowModuleByName`: %v\n", resp)
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

[**SingleModuleResult**](SingleModuleResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModule

> UpdateModule(ctx, moduleName).ModuleName2(moduleName2).Execute()

Update module name



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/funlessdev/fl-client-sdk-go"
)

func main() {
    moduleName := "moduleName_example" // string | The name of the module to retrieve
    moduleName2 := *openapiclient.NewModuleName() // ModuleName | New module name to use

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ModulesAPI.UpdateModule(context.Background(), moduleName).ModuleName2(moduleName2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModulesAPI.UpdateModule``: %v\n", err)
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

 **moduleName2** | [**ModuleName**](ModuleName.md) | New module name to use | 

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

