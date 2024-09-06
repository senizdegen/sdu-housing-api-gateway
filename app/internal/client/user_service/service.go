package user_service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/structs"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/apperror"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/rest"
)

type token struct {
	token string
}

type client struct {
	base     rest.BaseClient
	Resource string
}

func NewService(baseURL string, resource string, logger logging.Logger) UserService {
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

type UserService interface {
	GetByPhoneNumberAndPassword(ctx context.Context, phoneNumber, password string) (User, error)
	GetByUUID(ctx context.Context, uuid string) (User, error)
	Create(ctx context.Context, dto CreateUserDTO) (User, error)
}

func (c *client) GetByPhoneNumberAndPassword(ctx context.Context, phoneNumber, password string) (u User, err error) {
	c.base.Logger.Debug("add phone number and password to filter options")
	filters := []rest.FilterOptions{
		{
			Field:  "phone_number",
			Values: []string{phoneNumber},
		},
		{
			Field:  "password",
			Values: []string{password},
		},
	}

	c.base.Logger.Debug("build url with resource and filter")
	uri, err := c.base.BuildURL(c.Resource, filters)
	if err != nil {
		return u, fmt.Errorf("failed to build URL. error: %v", err)
	}

	c.base.Logger.Tracef("url: %s", uri)

	c.base.Logger.Debug("create new request")
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return u, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	c.base.Logger.Debug("send request")
	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := c.base.SendRequest(req)

	if err != nil {
		return u, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		if err = json.NewDecoder(response.Body()).Decode(&u); err != nil {
			return u, fmt.Errorf("failed to decode body due to error: %w", err)
		}
		c.base.Logger.Debug(response)
		return u, nil
	}

	return u, apperror.APIError(response.Error.ErrorCode, response.Error.Message, response.Error.DeveloperMessage)
}

func (c *client) GetByUUID(ctx context.Context, uuid string) (User, error) {
	var u User

	c.base.Logger.Debug("build url with resource and filter")
	uri, err := c.base.BuildURL(fmt.Sprintf("%s/%s", c.Resource, uuid), nil)
	if err != nil {
		return u, fmt.Errorf("failed to build URL. error: %v", err)
	}
	c.base.Logger.Tracef("url: %s", uri)

	c.base.Logger.Debug("create new request")
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return u, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	c.base.Logger.Debug("send request")
	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)

	response, err := c.base.SendRequest(req)
	if err != nil {
		return u, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		defer response.Body().Close()
		if err = json.NewDecoder(response.Body()).Decode(&u); err != nil {
			return u, fmt.Errorf("failed to decode body due to error: %w", err)
		}
		return u, nil
	}
	return u, apperror.APIError(response.Error.ErrorCode, response.Error.Message, response.Error.DeveloperMessage)
}

func (c *client) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	var u User

	c.base.Logger.Debug("build url with resource and filter")
	uri, err := c.base.BuildURL(c.Resource, nil)
	if err != nil {
		return u, fmt.Errorf("failed to build URL. error: %w", err)
	}
	c.base.Logger.Tracef("url: %s", uri)

	c.base.Logger.Debug("convert dto to map")
	structs.DefaultTagName = "json"
	data := structs.Map(dto)

	c.base.Logger.Debug("marshal map to bytes")
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return u, fmt.Errorf("failed to marshal dto")
	}

	c.base.Logger.Debug("create new request")
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(dataBytes))
	if err != nil {
		return u, fmt.Errorf("failed to create new  request due to error: %w", err)
	}

	c.base.Logger.Debug("send request")
	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req = req.WithContext(reqCtx)

	response, err := c.base.SendRequest(req)

	if err != nil {

		return u, fmt.Errorf("failed to send request due to error %w", err)
	}

	if response.IsOk {
		defer response.Body().Close()
		if err = json.NewDecoder(response.Body()).Decode(&u); err != nil {
			return u, fmt.Errorf("failed to decode body due to error: %w", err)
		}
		return u, nil
	}
	return u, apperror.APIError(response.Error.ErrorCode, response.Error.Message, response.Error.DeveloperMessage)
}
