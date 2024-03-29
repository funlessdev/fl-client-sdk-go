# Go API client for openapi

The API for the FunLess Platform

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.5
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/funlessdev/fl-client-sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
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
*AppAPI* | [**CreateApp**](docs/AppAPI.md#createapp) | **Post** /v1/app | Create new APP script
*AppAPI* | [**DeleteApp**](docs/AppAPI.md#deleteapp) | **Delete** /v1/app/{app_name} | Delete APP
*AppAPI* | [**ListApp**](docs/AppAPI.md#listapp) | **Get** /v1/app | List current APP scripts
*AppAPI* | [**ShowAppByName**](docs/AppAPI.md#showappbyname) | **Get** /v1/app/{app_name} | Show APP info
*FunctionsAPI* | [**CreateFunction**](docs/FunctionsAPI.md#createfunction) | **Post** /v1/fn/{module_name} | Create new function
*FunctionsAPI* | [**DeleteFunction**](docs/FunctionsAPI.md#deletefunction) | **Delete** /v1/fn/{module_name}/{function_name} | Delete function
*FunctionsAPI* | [**InvokeFunction**](docs/FunctionsAPI.md#invokefunction) | **Post** /v1/fn/{module_name}/{function_name} | Invoke function
*FunctionsAPI* | [**ShowFunctionByName**](docs/FunctionsAPI.md#showfunctionbyname) | **Get** /v1/fn/{module_name}/{function_name} | Show function info
*FunctionsAPI* | [**UpdateFunction**](docs/FunctionsAPI.md#updatefunction) | **Put** /v1/fn/{module_name}/{function_name} | Update function
*ModulesAPI* | [**CreateModule**](docs/ModulesAPI.md#createmodule) | **Post** /v1/fn | Create
*ModulesAPI* | [**DeleteModule**](docs/ModulesAPI.md#deletemodule) | **Delete** /v1/fn/{module_name} | Delete module
*ModulesAPI* | [**ListModules**](docs/ModulesAPI.md#listmodules) | **Get** /v1/fn | List modules
*ModulesAPI* | [**ShowModuleByName**](docs/ModulesAPI.md#showmodulebyname) | **Get** /v1/fn/{module_name} | Show module info
*ModulesAPI* | [**UpdateModule**](docs/ModulesAPI.md#updatemodule) | **Put** /v1/fn/{module_name} | Update module name
*SubjectsAPI* | [**CreateSubject**](docs/SubjectsAPI.md#createsubject) | **Post** /v1/admin/subjects | Create subject
*SubjectsAPI* | [**ListSubjects**](docs/SubjectsAPI.md#listsubjects) | **Get** /v1/admin/subjects | List all subjects


## Documentation For Models

 - [AppCreateUpdate](docs/AppCreateUpdate.md)
 - [ErrorErrors](docs/ErrorErrors.md)
 - [FunctionCreateUpdate](docs/FunctionCreateUpdate.md)
 - [FunctionCreateUpdateEventsInner](docs/FunctionCreateUpdateEventsInner.md)
 - [FunctionCreateUpdateSinksInner](docs/FunctionCreateUpdateSinksInner.md)
 - [InvokeInput](docs/InvokeInput.md)
 - [InvokeResult](docs/InvokeResult.md)
 - [MixedResults](docs/MixedResults.md)
 - [MixedResultsData](docs/MixedResultsData.md)
 - [MixedResultsDataEventsInner](docs/MixedResultsDataEventsInner.md)
 - [MixedResultsDataEventsMetadata](docs/MixedResultsDataEventsMetadata.md)
 - [MixedResultsDataSinksInner](docs/MixedResultsDataSinksInner.md)
 - [MixedResultsDataSinksMetadata](docs/MixedResultsDataSinksMetadata.md)
 - [ModelError](docs/ModelError.md)
 - [ModuleName](docs/ModuleName.md)
 - [ModuleNamesResult](docs/ModuleNamesResult.md)
 - [SingleAppResult](docs/SingleAppResult.md)
 - [SingleAppResultData](docs/SingleAppResultData.md)
 - [SingleFunctionResult](docs/SingleFunctionResult.md)
 - [SingleModuleResult](docs/SingleModuleResult.md)
 - [SingleModuleResultData](docs/SingleModuleResultData.md)
 - [SingleSubjectResult](docs/SingleSubjectResult.md)
 - [SingleSubjectResultData](docs/SingleSubjectResultData.md)
 - [SubjectName](docs/SubjectName.md)
 - [SubjectNameSubject](docs/SubjectNameSubject.md)


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



