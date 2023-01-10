# {{classname}}

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](FunctionsApi.md#CreateFunction) | **Post** /v1/fn/{module_name} | Create new function
[**DeleteFunction**](FunctionsApi.md#DeleteFunction) | **Delete** /v1/fn/{module_name}/{function_name} | Delete function
[**InvokeFunction**](FunctionsApi.md#InvokeFunction) | **Post** /v1/fn/{module_name}/{function_name} | Invoke function
[**ShowFunctionByName**](FunctionsApi.md#ShowFunctionByName) | **Get** /v1/fn/{module_name}/{function_name} | Show function info
[**UpdateFunction**](FunctionsApi.md#UpdateFunction) | **Put** /v1/fn/{module_name}/{function_name} | Update function

# **CreateFunction**
> CreateFunction(ctx, name, code, events, moduleName)
Create new function

Create a new function in the specified module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **code** | ***os.File*****os.File**|  | 
  **events** | [**[]V1fnmoduleNameEvents**](V1fnmoduleNameEvents.md)|  | 
  **moduleName** | **string**| The name of the module to retrieve | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFunction**
> DeleteFunction(ctx, moduleName, functionName)
Delete function

Delete single function in module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**| The name of the module | 
  **functionName** | **string**| The name of the function | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokeFunction**
> InlineResponse2003 InvokeFunction(ctx, body, moduleName, functionName)
Invoke function

Invoke function

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModuleNameFunctionNameBody1**](ModuleNameFunctionNameBody1.md)| Function input | 
  **moduleName** | **string**| The name of the module | 
  **functionName** | **string**| The name of the function | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFunctionByName**
> InlineResponse2002 ShowFunctionByName(ctx, moduleName, functionName)
Show function info

Get function data (name, module name, size of code)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**| The name of the module to retrieve | 
  **functionName** | **string**| The name of the function | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFunction**
> UpdateFunction(ctx, name, code, events, moduleName, functionName)
Update function

Update function

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **code** | ***os.File*****os.File**|  | 
  **events** | [**[]V1fnmoduleNameEvents**](V1fnmoduleNameEvents.md)|  | 
  **moduleName** | **string**| The name of the module | 
  **functionName** | **string**| The name of the function | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

