package dify

// DefaultBaseURL 默认的API基础URL
var DefaultBaseURL = "https://api.dify.ai/v1"

const (

	// API 端点
	EndpointCompletion = "/completion-messages" // 文本生成
	EndpointMessages   = "/messages"            // 消息相关
	EndpointFiles      = "/files"               // 文件上传
	EndpointAudio      = "/audio"               // 音频相关
	EndpointInfo       = "/info"                // 应用信息
	EndpointParameters = "/parameters"          // 应用参数
	EndpointFeedbacks  = "/feedbacks"           // 消息反馈
)

// API响应模式
const (
	ResponseModeBlocking  = "blocking"  // 阻塞模式
	ResponseModeStreaming = "streaming" // 流式模式
)

// 示例常量
var (
	// UserExample 用户ID示例
	UserExample = "user123"

	// QueryExample 查询示例
	QueryExample = "你好，请介绍一下你自己"
)
