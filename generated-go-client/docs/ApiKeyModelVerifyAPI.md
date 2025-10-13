# \ApiKeyModelVerifyAPI

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyModelIdPreVerifyPost**](ApiKeyModelVerifyAPI.md#ApiKeyModelIdPreVerifyPost) | **Post** /api-key/model/{id}/pre-verify | Check Valid Source code To Verify Ai Model By Api Key
[**ApiKeyModelIdVerifyCostPost**](ApiKeyModelVerifyAPI.md#ApiKeyModelIdVerifyCostPost) | **Post** /api-key/model/{id}/verify/cost | Calculate Cost To Verify Ai Model By Api Key
[**ApiKeyModelIdVerifyPost**](ApiKeyModelVerifyAPI.md#ApiKeyModelIdVerifyPost) | **Post** /api-key/model/{id}/verify | Verify Ai Model By Api Key
[**ApiKeyModelIdVerifyTaskGet**](ApiKeyModelVerifyAPI.md#ApiKeyModelIdVerifyTaskGet) | **Get** /api-key/model/{id}/verify/task | Get List Verify Model Task By Commit Hash And Status
[**ApiKeyModelVerifyHubTaskIdGet**](ApiKeyModelVerifyAPI.md#ApiKeyModelVerifyHubTaskIdGet) | **Get** /api-key/model/verify/hub/task/{id} | Get Model Versioning By Hub Task Id By Api Key
[**ApiKeyModelVerifyPlatformTaskIdGet**](ApiKeyModelVerifyAPI.md#ApiKeyModelVerifyPlatformTaskIdGet) | **Get** /api-key/model/verify/platform/task/{id} | Get Verify Platform Task By Id By Api Key



## ApiKeyModelIdPreVerifyPost

> ResponseCheckValidToVerifyAiModelResponse ApiKeyModelIdPreVerifyPost(ctx, id).Input(input).XApiKey(xApiKey).Execute()

Check Valid Source code To Verify Ai Model By Api Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Model's id
	input := *openapiclient.NewRequestCheckValidToVerifyAiModelRequest("CommitHash_example", []string{"Platforms_example"}) // RequestCheckValidToVerifyAiModelRequest | Verify Ai Model Request
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVerifyAPI.ApiKeyModelIdPreVerifyPost(context.Background(), id).Input(input).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVerifyAPI.ApiKeyModelIdPreVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdPreVerifyPost`: ResponseCheckValidToVerifyAiModelResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVerifyAPI.ApiKeyModelIdPreVerifyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdPreVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **input** | [**RequestCheckValidToVerifyAiModelRequest**](RequestCheckValidToVerifyAiModelRequest.md) | Verify Ai Model Request | 
 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseCheckValidToVerifyAiModelResponse**](ResponseCheckValidToVerifyAiModelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVerifyCostPost

> ResponseEstimateCostResponse ApiKeyModelIdVerifyCostPost(ctx, id).Input(input).XApiKey(xApiKey).Execute()

Calculate Cost To Verify Ai Model By Api Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Model's id
	input := *openapiclient.NewRequestCalculateCostToVerifyAiModelRequest("CommitHash_example", []string{"Platforms_example"}) // RequestCalculateCostToVerifyAiModelRequest | Verify Ai Model Request
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyCostPost(context.Background(), id).Input(input).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyCostPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdVerifyCostPost`: ResponseEstimateCostResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyCostPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVerifyCostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **input** | [**RequestCalculateCostToVerifyAiModelRequest**](RequestCalculateCostToVerifyAiModelRequest.md) | Verify Ai Model Request | 
 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseEstimateCostResponse**](ResponseEstimateCostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVerifyPost

> ResponseVerifyAiModelResponse ApiKeyModelIdVerifyPost(ctx, id).Input(input).XApiKey(xApiKey).Execute()

Verify Ai Model By Api Key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Model's id
	input := *openapiclient.NewRequestVerifyAiModelRequest([]string{"Platforms_example"}) // RequestVerifyAiModelRequest | Verify Ai Model Request
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyPost(context.Background(), id).Input(input).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdVerifyPost`: ResponseVerifyAiModelResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **input** | [**RequestVerifyAiModelRequest**](RequestVerifyAiModelRequest.md) | Verify Ai Model Request | 
 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseVerifyAiModelResponse**](ResponseVerifyAiModelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVerifyTaskGet

> ResponseModelVersioningGroupLiteListResponse ApiKeyModelIdVerifyTaskGet(ctx, id).CommitHash(commitHash).XApiKey(xApiKey).VerifyStatus(verifyStatus).Execute()

Get List Verify Model Task By Commit Hash And Status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Model's id
	commitHash := "commitHash_example" // string | 
	xApiKey := "xApiKey_example" // string | api-key (optional)
	verifyStatus := "verifyStatus_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyTaskGet(context.Background(), id).CommitHash(commitHash).XApiKey(xApiKey).VerifyStatus(verifyStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdVerifyTaskGet`: ResponseModelVersioningGroupLiteListResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVerifyAPI.ApiKeyModelIdVerifyTaskGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVerifyTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitHash** | **string** |  | 
 **xApiKey** | **string** | api-key | 
 **verifyStatus** | **string** |  | 

### Return type

[**ResponseModelVersioningGroupLiteListResponse**](ResponseModelVersioningGroupLiteListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelVerifyHubTaskIdGet

> ResponseModelVersioningResponse ApiKeyModelVerifyHubTaskIdGet(ctx, id).XApiKey(xApiKey).Execute()

Get Model Versioning By Hub Task Id By Api Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Hub Task's id
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVerifyAPI.ApiKeyModelVerifyHubTaskIdGet(context.Background(), id).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVerifyAPI.ApiKeyModelVerifyHubTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelVerifyHubTaskIdGet`: ResponseModelVersioningResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVerifyAPI.ApiKeyModelVerifyHubTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Hub Task&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelVerifyHubTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseModelVersioningResponse**](ResponseModelVersioningResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelVerifyPlatformTaskIdGet

> ResponseQueueTaskResponse ApiKeyModelVerifyPlatformTaskIdGet(ctx, id).XApiKey(xApiKey).Execute()

Get Verify Platform Task By Id By Api Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Task's Id
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVerifyAPI.ApiKeyModelVerifyPlatformTaskIdGet(context.Background(), id).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVerifyAPI.ApiKeyModelVerifyPlatformTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelVerifyPlatformTaskIdGet`: ResponseQueueTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVerifyAPI.ApiKeyModelVerifyPlatformTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Task&#39;s Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelVerifyPlatformTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseQueueTaskResponse**](ResponseQueueTaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

