package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/hb1707/dify-go-sdk/dify"
	"github.com/hb1707/dify-go-sdk/knowledge"
)

func main() {
	// 检查环境变量
	apiKey := "dataset-y0sgdQcTdcPMqkRC9y46QAh7"
	if apiKey == "" {
		log.Fatal("请设置 DIFY_API_KEY 环境变量")
	}

	// ================ 知识库功能示例 ================
	// 1. 首先创建知识库
	// testCreateKnowledge()

	// 2. 列出知识库以获取 ID
	// testListKnowledge()

	// 3. 使用获取到的知识库 ID 创建文档
	// testCreateDocumentByFile()

	// testCreateDocumentByText()

	// testGetDocumentIndexingStatus()

	testUpdateDocumentByText()

	// testDeleteDocument()
}

// StreamHandler 实现了 dify.StreamHandler 接口
type StreamHandler struct {
	OnMessageFunc    func(*dify.MessageStreamResponse) error
	OnMessageEndFunc func(*dify.MessageEndStreamResponse) error
	OnTTSFunc        func(*dify.TTSStreamResponse) error
	OnErrorFunc      func(error) error
}

// OnMessageWorkflow implements dify.StreamHandler.
func (h *StreamHandler) OnMessageWorkflow(response *dify.WorkflowStreamResponse) error {
	panic("unimplemented")
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
}

// ExampleHandler 实现 StreamHandler 接口
type ExampleHandler struct {
	mu sync.Mutex
}

func (h *ExampleHandler) OnMessageWorkflow(response *dify.WorkflowStreamResponse) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return nil
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
}

// testCreateKnowledge 测试创建知识库
func testCreateKnowledge() {
	client := knowledge.NewClient(os.Getenv("DIFY_API_KEY"))
	ctx := context.Background()

	// 创建知识库请求
	req := &knowledge.CreateKnowledgeRequest{
		Name:       "测试知识库",
		Permission: "all_team_members", // 默认值
	}

	// 可选参数示例
	// req.Description = "这是一个测试知识库"
	// req.IndexingTechnique = "high_quality" // 或 "economy"
	// req.Provider = "vendor" // 或 "external"
	// req.ExternalKnowledgeAPIID = "your-api-id"
	// req.ExternalKnowledgeID = "your-knowledge-id"

	k, err := client.CreateKnowledge(ctx, req)
	if err != nil {
		log.Fatalf("创建知识库失败: %v", err)
	}
	fmt.Printf("创建知识库成功: %+v\n", k)
}

// testListKnowledge 测试列出知识库
func testListKnowledge() {
	client := knowledge.NewClient(os.Getenv("DIFY_API_KEY"))
	ctx := context.Background()

	req := &knowledge.ListKnowledgeRequest{
		Page:      1,
		Limit:     10,
		Keyword:   "测试",
		SortBy:    "created_at",
		SortOrder: "desc",
	}

	resp, err := client.ListKnowledge(ctx, req)
	if err != nil {
		log.Fatalf("列出知识库失败: %v", err)
	}
	fmt.Printf("列出知识库成功: %+v\n", resp)
}

// testDeleteKnowledge 测试删除知识库
func testDeleteKnowledge() {
	client := knowledge.NewClient(os.Getenv("DIFY_API_KEY"))
	ctx := context.Background()

	// 替换为实际的知识库ID
	knowledgeID := "39efc3a6-4377-459c-920a-29cb8a79ec73"
	err := client.DeleteKnowledge(ctx, knowledgeID)
	if err != nil {
		log.Printf("删除知识库失败: %v", err)
	} else {
		fmt.Println("删除知识库成功")
	}
}

// testCreateDocumentByText 测试通过文本创建文档
func testCreateDocumentByText() {
	client := knowledge.NewClient("dataset-y0sgdQcTdcPMqkRC9y46QAh7")
	ctx := context.Background()

	// 替换为实际的知识库ID
	datasetID := "9dc6342e-0c81-4f4e-9ef8-07d7c94ae0fe"

	// 创建文档请求
	req := &knowledge.CreateDocumentByTextRequest{
		Name:    "测试文档3",
		Text:    "这是一个测试文档的内容。\n这是第二行内容。\n这是第三行内容。",
		DocType: "personal_document",
		DocMetadata: map[string]interface{}{
			"author":     "测试作者",
			"created_at": "2024-03-22",
			"custom_data": map[string]interface{}{
				"category": "测试",
				"tags":     []string{"测试", "文档"},
			},
		},
		IndexingTechnique: "high_quality",
		DocForm:           "text_model",
		ProcessRule: &knowledge.ProcessRule{
			Mode:  "automatic",
			Rules: map[string]interface{}{},
			PreProcessingRules: []knowledge.PreProcessRule{
				{
					ID:      "remove_extra_spaces",
					Enabled: true,
				},
				{
					ID:      "remove_urls_emails",
					Enabled: true,
				},
			},
			Segmentation: &knowledge.SegmentationRule{
				Separator:  "\n",
				MaxTokens:  1000,
				ParentMode: "full-doc",
				SubchunkSegmentation: &knowledge.SubchunkSegmentation{
					Separator:    "***",
					MaxTokens:    500,
					ChunkOverlap: 50,
				},
			},
		},
		RetrievalModel: &knowledge.RetrievalModel{
			SearchMethod:    "hybrid_search",
			RerankingEnable: true,
			RerankingModel: &knowledge.RerankModel{
				ProviderName: "cohere",
				ModelName:    "rerank-english-v2.0",
			},
			TopK:                  3,
			ScoreThresholdEnabled: true,
			ScoreThreshold:        0.7,
		},
		EmbeddingModel:    "text-embedding-ada-002",
		EmbeddingProvider: "openai",
	}

	doc, err := client.CreateDocumentByText(ctx, datasetID, req)
	if err != nil {
		log.Fatalf("创建文档失败: %v", err)
	}
	fmt.Printf("创建文档成功: %+v\n", doc)
}

// testCreateDocumentByFile 测试通过文件创建文档
func testCreateDocumentByFile() {
	// 创建客户端
	client := knowledge.NewClient("dataset-y0sgdQcTdcPMqkRC9y46QAh7")

	// 创建上下文
	ctx := context.Background()

	// 替换为实际的知识库ID
	datasetID := "aa3e0da7-b090-45ca-b03c-70263ca4c496"

	// 打开文件
	file, err := os.Open("examples/test.txt")
	if err != nil {
		log.Fatalf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 创建文档请求
	req := &knowledge.CreateDocumentByFileRequest{
		Name:              "测试文件文档",
		IndexingTechnique: "high_quality",
		DocForm:           "text_model",
		DocType:           "personal_document",
		DocMetadata: map[string]interface{}{
			"author":     "测试作者",
			"created_at": time.Now().Format(time.RFC3339),
		},
		ProcessRule: &knowledge.ProcessRule{
			Mode: "custom",
			PreProcessingRules: []knowledge.PreProcessRule{
				{
					ID:      "remove_extra_spaces",
					Enabled: true,
				},
				{
					ID:      "remove_urls_emails",
					Enabled: true,
				},
			},
			Segmentation: &knowledge.SegmentationRule{
				Separator:  "###",
				MaxTokens:  500,
				ParentMode: "full-doc",
			},
		},
	}

	// 打印完整的请求结构
	log.Printf("请求内容:\n")
	log.Printf("  Name: %s\n", req.Name)
	log.Printf("  IndexingTechnique: %s\n", req.IndexingTechnique)
	log.Printf("  DocForm: %s\n", req.DocForm)
	log.Printf("  DocType: %s\n", req.DocType)
	log.Printf("  DocMetadata: %+v\n", req.DocMetadata)
	log.Printf("  ProcessRule:\n")
	log.Printf("    Mode: %s\n", req.ProcessRule.Mode)
	log.Printf("    PreProcessingRules: %+v\n", req.ProcessRule.PreProcessingRules)
	log.Printf("    Segmentation: %+v\n", req.ProcessRule.Segmentation)

	// 调用创建文档方法，传入 io.Reader 接口
	doc, err := client.CreateDocumentByFile(ctx, datasetID, req, file)
	if err != nil {
		log.Printf("创建文档失败: %v", err)
	} else {
		log.Printf("创建文档成功: %+v\n", doc)
	}
}

// testGetDocumentIndexingStatus 测试获取文档嵌入状态
func testGetDocumentIndexingStatus() {
	client := knowledge.NewClient(os.Getenv("DIFY_API_KEY"))
	ctx := context.Background()

	// 替换为实际的知识库ID和批次号
	datasetID := "aa3e0da7-b090-45ca-b03c-70263ca4c496"
	batch := "your-batch-id" // 从创建文档的响应中获取

	status, err := client.GetDocumentIndexingStatus(ctx, datasetID, batch)
	if err != nil {
		log.Printf("获取文档嵌入状态失败: %v", err)
	} else {
		fmt.Printf("文档嵌入状态: %+v\n", status)
	}
}

// testUpdateDocumentByText 测试通过文本更新文档
func testUpdateDocumentByText() {
	client := knowledge.NewClient(os.Getenv("DIFY_API_KEY"))
	ctx := context.Background()

	datasetID := "aa3e0da7-b090-45ca-b03c-70263ca4c496"
	documentID := "a8a92583-1c49-43fa-90e6-95399f1030da" // 需要替换为实际的文档ID

	req := &knowledge.UpdateDocumentByTextRequest{
		Name:    "更新后的文档名称",
		Text:    "这是更新后的文档内容",
		DocType: "personal_document",
		DocMetadata: map[string]interface{}{
			"title":      "更新后的标题",
			"language":   "zh",
			"author":     "测试作者",
			"created_at": time.Now().Format(time.RFC3339),
		},
		ProcessRule: &knowledge.ProcessRule{
			Mode:  "automatic",
			Rules: map[string]interface{}{},
			PreProcessingRules: []knowledge.PreProcessRule{
				{
					ID:      "remove_extra_spaces",
					Enabled: true,
				},
				{
					ID:      "remove_urls_emails",
					Enabled: true,
				},
			},
			Segmentation: &knowledge.SegmentationRule{
				Separator:  "###",
				MaxTokens:  500,
				ParentMode: "full-doc",
			},
		},
	}

	doc, err := client.UpdateDocumentByText(ctx, datasetID, documentID, req)
	if err != nil {
		fmt.Printf("更新文档失败: %v\n", err)
		return
	}

	fmt.Printf("文档更新成功:\n")
	fmt.Printf("ID: %s\n", doc.ID)
	fmt.Printf("名称: %s\n", doc.Name)
	fmt.Printf("数据源类型: %s\n", doc.DataSourceType)
	fmt.Printf("创建时间: %d\n", doc.CreatedAt)
	fmt.Printf("Token数量: %d\n", doc.Tokens)
	fmt.Printf("索引状态: %s\n", doc.IndexingStatus)
	fmt.Printf("是否启用: %v\n", doc.Enabled)
	fmt.Printf("是否归档: %v\n", doc.Archived)
	fmt.Printf("显示状态: %s\n", doc.DisplayStatus)
	fmt.Printf("字数统计: %d\n", doc.WordCount)
	fmt.Printf("命中次数: %d\n", doc.HitCount)
	fmt.Printf("文档形式: %s\n", doc.DocForm)
}

// testDeleteDocument 测试删除文档
func testDeleteDocument() {
	client := knowledge.NewClient(os.Getenv("DIFY_API_KEY"))
	ctx := context.Background()

	datasetID := "aa3e0da7-b090-45ca-b03c-70263ca4c496"
	documentID := "9958c619-2b03-4826-9c48-73d4786f8d3e" // 需要替换为实际的文档ID

	err := client.DeleteDocument(ctx, datasetID, documentID)
	if err != nil {
		fmt.Printf("删除文档失败: %v\n", err)
		return
	}

	fmt.Println("文档删除成功")
}
