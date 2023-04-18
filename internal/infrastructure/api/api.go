package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL string // The base URL of the API
}

// NewAPIClient creates a new instance of APIClient
func NewAPIClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

func (c *Client) Get(path string, response interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to call API: %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}
	return nil
}
