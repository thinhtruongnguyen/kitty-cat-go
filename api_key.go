package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyServiceI interface {
	ApiKeyBalanceGetWithContext(ctx context.Context) (*ResponseWalletWithAddressResponse, error)
	ApiKeyBalanceGet() (*ResponseWalletWithAddressResponse, error)
	ApiKeyPermissionGetWithContext(ctx context.Context) (*ResponseGetApiKeyPermissionResponse, error)
	ApiKeyPermissionGet() (*ResponseGetApiKeyPermissionResponse, error)
	ApiKeyTaskHistoriesGet(limit int, offset int) (*ResponseApiKeyHistoryListResponse, error)
	ApiKeyTaskHistoriesGetWithContext(ctx context.Context, limit int, offset int) (*ResponseApiKeyHistoryListResponse, error)
	ApiKeyTaskIdResultGet(taskId string) (*ResponseGetTaskResultResponse, error)
	ApiKeyTaskIdResultGetWithContext(ctx context.Context, taskId string) (*ResponseGetTaskResultResponse, error)
	ApiKeyStatisticsPostWithContext(ctx context.Context, request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error)
	ApiKeyStatisticsPost(request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error)
	ApiKeyTaskPost(request RequestDistributeTaskRequest) (*ResponseDistributeTaskResponse, error)
	ApiKeyTaskPostWithContext(ctx context.Context, request RequestDistributeTaskRequest) (*ResponseDistributeTaskResponse, error)
	ApiKeyTaskIdCancelDeleteWithContext(ctx context.Context, taskId string) (*ResponseSuccessResponse, error)
	ApiKeyTaskIdCancelDelete(taskId string) (*ResponseSuccessResponse, error)
}

type ApiKeyService struct {
	Client *ClientV2
}

func (service *ApiKeyService) ApiKeyTaskIdCancelDelete(taskId string) (*ResponseSuccessResponse, error) {
	return service.ApiKeyTaskIdCancelDeleteWithContext(context.Background(), taskId)
}
func (service *ApiKeyService) ApiKeyTaskIdCancelDeleteWithContext(ctx context.Context, taskId string) (*ResponseSuccessResponse, error) {
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

func (service *ApiKeyService) ApiKeyBalanceGet() (*ResponseWalletWithAddressResponse, error) {
	return service.ApiKeyBalanceGetWithContext(context.Background())
}
func (service *ApiKeyService) ApiKeyBalanceGetWithContext(ctx context.Context) (*ResponseWalletWithAddressResponse, error) {
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

func (service *ApiKeyService) ApiKeyPermissionGet() (*ResponseGetApiKeyPermissionResponse, error) {
	return service.ApiKeyPermissionGetWithContext(context.Background())
}
func (service *ApiKeyService) ApiKeyPermissionGetWithContext(ctx context.Context) (*ResponseGetApiKeyPermissionResponse, error) {
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

func (service *ApiKeyService) ApiKeyTaskPost(request RequestDistributeTaskRequest) (*ResponseDistributeTaskResponse, error) {
	return service.ApiKeyTaskPostWithContext(context.Background(), request)
}
func (service *ApiKeyService) ApiKeyTaskPostWithContext(ctx context.Context, request RequestDistributeTaskRequest) (*ResponseDistributeTaskResponse, error) {
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

func (service *ApiKeyService) ApiKeyStatisticsPost(request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error) {
	return service.ApiKeyStatisticsPostWithContext(context.Background(), request)
}
func (service *ApiKeyService) ApiKeyStatisticsPostWithContext(
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

func (service *ApiKeyService) ApiKeyTaskIdResultGet(taskId string) (*ResponseGetTaskResultResponse, error) {
	return service.ApiKeyTaskIdResultGetWithContext(context.Background(), taskId)
}
func (service *ApiKeyService) ApiKeyTaskIdResultGetWithContext(ctx context.Context, taskId string) (*ResponseGetTaskResultResponse, error) {
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
func (service *ApiKeyService) ApiKeyTaskHistoriesGet(limit int, offset int) (*ResponseApiKeyHistoryListResponse, error) {
	return service.ApiKeyTaskHistoriesGetWithContext(context.Background(), limit, offset)
}
func (service *ApiKeyService) ApiKeyTaskHistoriesGetWithContext(ctx context.Context, limit int, offset int) (*ResponseApiKeyHistoryListResponse, error) {
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
