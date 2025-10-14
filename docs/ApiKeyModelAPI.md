# \ApiKeyModelAPI

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyModelIdInfoGet**](ApiKeyModelAPI.md#ApiKeyModelIdInfoGet) | **Get** /api-key/model/{id}/info | Get Api Key Model Info
[**ApiKeyModelIdServingGet**](ApiKeyModelAPI.md#ApiKeyModelIdServingGet) | **Get** /api-key/model/{id}/serving | Check Model Is Serving
[**ApiKeyModelIdStatisticsPost**](ApiKeyModelAPI.md#ApiKeyModelIdStatisticsPost) | **Post** /api-key/model/{id}/statistics | Get Model Statistics
[**ApiKeyModelIdTaskCostGet**](ApiKeyModelAPI.md#ApiKeyModelIdTaskCostGet) | **Get** /api-key/model/{id}/task/cost | Get cost to compute task by model api key
[**ApiKeyModelVerifySupportPlatformsGet**](ApiKeyModelAPI.md#ApiKeyModelVerifySupportPlatformsGet) | **Get** /api-key/model/verify/support/platforms | Get List Platforms Support By Api Key



## ApiKeyModelIdInfoGet

> ResponseApiKeyInfoResponse ApiKeyModelIdInfoGet(ctx, id).XApiKey(xApiKey).Execute()

Get Api Key Model Info

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
	resp, r, err := apiClient.ApiKeyModelAPI.ApiKeyModelIdInfoGet(context.Background(), id).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelAPI.ApiKeyModelIdInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdInfoGet`: ResponseApiKeyInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelAPI.ApiKeyModelIdInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseApiKeyInfoResponse**](ResponseApiKeyInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdServingGet

> ResponseCheckModelIsServingResponse ApiKeyModelIdServingGet(ctx, id).XApiKey(xApiKey).Execute()

Check Model Is Serving

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
	resp, r, err := apiClient.ApiKeyModelAPI.ApiKeyModelIdServingGet(context.Background(), id).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelAPI.ApiKeyModelIdServingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdServingGet`: ResponseCheckModelIsServingResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelAPI.ApiKeyModelIdServingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdServingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseCheckModelIsServingResponse**](ResponseCheckModelIsServingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdStatisticsPost

> ResponseGetTaskStatisticsResponse ApiKeyModelIdStatisticsPost(ctx, id).Input(input).XApiKey(xApiKey).Execute()

Get Model Statistics

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
	input := *openapiclient.NewRequestGetApiKeyStatisticsByModelIdRequest() // RequestGetApiKeyStatisticsByModelIdRequest | Get Api Key Statistics By Model Id Request
	xApiKey := "xApiKey_example" // string | api-key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyModelAPI.ApiKeyModelIdStatisticsPost(context.Background(), id).Input(input).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelAPI.ApiKeyModelIdStatisticsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdStatisticsPost`: ResponseGetTaskStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelAPI.ApiKeyModelIdStatisticsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdStatisticsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **input** | [**RequestGetApiKeyStatisticsByModelIdRequest**](RequestGetApiKeyStatisticsByModelIdRequest.md) | Get Api Key Statistics By Model Id Request | 
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


## ApiKeyModelIdTaskCostGet

> ResponseEstimateCostResponse ApiKeyModelIdTaskCostGet(ctx, id).XApiKey(xApiKey).Execute()

Get cost to compute task by model api key

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
	resp, r, err := apiClient.ApiKeyModelAPI.ApiKeyModelIdTaskCostGet(context.Background(), id).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelAPI.ApiKeyModelIdTaskCostGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelIdTaskCostGet`: ResponseEstimateCostResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelAPI.ApiKeyModelIdTaskCostGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdTaskCostGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseEstimateCostResponse**](ResponseEstimateCostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelVerifySupportPlatformsGet

> ResponseGetListPlatformSupportResponse ApiKeyModelVerifySupportPlatformsGet(ctx).XApiKey(xApiKey).Execute()

Get List Platforms Support By Api Key

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
	resp, r, err := apiClient.ApiKeyModelAPI.ApiKeyModelVerifySupportPlatformsGet(context.Background()).XApiKey(xApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyModelAPI.ApiKeyModelVerifySupportPlatformsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyModelVerifySupportPlatformsGet`: ResponseGetListPlatformSupportResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyModelAPI.ApiKeyModelVerifySupportPlatformsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelVerifySupportPlatformsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** | api-key | 

### Return type

[**ResponseGetListPlatformSupportResponse**](ResponseGetListPlatformSupportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

