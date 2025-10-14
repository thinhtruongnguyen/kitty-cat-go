package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyModelVerifyServiceI interface {
	GetModelVersioningByTaskIdWithContext(ctx context.Context, taskId string) (*ResponseModelVersioningResponse, error)
	GetModelVersioningByTaskId(taskId string) (*ResponseModelVersioningResponse, error)
	GetVerifyPlatformTaskByIdWithContext(ctx context.Context, taskId string) (*ResponseQueueTaskResponse, error)
	GetVerifyPlatformTaskById(taskId string) (*ResponseQueueTaskResponse, error)
	GetListVerifyModelTaskByCommitHashAndStatusWithContext(ctx context.Context, modelId, commitHash, status string) (*ResponseModelVersioningGroupLiteListResponse, error)
	GetListVerifyModelTaskByCommitHashAndStatus(modelId, commitHash, status string) (*ResponseModelVersioningGroupLiteListResponse, error)
	PreCheckToVerifyModel(modelId string, request RequestCheckValidToVerifyAiModelRequest) (*ResponseCheckValidToVerifyAiModelResponse, error)
	PreCheckToVerifyModelWithContext(ctx context.Context, modelId string, request RequestCheckValidToVerifyAiModelRequest) (*ResponseCheckValidToVerifyAiModelResponse, error)
	VerifyModelWithContext(ctx context.Context, modelId string, request RequestVerifyAiModelRequest) (*ResponseVerifyAiModelResponse, error)
	VerifyModel(modelId string, request RequestVerifyAiModelRequest) (*ResponseVerifyAiModelResponse, error)
	CalculateCostToVerifyModelWithContext(ctx context.Context, modelId string, request RequestCalculateCostToVerifyAiModelRequest) (*ResponseEstimateCostResponse, error)
	CalculateCostToVerifyModel(modelId string, request RequestCalculateCostToVerifyAiModelRequest) (*ResponseEstimateCostResponse, error)
}

type ApiKeyModelVerifyService struct {
	Client *ClientV2
}

func (service *ApiKeyModelVerifyService) CalculateCostToVerifyModel(
	modelId string,
	request RequestCalculateCostToVerifyAiModelRequest) (*ResponseEstimateCostResponse, error) {
	return service.CalculateCostToVerifyModelWithContext(context.Background(), modelId, request)
}
func (service *ApiKeyModelVerifyService) CalculateCostToVerifyModelWithContext(
	ctx context.Context, modelId string,
	request RequestCalculateCostToVerifyAiModelRequest) (*ResponseEstimateCostResponse, error) {
	localVarPath := fmt.Sprintf("/api-key/model/%s/verify/cost", modelId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarPostBody := request

	req, err := service.Client.prepareRequest(
		ctx,
		http.MethodPost,
		localVarPath,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
	)
	if err != nil {
		return nil, err
	}

	res := new(ResponseEstimateCostResponse)
	err = service.Client.do(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (service *ApiKeyModelVerifyService) VerifyModel(
	modelId string, request RequestVerifyAiModelRequest) (*ResponseVerifyAiModelResponse, error) {
	return service.VerifyModelWithContext(context.Background(), modelId, request)
}
func (service *ApiKeyModelVerifyService) VerifyModelWithContext(
	ctx context.Context, modelId string, request RequestVerifyAiModelRequest) (*ResponseVerifyAiModelResponse, error) {
	localVarPath := fmt.Sprintf("/api-key/model/%s/verify", modelId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarPostBody := request

	req, err := service.Client.prepareRequest(
		ctx,
		http.MethodPost,
		localVarPath,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
	)
	if err != nil {
		return nil, err
	}

	res := new(ResponseVerifyAiModelResponse)
	err = service.Client.do(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyModelVerifyService) PreCheckToVerifyModel(
	modelId string,
	request RequestCheckValidToVerifyAiModelRequest,
) (*ResponseCheckValidToVerifyAiModelResponse, error) {
	return service.PreCheckToVerifyModelWithContext(context.Background(), modelId, request)
}
func (service *ApiKeyModelVerifyService) PreCheckToVerifyModelWithContext(
	ctx context.Context,
	modelId string,
	request RequestCheckValidToVerifyAiModelRequest,
) (*ResponseCheckValidToVerifyAiModelResponse, error) {

	localVarPath := fmt.Sprintf("/api-key/model/%s/pre-verify", modelId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarPostBody := request

	req, err := service.Client.prepareRequest(
		ctx,
		http.MethodPost,
		localVarPath,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
	)
	if err != nil {
		return nil, err
	}

	res := new(ResponseCheckValidToVerifyAiModelResponse)
	err = service.Client.do(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyModelVerifyService) GetListVerifyModelTaskByCommitHashAndStatus(
	modelId, commitHash, status string,
) (*ResponseModelVersioningGroupLiteListResponse, error) {
	return service.GetListVerifyModelTaskByCommitHashAndStatusWithContext(context.Background(), modelId, commitHash, status)
}
func (service *ApiKeyModelVerifyService) GetListVerifyModelTaskByCommitHashAndStatusWithContext(
	ctx context.Context,
	modelId, commitHash, status string,
) (*ResponseModelVersioningGroupLiteListResponse, error) {
	localVarPath := fmt.Sprintf("/api-key/model/%s/verify/task", modelId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("commitHash", commitHash)
	localVarQueryParams.Add("verifyStatus", status)

	var localVarPostBody interface{} = nil

	req, err := service.Client.prepareRequest(
		ctx,
		http.MethodGet,
		localVarPath,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
	)
	if err != nil {
		return nil, err
	}

	res := new(ResponseModelVersioningGroupLiteListResponse)
	err = service.Client.do(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (service *ApiKeyModelVerifyService) GetVerifyPlatformTaskById(taskId string) (*ResponseQueueTaskResponse, error) {
	return service.GetVerifyPlatformTaskByIdWithContext(context.Background(), taskId)
}
func (service *ApiKeyModelVerifyService) GetVerifyPlatformTaskByIdWithContext(ctx context.Context, taskId string) (*ResponseQueueTaskResponse, error) {
	endpoint := fmt.Sprintf("/api-key/model/verify/platform/task/%s", taskId)
	var localVarPostBody any
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, endpoint, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseQueueTaskResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
func (service *ApiKeyModelVerifyService) GetModelVersioningByTaskId(taskId string) (*ResponseModelVersioningResponse, error) {
	return service.GetModelVersioningByTaskIdWithContext(context.Background(), taskId)
}
func (service *ApiKeyModelVerifyService) GetModelVersioningByTaskIdWithContext(ctx context.Context, taskId string) (*ResponseModelVersioningResponse, error) {
	endpoint := fmt.Sprintf("/api-key/model/verify/hub/task/%s", taskId)
	var localVarPostBody any
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, endpoint, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseModelVersioningResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
