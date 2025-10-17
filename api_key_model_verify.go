package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyModelVerifyServiceI interface {
	ApiKeyModelVerifyHubTaskIdGetWithContext(ctx context.Context, taskId string) (*ResponseModelVersioningResponse, error)
	ApiKeyModelVerifyHubTaskIdGet(taskId string) (*ResponseModelVersioningResponse, error)
	ApiKeyModelVerifyPlatformTaskIdGetWithContext(ctx context.Context, taskId string) (*ResponseQueueTaskResponse, error)
	ApiKeyModelVerifyPlatformTaskIdGet(taskId string) (*ResponseQueueTaskResponse, error)
	ApiKeyModelIdVerifyTaskGetWithContext(ctx context.Context, modelId, commitHash, status string) (*ResponseModelVersioningGroupLiteListResponse, error)
	ApiKeyModelIdVerifyTaskGet(modelId, commitHash, status string) (*ResponseModelVersioningGroupLiteListResponse, error)
	ApiKeyModelIdPreVerifyPost(modelId string, request RequestCheckValidToVerifyAiModelRequest) (*ResponseCheckValidToVerifyAiModelResponse, error)
	ApiKeyModelIdPreVerifyPostWithContext(ctx context.Context, modelId string, request RequestCheckValidToVerifyAiModelRequest) (*ResponseCheckValidToVerifyAiModelResponse, error)
	ApiKeyModelIdVerifyPostWithContext(ctx context.Context, modelId string, request RequestVerifyAiModelRequest) (*ResponseVerifyAiModelResponse, error)
	ApiKeyModelIdVerifyPost(modelId string, request RequestVerifyAiModelRequest) (*ResponseVerifyAiModelResponse, error)
	ApiKeyModelIdVerifyCostPostWithContext(ctx context.Context, modelId string, request RequestCalculateCostToVerifyAiModelRequest) (*ResponseEstimateCostResponse, error)
	ApiKeyModelIdVerifyCostPost(modelId string, request RequestCalculateCostToVerifyAiModelRequest) (*ResponseEstimateCostResponse, error)
}

type ApiKeyModelVerifyService struct {
	Client *ClientV2
}

func (service *ApiKeyModelVerifyService) ApiKeyModelIdVerifyCostPost(
	modelId string,
	request RequestCalculateCostToVerifyAiModelRequest) (*ResponseEstimateCostResponse, error) {
	return service.ApiKeyModelIdVerifyCostPostWithContext(context.Background(), modelId, request)
}
func (service *ApiKeyModelVerifyService) ApiKeyModelIdVerifyCostPostWithContext(
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
func (service *ApiKeyModelVerifyService) ApiKeyModelIdVerifyPost(
	modelId string, request RequestVerifyAiModelRequest) (*ResponseVerifyAiModelResponse, error) {
	return service.ApiKeyModelIdVerifyPostWithContext(context.Background(), modelId, request)
}
func (service *ApiKeyModelVerifyService) ApiKeyModelIdVerifyPostWithContext(
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

func (service *ApiKeyModelVerifyService) ApiKeyModelIdPreVerifyPost(
	modelId string,
	request RequestCheckValidToVerifyAiModelRequest,
) (*ResponseCheckValidToVerifyAiModelResponse, error) {
	return service.ApiKeyModelIdPreVerifyPostWithContext(context.Background(), modelId, request)
}
func (service *ApiKeyModelVerifyService) ApiKeyModelIdPreVerifyPostWithContext(
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

func (service *ApiKeyModelVerifyService) ApiKeyModelIdVerifyTaskGet(
	modelId, commitHash, status string,
) (*ResponseModelVersioningGroupLiteListResponse, error) {
	return service.ApiKeyModelIdVerifyTaskGetWithContext(context.Background(), modelId, commitHash, status)
}
func (service *ApiKeyModelVerifyService) ApiKeyModelIdVerifyTaskGetWithContext(
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
func (service *ApiKeyModelVerifyService) ApiKeyModelVerifyPlatformTaskIdGet(taskId string) (*ResponseQueueTaskResponse, error) {
	return service.ApiKeyModelVerifyPlatformTaskIdGetWithContext(context.Background(), taskId)
}
func (service *ApiKeyModelVerifyService) ApiKeyModelVerifyPlatformTaskIdGetWithContext(ctx context.Context, taskId string) (*ResponseQueueTaskResponse, error) {
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
func (service *ApiKeyModelVerifyService) ApiKeyModelVerifyHubTaskIdGet(taskId string) (*ResponseModelVersioningResponse, error) {
	return service.ApiKeyModelVerifyHubTaskIdGetWithContext(context.Background(), taskId)
}
func (service *ApiKeyModelVerifyService) ApiKeyModelVerifyHubTaskIdGetWithContext(ctx context.Context, taskId string) (*ResponseModelVersioningResponse, error) {
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
