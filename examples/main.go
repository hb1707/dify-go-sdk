package main

import (
<<<<<<< HEAD
    "fmt"
    "log"
    "os"
    "sync"

    "github.com/hb170/dify-go-sdk/dify"
)

func main() {
    // 创建客户端实例
    client := dify.NewClient(os.Getenv("DIFY_API_KEY"))

    // 获取应用信息
    appInfo, err := client.GetAppInfo()
    if err != nil {
        log.Fatalf("获取应用信息失败: %v", err)
    }
    fmt.Printf("应用名称: %s\n描述: %s\n标签: %v\n\n", appInfo.Name, appInfo.Description, appInfo.Tags)

    // 获取应用参数
    params, err := client.GetAppParameters()
    if err != nil {
        log.Fatalf("获取应用参数失败: %v", err)
    }
    fmt.Printf("开场白: %s\n推荐问题: %v\n\n", params.OpeningStatement, params.SuggestedQuestions)

    // 发送文本生成请求（阻塞模式）
    req := &dify.CompletionRequest{
        Inputs: map[string]string{
            "query": "你好，请介绍一下自己",
        },
        User: "test-user",
    }

    resp, err := client.CreateCompletion(req)
    if err != nil {
        log.Fatalf("发送请求失败: %v", err)
    }
    fmt.Printf("回复: %s\n\n", resp.Answer)

    // 发送反馈
    feedback := &dify.FeedbackRequest{
        Rating: "like",
        User:   "test-user",
    }
    err = client.SendFeedback(resp.MessageID, feedback)
    if err != nil {
        log.Printf("发送反馈失败: %v", err)
    }

    // 上传文件示例
    file, err := client.UploadFile("test.jpg", "test-user")
    if err != nil {
        log.Printf("上传文件失败: %v", err)
    } else {
        fmt.Printf("文件上传成功: %s\n", file.ID)
    }

    // 流式请求示例
    streamReq := &dify.CompletionRequest{
        Inputs: map[string]string{
            "query": "给我讲个故事",
        },
        ResponseMode: "streaming",
        User:        "test-user",
    }

    // 实现流式处理接口
    handler := &StreamHandler{
        OnMessageFunc: func(resp *dify.MessageStreamResponse) error {
            fmt.Print(resp.Answer)
            return nil
        },
        OnMessageEndFunc: func(resp *dify.MessageEndStreamResponse) error {
            fmt.Println("\n--- 故事结束 ---")
            return nil
        },
        OnErrorFunc: func(err error) error {
            fmt.Printf("\n发生错误: %v\n", err)
            return nil
        },
    }

    err = client.CreateStreamingCompletion(streamReq, handler)
    if err != nil {
        log.Fatalf("流式请求失败: %v", err)
    }

    // 文字转语音示例
    ttsReq := &dify.TTSRequest{
        Text: "你好，这是一条测试语音消息",
        User: "test-user",
    }
    audioData, err := client.TextToSpeech(ttsReq)
    if err != nil {
        log.Printf("文字转语音失败: %v", err)
    } else {
        // 保存音频文件
        err = os.WriteFile("output.mp3", audioData, 0644)
        if err != nil {
            log.Printf("保存音频文件失败: %v", err)
        }
    }

    // 示例1: 阻塞模式
    fmt.Println("=== 阻塞模式示例 ===")
    blockingExample(client)

    // 示例2: 流式模式
    fmt.Println("\n=== 流式模式示例 ===")
    streamingExample(client)
=======
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/hb1707/dify-go-sdk/dify"
)

func main() {
	// 创建客户端实例
	client := dify.NewClient(os.Getenv("DIFY_API_KEY"))

	// 获取应用信息
	appInfo, err := client.GetAppInfo()
	if err != nil {
		log.Fatalf("获取应用信息失败: %v", err)
	}
	fmt.Printf("应用名称: %s\n描述: %s\n标签: %v\n\n", appInfo.Name, appInfo.Description, appInfo.Tags)

	// 获取应用参数
	params, err := client.GetAppParameters()
	if err != nil {
		log.Fatalf("获取应用参数失败: %v", err)
	}
	fmt.Printf("开场白: %s\n推荐问题: %v\n\n", params.OpeningStatement, params.SuggestedQuestions)

	// 发送文本生成请求（阻塞模式）
	req := &dify.CompletionRequest{
		Inputs: map[string]string{
			"query": "你好，请介绍一下自己",
		},
		User: "test-user",
	}

	resp, err := client.CreateCompletion(req)
	if err != nil {
		log.Fatalf("发送请求失败: %v", err)
	}
	fmt.Printf("回复: %s\n\n", resp.Answer)

	// 发送反馈
	feedback := &dify.FeedbackRequest{
		Rating: "like",
		User:   "test-user",
	}
	err = client.SendFeedback(resp.MessageID, feedback)
	if err != nil {
		log.Printf("发送反馈失败: %v", err)
	}

	// 上传文件示例
	file, err := client.UploadFile("test.jpg", "test-user")
	if err != nil {
		log.Printf("上传文件失败: %v", err)
	} else {
		fmt.Printf("文件上传成功: %s\n", file.ID)
	}

	// 流式请求示例
	streamReq := &dify.CompletionRequest{
		Inputs: map[string]string{
			"query": "给我讲个故事",
		},
		ResponseMode: "streaming",
		User:         "test-user",
	}

	// 实现流式处理接口
	handler := &StreamHandler{
		OnMessageFunc: func(resp *dify.MessageStreamResponse) error {
			fmt.Print(resp.Answer)
			return nil
		},
		OnMessageEndFunc: func(resp *dify.MessageEndStreamResponse) error {
			fmt.Println("\n--- 故事结束 ---")
			return nil
		},
		OnErrorFunc: func(err error) error {
			fmt.Printf("\n发生错误: %v\n", err)
			return nil
		},
	}

	err = client.CreateStreamingCompletion(streamReq, handler)
	if err != nil {
		log.Fatalf("流式请求失败: %v", err)
	}

	// 文字转语音示例
	ttsReq := &dify.TTSRequest{
		Text: "你好，这是一条测试语音消息",
		User: "test-user",
	}
	audioData, err := client.TextToSpeech(ttsReq)
	if err != nil {
		log.Printf("文字转语音失败: %v", err)
	} else {
		// 保存音频文件
		err = os.WriteFile("output.mp3", audioData, 0644)
		if err != nil {
			log.Printf("保存音频文件失败: %v", err)
		}
	}

	// 示例1: 阻塞模式
	fmt.Println("=== 阻塞模式示例 ===")
	blockingExample(client)

	// 示例2: 流式模式
	fmt.Println("\n=== 流式模式示例 ===")
	streamingExample(client)
>>>>>>> aa2eee1 (init)
}

// StreamHandler 实现了 dify.StreamHandler 接口
type StreamHandler struct {
<<<<<<< HEAD
    OnMessageFunc    func(*dify.MessageStreamResponse) error
    OnMessageEndFunc func(*dify.MessageEndStreamResponse) error
    OnTTSFunc       func(*dify.TTSStreamResponse) error
    OnErrorFunc     func(error) error
}

func (h *StreamHandler) OnMessage(resp *dify.MessageStreamResponse) error {
    if h.OnMessageFunc != nil {
        return h.OnMessageFunc(resp)
    }
    return nil
}

func (h *StreamHandler) OnMessageEnd(resp *dify.MessageEndStreamResponse) error {
    if h.OnMessageEndFunc != nil {
        return h.OnMessageEndFunc(resp)
    }
    return nil
}

func (h *StreamHandler) OnTTS(resp *dify.TTSStreamResponse) error {
    if h.OnTTSFunc != nil {
        return h.OnTTSFunc(resp)
    }
    return nil
}

func (h *StreamHandler) OnError(err error) error {
    if h.OnErrorFunc != nil {
        return h.OnErrorFunc(err)
    }
    return nil
}

func blockingExample(client *dify.Client) {
    req := &dify.CompletionRequest{
        Inputs: map[string]string{
            "query": "你好，请介绍一下你自己",
        },
        User: "user123",
    }

    resp, err := client.CreateCompletion(req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Message ID: %s\n", resp.MessageID)
    fmt.Printf("Answer: %s\n", resp.Answer)
    fmt.Printf("Created At: %d\n", resp.CreatedAt)
    fmt.Printf("Usage: %+v\n", resp.Metadata.Usage)
=======
	OnMessageFunc    func(*dify.MessageStreamResponse) error
	OnMessageEndFunc func(*dify.MessageEndStreamResponse) error
	OnTTSFunc        func(*dify.TTSStreamResponse) error
	OnErrorFunc      func(error) error
}

func (h *StreamHandler) OnMessage(resp *dify.MessageStreamResponse) error {
	if h.OnMessageFunc != nil {
		return h.OnMessageFunc(resp)
	}
	return nil
}

func (h *StreamHandler) OnMessageEnd(resp *dify.MessageEndStreamResponse) error {
	if h.OnMessageEndFunc != nil {
		return h.OnMessageEndFunc(resp)
	}
	return nil
}

func (h *StreamHandler) OnTTS(resp *dify.TTSStreamResponse) error {
	if h.OnTTSFunc != nil {
		return h.OnTTSFunc(resp)
	}
	return nil
}

func (h *StreamHandler) OnError(err error) error {
	if h.OnErrorFunc != nil {
		return h.OnErrorFunc(err)
	}
	return nil
}

func blockingExample(client *dify.Client) {
	req := &dify.CompletionRequest{
		Inputs: map[string]string{
			"query": "你好，请介绍一下你自己",
		},
		User: "user123",
	}

	resp, err := client.CreateCompletion(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message ID: %s\n", resp.MessageID)
	fmt.Printf("Answer: %s\n", resp.Answer)
	fmt.Printf("Created At: %d\n", resp.CreatedAt)
	fmt.Printf("Usage: %+v\n", resp.Metadata.Usage)
>>>>>>> aa2eee1 (init)
}

// ExampleHandler 实现 StreamHandler 接口
type ExampleHandler struct {
<<<<<<< HEAD
    mu sync.Mutex
}

func (h *ExampleHandler) OnMessage(resp *dify.MessageStreamResponse) error {
    h.mu.Lock()
    defer h.mu.Unlock()
    fmt.Print(resp.Answer)
    return nil
}

func (h *ExampleHandler) OnMessageEnd(resp *dify.MessageEndStreamResponse) error {
    h.mu.Lock()
    defer h.mu.Unlock()
    fmt.Printf("\n\n=== 消息结束 ===\n")
    fmt.Printf("Message ID: %s\n", resp.MessageID)
    fmt.Printf("Usage: %+v\n", resp.Metadata.Usage)
    return nil
}

func (h *ExampleHandler) OnTTS(resp *dify.TTSStreamResponse) error {
    h.mu.Lock()
    defer h.mu.Unlock()
    fmt.Printf("\nReceived TTS audio data, length: %d bytes\n", len(resp.Audio))
    return nil
}

func (h *ExampleHandler) OnError(err error) error {
    h.mu.Lock()
    defer h.mu.Unlock()
    fmt.Printf("\nError: %v\n", err)
    return nil
}

func streamingExample(client *dify.Client) {
    req := &dify.CompletionRequest{
        Inputs: map[string]string{
            "query": "请用中文写一首关于春天的诗",
        },
        User: "user123",
    }

    handler := &ExampleHandler{}
    if err := client.CreateStreamingCompletion(req, handler); err != nil {
        log.Fatal(err)
    }
=======
	mu sync.Mutex
}

func (h *ExampleHandler) OnMessage(resp *dify.MessageStreamResponse) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	fmt.Print(resp.Answer)
	return nil
}

func (h *ExampleHandler) OnMessageEnd(resp *dify.MessageEndStreamResponse) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	fmt.Printf("\n\n=== 消息结束 ===\n")
	fmt.Printf("Message ID: %s\n", resp.MessageID)
	fmt.Printf("Usage: %+v\n", resp.Metadata.Usage)
	return nil
}

func (h *ExampleHandler) OnTTS(resp *dify.TTSStreamResponse) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	fmt.Printf("\nReceived TTS audio data, length: %d bytes\n", len(resp.Audio))
	return nil
}

func (h *ExampleHandler) OnError(err error) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	fmt.Printf("\nError: %v\n", err)
	return nil
}

func streamingExample(client *dify.Client) {
	req := &dify.CompletionRequest{
		Inputs: map[string]string{
			"query": "请用中文写一首关于春天的诗",
		},
		User: "user123",
	}

	handler := &ExampleHandler{}
	if err := client.CreateStreamingCompletion(req, handler); err != nil {
		log.Fatal(err)
	}
>>>>>>> aa2eee1 (init)
}
