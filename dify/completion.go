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

// CreateCompletion 发送阻塞模式的完成请求
func (c *Client) CreateCompletion(req *CompletionRequest) (*CompletionResponse, error) {

	endpoint := fmt.Sprintf("%s%s", c.BaseURL, EndpointCompletion)

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

	var result CompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// StreamHandler 流式响应处理函数类型
type StreamHandler interface {
	OnMessage(response *MessageStreamResponse) error
	OnMessageWorkflow(response *WorkflowStreamResponse) error
	OnMessageEnd(response *MessageEndStreamResponse) error
	OnTTS(response *TTSStreamResponse) error
	OnTTSEnd(response *TTSStreamResponse) error
	OnError(err error) error
}

// CreateStreamingCompletion 发送流式模式的完成请求
func (c *Client) CreateStreamingCompletion(req *CompletionRequest, handler StreamHandler) error {

	endpoint := fmt.Sprintf("%s%s", c.BaseURL, EndpointCompletion)

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
			if err.Error() == "context canceled" || err == io.EOF {
				var resp MessageEndStreamResponse
				resp.StreamResponse.Event = "message_end"
				if err := handler.OnMessageEnd(&resp); err != nil {
					return err
				}
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
		}

	}

	return nil
}
