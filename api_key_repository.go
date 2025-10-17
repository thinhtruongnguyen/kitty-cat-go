package aiozaiplatformgosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ApiKeyRepositoryServiceI interface {
	ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGetWithContext(ctx context.Context, req GetCommitHistoryRequest) (*ResponseGetCommitHistoryResponse, error)
	ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet(req GetCommitHistoryRequest) (*ResponseGetCommitHistoryResponse, error)
}

type ApiKeyRepositoryService struct {
	Client *ClientV2
}

func (service *ApiKeyRepositoryService) ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGet(
	req GetCommitHistoryRequest,
) (*ResponseGetCommitHistoryResponse, error) {
	return service.ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGetWithContext(context.Background(), req)
}

func (service *ApiKeyRepositoryService) ApiKeyRepositoryOwnerUsernameRepositoryNameCommitHistoryGetWithContext(
	ctx context.Context,
	req GetCommitHistoryRequest,
) (*ResponseGetCommitHistoryResponse, error) {

	endpoint := fmt.Sprintf("/api-key/repository/%s/%s/commit/history", req.OwnerUsername, req.RepositoryName)
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("path", req.Path)
	localVarQueryParams.Add("sha", req.Sha)
	localVarQueryParams.Add("repoType", req.RepoType)
	localVarQueryParams.Add("page", fmt.Sprintf("%d", req.Page))
	localVarQueryParams.Add("pageSize", fmt.Sprintf("%d", req.PageSize))

	httpReq, err := service.Client.prepareRequest(ctx, http.MethodGet, endpoint, nil, nil, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseGetCommitHistoryResponse)
	if err := service.Client.do(httpReq, res); err != nil {
		return nil, err
	}

	return res, nil
}
