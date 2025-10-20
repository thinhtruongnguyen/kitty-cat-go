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

Get Api Key Model Info



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdInfoGetRequest struct via the builder pattern


### Return type

[**ResponseApiKeyInfoResponse**](ResponseApiKeyInfoResponse.md)



### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdServingGet


Check Model Is Serving



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdServingGetRequest struct via the builder pattern


### Return type

[**ResponseCheckModelIsServingResponse**](ResponseCheckModelIsServingResponse.md)


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdStatisticsPost


Get Model Statistics



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdStatisticsPostRequest struct via the builder pattern


Name | Type | Description  | Require
------------- | ------------- | ------------- | -------------

 **input** | [**RequestGetApiKeyStatisticsByModelIdRequest**](RequestGetApiKeyStatisticsByModelIdRequest.md) | Get Api Key Statistics By Model Id Request | **yes**

### Return type

[**ResponseGetTaskStatisticsResponse**](ResponseGetTaskStatisticsResponse.md)


### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdTaskCostGet

Get cost to compute task by model api key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdTaskCostGetRequest struct via the builder pattern


### Return type

[**ResponseEstimateCostResponse**](ResponseEstimateCostResponse.md)


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelVerifySupportPlatformsGet


Get List Platforms Support By Api Key



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelVerifySupportPlatformsGetRequest struct via the builder pattern


### Return type

[**ResponseGetListPlatformSupportResponse**](ResponseGetListPlatformSupportResponse.md)


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

