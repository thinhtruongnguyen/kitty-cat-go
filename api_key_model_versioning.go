package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyModelVersioningServiceI interface {
	GetCurrentModelVersioningByModelIdWithContext(ctx context.Context, modelId string) (*ResponseModelVersioningGroupLiteResponse, error)
	GetCurrentModelVersioningByModelId(modelId string) (*ResponseModelVersioningGroupLiteResponse, error)
	GetListVerifiedModelVersioningWithContext(ctx context.Context, modelId string, offset, limit int, verifyStatus string) (*ResponseGetListModelVersioningLiteResponse, error)
	GetListVerifiedModelVersioning(modelId string, offset, limit int, verifyStatus string) (*ResponseGetListModelVersioningLiteResponse, error)
	ChangeCurrentModelVersioningByModelIdAndCommitHash(modelId, commitHash string) (*ResponseSuccessResponse, error)
	ChangeCurrentModelVersioningByModelIdAndCommitHashWithContext(ctx context.Context, modelId, commitHash string) (*ResponseSuccessResponse, error)
	DeleteModelVersioningByModelIdAndCommitHashWithContext(ctx context.Context, modelId, commitHash string) (*ResponseSuccessResponse, error)
	DeleteModelVersioningByModelIdAndCommitHash(modelId, commitHash string) (*ResponseSuccessResponse, error)
}

type ApiKeyModelVersioningService struct {
	Client *ClientV2
}

const modelVersioningPathFmt = "/api-key/model/%s/versioning"

func (service *ApiKeyModelVersioningService) DeleteModelVersioningByModelIdAndCommitHash(
	modelId, commitHash string,
) (*ResponseSuccessResponse, error) {
	return service.DeleteModelVersioningByModelIdAndCommitHashWithContext(context.Background(), modelId, commitHash)
}
func (service *ApiKeyModelVersioningService) DeleteModelVersioningByModelIdAndCommitHashWithContext(
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
func (service *ApiKeyModelVersioningService) ChangeCurrentModelVersioningByModelIdAndCommitHash(
	modelId, commitHash string,
) (*ResponseSuccessResponse, error) {
	return service.ChangeCurrentModelVersioningByModelIdAndCommitHashWithContext(context.Background(), modelId, commitHash)
}
func (service *ApiKeyModelVersioningService) ChangeCurrentModelVersioningByModelIdAndCommitHashWithContext(
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
func (service *ApiKeyModelVersioningService) GetListVerifiedModelVersioning(
	modelId string,
	offset, limit int,
	verifyStatus string,
) (*ResponseGetListModelVersioningLiteResponse, error) {
	return service.GetListVerifiedModelVersioningWithContext(context.Background(), modelId, offset, limit, verifyStatus)
}
func (service *ApiKeyModelVersioningService) GetListVerifiedModelVersioningWithContext(
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
func (service *ApiKeyModelVersioningService) GetCurrentModelVersioningByModelId(
	modelId string) (*ResponseModelVersioningGroupLiteResponse, error) {
	return service.GetCurrentModelVersioningByModelIdWithContext(context.Background(), modelId)
}
func (service *ApiKeyModelVersioningService) GetCurrentModelVersioningByModelIdWithContext(
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
