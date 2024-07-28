package user_service

import (
	"context"
	"net/http"
	"time"

	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/rest"
)

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

func (c *client) GetByPhoneNumberAndPassword(ctx context.Context, phoneNumber, password string) (User, error) {
	var u User
	c.base.Logger.Debug("return from GetByPhoneNumberAndPassword")
	return u, nil
}

func (c *client) GetByUUID(ctx context.Context, uuid string) (User, error) {
	var u User
	c.base.Logger.Debug("return from GetByUUID")
	return u, nil
}

func (c *client) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	var u User
	c.base.Logger.Debug("return from Create")
	return u, nil
}
