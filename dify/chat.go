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

// CreateChat 发送阻塞模式的完成请求
func (c *Client) CreateChat(req *ChatRequest) (*ChatResponse, error) {

	endpoint := fmt.Sprintf("%s%s", c.BaseURL, EndpointChat)

	// 设置响应模式为阻塞模式
	req.ResponseMode = ResponseModeBlocking

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// CreateStreamingChat 发送流式模式的完成请求
func (c *Client) CreateStreamingChat(req *ChatRequest, handler StreamHandler) error {

	endpoint := fmt.Sprintf("%s%s", c.BaseURL, EndpointChat)

	// 设置响应模式为流式模式
	req.ResponseMode = ResponseModeStreaming

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "text/event-stream")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
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
		case "message":
			var resp MessageStreamResponse
			if err := json.Unmarshal([]byte(data), &resp); err != nil {
				if err := handler.OnError(err); err != nil {
					return err
				}
				continue
			}
			if err := handler.OnMessage(&resp); err != nil {
				return err
			}

		case "message_end":
			var resp MessageEndStreamResponse
			if err := json.Unmarshal([]byte(data), &resp); err != nil {
				if err := handler.OnError(err); err != nil {
					return err
				}
				continue
			}
			if err := handler.OnMessageEnd(&resp); err != nil {
				return err
			}

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
