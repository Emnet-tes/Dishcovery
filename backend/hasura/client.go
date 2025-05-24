package hasura

import (
	"backend/config"
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type Client struct {
	client *graphql.Client
	config *config.Config
}

func NewClient(cfg *config.Config) *Client {
	client := graphql.NewClient(cfg.HasuraEndpoint)
	client.Log = func(s string) { fmt.Println(s) }
	return &Client{
		client: client,
		config: cfg,
	}
}

func (c *Client) Execute(ctx context.Context, query string, variables map[string]interface{}, response interface{}) error {
	req := graphql.NewRequest(query)
	req.Header.Set("x-hasura-admin-secret", c.config.HasuraAdminKey)

	for key, value := range variables {
		req.Var(key, value)
	}

	if err := c.client.Run(ctx, req, response); err != nil {
		return fmt.Errorf("hasura query failed: %v", err)
	}

	return nil
}

// HasuraActionResponse represents the standard response format for Hasura actions
type HasuraActionResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewActionResponse(code string, message string, data interface{}) HasuraActionResponse {
	return HasuraActionResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
