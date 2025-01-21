package dify

import (
	"net/http"
	"time"
)

// Client represents a client for the Dify API
type Client struct {
	// BaseURL is the base URL for API requests
	BaseURL string
	// APIKey is the API key for authentication
	APIKey string
	// HTTPClient is the HTTP client used for making requests
	HTTPClient *http.Client
}

// NewClient creates a new Dify API client
func NewClient(apiKey string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	return &Client{
		BaseURL:    DefaultBaseURL,
		APIKey:     apiKey,
		HTTPClient: httpClient,
	}
}
