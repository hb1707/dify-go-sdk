package dify

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// WorkflowRequest 工作流请求结构体
type WorkflowRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`
	ResponseMode string                 `json:"response_mode,omitempty" validate:"omitempty,oneof=blocking streaming"`
	User         string                 `json:"user,omitempty" validate:"omitempty,min=1"`
	Files        []FileInput            `json:"files,omitempty"`
}

type WorkflowDataResp struct {
	Id          string         `json:"id"`
	WorkflowId  string         `json:"workflow_id"`
	Status      string         `json:"status"`
	Outputs     map[string]any `json:"outputs"`
	Error       any            `json:"error"`
	ElapsedTime float64        `json:"elapsed_time"`
	TotalTokens int            `json:"total_tokens"`
	TotalSteps  int            `json:"total_steps"`
	CreatedAt   int            `json:"created_at"`
	FinishedAt  int            `json:"finished_at"`
}

// WorkflowResponse 工作流响应结构体
type WorkflowResponse struct {
	TaskID        string           `json:"task_id"`
	WorkflowRunId string           `json:"workflow_run_id"`
	Data          WorkflowDataResp `json:"data,omitempty"`
}

// WorkflowRun 执行工作流的方法
func (c *Client) WorkflowRun(request WorkflowRequest) (*WorkflowResponse, error) {
	url := fmt.Sprintf("%s%s/run", c.BaseURL, EndpointWorkflows)
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("workflow execution failed with status %d: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var workflowResp WorkflowResponse
	err = json.Unmarshal(body, &workflowResp)
	if err != nil {
		return nil, err
	}

	return &workflowResp, nil
}

// WorkflowRunStreaming 执行流式工作流的方法
func (c *Client) WorkflowRunStreaming(request WorkflowRequest, handler StreamHandler) error {
	url := fmt.Sprintf("%s%s/run", c.BaseURL, EndpointWorkflows)
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "text/event-stream")
	httpReq.Header.Set("Connection", "keep-alive")
	httpReq.Header.Set("Cache-Control", "no-cache")
	httpReq.Header.Set("Transfer-Encoding", "chunked")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("workflow streaming execution failed with status %d: %s", resp.StatusCode, string(body))
	}

	reader := bufio.NewReader(resp.Body)
	for {
		// 读取一行数据直到遇到 \n\n
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read stream: %w", err)
		}

		// 跳过空行
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 检查是否是 SSE 数据行
		if !strings.HasPrefix(line, "data: ") {
			continue
		}

		// 解析 JSON 数据
		data := strings.TrimPrefix(line, "data: ")
		var baseResp StreamResponse
		if err := json.Unmarshal([]byte(data), &baseResp); err != nil {
			if err := handler.OnError(fmt.Errorf("failed to parse stream response: %w", err)); err != nil {
				return err
			}
			continue
		}

		// 根据事件类型处理不同的响应
		switch baseResp.Event {
		case "tts_message":
			var resp TTSStreamResponse
			if err := json.Unmarshal([]byte(data), &resp); err != nil {
				if err := handler.OnError(err); err != nil {
					return err
				}
				continue
			}
			if err := handler.OnTTS(&resp); err != nil {
				return err
			}
		case "tts_message_end":
			var resp TTSStreamResponse
			if err := json.Unmarshal([]byte(data), &resp); err != nil {
				if err := handler.OnError(err); err != nil {
					return err
				}
				continue
			}
			if err := handler.OnTTSEnd(&resp); err != nil {
				return err
			}
		default:
			var resp WorkflowStreamResponse
			if err := json.Unmarshal([]byte(data), &resp); err != nil {
				if err := handler.OnError(err); err != nil {
					return err
				}
				continue
			}
			if err := handler.OnMessageWorkflow(&resp); err != nil {
				return err
			}
		}
	}

	return nil
}
