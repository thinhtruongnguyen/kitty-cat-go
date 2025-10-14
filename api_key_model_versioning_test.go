package aiozaiplatformgosdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiKeyModelVersioningServiceDeleteModelVersioningByModelIdAndCommitHash(t *testing.T) {
	tests := []struct {
		name       string
		modelId    string
		commitHash string
		wantErr    bool
	}{
		{
			name:       "Non-existing commitHash - expect fail",
			modelId:    validModelID,
			commitHash: "ba271f290d732d6aa5902bf4d42103b516fcb7c2",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVersioning.DeleteModelVersioningByModelIdAndCommitHash(tt.modelId, tt.commitHash)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			if resp.Status != nil {
				assert.Equal(t, "success", *resp.Status)
			}
		})
	}
}

func TestApiKeyModelVersioningServiceChangeCurrentModelVersioningByModelIdAndCommitHash(t *testing.T) {
	tests := []struct {
		name       string
		modelId    string
		commitHash string
		wantErr    bool
	}{
		{
			name:       "Invalid commitHash - expect error",
			modelId:    validModelID,
			commitHash: "ba271f290d732d6aa5902bf4d42103b516fcb7c2",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVersioning.ChangeCurrentModelVersioningByModelIdAndCommitHash(
				tt.modelId, tt.commitHash,
			)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "success", *resp.Status)
		})
	}
}

func TestApiKeyServiceGetListVerifiedModelVersioning(t *testing.T) {
	tests := []struct {
		name          string
		modelId       string
		offset, limit int
		verifyStatus  string
		wantErr       bool
	}{
		{
			name:         "Valid model - expect success even if empty list",
			modelId:      validModelID,
			offset:       0,
			limit:        10,
			verifyStatus: "verified",
			wantErr:      false,
		},
		{
			name:         "Invalid offset - expect 422 error",
			modelId:      validModelID,
			offset:       -99,
			limit:        10,
			verifyStatus: "verified",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVersioning.GetListVerifiedModelVersioning(
				tt.modelId, tt.offset, tt.limit, tt.verifyStatus,
			)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)

			if resp.Status != nil {
				assert.Equal(t, "success", *resp.Status)
			}

			if resp.Data != nil && resp.Data.Total != nil {
				assert.GreaterOrEqual(t, int64(*resp.Data.Total), int64(0), "Total should be >= 0")
			}
		})
	}
}

func TestApiKeyModelVersioningServiceGetCurrentModelVersioningByModelId(t *testing.T) {
	tests := []struct {
		name    string
		modelId string
		wantErr bool
	}{
		{
			name:    "Model has no versioning - expect 404",
			modelId: "7332d265-171e-4055-9f2f-af51a52a82c4",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyModelVersioning.GetCurrentModelVersioningByModelId(tt.modelId)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Status)
			assert.Equal(t, "success", *resp.Status)
			assert.NotNil(t, resp.Data)
			t.Logf("Current versioning: %+v", resp.Data)
		})
	}
}
