package aiozaiplatformgosdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	TypeJSON                 = "application/json"
	DEFAULT_ENDPOINT_URL     = "http://143.198.81.10:9000/"
	DEFAULT_ENDPOINT_VERSION = "api/v1"
)

type ClientV2 struct {
	BaseURL               *url.URL
	EndpointVersion       string
	ApiKey                string
	HttpClient            Doer
	ApiKeyModel           ApiKeyModelServiceI
	ApiApiKey             ApiKeyServiceI
	ApiKeyModelVerify     ApiKeyModelVerifyServiceI
	ApiKeyModelVersioning ApiKeyModelVersioningServiceI
	ApiKeyRepository      ApiKeyRepositoryServiceI
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type AuthCredentials struct {
	ApiKey string `json:"api_key"`
}

type Builder struct {
	BaseURL         string
	EndpointVersion string
	ApiKey          string
	HttpClient      Doer
}

type ErrorResponse struct {
	Response *http.Response
	Body     []byte
}

func (r *ErrorResponse) Error() string {
	var bodyStr string
	if len(r.Body) > 0 {
		var jsonBody interface{}
		if err := json.Unmarshal(r.Body, &jsonBody); err == nil {
			if prettyJSON, err := json.MarshalIndent(jsonBody, "", "  "); err == nil {
				bodyStr = string(prettyJSON)
			}
		}
		if bodyStr == "" {
			bodyStr = string(r.Body)
		}
	} else {
		bodyStr = "Empty body"
	}

	return fmt.Sprintf("\n%s\n",
		bodyStr,
	)
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	return &ErrorResponse{Response: r, Body: body}
}

func (c *ClientV2) prepareRequest(
	ctx context.Context,
	method, urlStr string,
	body any,
	headerParams map[string]string,
	queryParams url.Values,
) (*http.Request, error) {

	u, err := url.Parse(c.BaseURL.String())
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(strings.TrimRight(u.Path, "/"), strings.TrimLeft(urlStr, "/"))

	query := u.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	u.RawQuery = query.Encode()

	var buf *bytes.Buffer
	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	} else {
		buf = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", TypeJSON)
	if body != nil {
		req.Header.Set("Content-Type", TypeJSON)
	}
	req.Header.Set("x-api-key", c.ApiKey)

	for k, v := range headerParams {
		req.Header.Set(k, v)
	}

	return req, nil
}

func (c *ClientV2) do(req *http.Request, v interface{}) error {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return err
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}
	return nil
}

func ClientBuilder(credentials AuthCredentials) *Builder {
	return &Builder{
		BaseURL:         DEFAULT_ENDPOINT_URL,
		EndpointVersion: DEFAULT_ENDPOINT_VERSION,
		ApiKey:          credentials.ApiKey,
	}
}

func (cb *Builder) WithBaseURL(url string) *Builder {
	cb.BaseURL = url
	return cb
}

func (cb *Builder) WithEndpointVersion(version string) *Builder {
	cb.EndpointVersion = version
	return cb
}

func (cb *Builder) WithHttpClient(client Doer) *Builder {
	cb.HttpClient = client
	return cb
}

func (cb *Builder) Build() *ClientV2 {
	fullURL := strings.TrimRight(cb.BaseURL, "/")
	if cb.EndpointVersion != "" {
		fullURL = fullURL + "/" + strings.TrimLeft(cb.EndpointVersion, "/")
	}
	baseURL, _ := url.Parse(fullURL)

	httpClient := cb.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &ClientV2{
		BaseURL:         baseURL,
		EndpointVersion: cb.EndpointVersion,
		ApiKey:          cb.ApiKey,
		HttpClient:      httpClient,
	}

	client.ApiKeyModel = &ApiKeyModelService{Client: client}
	client.ApiApiKey = &ApiKeyService{Client: client}
	client.ApiKeyModelVerify = &ApiKeyModelVerifyService{Client: client}
	client.ApiKeyModelVersioning = &ApiKeyModelVersioningService{Client: client}
	client.ApiKeyRepository = &ApiKeyRepositoryService{Client: client}

	return client
}
