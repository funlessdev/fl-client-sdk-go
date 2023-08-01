# \AppAPI

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppAPI.md#CreateApp) | **Post** /v1/app | Create new APP script
[**DeleteApp**](AppAPI.md#DeleteApp) | **Delete** /v1/app/{app_name} | Delete APP
[**ListApp**](AppAPI.md#ListApp) | **Get** /v1/app | List current APP scripts
[**ShowAppByName**](AppAPI.md#ShowAppByName) | **Get** /v1/app/{app_name} | Show APP info



## CreateApp

> CreateApp(ctx).Name(name).Code(code).Events(events).Sinks(sinks).Execute()

Create new APP script



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
    name := "name_example" // string | Name of the function (optional)
    code := os.NewFile(1234, "some_file") // *os.File | File with the code of the function (optional)
    events := []openapiclient.FunctionCreateUpdateEventsInner{*openapiclient.NewFunctionCreateUpdateEventsInner()} // []FunctionCreateUpdateEventsInner | Events that can trigger the function (optional)
    sinks := []openapiclient.FunctionCreateUpdateSinksInner{*openapiclient.NewFunctionCreateUpdateSinksInner()} // []FunctionCreateUpdateSinksInner | Data sinks that receive invocation's results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppAPI.CreateApp(context.Background()).Name(name).Code(code).Events(events).Sinks(sinks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.CreateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


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


## DeleteApp

> DeleteApp(ctx, moduleName).Execute()

Delete APP



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
    r, err := apiClient.AppAPI.DeleteApp(context.Background(), moduleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.DeleteApp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


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


## ListApp

> ModuleNamesResult ListApp(ctx).Execute()

List current APP scripts



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
    resp, r, err := apiClient.AppAPI.ListApp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.ListApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApp`: ModuleNamesResult
    fmt.Fprintf(os.Stdout, "Response from `AppAPI.ListApp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppRequest struct via the builder pattern


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


## ShowAppByName

> SingleAppResult ShowAppByName(ctx, moduleName).Execute()

Show APP info



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
    resp, r, err := apiClient.AppAPI.ShowAppByName(context.Background(), moduleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.ShowAppByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowAppByName`: SingleAppResult
    fmt.Fprintf(os.Stdout, "Response from `AppAPI.ShowAppByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** | The name of the module to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAppByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleAppResult**](SingleAppResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

