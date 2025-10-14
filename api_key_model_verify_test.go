package aiozaiplatformgosdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	invalidTaskID = "aea6944c-a808-4c11-b8de-557fe02d505f"
	MODEL_ID      = "7332d265-171e-4055-9f2f-af51a52a82c4"
)

var (
	validCommitHash = "master"
)

func TestApiKeyModelVerifyServiceCalculateCostToVerifyModel(t *testing.T) {
	tests := []struct {
		name    string
		modelId string
		request RequestCalculateCostToVerifyAiModelRequest
		wantErr bool
	}{
		{
			name:    "Calculate cost for model not setup",
			modelId: invalidModelID,
			request: RequestCalculateCostToVerifyAiModelRequest{
				CommitHash: "master",
				Platforms:  []string{"window"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVerify.CalculateCostToVerifyModel(tt.modelId, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			if err != nil {
				return
			}

			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
			t.Logf("Estimated cost: %v", resp.Data)
		})
	}
}

func TestApiKeyModelVerifyServiceVerifyModel(t *testing.T) {
	tests := []struct {
		name    string
		modelId string
		request RequestVerifyAiModelRequest
		wantErr bool
	}{
		{
			name:    "Verify valid model",
			modelId: MODEL_ID,
			request: RequestVerifyAiModelRequest{
				CommitHash: &validCommitHash,
				Platforms:  []string{"window"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVerify.VerifyModel(tt.modelId, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			if err != nil {
				return
			}

			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
			t.Logf("Verify task ID: %v", *resp.Data)
		})
	}
}

func TestApiKeyModelVerifyServicePreCheckToVerifyModel(t *testing.T) {
	tests := []struct {
		name    string
		modelId string
		request RequestCheckValidToVerifyAiModelRequest
		wantErr bool
	}{
		{
			name:    "PreCheck on model not set up",
			modelId: MODEL_ID,
			request: RequestCheckValidToVerifyAiModelRequest{
				CommitHash: "master",
				Platforms:  []string{"window"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVerify.PreCheckToVerifyModel(tt.modelId, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				if resp != nil {
					t.Logf("Response status: %v, message: %v", resp.Status, resp.Message)
				}
				return
			}

			if err != nil {
				return
			}

			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
		})
	}
}

func TestApiKeyModelVerifyServiceGetListVerifyModelTaskByCommitHashAndStatus(t *testing.T) {
	tests := []struct {
		name       string
		modelId    string
		commitHash string
		status     string
		wantErr    bool
	}{
		{
			name:       "Valid commit hash with verified status",
			modelId:    "YOUR_MODEL_ID_WITH_VERSION",
			commitHash: "VALID_COMMIT_HASH",
			status:     "verified",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVerify.GetListVerifyModelTaskByCommitHashAndStatus(
				tt.modelId, tt.commitHash, tt.status,
			)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
		})
	}
}

func TestApiKeyModelVerifyServiceGetVerifyPlatformTaskById(t *testing.T) {
	tests := []struct {
		name    string
		taskID  string
		wantErr bool
	}{
		{
			name:    "Invalid task ID",
			taskID:  invalidTaskID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVerify.GetVerifyPlatformTaskById(tt.taskID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
		})
	}
}

func TestApiKeyModelVerifyServiceGetModelVersioningByTaskId(t *testing.T) {
	tests := []struct {
		name    string
		taskID  string
		wantErr bool
	}{
		{
			name:    "Invalid task ID",
			taskID:  invalidTaskID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVerify.GetModelVersioningByTaskId(tt.taskID)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
		})
	}
}
