<h1 align="center">AIOZ SDK Go client</h1>



## Project description
Using aioz-sdk golang package to querry data from AIOZ Flatform by using APIKey. For example , Getting user balance, creating a task, getting task result, getting statistics and history of models, checking API permission ...

To start, simply require the AIOZ-SDK and set up an instance with your W3AI API Keys.Please checking out your AIOZ Page. In the example below show you how to using the SDK.
## Installation

Create new Go project:

```sh
go mod init package_name
```

Put the package under your project folder and add the following in import:

```go
import myclient "github.com/GIT_USER_ID/GIT_REPO_ID"
```



## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `myclient.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), myclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `myclient.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), myclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `myclient.ContextOperationServerIndices` and `myclient.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), myclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), myclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiKeyAPI* | [**ApiKeyBalanceGet**](docs/ApiKeyAPI.md#apikeybalanceget) | **Get** /api-key/balance | Get Api Key Balance
*ApiKeyAPI* | [**ApiKeyPermissionGet**](docs/ApiKeyAPI.md#apikeypermissionget) | **Get** /api-key/permission | Get Api Key Permission
*ApiKeyAPI* | [**ApiKeyStatisticsPost**](docs/ApiKeyAPI.md#apikeystatisticspost) | **Post** /api-key/statistics | Get Api Key Statistics
*ApiKeyAPI* | [**ApiKeyTaskHistoriesGet**](docs/ApiKeyAPI.md#apikeytaskhistoriesget) | **Get** /api-key/task/histories | Get Tasks Histories
*ApiKeyAPI* | [**ApiKeyTaskIdCancelDelete**](docs/ApiKeyAPI.md#apikeytaskidcanceldelete) | **Delete** /api-key/task/{id}/cancel | Cancel Task By Api Key
*ApiKeyAPI* | [**ApiKeyTaskIdResultGet**](docs/ApiKeyAPI.md#apikeytaskidresultget) | **Get** /api-key/task/{id}/result | Get Task Result
*ApiKeyAPI* | [**ApiKeyTaskPost**](docs/ApiKeyAPI.md#apikeytaskpost) | **Post** /api-key/task | Distribute Task (Api-Key)
*ApiKeyModelAPI* | [**ApiKeyModelIdInfoGet**](docs/ApiKeyModelAPI.md#apikeymodelidinfoget) | **Get** /api-key/model/{id}/info | Get Api Key Model Info
*ApiKeyModelAPI* | [**ApiKeyModelIdServingGet**](docs/ApiKeyModelAPI.md#apikeymodelidservingget) | **Get** /api-key/model/{id}/serving | Check Model Is Serving
*ApiKeyModelAPI* | [**ApiKeyModelIdStatisticsPost**](docs/ApiKeyModelAPI.md#apikeymodelidstatisticspost) | **Post** /api-key/model/{id}/statistics | Get Model Statistics
*ApiKeyModelAPI* | [**ApiKeyModelIdTaskCostGet**](docs/ApiKeyModelAPI.md#apikeymodelidtaskcostget) | **Get** /api-key/model/{id}/task/cost | Get cost to compute task by model api key
*ApiKeyModelAPI* | [**ApiKeyModelVerifySupportPlatformsGet**](docs/ApiKeyModelAPI.md#apikeymodelverifysupportplatformsget) | **Get** /api-key/model/verify/support/platforms | Get List Platforms Support By Api Key
*ApiKeyModelVerifyAPI* | [**ApiKeyModelIdPreVerifyPost**](docs/ApiKeyModelVerifyAPI.md#apikeymodelidpreverifypost) | **Post** /api-key/model/{id}/pre-verify | Check Valid Source code To Verify Ai Model By Api Key
*ApiKeyModelVerifyAPI* | [**ApiKeyModelIdVerifyCostPost**](docs/ApiKeyModelVerifyAPI.md#apikeymodelidverifycostpost) | **Post** /api-key/model/{id}/verify/cost | Calculate Cost To Verify Ai Model By Api Key
*ApiKeyModelVerifyAPI* | [**ApiKeyModelIdVerifyPost**](docs/ApiKeyModelVerifyAPI.md#apikeymodelidverifypost) | **Post** /api-key/model/{id}/verify | Verify Ai Model By Api Key
*ApiKeyModelVerifyAPI* | [**ApiKeyModelIdVerifyTaskGet**](docs/ApiKeyModelVerifyAPI.md#apikeymodelidverifytaskget) | **Get** /api-key/model/{id}/verify/task | Get List Verify Model Task By Commit Hash And Status
*ApiKeyModelVerifyAPI* | [**ApiKeyModelVerifyHubTaskIdGet**](docs/ApiKeyModelVerifyAPI.md#apikeymodelverifyhubtaskidget) | **Get** /api-key/model/verify/hub/task/{id} | Get Model Versioning By Hub Task Id By Api Key
*ApiKeyModelVerifyAPI* | [**ApiKeyModelVerifyPlatformTaskIdGet**](docs/ApiKeyModelVerifyAPI.md#apikeymodelverifyplatformtaskidget) | **Get** /api-key/model/verify/platform/task/{id} | Get Verify Platform Task By Id By Api Key
*ApiKeyModelVersioningAPI* | [**ApiKeyModelIdVersioningDelete**](docs/ApiKeyModelVersioningAPI.md#apikeymodelidversioningdelete) | **Delete** /api-key/model/{id}/versioning | Delete Model Versioning By Commit Hash By Api Key
*ApiKeyModelVersioningAPI* | [**ApiKeyModelIdVersioningGet**](docs/ApiKeyModelVersioningAPI.md#apikeymodelidversioningget) | **Get** /api-key/model/{id}/versioning | Get Current Model Versioning By Model Id By ApiKey
*ApiKeyModelVersioningAPI* | [**ApiKeyModelIdVersioningListGet**](docs/ApiKeyModelVersioningAPI.md#apikeymodelidversioninglistget) | **Get** /api-key/model/{id}/versioning/list | Get Verified List Model Versioning By Api Key
*ApiKeyModelVersioningAPI* | [**ApiKeyModelIdVersioningPut**](docs/ApiKeyModelVersioningAPI.md#apikeymodelidversioningput) | **Put** /api-key/model/{id}/versioning | Change Model Versioning By Commit Hash By Api Key
*ApiKeyRepositoryAPI* | [**ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet**](docs/ApiKeyRepositoryAPI.md#apikeyrepositoryownerusernamerepositorynamecommithistoryget) | **Get** /api-key/repository/{ownerUsername}/{repositoryName}/commit/history | Get commit history by repository name and branch name by api key


## Documentation For Models

 - [ModelsApiKeyHistories](docs/ModelsApiKeyHistories.md)
 - [ModelsApiKeyInfo](docs/ModelsApiKeyInfo.md)
 - [ModelsApiKeyPackage](docs/ModelsApiKeyPackage.md)
 - [ModelsCommit](docs/ModelsCommit.md)
 - [ModelsCommitMeta](docs/ModelsCommitMeta.md)
 - [ModelsCommitStats](docs/ModelsCommitStats.md)
 - [ModelsCommitUser](docs/ModelsCommitUser.md)
 - [ModelsFollow](docs/ModelsFollow.md)
 - [ModelsMember](docs/ModelsMember.md)
 - [ModelsModelVersioning](docs/ModelsModelVersioning.md)
 - [ModelsModelVersioningGroupLite](docs/ModelsModelVersioningGroupLite.md)
 - [ModelsOffer](docs/ModelsOffer.md)
 - [ModelsPlatformTask](docs/ModelsPlatformTask.md)
 - [ModelsQueueTask](docs/ModelsQueueTask.md)
 - [ModelsRepoCommit](docs/ModelsRepoCommit.md)
 - [ModelsUser](docs/ModelsUser.md)
 - [ModelsWallet](docs/ModelsWallet.md)
 - [ModelsWalletWithAddress](docs/ModelsWalletWithAddress.md)
 - [RequestCalculateCostToVerifyAiModelRequest](docs/RequestCalculateCostToVerifyAiModelRequest.md)
 - [RequestCheckValidToVerifyAiModelRequest](docs/RequestCheckValidToVerifyAiModelRequest.md)
 - [RequestDistributeTaskRequest](docs/RequestDistributeTaskRequest.md)
 - [RequestGetApiKeyStatisticsByModelIdRequest](docs/RequestGetApiKeyStatisticsByModelIdRequest.md)
 - [RequestVerifyAiModelRequest](docs/RequestVerifyAiModelRequest.md)
 - [ResponseApiKeyHistoryListData](docs/ResponseApiKeyHistoryListData.md)
 - [ResponseApiKeyHistoryListResponse](docs/ResponseApiKeyHistoryListResponse.md)
 - [ResponseApiKeyInfoResponse](docs/ResponseApiKeyInfoResponse.md)
 - [ResponseCheckModelIsServingData](docs/ResponseCheckModelIsServingData.md)
 - [ResponseCheckModelIsServingResponse](docs/ResponseCheckModelIsServingResponse.md)
 - [ResponseCheckValidToVerifyAiModelData](docs/ResponseCheckValidToVerifyAiModelData.md)
 - [ResponseCheckValidToVerifyAiModelResponse](docs/ResponseCheckValidToVerifyAiModelResponse.md)
 - [ResponseDistributeTaskResponse](docs/ResponseDistributeTaskResponse.md)
 - [ResponseErrorResponse](docs/ResponseErrorResponse.md)
 - [ResponseEstimateCostData](docs/ResponseEstimateCostData.md)
 - [ResponseEstimateCostResponse](docs/ResponseEstimateCostResponse.md)
 - [ResponseFailResponse](docs/ResponseFailResponse.md)
 - [ResponseGetApiKeyPermissionData](docs/ResponseGetApiKeyPermissionData.md)
 - [ResponseGetApiKeyPermissionResponse](docs/ResponseGetApiKeyPermissionResponse.md)
 - [ResponseGetCommitHistoryData](docs/ResponseGetCommitHistoryData.md)
 - [ResponseGetCommitHistoryResponse](docs/ResponseGetCommitHistoryResponse.md)
 - [ResponseGetListModelVersioningLiteResponse](docs/ResponseGetListModelVersioningLiteResponse.md)
 - [ResponseGetListModelVersioningsLiteData](docs/ResponseGetListModelVersioningsLiteData.md)
 - [ResponseGetListPlatformSupportResponse](docs/ResponseGetListPlatformSupportResponse.md)
 - [ResponseGetTaskResultData](docs/ResponseGetTaskResultData.md)
 - [ResponseGetTaskResultResponse](docs/ResponseGetTaskResultResponse.md)
 - [ResponseGetTaskStatisticsData](docs/ResponseGetTaskStatisticsData.md)
 - [ResponseGetTaskStatisticsResponse](docs/ResponseGetTaskStatisticsResponse.md)
 - [ResponseModelVersioningGroupLiteListData](docs/ResponseModelVersioningGroupLiteListData.md)
 - [ResponseModelVersioningGroupLiteListResponse](docs/ResponseModelVersioningGroupLiteListResponse.md)
 - [ResponseModelVersioningGroupLiteResponse](docs/ResponseModelVersioningGroupLiteResponse.md)
 - [ResponseModelVersioningResponse](docs/ResponseModelVersioningResponse.md)
 - [ResponseQueueTaskResponse](docs/ResponseQueueTaskResponse.md)
 - [ResponseResult](docs/ResponseResult.md)
 - [ResponseSuccessResponse](docs/ResponseSuccessResponse.md)
 - [ResponseVerifyAiModelResponse](docs/ResponseVerifyAiModelResponse.md)
 - [ResponseWalletWithAddressResponse](docs/ResponseWalletWithAddressResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: BearerAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		myclient.ContextAPIKeys,
		map[string]myclient.APIKey{
			"BearerAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@swagger.io

