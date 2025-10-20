# \ApiKeyModelVersioningAPI

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyModelIdVersioningDelete**](ApiKeyModelVersioningAPI.md#ApiKeyModelIdVersioningDelete) | **Delete** /api-key/model/{id}/versioning | Delete Model Versioning By Commit Hash By Api Key
[**ApiKeyModelIdVersioningGet**](ApiKeyModelVersioningAPI.md#ApiKeyModelIdVersioningGet) | **Get** /api-key/model/{id}/versioning | Get Current Model Versioning By Model Id By ApiKey
[**ApiKeyModelIdVersioningListGet**](ApiKeyModelVersioningAPI.md#ApiKeyModelIdVersioningListGet) | **Get** /api-key/model/{id}/versioning/list | Get Verified List Model Versioning By Api Key
[**ApiKeyModelIdVersioningPut**](ApiKeyModelVersioningAPI.md#ApiKeyModelIdVersioningPut) | **Put** /api-key/model/{id}/versioning | Change Model Versioning By Commit Hash By Api Key



## ApiKeyModelIdVersioningDelete


Delete Model Versioning By Commit Hash By Api Key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVersioningDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitHash** | **string** |  | 

### Return type

[**ResponseSuccessResponse**](ResponseSuccessResponse.md)



### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVersioningGet


Get Current Model Versioning By Model Id By ApiKey

### Example



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVersioningGetRequest struct via the builder pattern



### Return type

[**ResponseModelVersioningGroupLiteResponse**](ResponseModelVersioningGroupLiteResponse.md)


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVersioningListGet


Get Verified List Model Versioning By Api Key




### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVersioningListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 10]
 **offset** | **int32** |  | [default to 0]
 **verifyStatus** | **string** |  | 

### Return type

[**ResponseGetListModelVersioningLiteResponse**](ResponseGetListModelVersioningLiteResponse.md)


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyModelIdVersioningPut

Change Model Versioning By Commit Hash By Api Key


### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**id** | **string** | Model&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyModelIdVersioningPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitHash** | **string** |  | 

### Return type

[**ResponseSuccessResponse**](ResponseSuccessResponse.md)


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

