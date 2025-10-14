package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyModelServiceI interface {
	GetModelInfo(modelId string) (*ResponseApiKeyInfoResponse, error)
	GetModelInfoWithContext(ctx context.Context, modelId string) (*ResponseApiKeyInfoResponse, error)
	GetListPlatformsSupportWithContext(ctx context.Context) (*ResponseGetListPlatformSupportResponse, error)
	GetListPlatformsSupport() (*ResponseGetListPlatformSupportResponse, error)
	CheckModelIsServing(modelId string) (*ResponseCheckModelIsServingResponse, error)
	CheckModelIsServingWithContext(ctx context.Context, modelId string) (*ResponseCheckModelIsServingResponse, error)
	GetModelStatistics(modelId string, request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error)
	GetModelStatisticsWithContext(ctx context.Context, modelId string, request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error)
	GetModelTaskCostWithContext(ctx context.Context, modelId string) (*ResponseEstimateCostResponse, error)
	GetModelTaskCost(modelId string) (*ResponseEstimateCostResponse, error)
}

type ApiKeyModelService struct {
	Client *ClientV2
}

func (service *ApiKeyModelService) GetModelTaskCost(modelId string) (*ResponseEstimateCostResponse, error) {
	return service.GetModelTaskCostWithContext(context.Background(), modelId)
}
func (service *ApiKeyModelService) GetModelTaskCostWithContext(ctx context.Context, modelId string) (*ResponseEstimateCostResponse, error) {
	localVarPath := fmt.Sprintf("/api-key/model/%s/task/cost", url.PathEscape(modelId))
	req, err := service.Client.prepareRequest(ctx, http.MethodGet, localVarPath, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	res := new(ResponseEstimateCostResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyModelService) GetModelStatistics(
	modelId string,
	request RequestGetApiKeyStatisticsByModelIdRequest,
) (*ResponseGetTaskStatisticsResponse, error) {
	return service.GetModelStatisticsWithContext(context.Background(), modelId, request)
}
func (service *ApiKeyModelService) GetModelStatisticsWithContext(
	ctx context.Context,
	modelId string,
	request RequestGetApiKeyStatisticsByModelIdRequest,
) (*ResponseGetTaskStatisticsResponse, error) {

	localVarPath := fmt.Sprintf("/api-key/model/%s/statistics", url.PathEscape(modelId))
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

	res := new(ResponseGetTaskStatisticsResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyModelService) CheckModelIsServing(modelId string) (*ResponseCheckModelIsServingResponse, error) {
	return service.CheckModelIsServingWithContext(context.Background(), modelId)
}
func (service *ApiKeyModelService) CheckModelIsServingWithContext(ctx context.Context, modelId string) (*ResponseCheckModelIsServingResponse, error) {
	localVarPath := fmt.Sprintf("/api-key/model/%s/serving", url.PathEscape(modelId))

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, localVarPath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	res := new(ResponseCheckModelIsServingResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyModelService) GetListPlatformsSupport() (*ResponseGetListPlatformSupportResponse, error) {
	return service.GetListPlatformsSupportWithContext(context.Background())
}
func (service *ApiKeyModelService) GetListPlatformsSupportWithContext(ctx context.Context) (*ResponseGetListPlatformSupportResponse, error) {
	localVarPath := "/api-key/model/verify/support/platforms"

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, localVarPath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	res := new(ResponseGetListPlatformSupportResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (service *ApiKeyModelService) GetModelInfo(modelId string) (*ResponseApiKeyInfoResponse, error) {
	return service.GetModelInfoWithContext(context.Background(), modelId)
}
func (service *ApiKeyModelService) GetModelInfoWithContext(ctx context.Context, modelId string) (*ResponseApiKeyInfoResponse, error) {
	localVarPath := fmt.Sprintf("/api-key/model/%s/info", url.PathEscape(modelId))

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, localVarPath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	res := new(ResponseApiKeyInfoResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
