package aiozaiplatformgosdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiKeyRepositoryServiceGetCommitHistory(t *testing.T) {
	tests := []struct {
		name    string
		req     GetCommitHistoryRequest
		wantErr bool
	}{
		{
			name: "Repo not found still returns valid error",
			req: GetCommitHistoryRequest{
				OwnerUsername:  "bohaha123-3333",
				RepositoryName: "HEHEE",
				Path:           "",
				Sha:            "master",
				RepoType:       "model",
				Page:           1,
				PageSize:       10,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiKeyRepository.GetCommitHistory(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)

		})
	}
}
