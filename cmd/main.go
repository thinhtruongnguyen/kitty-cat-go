package main

import (
	"fmt"

	aiozaiplatformgosdk "github.com/thinhtruongnguyen/kitty-cat-go"
)

const (
	API_KEY       = "d9eb2842bb95039bfc31196432a82998bc5930f4"
	ownerUsername = "bohaha123-3333"
	repoName      = "HEHEE"
	taskId        = "aea6944c-a808-4c11-b8de-557fe02d505f"
	verifyStatus  = "verified"
	commitHash    = "ba271f290d732d6aa5902bf4d42103b516fcb7c2"
	offset        = 0
	limit         = 10
)

var (
	FROM     = "2025-09-01"
	TO       = "2025-09-10"
	MODEL_ID = "7332d265-171e-4055-9f2f-af51a52a82c4"
)

func main() {
	builder := aiozaiplatformgosdk.ClientBuilder(
		aiozaiplatformgosdk.AuthCredentials{ApiKey: API_KEY},
	)
	client := builder.Build()

	// requestGetApiKeyStatisticsByModelIdRequest := aiozaiplatformgosdk.RequestGetApiKeyStatisticsByModelIdRequest{ //req
	// 	From: &FROM,
	// 	To:   &TO,
	// }

	// requestCalculateCostToVerifyAiModelRequest := aiozaiplatformgosdk.RequestCalculateCostToVerifyAiModelRequest{ //req3
	// 	CommitHash: commitHash,
	// 	Platforms:  []string{"window"},
	// }

	// commitHash := "master"
	// requestCheckValidToVerifyAiModelRequest := aiozaiplatformgosdk.RequestCheckValidToVerifyAiModelRequest{ // req5
	// 	CommitHash: commitHash,
	// 	Platforms:  []string{"window"},
	// }

	// requestVerifyAiModelRequest := aiozaiplatformgosdk.RequestVerifyAiModelRequest{ // req2
	// 	CommitHash: &commitHash,
	// 	Platforms:  []string{"window"},
	// }

	// respGetListVerifiedModelVersioning, err := client.ApiKeyModelVersioning.GetListVerifiedModelVersioning(MODEL_ID, offset, limit, verifyStatus)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetListVerifiedModelVersioning): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetListVerifiedModelVersioning)

	// respVerifyModel, err := client.ApiKeyModelVerify.VerifyModel(MODEL_ID, requestVerifyAiModelRequest)
	// if err != nil {
	// 	fmt.Printf("API call failed (PreCheckToVerifyModel): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respVerifyModel)

	// respGetModelInfo, err := client.ApiKeyModel.GetModelInfo(MODEL_ID)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetModelInfo): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetModelInfo)

	// respGetListPlatformsSupport, err := client.ApiKeyModel.GetListPlatformsSupport()
	// if err != nil {
	// 	fmt.Printf("API call failed (GetListPlatformsSupport): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetListPlatformsSupport)

	// respGetModelStatistics, err := client.ApiKeyModel.GetModelStatistics(MODEL_ID, requestGetApiKeyStatisticsByModelIdRequest)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetModelStatistics): %v", err)
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetModelStatistics)

	// respGetTaskHistory, err := client.ApiApiKey.GetTaskHistory(limit, offset)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetTaskHistory): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetTaskHistory)

	respGetBalance, err := client.ApiApiKey.ApiKeyBalanceGet()
	if err != nil {
		fmt.Printf("API call failed (GetBalance): %v", err)
		return
	}
	aiozaiplatformgosdk.PrettyPrint(respGetBalance)
	fmt.Println(*respGetBalance.Status)
	fmt.Println(*respGetBalance.Data.WalletAddress)

	// respGetApiKeyPermission, err := client.ApiApiKey.GetApiKeyPermission()
	// if err != nil {
	// 	fmt.Printf("API call failed (GetApiKeyPermission): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetApiKeyPermission)

	// end success

	// respGetCurrentModelVersioningByModelId, err := client.ApiKeyModelVersioning.GetCurrentModelVersioningByModelId(MODEL_ID)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetCurrentModelVersioningByModelId): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetCurrentModelVersioningByModelId)

	// respChangeCurrentModelVersioningByModelIdAndCommitHash, err := client.ApiKeyModelVersioning.ChangeCurrentModelVersioningByModelIdAndCommitHash(MODEL_ID, commitHash)
	// if err != nil {
	// 	fmt.Printf("API call failed (ChangeCurrentModelVersioningByModelIdAndCommitHash): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respChangeCurrentModelVersioningByModelIdAndCommitHash)

	// respDeleteModelVersioningByModelIdAndCommitHash, err := client.ApiKeyModelVersioning.DeleteModelVersioningByModelIdAndCommitHash(MODEL_ID, commitHash)
	// if err != nil {
	// 	fmt.Printf("API call failed (DeleteModelVersioningByModelIdAndCommitHash): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respDeleteModelVersioningByModelIdAndCommitHash)

	// respGetModelVersioningByTaskId, err := client.ApiKeyModelVerify.GetModelVersioningByTaskId(taskId)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetModelVersioningByTaskId): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetModelVersioningByTaskId)

	// respGetVerifyPlatformTaskById, err := client.ApiKeyModelVerify.GetVerifyPlatformTaskById(taskId)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetVerifyPlatformTaskById): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetVerifyPlatformTaskById)

	// respGetListVerifyModelTaskByCommitHashAndStatus, err := client.ApiKeyModelVerify.GetListVerifyModelTaskByCommitHashAndStatus(MODEL_ID, commitHash, verifyStatus)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetListVerifyModelTaskByCommitHashAndStatus): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetListVerifyModelTaskByCommitHashAndStatus)

	// respPreCheckToVerifyModel, err := client.ApiKeyModelVerify.PreCheckToVerifyModel(MODEL_ID, requestCheckValidToVerifyAiModelRequest)
	// if err != nil {
	// 	fmt.Printf("API call failed (PreCheckToVerifyModel): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respPreCheckToVerifyModel)

	// respCalculateCostToVerifyModel, err := client.ApiKeyModelVerify.CalculateCostToVerifyModel(MODEL_ID, requestCalculateCostToVerifyAiModelRequest)
	// if err != nil {
	// 	fmt.Printf("API call failed (CalculateCostToVerifyModel): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respCalculateCostToVerifyModel)

	// respCheckModelIsServing, err := client.ApiKeyModel.CheckModelIsServing(MODEL_ID)
	// if err != nil {
	// 	fmt.Printf("API call failed (CheckModelIsServing): %v\n", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respCheckModelIsServing)

	// cost, err := client.ApiKeyModel.GetModelTaskCost(MODEL_ID)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetModelTaskCost): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(cost)

	// respCancelTask, err := client.ApiApiKey.CancelTask(taskId)
	// if err != nil {
	// 	fmt.Printf("API call failed (CancelTask): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respCancelTask)

	// requestDistributeTaskRequest := aiozaiplatformgosdk.RequestDistributeTaskRequest{
	// 	ModelId: &MODEL_ID,
	// 	Files: []map[string]interface{}{
	// 		{
	// 			"key":  "input",
	// 			"data": "https://s3.w3s.aioz.network/w3ai-platform-v2/uploads/samples/48c4943f-f6e9-4245-a5d2-211633231ba8/2024/10/08/1728359037-bJxSfUAz6CLJ2GUrx9Xy9Q.jpg?AWSAccessKeyId=FT7EO3IGQNMIILHXIDZRVTJHWE&Signature=g2sjadorxmjBW7AQN1a3cXzCkwE%3D&Expires=2359079037",
	// 			"name": "animetosketch.jpg",
	// 		},
	// 	},
	// 	InputParams: map[string]interface{}{},
	// }
	// respCreateTask, err := client.ApiApiKey.CreateTask(requestDistributeTaskRequest)
	// if err != nil {
	// 	fmt.Printf("API call failed (CreateTask): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respCreateTask)

	// respGetApiKeyStatistics, err := client.ApiApiKey.GetApiKeyStatistics(requestGetApiKeyStatisticsByModelIdRequest)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetApiKeyStatistics): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetApiKeyStatistics)

	// respGetTaskResult, err := client.ApiApiKey.GetTaskResult(taskId)
	// if err != nil {
	// 	fmt.Printf("API call failed (GetTaskResult): %v", err)
	// 	return
	// }
	// aiozaiplatformgosdk.PrettyPrint(respGetTaskResult)

	//end

	requestGetCommitHistoryRequest := aiozaiplatformgosdk.GetCommitHistoryRequest{
		OwnerUsername:  ownerUsername,
		RepositoryName: repoName,
		Path:           "",
		Sha:            "master",
		RepoType:       "model",
		Page:           1,
		PageSize:       10,
	}

	respGetCommitHistory, err := client.ApiKeyRepository.ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet(requestGetCommitHistoryRequest)
	if err != nil {
		fmt.Printf("API call failed(GetCommitHistoryRequest): %v", err)
		return
	}
	aiozaiplatformgosdk.PrettyPrint(respGetCommitHistory)

}
