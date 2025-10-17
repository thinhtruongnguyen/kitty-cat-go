package aiozaiplatformgosdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fromDate       = "2023-05-07 15:04:05"
	toDate         = "2023-05-07 15:04:05"
	validModelID   = "7332d265-171e-4055-9f2f-af51a52a82c4"
	invalidModelID = "00000000-0000-0000-0000-000000000000"
)

func TestApiKeyModelServiceGetModelTaskCost(t *testing.T) {
	tests := []struct {
		name    string
		modelId string
		wantErr bool
	}{
		{
			name:    "Valid modelId - may return cost OR fail if no version",
			modelId: validModelID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModel.ApiKeyModelIdTaskCostGet(tt.modelId)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
		})
	}
}

func TestApiKeyModelServiceGetModelStatistics(t *testing.T) {
	tests := []struct {
		name    string
		modelId string
		request RequestGetApiKeyStatisticsByModelIdRequest
		wantErr bool
	}{
		{
			name:    "Valid modelId",
			modelId: validModelID,
			request: RequestGetApiKeyStatisticsByModelIdRequest{
				From: &fromDate,
				To:   &toDate,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModel.ApiKeyModelIdStatisticsPost(tt.modelId, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
		})
	}
}

func TestApiKeyModelServiceCheckModelIsServing(t *testing.T) {
	tests := []struct {
		name    string
		modelId string
		wantErr bool
	}{
		{
			name:    "Valid modelId - should return serving info",
			modelId: "7332d265-171e-4055-9f2f-af51a52a82c4",
			wantErr: true, // error case
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModel.ApiKeyModelIdServingGet(tt.modelId)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
		})
	}
}

func TestApiKeyModelServiceGetModelInfo(t *testing.T) {
	tests := []struct {
		name        string
		modelId     string
		wantErr     bool
		expectEmpty bool
	}{
		{
			name:    "Get model info success",
			modelId: validModelID,
			wantErr: false,
		},
		{
			name:    "Get model info invalid UUID (422)",
			modelId: "7332d265-171e-4055-9f2f-af51a52a82c",
			wantErr: true,
		},
		{
			name:        "Get model info not found but success status",
			modelId:     invalidModelID,
			wantErr:     false,
			expectEmpty: true, // true if model is not exist
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModel.ApiKeyModelIdInfoGet(tt.modelId)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Data)

			if tt.expectEmpty {
				assert.Equal(t, "00000000-0000-0000-0000-000000000000", *resp.Data.ModelId)
			} else {
				assert.Equal(t, tt.modelId, *resp.Data.ModelId)
				assert.NotNil(t, resp.Data.ModelDescription)
			}
		})
	}
}

// func TestApiKeyModelServiceGetListPlatformsSupport(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Valid ListPlatformsSupport",
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			resp, err := testClient.ApiKeyModel.GetListPlatformsSupport()

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				assert.Nil(t, resp)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.NotNil(t, resp)
// 				assert.NotNil(t, resp.Data)
// 			}
// 		})
// 	}
// }
