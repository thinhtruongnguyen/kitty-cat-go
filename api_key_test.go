package aiozaiplatformgosdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiKeyServiceCreateTask(t *testing.T) {
	tests := []struct {
		name    string
		request RequestDistributeTaskRequest
		wantErr bool
	}{
		{
			name: "Valid model with version",
			request: RequestDistributeTaskRequest{
				ModelId: &validModelID,
				Files: []map[string]interface{}{
					{
						"key":  "input",
						"data": "https://s3.w3s.aioz.network/w3ai-platform-v2/uploads/samples/sample.jpg",
						"name": "animetosketch.jpg",
					},
				},
				InputParams: map[string]interface{}{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiApiKey.CreateTask(tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			if err != nil {
				return
			}

			assert.NotNil(t, resp)
			if resp.Status != nil {
				assert.Equal(t, "success", *resp.Status)
			}
			assert.NotNil(t, resp.Data)
		})
	}
}

func TestApiKeyServiceGetTaskResult(t *testing.T) {
	tests := []struct {
		name      string
		taskID    string
		wantError bool
	}{
		{
			name:      "Invalid TaskID returns fail status",
			taskID:    "aea6944c-a808-4c11-b8de-557fe02d505f",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiApiKey.GetTaskResult(tt.taskID)

			if tt.wantError {
				if err != nil {
					return
				}
				assert.NotNil(t, resp, "Response must not be nil")
				assert.Equal(t, "fail", *resp.Status)
				assert.NotNil(t, resp.Message, "Message should explain the error")
			} else {
				assert.NoError(t, err, "API call should not return client error")
				assert.NotNil(t, resp, "Response must not be nil")
				assert.Equal(t, "success", *resp.Status)
				assert.NotNil(t, resp.Data, "Task result must not be nil on success")
			}
		})
	}

}

func TestApiKeyServiceGetTaskHistory(t *testing.T) {
	tests := []struct {
		name    string
		limit   int
		offset  int
		wantErr bool
	}{
		{
			name:    "Valid Task History",
			limit:   5,
			offset:  0,
			wantErr: false,
		},
		{
			name:    "Invalid Offset (Expect Error)",
			limit:   5,
			offset:  -2,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiApiKey.GetTaskHistory(tt.limit, tt.offset)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.Data)
			}
		})
	}
}

func TestApiKeyServiceGetApiKeyPermission(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Valid ApiKeyPermission",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiApiKey.GetApiKeyPermission()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotNil(t, resp.Data)
				assert.NotNil(t, resp.Data.LimitModels)
				assert.Equal(t, false, *resp.Data.LimitModels)
			}
		})
	}
}

func TestApiKeyServiceGetBalance(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Valid Balance Request",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.ApiApiKey.GetBalance()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)

			if resp != nil {
				assert.NotNil(t, resp.Status)
				assert.Equal(t, "success", *resp.Status)

				if resp.Data != nil {
					assert.NotNil(t, resp.Data.Balance)
					assert.NotNil(t, resp.Data.FreeBalance)
					assert.NotNil(t, resp.Data.WalletAddress)
				}
			}
		})
	}
}
