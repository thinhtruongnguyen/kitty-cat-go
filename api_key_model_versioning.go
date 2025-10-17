package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyModelVersioningServiceI interface {
	ApiKeyModelIdVersioningGetWithContext(ctx context.Context, modelId string) (*ResponseModelVersioningGroupLiteResponse, error)
	ApiKeyModelIdVersioningGet(modelId string) (*ResponseModelVersioningGroupLiteResponse, error)
	ApiKeyModelIdVersioningListGetWithContext(ctx context.Context, modelId string, offset, limit int, verifyStatus string) (*ResponseGetListModelVersioningLiteResponse, error)
	ApiKeyModelIdVersioningListGet(modelId string, offset, limit int, verifyStatus string) (*ResponseGetListModelVersioningLiteResponse, error)
	ApiKeyModelIdVersioningPut(modelId, commitHash string) (*ResponseSuccessResponse, error)
	ApiKeyModelIdVersioningPutWithContext(ctx context.Context, modelId, commitHash string) (*ResponseSuccessResponse, error)
	ApiKeyModelIdVersioningDeleteWithContext(ctx context.Context, modelId, commitHash string) (*ResponseSuccessResponse, error)
	ApiKeyModelIdVersioningDelete(modelId, commitHash string) (*ResponseSuccessResponse, error)
}

type ApiKeyModelVersioningService struct {
	Client *ClientV2
}

const modelVersioningPathFmt = "/api-key/model/%s/versioning"

func (service *ApiKeyModelVersioningService) ApiKeyModelIdVersioningDelete(
	modelId, commitHash string,
) (*ResponseSuccessResponse, error) {
	return service.ApiKeyModelIdVersioningDeleteWithContext(context.Background(), modelId, commitHash)
}
func (service *ApiKeyModelVersioningService) ApiKeyModelIdVersioningDeleteWithContext(
	ctx context.Context,
	modelId, commitHash string,
) (*ResponseSuccessResponse, error) {

	localVarPath := fmt.Sprintf(modelVersioningPathFmt, modelId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("commitHash", commitHash)

	var localVarPostBody any

	req, err := service.Client.prepareRequest(
		ctx,
		http.MethodDelete,
		localVarPath,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
	)
	if err != nil {
		return nil, err
	}

	res := new(ResponseSuccessResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
func (service *ApiKeyModelVersioningService) ApiKeyModelIdVersioningPut(
	modelId, commitHash string,
) (*ResponseSuccessResponse, error) {
	return service.ApiKeyModelIdVersioningPutWithContext(context.Background(), modelId, commitHash)
}
func (service *ApiKeyModelVersioningService) ApiKeyModelIdVersioningPutWithContext(
	ctx context.Context,
	modelId, commitHash string,
) (*ResponseSuccessResponse, error) {
	localVarPath := fmt.Sprintf(modelVersioningPathFmt, modelId)

	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("commitHash", commitHash)

	localVarHeaderParams := make(map[string]string)

	var localVarPostBody any

	req, err := service.Client.prepareRequest(
		ctx,
		http.MethodPut,
		localVarPath,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
	)
	if err != nil {
		return nil, err
	}

	res := new(ResponseSuccessResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
func (service *ApiKeyModelVersioningService) ApiKeyModelIdVersioningListGet(
	modelId string,
	offset, limit int,
	verifyStatus string,
) (*ResponseGetListModelVersioningLiteResponse, error) {
	return service.ApiKeyModelIdVersioningListGetWithContext(context.Background(), modelId, offset, limit, verifyStatus)
}
func (service *ApiKeyModelVersioningService) ApiKeyModelIdVersioningListGetWithContext(
	ctx context.Context,
	modelId string,
	offset, limit int,
	verifyStatus string,
) (*ResponseGetListModelVersioningLiteResponse, error) {

	endpoint := fmt.Sprintf("/api-key/model/%s/versioning/list", modelId)

	var localVarPostBody any
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("offset", fmt.Sprintf("%d", offset))
	localVarQueryParams.Add("limit", fmt.Sprintf("%d", limit))
	if verifyStatus != "" {
		localVarQueryParams.Add("verifyStatus", verifyStatus)
	}

	req, err := service.Client.prepareRequest(
		ctx,
		http.MethodGet,
		endpoint,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
	)
	if err != nil {
		return nil, err
	}

	res := new(ResponseGetListModelVersioningLiteResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
func (service *ApiKeyModelVersioningService) ApiKeyModelIdVersioningGet(
	modelId string) (*ResponseModelVersioningGroupLiteResponse, error) {
	return service.ApiKeyModelIdVersioningGetWithContext(context.Background(), modelId)
}
func (service *ApiKeyModelVersioningService) ApiKeyModelIdVersioningGetWithContext(
	ctx context.Context,
	modelId string) (*ResponseModelVersioningGroupLiteResponse, error) {
	endpoint := fmt.Sprintf("/api-key/model/%s/versioning", modelId)
	var localVarPostBody any
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	req, err := service.Client.prepareRequest(ctx, http.MethodGet, endpoint, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseModelVersioningGroupLiteResponse)
	if err := service.Client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}
