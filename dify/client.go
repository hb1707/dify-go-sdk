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

// ClientOption 定义客户端选项接口
type ClientOption interface {
	apply(*Client)
}

// clientOptionFunc 是一个适配器，允许使用普通函数作为 ClientOption
type clientOptionFunc func(*Client)

func (f clientOptionFunc) apply(c *Client) {
	f(c)
}

// WithBaseURL 设置自定义的基础 URL
func WithBaseURL(baseURL string) ClientOption {
	return clientOptionFunc(func(c *Client) {
		c.BaseURL = baseURL
	})
}

// WithHTTPClient 设置自定义的 HTTP 客户端
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return clientOptionFunc(func(c *Client) {
		c.HTTPClient = httpClient
	})
}

// NewClient creates a new Dify API client
func NewClient(apiKey string, opts ...ClientOption) *Client {
	httpClient := &http.Client{
		Timeout: time.Minute * 10,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	c := &Client{
		BaseURL:    DefaultBaseURL,
		APIKey:     apiKey,
		HTTPClient: httpClient,
	}

	// 应用选项
	for _, opt := range opts {
		opt.apply(c)
	}

	return c
}
