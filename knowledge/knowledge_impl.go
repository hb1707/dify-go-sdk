package knowledge

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CreateKnowledge 创建知识库
func (c *Client) CreateKnowledge(ctx context.Context, req *CreateKnowledgeRequest) (*Knowledge, error) {
	url := fmt.Sprintf("%s/datasets", c.baseURL)
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	httpReq.Header.Set("Authorization", c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var knowledge Knowledge
	if err := json.NewDecoder(resp.Body).Decode(&knowledge); err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	return &knowledge, nil
}

// ListKnowledge 列出知识库
func (c *Client) ListKnowledge(ctx context.Context, req *ListKnowledgeRequest) (*ListKnowledgeResponse, error) {
	url := fmt.Sprintf("%s/datasets", c.baseURL)

	// 构建查询参数
	query := make(map[string]interface{})
	if req.Page > 0 {
		query["page"] = req.Page
	}
	if req.Limit > 0 {
		query["limit"] = req.Limit
	}
	if req.Keyword != "" {
		query["keyword"] = req.Keyword
	}
	if req.SortBy != "" {
		query["sort_by"] = req.SortBy
	}
	if req.SortOrder != "" {
		query["sort_order"] = req.SortOrder
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	httpReq.Header.Set("Authorization", c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response ListKnowledgeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	return &response, nil
}

// DeleteKnowledge 删除知识库
func (c *Client) DeleteKnowledge(ctx context.Context, knowledgeID string) error {
	url := fmt.Sprintf("%s/datasets/%s", c.baseURL, knowledgeID)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	httpReq.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// Retrieve 检索知识库
func (c *Client) Retrieve(ctx context.Context, datasetID string, req *RetrieveRequest) (*RetrieveResponse, error) {
	url := fmt.Sprintf("%s/datasets/%s/retrieve", c.baseURL, datasetID)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %v", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	httpReq.Header.Set("Authorization", c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d，响应: %s", resp.StatusCode, string(body))
	}
	var result RetrieveResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}
	return &result, nil
}
