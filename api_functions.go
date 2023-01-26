/*
FunLess Platfom API

The API for the FunLess Platform

API version: 0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// FunctionsApiService FunctionsApi service
type FunctionsApiService service

type ApiCreateFunctionRequest struct {
	ctx context.Context
	ApiService *FunctionsApiService
	moduleName string
	name *string
	code *os.File
	events *[]FunctionCreateUpdateEventsInner
	sinks *[]FunctionCreateUpdateSinksInner
}

// Name of the function
func (r ApiCreateFunctionRequest) Name(name string) ApiCreateFunctionRequest {
	r.name = &name
	return r
}

// File with the code of the function
func (r ApiCreateFunctionRequest) Code(code *os.File) ApiCreateFunctionRequest {
	r.code = code
	return r
}

// Events that can trigger the function
func (r ApiCreateFunctionRequest) Events(events []FunctionCreateUpdateEventsInner) ApiCreateFunctionRequest {
	r.events = &events
	return r
}

// Data sinks that receive invocation&#39;s results
func (r ApiCreateFunctionRequest) Sinks(sinks []FunctionCreateUpdateSinksInner) ApiCreateFunctionRequest {
	r.sinks = &sinks
	return r
}

func (r ApiCreateFunctionRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateFunctionExecute(r)
}

/*
CreateFunction Create new function

Create a new function in the specified module

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param moduleName The name of the module to retrieve
 @return ApiCreateFunctionRequest
*/
func (a *FunctionsApiService) CreateFunction(ctx context.Context, moduleName string) ApiCreateFunctionRequest {
	return ApiCreateFunctionRequest{
		ApiService: a,
		ctx: ctx,
		moduleName: moduleName,
	}
}

// Execute executes the request
func (a *FunctionsApiService) CreateFunctionExecute(r ApiCreateFunctionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.CreateFunction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fn/{module_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"module_name"+"}", url.PathEscape(parameterValueToString(r.moduleName, "moduleName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		parameterAddToQuery(localVarFormParams, "name", r.name, "")
	}
	var codeLocalVarFormFileName string
	var codeLocalVarFileName     string
	var codeLocalVarFileBytes    []byte

	codeLocalVarFormFileName = "code"


	codeLocalVarFile := r.code

	if codeLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(codeLocalVarFile)

		codeLocalVarFileBytes = fbs
		codeLocalVarFileName = codeLocalVarFile.Name()
		codeLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: codeLocalVarFileBytes, fileName: codeLocalVarFileName, formFileName: codeLocalVarFormFileName})
	}
	if r.events != nil {
		parameterAddToQuery(localVarFormParams, "events", r.events, "csv")
	}
	if r.sinks != nil {
		parameterAddToQuery(localVarFormParams, "sinks", r.sinks, "csv")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteFunctionRequest struct {
	ctx context.Context
	ApiService *FunctionsApiService
	moduleName string
	functionName string
}

func (r ApiDeleteFunctionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFunctionExecute(r)
}

/*
DeleteFunction Delete function

Delete single function in module

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param moduleName The name of the module
 @param functionName The name of the function
 @return ApiDeleteFunctionRequest
*/
func (a *FunctionsApiService) DeleteFunction(ctx context.Context, moduleName string, functionName string) ApiDeleteFunctionRequest {
	return ApiDeleteFunctionRequest{
		ApiService: a,
		ctx: ctx,
		moduleName: moduleName,
		functionName: functionName,
	}
}

// Execute executes the request
func (a *FunctionsApiService) DeleteFunctionExecute(r ApiDeleteFunctionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.DeleteFunction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fn/{module_name}/{function_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"module_name"+"}", url.PathEscape(parameterValueToString(r.moduleName, "moduleName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"function_name"+"}", url.PathEscape(parameterValueToString(r.functionName, "functionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInvokeFunctionRequest struct {
	ctx context.Context
	ApiService *FunctionsApiService
	moduleName string
	functionName string
	invokeInput *InvokeInput
}

// Function input
func (r ApiInvokeFunctionRequest) InvokeInput(invokeInput InvokeInput) ApiInvokeFunctionRequest {
	r.invokeInput = &invokeInput
	return r
}

func (r ApiInvokeFunctionRequest) Execute() (*InvokeResult, *http.Response, error) {
	return r.ApiService.InvokeFunctionExecute(r)
}

/*
InvokeFunction Invoke function

Invoke function

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param moduleName The name of the module
 @param functionName The name of the function
 @return ApiInvokeFunctionRequest
*/
func (a *FunctionsApiService) InvokeFunction(ctx context.Context, moduleName string, functionName string) ApiInvokeFunctionRequest {
	return ApiInvokeFunctionRequest{
		ApiService: a,
		ctx: ctx,
		moduleName: moduleName,
		functionName: functionName,
	}
}

// Execute executes the request
//  @return InvokeResult
func (a *FunctionsApiService) InvokeFunctionExecute(r ApiInvokeFunctionRequest) (*InvokeResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvokeResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.InvokeFunction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fn/{module_name}/{function_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"module_name"+"}", url.PathEscape(parameterValueToString(r.moduleName, "moduleName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"function_name"+"}", url.PathEscape(parameterValueToString(r.functionName, "functionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.invokeInput == nil {
		return localVarReturnValue, nil, reportError("invokeInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.invokeInput
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiShowFunctionByNameRequest struct {
	ctx context.Context
	ApiService *FunctionsApiService
	moduleName string
	functionName string
}

func (r ApiShowFunctionByNameRequest) Execute() (*SingleFunctionResult, *http.Response, error) {
	return r.ApiService.ShowFunctionByNameExecute(r)
}

/*
ShowFunctionByName Show function info

Get function data (name, module name, size of code)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param moduleName The name of the module to retrieve
 @param functionName The name of the function
 @return ApiShowFunctionByNameRequest
*/
func (a *FunctionsApiService) ShowFunctionByName(ctx context.Context, moduleName string, functionName string) ApiShowFunctionByNameRequest {
	return ApiShowFunctionByNameRequest{
		ApiService: a,
		ctx: ctx,
		moduleName: moduleName,
		functionName: functionName,
	}
}

// Execute executes the request
//  @return SingleFunctionResult
func (a *FunctionsApiService) ShowFunctionByNameExecute(r ApiShowFunctionByNameRequest) (*SingleFunctionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SingleFunctionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.ShowFunctionByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fn/{module_name}/{function_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"module_name"+"}", url.PathEscape(parameterValueToString(r.moduleName, "moduleName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"function_name"+"}", url.PathEscape(parameterValueToString(r.functionName, "functionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateFunctionRequest struct {
	ctx context.Context
	ApiService *FunctionsApiService
	moduleName string
	functionName string
	name *string
	code *os.File
	events *[]FunctionCreateUpdateEventsInner
	sinks *[]FunctionCreateUpdateSinksInner
}

// Name of the function
func (r ApiUpdateFunctionRequest) Name(name string) ApiUpdateFunctionRequest {
	r.name = &name
	return r
}

// File with the code of the function
func (r ApiUpdateFunctionRequest) Code(code *os.File) ApiUpdateFunctionRequest {
	r.code = code
	return r
}

// Events that can trigger the function
func (r ApiUpdateFunctionRequest) Events(events []FunctionCreateUpdateEventsInner) ApiUpdateFunctionRequest {
	r.events = &events
	return r
}

// Data sinks that receive invocation&#39;s results
func (r ApiUpdateFunctionRequest) Sinks(sinks []FunctionCreateUpdateSinksInner) ApiUpdateFunctionRequest {
	r.sinks = &sinks
	return r
}

func (r ApiUpdateFunctionRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateFunctionExecute(r)
}

/*
UpdateFunction Update function

Update function

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param moduleName The name of the module
 @param functionName The name of the function
 @return ApiUpdateFunctionRequest
*/
func (a *FunctionsApiService) UpdateFunction(ctx context.Context, moduleName string, functionName string) ApiUpdateFunctionRequest {
	return ApiUpdateFunctionRequest{
		ApiService: a,
		ctx: ctx,
		moduleName: moduleName,
		functionName: functionName,
	}
}

// Execute executes the request
func (a *FunctionsApiService) UpdateFunctionExecute(r ApiUpdateFunctionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.UpdateFunction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fn/{module_name}/{function_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"module_name"+"}", url.PathEscape(parameterValueToString(r.moduleName, "moduleName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"function_name"+"}", url.PathEscape(parameterValueToString(r.functionName, "functionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		parameterAddToQuery(localVarFormParams, "name", r.name, "")
	}
	var codeLocalVarFormFileName string
	var codeLocalVarFileName     string
	var codeLocalVarFileBytes    []byte

	codeLocalVarFormFileName = "code"


	codeLocalVarFile := r.code

	if codeLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(codeLocalVarFile)

		codeLocalVarFileBytes = fbs
		codeLocalVarFileName = codeLocalVarFile.Name()
		codeLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: codeLocalVarFileBytes, fileName: codeLocalVarFileName, formFileName: codeLocalVarFormFileName})
	}
	if r.events != nil {
		parameterAddToQuery(localVarFormParams, "events", r.events, "csv")
	}
	if r.sinks != nil {
		parameterAddToQuery(localVarFormParams, "sinks", r.sinks, "csv")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
