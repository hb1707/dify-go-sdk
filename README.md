# Dify Go SDK

这是一个用于访问 Dify API 的 Go 语言 SDK。

## 安装

```bash
<<<<<<< HEAD
go get github.com/hb170/dify-go-sdk
=======
go get github.com/hb1707/dify-go-sdk
>>>>>>> aa2eee1 (init)
```

## 快速开始

### 初始化客户端

```go
<<<<<<< HEAD
import "github.com/hb170/dify-go-sdk/dify"
=======
import "github.com/hb1707/dify-go-sdk/dify"
>>>>>>> aa2eee1 (init)

client := dify.NewClient("your-api-key")
```

### 阻塞模式示例

```go
req := &dify.CompletionRequest{
    Inputs: map[string]string{
        "query": "你好，请介绍一下你自己",
    },
    ResponseMode: "blocking",
    User:        "user123",
}

resp, err := client.CreateCompletion(req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Response: %+v\n", resp)
```

### 流式模式示例

```go
req := &dify.CompletionRequest{
    Inputs: map[string]string{
        "query": "请用中文写一首关于春天的诗",
    },
    User: "user123",
}

resp, err := client.CreateStreamingCompletion(req)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

reader := bufio.NewReader(resp.Body)
for {
    line, err := reader.ReadString('\n')
    if err != nil {
        break
    }
    fmt.Print(line)
}
```

### 自定义配置

```go
// 自定义基础 URL
client := dify.NewClient("your-api-key", dify.WithBaseURL("https://your-custom-url.com"))

// 自定义 HTTP 客户端
httpClient := &http.Client{
    Timeout: time.Second * 60,
}
client := dify.NewClient("your-api-key", dify.WithHTTPClient(httpClient))
```

## 特性

- 支持阻塞和流式响应模式
- 支持文件上传
- 可自定义 HTTP 客户端
- 可自定义基础 URL
- 完整的错误处理
- 类型安全的请求/响应结构

## 错误处理

SDK 定义了以下错误类型：

- ErrInvalidParam: 参数无效
- ErrAppUnavailable: 应用不可用
- ErrProviderNotInitialize: 提供者未初始化
- ErrTooManyFiles: 文件数量过多
- ErrUnsupportedPreview: 不支持预览
- ErrFileTooLarge: 文件太大
- ErrUnsupportedFileType: 不支持的文件类型

## 许可证

MIT License
