package knowledge

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// CreateDocumentByText 通过文本创建文档
func (c *client) CreateDocumentByText(ctx context.Context, datasetID string, req *CreateDocumentByTextRequest) (*Document, error) {
	url := fmt.Sprintf("%s/datasets/%s/document/create-by-text", c.baseURL, datasetID)

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

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data  *Document `json:"data"`
		Batch string    `json:"batch"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}
	fmt.Printf("batch: %v", result.Batch)

	return result.Data, nil
}

// CreateDocumentByFile 通过文件创建文档
func (c *client) CreateDocumentByFile(ctx context.Context, datasetID string, req *CreateDocumentByFileRequest, file io.Reader) (*Document, error) {
	// 构造 API URL
	url := fmt.Sprintf("%s/datasets/%s/document/create-by-file", c.baseURL, datasetID)

	// 创建 multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件
	part, err := writer.CreateFormFile("file", "document.txt")
	if err != nil {
		return nil, fmt.Errorf("创建文件表单失败: %v", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("复制文件内容失败: %v", err)
	}

	// 构造符合 API 要求的请求结构
	apiReq := struct {
		Name              string                 `json:"name"`
		IndexingTechnique string                 `json:"indexing_technique"`
		DocForm           string                 `json:"doc_form"`
		DocType           string                 `json:"doc_type,omitempty"`
		DocMetadata       map[string]interface{} `json:"doc_metadata,omitempty"`
		ProcessRule       struct {
			Mode               string            `json:"mode"`
			PreProcessingRules []PreProcessRule  `json:"pre_processing_rules"`
			Segmentation       *SegmentationRule `json:"segmentation"`
		} `json:"process_rule"`
	}{
		Name:              req.Name,
		IndexingTechnique: req.IndexingTechnique,
		DocForm:           req.DocForm,
		DocType:           req.DocType,
		DocMetadata:       req.DocMetadata,
		ProcessRule: struct {
			Mode               string            `json:"mode"`
			PreProcessingRules []PreProcessRule  `json:"pre_processing_rules"`
			Segmentation       *SegmentationRule `json:"segmentation"`
		}{
			Mode:               req.ProcessRule.Mode,
			PreProcessingRules: req.ProcessRule.PreProcessingRules,
			Segmentation:       req.ProcessRule.Segmentation,
		},
	}

	// 序列化请求
	reqJSON, err := json.Marshal(apiReq)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %v", err)
	}

	// 添加 data 字段，设置 Content-Type 为 text/plain
	formWriter, err := writer.CreateFormField("data")
	if err != nil {
		return nil, fmt.Errorf("创建 data 字段失败: %v", err)
	}
	formWriter.Write(reqJSON)

	// 关闭 writer
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("关闭 writer 失败: %v", err)
	}

	// 创建请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	httpReq.Header.Set("Authorization", fmt.Sprintf(c.apiKey))
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// 解析响应
	var result struct {
		Data *Document `json:"data"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	return result.Data, nil
}

// GetDocumentIndexingStatus 获取文档嵌入状态
func (c *client) GetDocumentIndexingStatus(ctx context.Context, datasetID string, batch string) (*DocumentIndexingStatus, error) {
	url := fmt.Sprintf("%s/datasets/%s/documents/%s/indexing-status", c.baseURL, datasetID, batch)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	httpReq.Header.Set("Authorization", c.apiKey)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// 读取响应体
	var statuses *DocumentIndexingStatus
	if err := json.NewDecoder(resp.Body).Decode(&statuses); err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	return statuses, nil
}

// UpdateDocumentByText 通过文本更新文档
func (c *client) UpdateDocumentByText(ctx context.Context, datasetID string, documentID string, req *UpdateDocumentByTextRequest) (*Document, error) {
	url := fmt.Sprintf("%s/datasets/%s/documents/%s/update-by-text", c.baseURL, datasetID, documentID)

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

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data *Document `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	return result.Data, nil
}

// UpdateDocumentByFile 通过文件更新文档
func (c *client) UpdateDocumentByFile(ctx context.Context, datasetID string, documentID string, req *UpdateDocumentByFileRequest, file io.Reader) (*Document, error) {
	url := fmt.Sprintf("%s/datasets/%s/documents/%s/update-by-file", c.baseURL, datasetID, documentID)

	// 创建 multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件
	part, err := writer.CreateFormFile("file", "document.txt")
	if err != nil {
		return nil, fmt.Errorf("创建文件表单失败: %v", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("复制文件内容失败: %v", err)
	}

	// 构造请求数据
	reqData := struct {
		Name        string                 `json:"name,omitempty"`
		DocType     string                 `json:"doc_type,omitempty"`
		DocMetadata map[string]interface{} `json:"doc_metadata,omitempty"`
		ProcessRule *ProcessRule           `json:"process_rule,omitempty"`
	}{
		Name:        req.Name,
		DocType:     req.DocType,
		DocMetadata: req.DocMetadata,
		ProcessRule: req.ProcessRule,
	}

	// 序列化请求数据
	reqJSON, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %v", err)
	}

	// 添加 data 字段，设置 Content-Type 为 text/plain
	if err := writer.WriteField("data", string(reqJSON)); err != nil {
		return nil, fmt.Errorf("添加 data 字段失败: %v", err)
	}

	// 完成 multipart form
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("关闭 multipart writer 失败: %v", err)
	}

	// 创建请求
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	httpReq.Header.Set("Authorization", c.apiKey)
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("请求失败，状态码: %d，响应: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result struct {
		Data *Document `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	return result.Data, nil
}

// DeleteDocument 删除文档
func (c *client) DeleteDocument(ctx context.Context, datasetID string, documentID string) error {
	url := fmt.Sprintf("%s/datasets/%s/documents/%s", c.baseURL, datasetID, documentID)

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

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
