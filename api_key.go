package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyServiceI interface {
	GetBalanceWithContext(ctx context.Context) (*ResponseWalletWithAddressResponse, error)
	GetBalance() (*ResponseWalletWithAddressResponse, error)
	GetApiKeyPermissionWithContext(ctx context.Context) (*ResponseGetApiKeyPermissionResponse, error)
	GetApiKeyPermission() (*ResponseGetApiKeyPermissionResponse, error)
	GetTaskHistory(limit int, offset int) (*ResponseApiKeyHistoryListResponse, error)
	GetTaskHistoryWithContext(ctx context.Context, limit int, offset int) (*ResponseApiKeyHistoryListResponse, error)
	GetTaskResult(taskId string) (*ResponseGetTaskResultResponse, error)
	GetTaskResultWithContext(ctx context.Context, taskId string) (*ResponseGetTaskResultResponse, error)
	GetApiKeyStatisticsWithContext(ctx context.Context, request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error)
	GetApiKeyStatistics(request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error)
	CreateTask(request RequestDistributeTaskRequest) (*ResponseDistributeTaskResponse, error)
	CreateTaskWithContext(ctx context.Context, request RequestDistributeTaskRequest) (*ResponseDistributeTaskResponse, error)
	CancelTaskWithContext(ctx context.Context, taskId string) (*ResponseSuccessResponse, error)
	CancelTask(taskId string) (*ResponseSuccessResponse, error)
}

type ApiKeyService struct {
	Client *ClientV2
}

func (service *ApiKeyService) CancelTask(taskId string) (*ResponseSuccessResponse, error) {
	return service.CancelTaskWithContext(context.Background(), taskId)
}
func (service *ApiKeyService) CancelTaskWithContext(ctx context.Context, taskId string) (*ResponseSuccessResponse, error) {
	endpoint := fmt.Sprintf("/api-key/task/%s/cancel", taskId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	var localVarPostBody any

	req, err := service.Client.prepareRequest(ctx, http.MethodDelete, endpoint, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseSuccessResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyService) GetBalance() (*ResponseWalletWithAddressResponse, error) {
	return service.GetBalanceWithContext(context.Background())
}
func (service *ApiKeyService) GetBalanceWithContext(ctx context.Context) (*ResponseWalletWithAddressResponse, error) {
	var localVarPostBody any
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, "/api-key/balance", localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseWalletWithAddressResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyService) GetApiKeyPermission() (*ResponseGetApiKeyPermissionResponse, error) {
	return service.GetApiKeyPermissionWithContext(context.Background())
}
func (service *ApiKeyService) GetApiKeyPermissionWithContext(ctx context.Context) (*ResponseGetApiKeyPermissionResponse, error) {
	var localVarPostBody any
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, "/api-key/permission", localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseGetApiKeyPermissionResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyService) CreateTask(request RequestDistributeTaskRequest) (*ResponseDistributeTaskResponse, error) {
	return service.CreateTaskWithContext(context.Background(), request)
}
func (service *ApiKeyService) CreateTaskWithContext(ctx context.Context, request RequestDistributeTaskRequest) (*ResponseDistributeTaskResponse, error) {
	localVarQueryParams := url.Values{}

	localVarHeaderParams := map[string]string{}
	localVarPostBody := request
	localVarPath := "/api-key/task"
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
	res := new(ResponseDistributeTaskResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (service *ApiKeyService) GetApiKeyStatistics(request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error) {
	return service.GetApiKeyStatisticsWithContext(context.Background(), request)
}
func (service *ApiKeyService) GetApiKeyStatisticsWithContext(
	ctx context.Context,
	request RequestGetApiKeyStatisticsByModelIdRequest,
) (*ResponseGetTaskStatisticsResponse, error) {
	localVarQueryParams := url.Values{}

	localVarHeaderParams := map[string]string{}
	localVarPostBody := request
	localVarPath := "/api-key/statistics"
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
	res := new(ResponseGetTaskStatisticsResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (service *ApiKeyService) GetTaskResult(taskId string) (*ResponseGetTaskResultResponse, error) {
	return service.GetTaskResultWithContext(context.Background(), taskId)
}
func (service *ApiKeyService) GetTaskResultWithContext(ctx context.Context, taskId string) (*ResponseGetTaskResultResponse, error) {
	endpoint := fmt.Sprintf("/api-key/task/%s/result", taskId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	var localVarPostBody any

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, endpoint, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseGetTaskResultResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
func (service *ApiKeyService) GetTaskHistory(limit int, offset int) (*ResponseApiKeyHistoryListResponse, error) {
	return service.GetTaskHistoryWithContext(context.Background(), limit, offset)
}
func (service *ApiKeyService) GetTaskHistoryWithContext(ctx context.Context, limit int, offset int) (*ResponseApiKeyHistoryListResponse, error) {
	localVarPath := "/api-key/task/histories"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("limit", fmt.Sprintf("%d", limit))
	localVarQueryParams.Add("offset", fmt.Sprintf("%d", offset))
	var localVarPostBody any
	req, err := service.Client.prepareRequest(ctx, http.MethodGet, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}
	res := new(ResponseApiKeyHistoryListResponse)
	err = service.Client.do(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
