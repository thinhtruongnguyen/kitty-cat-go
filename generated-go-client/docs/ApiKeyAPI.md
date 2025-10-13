# \ApiKeyAPI

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyBalanceGet**](ApiKeyAPI.md#ApiKeyBalanceGet) | **Get** /api-key/balance | Get Api Key Balance
[**ApiKeyPermissionGet**](ApiKeyAPI.md#ApiKeyPermissionGet) | **Get** /api-key/permission | Get Api Key Permission
[**ApiKeyStatisticsPost**](ApiKeyAPI.md#ApiKeyStatisticsPost) | **Post** /api-key/statistics | Get Api Key Statistics
[**ApiKeyTaskHistoriesGet**](ApiKeyAPI.md#ApiKeyTaskHistoriesGet) | **Get** /api-key/task/histories | Get Tasks Histories
[**ApiKeyTaskIdCancelDelete**](ApiKeyAPI.md#ApiKeyTaskIdCancelDelete) | **Delete** /api-key/task/{id}/cancel | Cancel Task By Api Key
[**ApiKeyTaskIdResultGet**](ApiKeyAPI.md#ApiKeyTaskIdResultGet) | **Get** /api-key/task/{id}/result | Get Task Result
[**ApiKeyTaskPost**](ApiKeyAPI.md#ApiKeyTaskPost) | **Post** /api-key/task | Distribute Task (Api-Key)



## ApiKeyBalanceGet

> ResponseWalletWithAddressResponse ApiKeyBalanceGet(ctx).XApiKey(xApiKey).Execute()

Get Api Key Balance

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
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.ApiKeyBalanceGet(context.Background()).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.ApiKeyBalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyBalanceGet`: ResponseWalletWithAddressResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.ApiKeyBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseWalletWithAddressResponse**](ResponseWalletWithAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyPermissionGet

> ResponseGetApiKeyPermissionResponse ApiKeyPermissionGet(ctx).XApiKey(xApiKey).Execute()

Get Api Key Permission

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
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.ApiKeyPermissionGet(context.Background()).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.ApiKeyPermissionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyPermissionGet`: ResponseGetApiKeyPermissionResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.ApiKeyPermissionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyPermissionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseGetApiKeyPermissionResponse**](ResponseGetApiKeyPermissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyStatisticsPost

> ResponseGetTaskStatisticsResponse ApiKeyStatisticsPost(ctx).Input(input).XApiKey(xApiKey).Execute()

Get Api Key Statistics

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
	input := *openapiclient.NewRequestGetApiKeyStatisticsByModelIdRequest() // RequestGetApiKeyStatisticsByModelIdRequest | Get Api Key Statistics Request
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.ApiKeyStatisticsPost(context.Background()).Input(input).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.ApiKeyStatisticsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyStatisticsPost`: ResponseGetTaskStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.ApiKeyStatisticsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyStatisticsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | [**RequestGetApiKeyStatisticsByModelIdRequest**](RequestGetApiKeyStatisticsByModelIdRequest.md) | Get Api Key Statistics Request | 
 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseGetTaskStatisticsResponse**](ResponseGetTaskStatisticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyTaskHistoriesGet

> ResponseApiKeyHistoryListResponse ApiKeyTaskHistoriesGet(ctx).XApiKey(xApiKey).Limit(limit).Offset(offset).Execute()

Get Tasks Histories

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
	xApiKey := "xApiKey_example" // string | api-key (optional)
	limit := int32(56) // int32 |  (optional) (default to 10)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.ApiKeyTaskHistoriesGet(context.Background()).XApiKey(xApiKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.ApiKeyTaskHistoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyTaskHistoriesGet`: ResponseApiKeyHistoryListResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.ApiKeyTaskHistoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyTaskHistoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | api-key | 
 **limit** | **int32** |  | [default to 10]
 **offset** | **int32** |  | [default to 0]

### Return type

[**ResponseApiKeyHistoryListResponse**](ResponseApiKeyHistoryListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyTaskIdCancelDelete

> ResponseSuccessResponse ApiKeyTaskIdCancelDelete(ctx, id).XApiKey(xApiKey).Execute()

Cancel Task By Api Key

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
	id := "id_example" // string | Task's id
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.ApiKeyTaskIdCancelDelete(context.Background(), id).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.ApiKeyTaskIdCancelDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyTaskIdCancelDelete`: ResponseSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.ApiKeyTaskIdCancelDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Task&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyTaskIdCancelDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseSuccessResponse**](ResponseSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyTaskIdResultGet

> ResponseGetTaskResultResponse ApiKeyTaskIdResultGet(ctx, id).XApiKey(xApiKey).Execute()

Get Task Result

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
	id := "id_example" // string | Task's id
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.ApiKeyTaskIdResultGet(context.Background(), id).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.ApiKeyTaskIdResultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyTaskIdResultGet`: ResponseGetTaskResultResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.ApiKeyTaskIdResultGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Task&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyTaskIdResultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseGetTaskResultResponse**](ResponseGetTaskResultResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyTaskPost

> ResponseDistributeTaskResponse ApiKeyTaskPost(ctx).Input(input).XApiKey(xApiKey).Execute()

Distribute Task (Api-Key)

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
	input := *openapiclient.NewRequestDistributeTaskRequest() // RequestDistributeTaskRequest | Distribute Task Request
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyAPI.ApiKeyTaskPost(context.Background()).Input(input).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyAPI.ApiKeyTaskPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyTaskPost`: ResponseDistributeTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyAPI.ApiKeyTaskPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyTaskPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | [**RequestDistributeTaskRequest**](RequestDistributeTaskRequest.md) | Distribute Task Request | 
 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseDistributeTaskResponse**](ResponseDistributeTaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

