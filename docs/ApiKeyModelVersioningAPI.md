# \ApiKeyModelVersioningAPI

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyModelIdVersioningDelete**](ApiKeyModelVersioningAPI.md#ApiKeyModelIdVersioningDelete) | **Delete** /api-key/model/{id}/versioning | Delete Model Versioning By Commit Hash By Api Key
[**ApiKeyModelIdVersioningGet**](ApiKeyModelVersioningAPI.md#ApiKeyModelIdVersioningGet) | **Get** /api-key/model/{id}/versioning | Get Current Model Versioning By Model Id By ApiKey
[**ApiKeyModelIdVersioningListGet**](ApiKeyModelVersioningAPI.md#ApiKeyModelIdVersioningListGet) | **Get** /api-key/model/{id}/versioning/list | Get Verified List Model Versioning By Api Key
[**ApiKeyModelIdVersioningPut**](ApiKeyModelVersioningAPI.md#ApiKeyModelIdVersioningPut) | **Put** /api-key/model/{id}/versioning | Change Model Versioning By Commit Hash By Api Key



## ApiKeyModelIdVersioningDelete

> ResponseSuccessResponse ApiKeyModelIdVersioningDelete(ctx, id).CommitHash(commitHash).XApiKey(xApiKey).Execute()

Delete Model Versioning By Commit Hash By Api Key

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningDelete(context.Background(), id).CommitHash(commitHash).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdVersioningDelete`: ResponseSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVersioningDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitHash** | **string** |  | 
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


## ApiKeyModelIdVersioningGet

> ResponseModelVersioningGroupLiteResponse ApiKeyModelIdVersioningGet(ctx, id).XApiKey(xApiKey).Execute()

Get Current Model Versioning By Model Id By ApiKey

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
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningGet(context.Background(), id).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdVersioningGet`: ResponseModelVersioningGroupLiteResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVersioningGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseModelVersioningGroupLiteResponse**](ResponseModelVersioningGroupLiteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVersioningListGet

> ResponseGetListModelVersioningLiteResponse ApiKeyModelIdVersioningListGet(ctx, id).XApiKey(xApiKey).Limit(limit).Offset(offset).VerifyStatus(verifyStatus).Execute()

Get Verified List Model Versioning By Api Key



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
	xApiKey := "xApiKey_example" // string | api-key (optional)
	limit := int32(56) // int32 |  (optional) (default to 10)
	offset := int32(56) // int32 |  (optional) (default to 0)
	verifyStatus := "verifyStatus_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningListGet(context.Background(), id).XApiKey(xApiKey).Limit(limit).Offset(offset).VerifyStatus(verifyStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdVersioningListGet`: ResponseGetListModelVersioningLiteResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningListGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVersioningListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 
 **limit** | **int32** |  | [default to 10]
 **offset** | **int32** |  | [default to 0]
 **verifyStatus** | **string** |  | 

### Return type

[**ResponseGetListModelVersioningLiteResponse**](ResponseGetListModelVersioningLiteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVersioningPut

> ResponseSuccessResponse ApiKeyModelIdVersioningPut(ctx, id).CommitHash(commitHash).XApiKey(xApiKey).Execute()

Change Model Versioning By Commit Hash By Api Key

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningPut(context.Background(), id).CommitHash(commitHash).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdVersioningPut`: ResponseSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelVersioningAPI.ApiKeyModelIdVersioningPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVersioningPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitHash** | **string** |  | 
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

