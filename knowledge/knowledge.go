package knowledge

import (
	"context"
	"io"
	"net/http"
	"strings"
)

// Client 知识库客户端接口
type Client interface {
	// 知识库基础操作
	CreateKnowledge(ctx context.Context, req *CreateKnowledgeRequest) (*Knowledge, error)
	ListKnowledge(ctx context.Context, req *ListKnowledgeRequest) (*ListKnowledgeResponse, error)
	DeleteKnowledge(ctx context.Context, knowledgeID string) error

	// 文档操作
	CreateDocumentByText(ctx context.Context, datasetID string, req *CreateDocumentByTextRequest) (*Document, error)
	CreateDocumentByFile(ctx context.Context, datasetID string, req *CreateDocumentByFileRequest, file io.Reader) (*Document, error)

	// 文档状态操作
	GetDocumentIndexingStatus(ctx context.Context, datasetID string, batch string) (*DocumentIndexingStatus, error)

	// 文档更新操作
	UpdateDocumentByText(ctx context.Context, datasetID string, documentID string, req *UpdateDocumentByTextRequest) (*Document, error)
	UpdateDocumentByFile(ctx context.Context, datasetID string, documentID string, req *UpdateDocumentByFileRequest, file io.Reader) (*Document, error)

	// 文档删除操作
	DeleteDocument(ctx context.Context, datasetID string, documentID string) error

	// 知识库检索操作
	Retrieve(ctx context.Context, datasetID string, req *RetrieveRequest) (*RetrieveResponse, error)
}

// client 实现 Client 接口
type client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// NewClient 创建新的知识库客户端
func NewClient(apiKey string, opts ...Option) Client {
	// 确保 API key 以 "Bearer " 开头
	if !strings.HasPrefix(apiKey, "Bearer ") {
		apiKey = "Bearer " + apiKey
	}

	c := &client{
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
type Option func(*client)

// WithHTTPClient 设置自定义 HTTP 客户端
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL 设置自定义基础 URL
func WithBaseURL(baseURL string) Option {
	return func(c *client) {
		c.baseURL = baseURL
	}
}
