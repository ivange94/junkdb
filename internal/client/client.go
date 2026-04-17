package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ivange94/junkdb/internal/config"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func New(cfg *config.Config) *Client {
	return &Client{
		baseURL:    fmt.Sprintf("http://%s", cfg.BindAddr),
		httpClient: &http.Client{},
	}
}

func (c *Client) Put(ctx context.Context, key, value string) error {
	endpoint := fmt.Sprintf("%s/api/v1/%s", c.baseURL, key)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader([]byte(value)))
	if err != nil {
		return fmt.Errorf("create post request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send put request: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("put failed with status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	endpoint := fmt.Sprintf("%s/api/v1/%s", c.baseURL, key)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("create get request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("fetch value: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response body: %w", err)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("get failed with status %d: %s", resp.StatusCode, string(body))
	}
	return string(body), nil
}
