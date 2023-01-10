# Go API client for openapi

The API for the FunLess Platform

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.5
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/funlessdev/fl-client-sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:4000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FunctionsApi* | [**CreateFunction**](docs/FunctionsApi.md#createfunction) | **Post** /v1/fn/{module_name} | Create new function
*FunctionsApi* | [**DeleteFunction**](docs/FunctionsApi.md#deletefunction) | **Delete** /v1/fn/{module_name}/{function_name} | Delete function
*FunctionsApi* | [**InvokeFunction**](docs/FunctionsApi.md#invokefunction) | **Post** /v1/fn/{module_name}/{function_name} | Invoke function
*FunctionsApi* | [**ShowFunctionByName**](docs/FunctionsApi.md#showfunctionbyname) | **Get** /v1/fn/{module_name}/{function_name} | Show function info
*FunctionsApi* | [**UpdateFunction**](docs/FunctionsApi.md#updatefunction) | **Put** /v1/fn/{module_name}/{function_name} | Update function
*ModulesApi* | [**CreateModule**](docs/ModulesApi.md#createmodule) | **Post** /v1/fn | Create
*ModulesApi* | [**DeleteModule**](docs/ModulesApi.md#deletemodule) | **Delete** /v1/fn/{module_name} | Delete module
*ModulesApi* | [**ListModules**](docs/ModulesApi.md#listmodules) | **Get** /v1/fn | List modules
*ModulesApi* | [**ShowModuleByName**](docs/ModulesApi.md#showmodulebyname) | **Get** /v1/fn/{module_name} | Show module info
*ModulesApi* | [**UpdateModule**](docs/ModulesApi.md#updatemodule) | **Put** /v1/fn/{module_name} | Update module name


## Documentation For Models

 - [CreateFunction207Response](docs/CreateFunction207Response.md)
 - [CreateFunction207ResponseData](docs/CreateFunction207ResponseData.md)
 - [CreateFunction207ResponseDataEventsInner](docs/CreateFunction207ResponseDataEventsInner.md)
 - [CreateFunction207ResponseDataMetadata](docs/CreateFunction207ResponseDataMetadata.md)
 - [CreateModuleRequest](docs/CreateModuleRequest.md)
 - [FunctionCreateUpdate](docs/FunctionCreateUpdate.md)
 - [InvokeFunction200Response](docs/InvokeFunction200Response.md)
 - [InvokeFunctionRequest](docs/InvokeFunctionRequest.md)
 - [InvokeInput](docs/InvokeInput.md)
 - [InvokeResult](docs/InvokeResult.md)
 - [ListModules200Response](docs/ListModules200Response.md)
 - [ListModulesDefaultResponse](docs/ListModulesDefaultResponse.md)
 - [ListModulesDefaultResponseErrors](docs/ListModulesDefaultResponseErrors.md)
 - [MixedEventResults](docs/MixedEventResults.md)
 - [ModelError](docs/ModelError.md)
 - [ModuleName](docs/ModuleName.md)
 - [ModuleNamesResult](docs/ModuleNamesResult.md)
 - [ShowFunctionByName200Response](docs/ShowFunctionByName200Response.md)
 - [ShowModuleByName200Response](docs/ShowModuleByName200Response.md)
 - [ShowModuleByName200ResponseData](docs/ShowModuleByName200ResponseData.md)
 - [SingleFunctionResult](docs/SingleFunctionResult.md)
 - [SingleModuleResult](docs/SingleModuleResult.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



