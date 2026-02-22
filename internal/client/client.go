package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func Put(key, value string) error {
	endpoint := fmt.Sprintf("http://localhost:9429/api/v1/%s", key)
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, endpoint, bytes.NewReader([]byte(value)))
	if err != nil {
		return fmt.Errorf("error creating post request: %w", err)
	}
	_, err = httpClient.Do(req)
	return err
}

func Get(key string) (string, error) {
	endpoint := fmt.Sprintf("http://localhost:9429/api/v1/%s", key)
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("error creating get request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error fetching value: %w", err)
	}
	var value string
	return value, json.NewDecoder(resp.Body).Decode(&value)
}
