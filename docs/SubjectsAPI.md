# \SubjectsAPI

All URIs are relative to *http://localhost:4000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubject**](SubjectsAPI.md#CreateSubject) | **Post** /v1/admin/subjects | Create subject
[**ListSubjects**](SubjectsAPI.md#ListSubjects) | **Get** /v1/admin/subjects | List all subjects



## CreateSubject

> SingleSubjectResult CreateSubject(ctx).SubjectName(subjectName).Execute()

Create subject



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
	subjectName := *openapiclient.NewSubjectName() // SubjectName | Subject to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectsAPI.CreateSubject(context.Background()).SubjectName(subjectName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectsAPI.CreateSubject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubject`: SingleSubjectResult
	fmt.Fprintf(os.Stdout, "Response from `SubjectsAPI.CreateSubject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectName** | [**SubjectName**](SubjectName.md) | Subject to create | 

### Return type

[**SingleSubjectResult**](SingleSubjectResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubjects

> ModuleNamesResult ListSubjects(ctx).Execute()

List all subjects



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
	resp, r, err := apiClient.SubjectsAPI.ListSubjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectsAPI.ListSubjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubjects`: ModuleNamesResult
	fmt.Fprintf(os.Stdout, "Response from `SubjectsAPI.ListSubjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectsRequest struct via the builder pattern


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

