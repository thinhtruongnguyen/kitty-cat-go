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


Get Api Key Balance


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


Get Api Key Permission



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


Get Api Key Statistics



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyStatisticsPostRequest struct via the builder pattern


Name | Type | Description  | Require
------------- | ------------- | ------------- | -------------
 **input** | [**RequestGetApiKeyStatisticsByModelIdRequest**](RequestGetApiKeyStatisticsByModelIdRequest.md) | Get Api Key Statistics Request | **yes**


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


Get Tasks Histories



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyTaskHistoriesGetRequest struct via the builder pattern


Name | Type | Description  | Require
------------- | ------------- | ------------- | -------------
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

Cancel Task By Api Key


### Path Parameters


Name | Type | Description  | Require
------------- | ------------- | ------------- | -------------
**id** | **string** | Task&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyTaskIdCancelDeleteRequest struct via the builder pattern


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


Get Task Result



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Task&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyTaskIdResultGetRequest struct via the builder pattern


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


Distribute Task (Api-Key)


### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyTaskPostRequest struct via the builder pattern


Name | Type | Description  | Require
------------- | ------------- | ------------- | -------------
 **input** | [**RequestDistributeTaskRequest**](RequestDistributeTaskRequest.md) | Distribute Task Request | **yes**

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

