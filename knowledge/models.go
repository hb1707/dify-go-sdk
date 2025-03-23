package knowledge

// Knowledge 知识库
type Knowledge struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Tags                 []string `json:"tags"`
	IndexingTechnique    string   `json:"indexing_technique"`
	Permission           string   `json:"permission"`
	Provider             string   `json:"provider"`
	CreatedAt            int64    `json:"created_at"`
	UpdatedAt            int64    `json:"updated_at"`
	DocumentCount        int      `json:"document_count"`
	EmbeddingModel       string   `json:"embedding_model"`
	EmbeddingModelConfig struct {
		Provider string `json:"provider"`
		Model    string `json:"model"`
	} `json:"embedding_model_config"`
	RetrievalModel       string `json:"retrieval_model"`
	RetrievalModelConfig struct {
		Provider string `json:"provider"`
		Model    string `json:"model"`
	} `json:"retrieval_model_config"`
	Status string `json:"status"`
}

// CreateKnowledgeRequest 创建知识库请求
type CreateKnowledgeRequest struct {
	Name                   string `json:"name"`                                // 知识库名称（必填）
	Description            string `json:"description,omitempty"`               // 知识库描述（选填）
	IndexingTechnique      string `json:"indexing_technique,omitempty"`        // 索引模式（选填）：high_quality/economy
	Permission             string `json:"permission,omitempty"`                // 权限（选填，默认 only_me）：only_me/all_team_members/partial_members
	Provider               string `json:"provider,omitempty"`                  // Provider（选填，默认 vendor）：vendor/external
	ExternalKnowledgeAPIID string `json:"external_knowledge_api_id,omitempty"` // 外部知识库 API_ID（选填）
	ExternalKnowledgeID    string `json:"external_knowledge_id,omitempty"`     // 外部知识库 ID（选填）
}

// ListKnowledgeRequest 列出知识库请求
type ListKnowledgeRequest struct {
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Keyword   string `json:"keyword,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
	SortOrder string `json:"sort_order,omitempty"`
}

// ListKnowledgeResponse 列出知识库响应
type ListKnowledgeResponse struct {
	Data    []Knowledge `json:"data"`
	Total   int         `json:"total"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	HasMore bool        `json:"has_more"`
}

// Document 文档
type Document struct {
	ID                   string                 `json:"id"`                      // 文档ID
	Position             int                    `json:"position"`                // 位置
	DataSourceType       string                 `json:"data_source_type"`        // 数据源类型
	DataSourceInfo       map[string]interface{} `json:"data_source_info"`        // 数据源信息
	DatasetProcessRuleID string                 `json:"dataset_process_rule_id"` // 数据集处理规则ID
	Name                 string                 `json:"name"`                    // 文档名称
	DocType              *string                `json:"doc_type"`                // 文档类型
	CreatedFrom          string                 `json:"created_from"`            // 创建来源
	CreatedBy            string                 `json:"created_by"`              // 创建者
	CreatedAt            int64                  `json:"created_at"`              // 创建时间
	Tokens               int                    `json:"tokens"`                  // token数量
	IndexingStatus       string                 `json:"indexing_status"`         // 索引状态
	Error                *string                `json:"error"`                   // 错误信息
	Enabled              bool                   `json:"enabled"`                 // 是否启用
	DisabledAt           *int64                 `json:"disabled_at"`             // 禁用时间
	DisabledBy           *string                `json:"disabled_by"`             // 禁用者
	Archived             bool                   `json:"archived"`                // 是否归档
	DisplayStatus        string                 `json:"display_status"`          // 显示状态
	WordCount            int                    `json:"word_count"`              // 字数统计
	HitCount             int                    `json:"hit_count"`               // 命中次数
	DocForm              string                 `json:"doc_form"`                // 文档形式
}

// Metadata 元数据
type Metadata struct {
	Source     string                 `json:"source,omitempty"`
	Author     string                 `json:"author,omitempty"`
	CreatedAt  string                 `json:"created_at,omitempty"`
	UpdatedAt  string                 `json:"updated_at,omitempty"`
	CustomData map[string]interface{} `json:"custom_data,omitempty"`
}

// CreateDocumentByTextRequest 通过文本创建文档请求
type CreateDocumentByTextRequest struct {
	Name              string                 `json:"name"`                               // 文档名称
	Text              string                 `json:"text"`                               // 文档内容
	DocType           string                 `json:"doc_type,omitempty"`                 // 文档类型（选填）
	DocMetadata       map[string]interface{} `json:"doc_metadata,omitempty"`             // 文档元数据（如提供文档类型则必填）
	IndexingTechnique string                 `json:"indexing_technique"`                 // 索引方式：high_quality/economy
	DocForm           string                 `json:"doc_form"`                           // 索引内容形式：text_model/hierarchical_model/qa_model
	DocLanguage       string                 `json:"doc_language,omitempty"`             // 文档语言（Q&A模式下必填）
	ProcessRule       *ProcessRule           `json:"process_rule,omitempty"`             // 处理规则
	RetrievalModel    *RetrievalModel        `json:"retrieval_model,omitempty"`          // 检索模式
	EmbeddingModel    string                 `json:"embedding_model,omitempty"`          // Embedding模型名称
	EmbeddingProvider string                 `json:"embedding_model_provider,omitempty"` // Embedding模型供应商
}

// CreateDocumentByFileRequest 通过文件创建文档请求
type CreateDocumentByFileRequest struct {
	OriginalDocumentID string                 `json:"original_document_id,omitempty"`     // 源文档ID（选填）
	Name               string                 `json:"name"`                               // 文档名称
	IndexingTechnique  string                 `json:"indexing_technique"`                 // 索引方式：high_quality/economy
	DocForm            string                 `json:"doc_form"`                           // 索引内容形式：text_model/hierarchical_model/qa_model
	DocType            string                 `json:"doc_type,omitempty"`                 // 文档类型（选填）：book/web_page/paper/social_media_post/wikipedia_entry/personal_document/business_document/im_chat_log/synced_from_notion/synced_from_github/others
	DocMetadata        map[string]interface{} `json:"doc_metadata,omitempty"`             // 文档元数据（如提供文档类型则必填）
	DocLanguage        string                 `json:"doc_language,omitempty"`             // 文档语言（Q&A模式下必填）
	ProcessRule        *ProcessRule           `json:"process_rule"`                       // 处理规则（未传入 original_document_id 时必填）
	RetrievalModel     *RetrievalModel        `json:"retrieval_model,omitempty"`          // 检索模式（首次上传时可选）
	EmbeddingModel     string                 `json:"embedding_model,omitempty"`          // Embedding模型名称（首次上传时可选）
	EmbeddingProvider  string                 `json:"embedding_model_provider,omitempty"` // Embedding模型供应商（首次上传时可选）
}

// ProcessRule 处理规则
type ProcessRule struct {
	Mode               string                 `json:"mode"`                 // 清洗、分段模式：automatic/custom
	Rules              map[string]interface{} `json:"rules"`                // 自定义规则（自动模式下为空）
	PreProcessingRules []PreProcessRule       `json:"pre_processing_rules"` // 预处理规则
	Segmentation       *SegmentationRule      `json:"segmentation"`         // 分段规则
}

// PreProcessRule 预处理规则
type PreProcessRule struct {
	ID      string `json:"id"`      // 预处理规则的唯一标识符：remove_extra_spaces/remove_urls_emails
	Enabled bool   `json:"enabled"` // 是否启用该规则
}

// SegmentationRule 分段规则
type SegmentationRule struct {
	Separator            string                `json:"separator"`                       // 分段标识符，默认为 \n
	MaxTokens            int                   `json:"max_tokens"`                      // 最大长度（token），默认为 1000
	ParentMode           string                `json:"parent_mode"`                     // 父分段的召回模式：full-doc/paragraph
	SubchunkSegmentation *SubchunkSegmentation `json:"subchunk_segmentation,omitempty"` // 子分段规则
}

// SubchunkSegmentation 子分段规则
type SubchunkSegmentation struct {
	Separator    string `json:"separator"`               // 分段标识符，默认为 ***
	MaxTokens    int    `json:"max_tokens"`              // 最大长度（token），需要小于父级长度
	ChunkOverlap int    `json:"chunk_overlap,omitempty"` // 分段重叠（选填）
}

// RetrievalModel 检索模式
type RetrievalModel struct {
	SearchMethod          string       `json:"search_method"`           // 检索方法：hybrid_search/semantic_search/full_text_search
	RerankingEnable       bool         `json:"reranking_enable"`        // 是否开启rerank
	RerankingModel        *RerankModel `json:"reranking_model"`         // Rerank模型配置
	TopK                  int          `json:"top_k"`                   // 召回条数
	ScoreThresholdEnabled bool         `json:"score_threshold_enabled"` // 是否开启召回分数限制
	ScoreThreshold        float64      `json:"score_threshold"`         // 召回分数限制
}

// RerankModel Rerank模型配置
type RerankModel struct {
	ProviderName string `json:"reranking_provider_name"` // Rerank模型的提供商
	ModelName    string `json:"reranking_model_name"`    // Rerank模型的名称
}

// DocumentIndexingStatus 文档嵌入状态
type DocumentIndexingStatus struct {
	ID                   string   `json:"id"`                     // 文档ID
	IndexingStatus       string   `json:"indexing_status"`        // 索引状态
	ProcessingStartedAt  float64  `json:"processing_started_at"`  // 处理开始时间
	ParsingCompletedAt   float64  `json:"parsing_completed_at"`   // 解析完成时间
	CleaningCompletedAt  float64  `json:"cleaning_completed_at"`  // 清洗完成时间
	SplittingCompletedAt float64  `json:"splitting_completed_at"` // 分段完成时间
	CompletedAt          *float64 `json:"completed_at"`           // 完成时间
	PausedAt             *float64 `json:"paused_at"`              // 暂停时间
	Error                *string  `json:"error"`                  // 错误信息
	StoppedAt            *float64 `json:"stopped_at"`             // 停止时间
	CompletedSegments    int      `json:"completed_segments"`     // 已完成段落数
	TotalSegments        int      `json:"total_segments"`         // 总段落数
}

// UpdateDocumentByTextRequest 通过文本更新文档请求
type UpdateDocumentByTextRequest struct {
	Name        string                 `json:"name,omitempty"`         // 文档名称（选填）
	Text        string                 `json:"text,omitempty"`         // 文档内容（选填）
	DocType     string                 `json:"doc_type,omitempty"`     // 文档类型（选填）
	DocMetadata map[string]interface{} `json:"doc_metadata,omitempty"` // 文档元数据（如提供文档类型则必填）
	ProcessRule *ProcessRule           `json:"process_rule,omitempty"` // 处理规则（选填）
}

// UpdateDocumentByFileRequest 通过文件更新文档请求
type UpdateDocumentByFileRequest struct {
	Name        string                 `json:"name,omitempty"`         // 文档名称（选填）
	DocType     string                 `json:"doc_type,omitempty"`     // 文档类型（选填）
	DocMetadata map[string]interface{} `json:"doc_metadata,omitempty"` // 文档元数据（如提供文档类型则必填）
	ProcessRule *ProcessRule           `json:"process_rule,omitempty"` // 处理规则（选填）
}

// RetrieveRequest 检索知识库请求
type RetrieveRequest struct {
	Query          string          `json:"query"`                     // 检索关键词
	RetrievalModel *RetrievalModel `json:"retrieval_model,omitempty"` // 检索参数（选填）
}

// RetrieveResponse 检索知识库响应
type RetrieveResponse struct {
	Query struct {
		Content string `json:"content"` // 查询内容
	} `json:"query"`
	Records []Record `json:"records"` // 检索结果记录列表
}

// Record 检索结果记录
type Record struct {
	Segment      Segment  `json:"segment"`       // 文档片段
	Score        float64  `json:"score"`         // 相关度分数
	TsnePosition *float64 `json:"tsne_position"` // TSNE 位置（可选）
}

// Segment 检索结果片段
type Segment struct {
	ID            string    `json:"id"`              // 片段ID
	Position      int       `json:"position"`        // 位置
	DocumentID    string    `json:"document_id"`     // 文档ID
	Content       string    `json:"content"`         // 内容
	Answer        *string   `json:"answer"`          // 答案（可选）
	WordCount     int       `json:"word_count"`      // 字数统计
	Tokens        int       `json:"tokens"`          // token数量
	Keywords      []string  `json:"keywords"`        // 关键词列表
	IndexNodeID   string    `json:"index_node_id"`   // 索引节点ID
	IndexNodeHash string    `json:"index_node_hash"` // 索引节点哈希值
	HitCount      int       `json:"hit_count"`       // 命中次数
	Enabled       bool      `json:"enabled"`         // 是否启用
	DisabledAt    *int64    `json:"disabled_at"`     // 禁用时间
	DisabledBy    *string   `json:"disabled_by"`     // 禁用者
	Status        string    `json:"status"`          // 状态
	CreatedBy     string    `json:"created_by"`      // 创建者
	CreatedAt     int64     `json:"created_at"`      // 创建时间
	IndexingAt    int64     `json:"indexing_at"`     // 索引时间
	CompletedAt   int64     `json:"completed_at"`    // 完成时间
	Error         *string   `json:"error"`           // 错误信息
	StoppedAt     *int64    `json:"stopped_at"`      // 停止时间
	Document      *Document `json:"document"`        // 文档信息
}
