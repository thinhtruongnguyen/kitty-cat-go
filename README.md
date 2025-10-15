<h1 align="center">AIOZ SDK Go client</h1>



## Project description
Using aioz-sdk golang package to querry data from AIOZ Flatform by using APIKey. For example , Getting user balance, creating a task, getting task result, getting statistics and history of models, checking API permission ...

To start, simply require the AIOZ-SDK and set up an instance with your W3AI API Keys.Please checking out your AIOZ Page. In the example below show you how to using the SDK.
### Installation
```bash
go get github.com/thinhtruongnguyen/kitty-cat-go
```



## Code sample
```golang
package main

import (
	"fmt"

	aiozaiplatformgosdk "github.com/thinhtruongnguyen/kitty-cat-go"
)
const API_KEY = "YOUR_API_KEY"

func main() {
	builder := aiozaiplatformgosdk.ClientBuilder(
		aiozaiplatformgosdk.AuthCredentials{ApiKey: API_KEY},
	)
	client := builder.Build()

	respGetBalance, err := client.ApiApiKey.GetBalance()
	if err != nil {
		fmt.Printf("API call failed (GetBalance): %v", err)
		return
	}
	aiozaiplatformgosdk.PrettyPrint(respGetBalance)
	// fmt.Println(*respGetBalance.Status)
	// fmt.Println(*respGetBalance.Data.FreeBalance)
}

```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/v1*

#### ApiKey


##### Retrieve an instance of the ApiKey API:
```go
import (
	aiozaiplatformgosdk "github.com/thinhtruongnguyen/kitty-cat-go"
)
const API_KEY = "YOUR_API_KEY"

func main() {
	builder := aiozaiplatformgosdk.ClientBuilder(
		aiozaiplatformgosdk.AuthCredentials{ApiKey: API_KEY},
	)
	client := builder.Build()

	ApiKeyApi, err := client.ApiApiKey
}
```

##### Endpoints

Method | HTTP request | Description
 ------------- | ------------- | -------------
 [**ApiKeyBalanceGet**](docs/ApiKeyAPI.md#apikeybalanceget) | **Get** /api-key/balance | Get Api Key Balance
 [**ApiKeyPermissionGet**](docs/ApiKeyAPI.md#apikeypermissionget) | **Get** /api-key/permission | Get Api Key Permission
 [**ApiKeyStatisticsPost**](docs/ApiKeyAPI.md#apikeystatisticspost) | **Post** /api-key/statistics | Get Api Key Statistics
 [**ApiKeyTaskHistoriesGet**](docs/ApiKeyAPI.md#apikeytaskhistoriesget) | **Get** /api-key/task/histories | Get Tasks Histories
 [**ApiKeyTaskIdCancelDelete**](docs/ApiKeyAPI.md#apikeytaskidcanceldelete) | **Delete** /api-key/task/{id}/cancel | Cancel Task By Api Key
 [**ApiKeyTaskIdResultGet**](docs/ApiKeyAPI.md#apikeytaskidresultget) | **Get** /api-key/task/{id}/result | Get Task Result
 [**ApiKeyTaskPost**](docs/ApiKeyAPI.md#apikeytaskpost) | **Post** /api-key/task | Distribute Task (Api-Key)


#### ApiKeyModel


##### Retrieve an instance of the ApiKey API:
```go
import (
	aiozaiplatformgosdk "github.com/thinhtruongnguyen/kitty-cat-go"
)
const API_KEY = "YOUR_API_KEY"

func main() {
	builder := aiozaiplatformgosdk.ClientBuilder(
		aiozaiplatformgosdk.AuthCredentials{ApiKey: API_KEY},
	)
	client := builder.Build()

	ApiKeyModelApi, err := client.ApiKeyModel
}
```

##### Endpoints

Method | HTTP request | Description
 ------------- | ------------- | -------------
[**ApiKeyModelIdInfoGet**](docs/ApiKeyModelAPI.md#apikeymodelidinfoget) | **Get** /api-key/model/{id}/info | Get Api Key Model Info
[**ApiKeyModelIdServingGet**](docs/ApiKeyModelAPI.md#apikeymodelidservingget) | **Get** /api-key/model/{id}/serving | Check Model Is Serving
[**ApiKeyModelIdStatisticsPost**](docs/ApiKeyModelAPI.md#apikeymodelidstatisticspost) | **Post** /api-key/model/{id}/statistics | Get Model Statistics
[**ApiKeyModelIdTaskCostGet**](docs/ApiKeyModelAPI.md#apikeymodelidtaskcostget) | **Get** /api-key/model/{id}/task/cost | Get cost to compute task by model api key
[**ApiKeyModelVerifySupportPlatformsGet**](docs/ApiKeyModelAPI.md#apikeymodelverifysupportplatformsget) | **Get** /api-key/model/verify/support/platforms | Get List Platforms Support By Api Key


#### ApiKeyModelVerify


##### Retrieve an instance of the ApiKey API:
```go
import (
	aiozaiplatformgosdk "github.com/thinhtruongnguyen/kitty-cat-go"
)
const API_KEY = "YOUR_API_KEY"

func main() {
	builder := aiozaiplatformgosdk.ClientBuilder(
		aiozaiplatformgosdk.AuthCredentials{ApiKey: API_KEY},
	)
	client := builder.Build()

	ApiKeyModelVerifyAPI, err := client.ApiKeyModelVerify
}
```

##### Endpoints

Method | HTTP request | Description
 ------------- | ------------- | -------------
[**ApiKeyModelIdPreVerifyPost**](docs/ApiKeyModelVerifyAPI.md#apikeymodelidpreverifypost) | **Post** /api-key/model/{id}/pre-verify | Check Valid Source code To Verify Ai Model By Api Key
[**ApiKeyModelIdVerifyCostPost**](docs/ApiKeyModelVerifyAPI.md#apikeymodelidverifycostpost) | **Post** /api-key/model/{id}/verify/cost | Calculate Cost To Verify Ai Model By Api Key
[**ApiKeyModelIdVerifyPost**](docs/ApiKeyModelVerifyAPI.md#apikeymodelidverifypost) | **Post** /api-key/model/{id}/verify | Verify Ai Model By Api Key
[**ApiKeyModelIdVerifyTaskGet**](docs/ApiKeyModelVerifyAPI.md#apikeymodelidverifytaskget) | **Get** /api-key/model/{id}/verify/task | Get List Verify Model Task By Commit Hash And Status
[**ApiKeyModelVerifyHubTaskIdGet**](docs/ApiKeyModelVerifyAPI.md#apikeymodelverifyhubtaskidget) | **Get** /api-key/model/verify/hub/task/{id} | Get Model Versioning By Hub Task Id By Api Key
[**ApiKeyModelVerifyPlatformTaskIdGet**](docs/ApiKeyModelVerifyAPI.md#apikeymodelverifyplatformtaskidget) | **Get** /api-key/model/verify/platform/task/{id} | Get Verify Platform Task By Id By Api Key


#### ApiKeyModelVersioning


##### Retrieve an instance of the ApiKey API:
```go
import (
	aiozaiplatformgosdk "github.com/thinhtruongnguyen/kitty-cat-go"
)
const API_KEY = "YOUR_API_KEY"

func main() {
	builder := aiozaiplatformgosdk.ClientBuilder(
		aiozaiplatformgosdk.AuthCredentials{ApiKey: API_KEY},
	)
	client := builder.Build()

	ApiKeyModelVersioningAPI, err := client.ApiKeyModelVersioning
}
```

##### Endpoints

Method | HTTP request | Description
 ------------- | ------------- | -------------
[**ApiKeyModelIdVersioningDelete**](docs/ApiKeyModelVersioningAPI.md#apikeymodelidversioningdelete) | **Delete** /api-key/model/{id}/versioning | Delete Model Versioning By Commit Hash By Api Key
[**ApiKeyModelIdVersioningGet**](docs/ApiKeyModelVersioningAPI.md#apikeymodelidversioningget) | **Get** /api-key/model/{id}/versioning | Get Current Model Versioning By Model Id By ApiKey
[**ApiKeyModelIdVersioningListGet**](docs/ApiKeyModelVersioningAPI.md#apikeymodelidversioninglistget) | **Get** /api-key/model/{id}/versioning/list | Get Verified List Model Versioning By Api Key
[**ApiKeyModelIdVersioningPut**](docs/ApiKeyModelVersioningAPI.md#apikeymodelidversioningput) | **Put** /api-key/model/{id}/versioning | Change Model Versioning By Commit Hash By Api Key

#### ApiKeyRepositoryAPI


##### Retrieve an instance of the ApiKey API:
```go
import (
	aiozaiplatformgosdk "github.com/thinhtruongnguyen/kitty-cat-go"
)
const API_KEY = "YOUR_API_KEY"

func main() {
	builder := aiozaiplatformgosdk.ClientBuilder(
		aiozaiplatformgosdk.AuthCredentials{ApiKey: API_KEY},
	)
	client := builder.Build()

	ApiKeyRepositoryAPI, err := client.ApiKeyRepository
}
```

##### Endpoints

Method | HTTP request | Description
 ------------- | ------------- | -------------
[**ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet**](docs/ApiKeyRepositoryAPI.md#apikeyrepositoryownerusernamerepositorynamecommithistoryget) | **Get** /api-key/repository/{ownerUsername}/{repositoryName}/commit/history | Get commit history by repository name and branch name by api key


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

All of API request to be authenticated using the API key that be described in our [documentation](https://google.com).
First of all you have to do is provided an API key when instantiating the AIOZClient:

Example

```go
import (
	aiozaiplatformgosdk "github.com/thinhtruongnguyen/kitty-cat-go"
)
const API_KEY = "YOUR_API_KEY"

func main() {
	builder := aiozaiplatformgosdk.ClientBuilder(
		aiozaiplatformgosdk.AuthCredentials{ApiKey: API_KEY},
	)
	client := builder.Build()
}
```
