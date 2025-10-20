# \ApiKeyRepositoryAPI

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet**](ApiKeyRepositoryAPI.md#ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet) | **Get** /api-key/repository/{ownerUsername}/{repositoryName}/commit/history | Get commit history by repository name and branch name by api key



## ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet

Get commit history by repository name and branch name by api key


### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ownerUsername** | **string** | repository&#39;s owner | 
**repositoryName** | **string** | repository&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sha** | **string** | Sha is the sha of the commit (optional) | 
 **page** | **int32** | Page is the page number (default: 1) (optional) | [default to 1]
 **pageSize** | **int32** | PageSize is the page size (default: 10) (optional) | [default to 10]
 **path** | **string** | Path is the path of the file (optional) | 
 **repoType** | **string** |  | 

### Return type

[**ResponseGetCommitHistoryResponse**](ResponseGetCommitHistoryResponse.md)


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

