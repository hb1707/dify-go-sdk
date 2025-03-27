package knowledge

import (
	"net/http"
	"strings"
)

// Client 实现 Client 接口
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// NewClient 创建新的知识库客户端
func NewClient(apiKey string, opts ...Option) *Client {
	// 确保 API key 以 "Bearer " 开头
	if !strings.HasPrefix(apiKey, "Bearer ") {
		apiKey = "Bearer " + apiKey
	}

	c := &Client{
		httpClient: &http.Client{},
		baseURL:    "https://dify.aix101.com/v1",
		apiKey:     apiKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Option 客户端配置选项
type Option func(*Client)

// WithHTTPClient 设置自定义 HTTP 客户端
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL 设置自定义基础 URL
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}
