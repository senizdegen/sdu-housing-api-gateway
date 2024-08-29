package property_service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/senizdegen/sdu-housing/api-gateway/internal/apperror"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/rest"
)

type client struct {
	base     rest.BaseClient
	Resource string
}

func NewService(baseURL string, resource string, logger logging.Logger) PropertyService {
	c := client{
		Resource: resource,
		base: rest.BaseClient{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			Logger: logger,
		},
	}
	return &c
}

type PropertyService interface {
	GetAll(ctx context.Context) ([]Property, error)
	Create(ctx context.Context, dto CreatePropertyDTO) (Property, error)
}

func (c *client) GetAll(ctx context.Context) ([]Property, error) {
	c.base.Logger.Debug("build url")
	uri, err := c.base.BuildURL(c.Resource, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL. error: %w", err)
	}
	c.base.Logger.Trace("url: %s", uri)

	c.base.Logger.Debug("create new request")
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	c.base.Logger.Debug("send request")
	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req = req.WithContext(reqCtx)
	response, err := c.base.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		var property []Property
		bytes, err := response.ReadBody()
		if err != nil {
			return nil, fmt.Errorf("failed to read response body. error: %w", err)
		}
		if err := json.Unmarshal(bytes, &property); err != nil {
			return nil, fmt.Errorf("failed to unmarshal bytes. error: %w", err)
		}
		return property, nil
	}

	return nil, apperror.APIError(response.Error.ErrorCode, response.Error.Message, response.Error.DeveloperMessage)

}

func (c *client) Create(ctx context.Context, dto CreatePropertyDTO) (Property, error) {
	return Property{}, nil
}
