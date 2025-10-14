# \ApiKeyRepositoryAPI

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet**](ApiKeyRepositoryAPI.md#ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet) | **Get** /api-key/repository/{ownerUsername}/{repositoryName}/commit/history | Get commit history by repository name and branch name by api key



## ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet

> ResponseGetCommitHistoryResponse ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet(ctx, ownerUsername, repositoryName).Sha(sha).XApiKey(xApiKey).Page(page).PageSize(pageSize).Path(path).RepoType(repoType).Execute()

Get commit history by repository name and branch name by api key

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
	ownerUsername := "ownerUsername_example" // string | repository's owner
	repositoryName := "repositoryName_example" // string | repository's name
	sha := "sha_example" // string | Sha is the sha of the commit (optional)
	xApiKey := "xApiKey_example" // string | api-key (optional)
	page := int32(56) // int32 | Page is the page number (default: 1) (optional) (optional) (default to 1)
	pageSize := int32(56) // int32 | PageSize is the page size (default: 10) (optional) (optional) (default to 10)
	path := "path_example" // string | Path is the path of the file (optional) (optional)
	repoType := "repoType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyRepositoryAPI.ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet(context.Background(), ownerUsername, repositoryName).Sha(sha).XApiKey(xApiKey).Page(page).PageSize(pageSize).Path(path).RepoType(repoType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyRepositoryAPI.ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet`: ResponseGetCommitHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyRepositoryAPI.ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerUsername** | **string** | repository&#39;s owner | 
**repositoryName** | **string** | repository&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sha** | **string** | Sha is the sha of the commit (optional) | 
 **xApiKey** | **string** | api-key | 
 **page** | **int32** | Page is the page number (default: 1) (optional) | [default to 1]
 **pageSize** | **int32** | PageSize is the page size (default: 10) (optional) | [default to 10]
 **path** | **string** | Path is the path of the file (optional) | 
 **repoType** | **string** |  | 

### Return type

[**ResponseGetCommitHistoryResponse**](ResponseGetCommitHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

