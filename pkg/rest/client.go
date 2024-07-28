package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
)

type BaseClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Logger     logging.Logger
}

func (c *BaseClient) SendRequest(req *http.Request) (*APIResponse, error) {
	if c.HTTPClient == nil {
		return nil, errors.New("no http client")
	}
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request. error: %w", err)
	}

	apiResponse := APIResponse{
		IsOk:     true,
		response: response,
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		apiResponse.IsOk = false
		defer response.Body.Close()

		var apiErr APIError
		if err = json.NewDecoder(response.Body).Decode(&apiErr); err == nil {
			apiResponse.Error = apiErr
		}
	}
	return &apiResponse, nil
}
