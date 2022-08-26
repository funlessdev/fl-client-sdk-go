# {{classname}}

All URIs are relative to *http://localhost:4001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePost**](DefaultApi.md#CreatePost) | **Post** /create | Create a new function
[**DeletePost**](DefaultApi.md#DeletePost) | **Post** /delete | Delete a function
[**InvokePost**](DefaultApi.md#InvokePost) | **Post** /invoke | Invoke a function

# **CreatePost**
> FunctionCreationSuccess CreatePost(ctx, body)
Create a new function

Creates the given function, extracting the parameters from the POST body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FunctionCreation**](FunctionCreation.md)| Object containing the function to create, with name, optional namespace, code and runtime image identifier | 

### Return type

[**FunctionCreationSuccess**](function_creation_success.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePost**
> FunctionDeletionSuccess DeletePost(ctx, body)
Delete a function

Deletes the function with the given name and namespace, extracted from the POST body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FunctionDeletion**](FunctionDeletion.md)| Object containing the name and namespace of the function to delete | 

### Return type

[**FunctionDeletionSuccess**](function_deletion_success.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePost**
> FunctionInvocationSuccess InvokePost(ctx, body)
Invoke a function

Invokes the specified function from the given namespace with optional parameters from the POST body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FunctionInvocation**](FunctionInvocation.md)| Object containing the function to invoke, the namespace and optional parameters | 

### Return type

[**FunctionInvocationSuccess**](function_invocation_success.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

