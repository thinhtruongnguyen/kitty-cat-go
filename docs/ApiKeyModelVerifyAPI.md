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

Check Valid Source code To Verify Ai Model By Api Key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdPreVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | [**RequestCheckValidToVerifyAiModelRequest**](RequestCheckValidToVerifyAiModelRequest.md) | Verify Ai Model Request | **yes**

### Return type

[**ResponseCheckValidToVerifyAiModelResponse**](ResponseCheckValidToVerifyAiModelResponse.md)



### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVerifyCostPost

Calculate Cost To Verify Ai Model By Api Key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVerifyCostPostRequest struct via the builder pattern


Name | Type | Description  | Require
------------- | ------------- | ------------- | -------------
 **input** | [**RequestCalculateCostToVerifyAiModelRequest**](RequestCalculateCostToVerifyAiModelRequest.md) | Verify Ai Model Request | **yes**

### Return type

[**ResponseEstimateCostResponse**](ResponseEstimateCostResponse.md)


### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVerifyPost


Verify Ai Model By Api Key




### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Require
------------- | ------------- | ------------- | -------------
 **input** | [**RequestVerifyAiModelRequest**](RequestVerifyAiModelRequest.md) | Verify Ai Model Request | **yes**

### Return type

[**ResponseVerifyAiModelResponse**](ResponseVerifyAiModelResponse.md)



### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVerifyTaskGet


Get List Verify Model Task By Commit Hash And Status





### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVerifyTaskGetRequest struct via the builder pattern


Name | Type | Description  | Require
------------- | ------------- | ------------- | -------------
 **commitHash** | **string** |  | **yes**
 **verifyStatus** | **string** |  | no

### Return type

[**ResponseModelVersioningGroupLiteListResponse**](ResponseModelVersioningGroupLiteListResponse.md)


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelVerifyHubTaskIdGet

Get Model Versioning By Hub Task Id By Api Key


### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Hub Task&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelVerifyHubTaskIdGetRequest struct via the builder pattern



### Return type

[**ResponseModelVersioningResponse**](ResponseModelVersioningResponse.md)



### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelVerifyPlatformTaskIdGet


Get Verify Platform Task By Id By Api Key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Task&#39;s Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelVerifyPlatformTaskIdGetRequest struct via the builder pattern



### Return type

[**ResponseQueueTaskResponse**](ResponseQueueTaskResponse.md)



### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

