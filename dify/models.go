package dify

// ChatRequest 完成请求的结构体
type ChatRequest struct {
	Inputs       map[string]string `json:"inputs" validate:"required"`
	Query        string            `json:"query"`
	ResponseMode string            `json:"response_mode,omitempty" validate:"omitempty,oneof=blocking streaming"`
	User         string            `json:"user,omitempty" validate:"omitempty,min=1"`
	Files        []FileInput       `json:"files,omitempty" validate:"omitempty,dive"`
}

// CompletionRequest 完成请求的结构体
type CompletionRequest struct {
	Inputs       map[string]string `json:"inputs" validate:"required"`
	ResponseMode string            `json:"response_mode,omitempty" validate:"omitempty,oneof=blocking streaming"`
	User         string            `json:"user,omitempty" validate:"omitempty,min=1"`
	Files        []FileInput       `json:"files,omitempty" validate:"omitempty,dive"`
}

// FileInput 文件输入的结构体
type FileInput struct {
	Type           string `json:"type" validate:"required,eq=image"`                               // 目前仅支持 "image"
	TransferMethod string `json:"transfer_method" validate:"required,oneof=remote_url local_file"` // "remote_url" 或 "local_file"
	URL            string `json:"url,omitempty" validate:"required_if=TransferMethod remote_url,url"`
	UploadFileID   string `json:"upload_file_id,omitempty" validate:"required_if=TransferMethod local_file"`
}

// ChatResponse 完成响应的结构体（阻塞模式）
type ChatResponse struct {
	Event          string           `json:"event"`
	MessageID      string           `json:"message_id"`
	ConversationId string           `json:"conversation_id"`
	Mode           string           `json:"mode"` // 固定为 "chat"
	Answer         string           `json:"answer"`
	Metadata       ResponseMetadata `json:"metadata"`
	CreatedAt      int64            `json:"created_at"`
}

// CompletionResponse 完成响应的结构体（阻塞模式）
type CompletionResponse struct {
	MessageID string           `json:"message_id"`
	Mode      string           `json:"mode"` // 固定为 "chat"
	Answer    string           `json:"answer"`
	Metadata  ResponseMetadata `json:"metadata"`
	CreatedAt int64            `json:"created_at"`
}

// ResponseMetadata 响应元数据
type ResponseMetadata struct {
	Usage              Usage               `json:"usage"`
	RetrieverResources []RetrieverResource `json:"retriever_resources"`
}

// Usage 模型使用信息
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// RetrieverResource 引用和归属分段
type RetrieverResource struct {
	// 根据实际需要添加字段
}

// StreamResponse 流式响应的基础结构
type StreamResponse struct {
	Event     string `json:"event"`
	TaskID    string `json:"task_id"`
	MessageID string `json:"message_id"`
	CreatedAt int64  `json:"created_at"`
}

// MessageStreamResponse 消息事件响应
type MessageStreamResponse struct {
	StreamResponse
	Answer string `json:"answer"`
}

// MessageEndStreamResponse 消息结束事件响应
type MessageEndStreamResponse struct {
	StreamResponse
	Metadata ResponseMetadata `json:"metadata"`
}

// TTSStreamResponse TTS事件响应
type TTSStreamResponse struct {
	StreamResponse
	Audio string `json:"audio"` // base64编码的MP3音频数据
}

// FileUploadResponse 文件上传响应
type FileUploadResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	MimeType  string `json:"mime_type"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
}

// FeedbackRequest 消息反馈（点赞）
type FeedbackRequest struct {
	Rating  string `json:"rating"` // like, dislike, null
	User    string `json:"user"`
	Content string `json:"content,omitempty"`
}

// TTSRequest 文字转语音请求
type TTSRequest struct {
	MessageID string `json:"message_id,omitempty"`
	Text      string `json:"text,omitempty"`
	User      string `json:"user"`
}

// AppInfo 应用基本信息
type AppInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// AppParameters 应用参数
type AppParameters struct {
	OpeningStatement              string                   `json:"opening_statement"`
	SuggestedQuestions            []string                 `json:"suggested_questions"`
	SuggestedQuestionsAfterAnswer map[string]bool          `json:"suggested_questions_after_answer"`
	SpeechToText                  map[string]bool          `json:"speech_to_text"`
	RetrieverResource             map[string]bool          `json:"retriever_resource"`
	AnnotationReply               map[string]bool          `json:"annotation_reply"`
	UserInputForm                 []map[string]interface{} `json:"user_input_form"`
	FileUpload                    FileUploadConfig         `json:"file_upload"`
	SystemParameters              SystemParameters         `json:"system_parameters"`
}

// FileUploadConfig 文件上传配置
type FileUploadConfig struct {
	Image ImageUploadConfig `json:"image"`
}

// ImageUploadConfig 图片上传配置
type ImageUploadConfig struct {
	Enabled         bool     `json:"enabled"`
	NumberLimits    int      `json:"number_limits"`
	TransferMethods []string `json:"transfer_methods"`
}

// SystemParameters 系统参数
type SystemParameters struct {
	FileSizeLimit      int `json:"file_size_limit"`
	ImageFileSizeLimit int `json:"image_file_size_limit"`
	AudioFileSizeLimit int `json:"audio_file_size_limit"`
	VideoFileSizeLimit int `json:"video_file_size_limit"`
}
