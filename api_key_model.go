package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyModelServiceI interface {
	ApiKeyModelIdInfoGet(modelId string) (*ResponseApiKeyInfoResponse, error)
	ApiKeyModelIdInfoGetWithContext(ctx context.Context, modelId string) (*ResponseApiKeyInfoResponse, error)
	ApiKeyModelVerifySupportPlatformsGetWithContext(ctx context.Context) (*ResponseGetListPlatformSupportResponse, error)
	ApiKeyModelVerifySupportPlatformsGet() (*ResponseGetListPlatformSupportResponse, error)
	ApiKeyModelIdServingGet(modelId string) (*ResponseCheckModelIsServingResponse, error)
	ApiKeyModelIdServingGetWithContext(ctx context.Context, modelId string) (*ResponseCheckModelIsServingResponse, error)
	ApiKeyModelIdStatisticsPost(modelId string, request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error)
	ApiKeyModelIdStatisticsPostWithContext(ctx context.Context, modelId string, request RequestGetApiKeyStatisticsByModelIdRequest) (*ResponseGetTaskStatisticsResponse, error)
	ApiKeyModelIdTaskCostGetWithContext(ctx context.Context, modelId string) (*ResponseEstimateCostResponse, error)
	ApiKeyModelIdTaskCostGet(modelId string) (*ResponseEstimateCostResponse, error)
}

type ApiKeyModelService struct {
	Client *ClientV2
}

func (service *ApiKeyModelService) ApiKeyModelIdTaskCostGet(modelId string) (*ResponseEstimateCostResponse, error) {
	return service.ApiKeyModelIdTaskCostGetWithContext(context.Background(), modelId)
}
func (service *ApiKeyModelService) ApiKeyModelIdTaskCostGetWithContext(ctx context.Context, modelId string) (*ResponseEstimateCostResponse, error) {
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

func (service *ApiKeyModelService) ApiKeyModelIdStatisticsPost(
	modelId string,
	request RequestGetApiKeyStatisticsByModelIdRequest,
) (*ResponseGetTaskStatisticsResponse, error) {
	return service.ApiKeyModelIdStatisticsPostWithContext(context.Background(), modelId, request)
}
func (service *ApiKeyModelService) ApiKeyModelIdStatisticsPostWithContext(
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

func (service *ApiKeyModelService) ApiKeyModelIdServingGet(modelId string) (*ResponseCheckModelIsServingResponse, error) {
	return service.ApiKeyModelIdServingGetWithContext(context.Background(), modelId)
}
func (service *ApiKeyModelService) ApiKeyModelIdServingGetWithContext(ctx context.Context, modelId string) (*ResponseCheckModelIsServingResponse, error) {
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

func (service *ApiKeyModelService) ApiKeyModelVerifySupportPlatformsGet() (*ResponseGetListPlatformSupportResponse, error) {
	return service.ApiKeyModelVerifySupportPlatformsGetWithContext(context.Background())
}
func (service *ApiKeyModelService) ApiKeyModelVerifySupportPlatformsGetWithContext(ctx context.Context) (*ResponseGetListPlatformSupportResponse, error) {
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

func (service *ApiKeyModelService) ApiKeyModelIdInfoGet(modelId string) (*ResponseApiKeyInfoResponse, error) {
	return service.ApiKeyModelIdInfoGetWithContext(context.Background(), modelId)
}
func (service *ApiKeyModelService) ApiKeyModelIdInfoGetWithContext(ctx context.Context, modelId string) (*ResponseApiKeyInfoResponse, error) {
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
