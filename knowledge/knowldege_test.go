package knowledge

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"
)

var Key = "Your-API-Key"

// 测试创建知识库
func TestCreateKnowledge(t *testing.T) {
	client := NewClient(Key)
	knowledge, err := client.CreateKnowledge(context.Background(), &CreateKnowledgeRequest{
		Name:       "测试知识库2",
		Permission: "all_team_members",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("知识库: %v创建成功", knowledge.Name)
}

// 测试知识库列表
func TestListKnowledge(t *testing.T) {
	client := NewClient(Key)
	knowledge, err := client.ListKnowledge(context.Background(), &ListKnowledgeRequest{
		Page:      1,
		Limit:     10,
		Keyword:   "2025",
		SortBy:    "created_at",
		SortOrder: "desc",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("知识库: %v", knowledge)
}

// 测试删除知识库
func TestDeleteKnowledge(t *testing.T) {
	client := NewClient(Key)
	datasetID := "8f8dd638-459b-40af-ae20-0abd6f9b42d1"
	err := client.DeleteKnowledge(context.Background(), datasetID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("知识库: %v删除成功", datasetID)
}

// 测试通过文本创建文档
func TestCreateDocumentByText(t *testing.T) {
	client := NewClient(Key)
	datasetID := "9dc6342e-0c81-4f4e-9ef8-07d7c94ae0fe"
	doc, err := client.CreateDocumentByText(context.Background(), datasetID, &CreateDocumentByTextRequest{
		Name:    "测试文档4",
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
		ProcessRule: &ProcessRule{
			Mode:  "automatic",
			Rules: map[string]interface{}{},
			PreProcessingRules: []PreProcessRule{
				{
					ID:      "remove_extra_spaces",
					Enabled: true,
				},
				{
					ID:      "remove_urls_emails",
					Enabled: true,
				},
			},
			Segmentation: &SegmentationRule{
				Separator:  "\n",
				MaxTokens:  1000,
				ParentMode: "full-doc",
				SubchunkSegmentation: &SubchunkSegmentation{
					Separator:    "***",
					MaxTokens:    500,
					ChunkOverlap: 50,
				},
			},
		},
		RetrievalModel: &RetrievalModel{
			SearchMethod:    "hybrid_search",
			RerankingEnable: true,
			RerankingModel: &RerankModel{
				ProviderName: "cohere",
				ModelName:    "rerank-english-v2.0",
			},
			TopK:                  3,
			ScoreThresholdEnabled: true,
			ScoreThreshold:        0.7,
		},
		EmbeddingModel:    "text-embedding-ada-002",
		EmbeddingProvider: "openai",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("创建文档成功: %+v\n", doc)
}

// 测试通过文本更新文档
func TestUpdateDocumentByText(t *testing.T) {
	client := NewClient(Key)
	datasetID := "9dc6342e-0c81-4f4e-9ef8-07d7c94ae0fe"
	documentID := "c23180d5-a484-4df1-b6a7-d25664f3a44a"
	doc, err := client.UpdateDocumentByText(context.Background(), datasetID, documentID, &UpdateDocumentByTextRequest{
		Name:    "更新后的文档名称",
		Text:    "这是更新后的文档内容，哈哈哈哈哈哈哈哈哈哈",
		DocType: "personal_document",
		DocMetadata: map[string]interface{}{
			"title":      "更新后的标题",
			"language":   "zh",
			"author":     "测试作者",
			"created_at": time.Now().Format(time.RFC3339),
		},
		ProcessRule: &ProcessRule{
			Mode:  "automatic",
			Rules: map[string]interface{}{},
			PreProcessingRules: []PreProcessRule{
				{
					ID:      "remove_extra_spaces",
					Enabled: true,
				},
				{
					ID:      "remove_urls_emails",
					Enabled: true,
				},
			},
			Segmentation: &SegmentationRule{
				Separator:  "###",
				MaxTokens:  500,
				ParentMode: "full-doc",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("更新文档成功: %+v\n", doc)
}

// 测试删除文档
func TestDeleteDocument(t *testing.T) {
	client := NewClient(Key)
	datasetID := "9dc6342e-0c81-4f4e-9ef8-07d7c94ae0fe"
	documentID := "92f484ae-23c3-4eda-8573-907c4e8835b5"
	err := client.DeleteDocument(context.Background(), datasetID, documentID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("删除文档成功: %v", documentID)
}

// 测试获取文档嵌入状态（进度）
func TestGetDocumentIndexingStatus(t *testing.T) {
	client := NewClient(Key)
	datasetID := "9dc6342e-0c81-4f4e-9ef8-07d7c94ae0fe"
	batch := "20250323045136116479"
	status, err := client.GetDocumentIndexingStatus(context.Background(), datasetID, batch)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("文档索引状态:\n")
	fmt.Printf("文档ID: %s\n", status.ID)
	fmt.Printf("索引状态: %s\n", status.IndexingStatus)
	fmt.Printf("处理开始时间: %.2f\n", status.ProcessingStartedAt)
	fmt.Printf("解析完成时间: %.2f\n", status.ParsingCompletedAt)
	fmt.Printf("清洗完成时间: %.2f\n", status.CleaningCompletedAt)
	fmt.Printf("分段完成时间: %.2f\n", status.SplittingCompletedAt)
	if status.CompletedAt != nil {
		fmt.Printf("完成时间: %.2f\n", *status.CompletedAt)
	}
	if status.PausedAt != nil {
		fmt.Printf("暂停时间: %.2f\n", *status.PausedAt)
	}
	if status.Error != nil {
		fmt.Printf("错误: %s\n", *status.Error)
	}
	if status.StoppedAt != nil {
		fmt.Printf("停止时间: %.2f\n", *status.StoppedAt)
	}
	fmt.Printf("已完成段落数: %d\n", status.CompletedSegments)
	fmt.Printf("总段落数: %d\n", status.TotalSegments)
}

// 测试通过文件创建文档
func TestCreateDocumentByFile(t *testing.T) {
	client := NewClient(Key)
	datasetID := "9dc6342e-0c81-4f4e-9ef8-07d7c94ae0fe"
	// 创建文件内容
	fileContent := "这是一个测试文档内容"
	file := bytes.NewReader([]byte(fileContent))

	// 构造请求参数
	req := &CreateDocumentByFileRequest{
		Name:              "测试文档.txt",
		IndexingTechnique: "high_quality",
		DocForm:           "text_model",
		DocType:           "personal_document",
		DocMetadata: map[string]interface{}{
			"source": "test",
		},
		ProcessRule: &ProcessRule{
			Mode: "automatic",
			PreProcessingRules: []PreProcessRule{
				{
					ID:      "remove_extra_spaces",
					Enabled: true,
				},
			},
			Segmentation: &SegmentationRule{
				Separator:  "\n",
				MaxTokens:  1000,
				ParentMode: "paragraph",
			},
		},
	}

	doc, err := client.CreateDocumentByFile(context.Background(), datasetID, req, file)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("创建文档成功: %+v\n", doc)
}

// 测试通过文件更新文档
func TestUpdateDocumentByFile(t *testing.T) {
	client := NewClient(Key)
	datasetID := "9dc6342e-0c81-4f4e-9ef8-07d7c94ae0fe"
	documentID := "c23180d5-a484-4df1-b6a7-d25664f3a44a"

	// 创建文件内容
	fileContent := "这是更新后的测试文档内容"
	file := bytes.NewReader([]byte(fileContent))

	// 构造请求参数
	req := &UpdateDocumentByFileRequest{
		Name:    "更新后的测试文档.txt",
		DocType: "personal_document",
		DocMetadata: map[string]interface{}{
			"source":     "test",
			"language":   "zh",
			"author":     "测试作者",
			"created_at": time.Now().Format(time.RFC3339),
		},
		ProcessRule: &ProcessRule{
			Mode: "automatic",
			PreProcessingRules: []PreProcessRule{
				{
					ID:      "remove_extra_spaces",
					Enabled: true,
				},
				{
					ID:      "remove_urls_emails",
					Enabled: true,
				},
			},
			Segmentation: &SegmentationRule{
				Separator:  "\n",
				MaxTokens:  1000,
				ParentMode: "paragraph",
				SubchunkSegmentation: &SubchunkSegmentation{
					Separator:    "***",
					MaxTokens:    500,
					ChunkOverlap: 50,
				},
			},
		},
	}

	doc, err := client.UpdateDocumentByFile(context.Background(), datasetID, documentID, req, file)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("更新文档成功: %+v\n", doc)
}

// 测试检索知识库
func TestRetrieve(t *testing.T) {
	client := NewClient(Key)
	datasetID := "9dc6342e-0c81-4f4e-9ef8-07d7c94ae0fe"

	// 构造请求参数
	req := &RetrieveRequest{
		Query: "测试文档",
		RetrievalModel: &RetrievalModel{
			SearchMethod:          "hybrid_search",
			RerankingEnable:       false,
			TopK:                  3,
			ScoreThresholdEnabled: true,
			ScoreThreshold:        0.7,
		},
	}

	// 发送检索请求
	resp, err := client.Retrieve(context.Background(), datasetID, req)
	if err != nil {
		t.Fatal(err)
	}

	// 打印检索结果
	fmt.Printf("查询内容: %s\n", resp.Query.Content)
	fmt.Printf("\n检索结果:\n")
	for i, record := range resp.Records {
		fmt.Printf("\n记录 %d:\n", i+1)
		segment := record.Segment
		fmt.Printf("片段ID: %s\n", segment.ID)
		fmt.Printf("位置: %d\n", segment.Position)
		fmt.Printf("文档ID: %s\n", segment.DocumentID)
		fmt.Printf("内容: %s\n", segment.Content)
		if segment.Answer != nil {
			fmt.Printf("答案: %s\n", *segment.Answer)
		}
		fmt.Printf("字数: %d\n", segment.WordCount)
		fmt.Printf("Token数: %d\n", segment.Tokens)
		fmt.Printf("关键词: %v\n", segment.Keywords)
		fmt.Printf("索引节点ID: %s\n", segment.IndexNodeID)
		fmt.Printf("索引节点哈希: %s\n", segment.IndexNodeHash)
		fmt.Printf("命中次数: %d\n", segment.HitCount)
		fmt.Printf("是否启用: %v\n", segment.Enabled)
		if segment.DisabledAt != nil {
			fmt.Printf("禁用时间: %d\n", *segment.DisabledAt)
		}
		if segment.DisabledBy != nil {
			fmt.Printf("禁用者: %s\n", *segment.DisabledBy)
		}
		fmt.Printf("状态: %s\n", segment.Status)
		fmt.Printf("创建者: %s\n", segment.CreatedBy)
		fmt.Printf("创建时间: %d\n", segment.CreatedAt)
		fmt.Printf("索引时间: %d\n", segment.IndexingAt)
		fmt.Printf("完成时间: %d\n", segment.CompletedAt)
		if segment.Error != nil {
			fmt.Printf("错误: %s\n", *segment.Error)
		}
		if segment.StoppedAt != nil {
			fmt.Printf("停止时间: %d\n", *segment.StoppedAt)
		}
		if segment.Document != nil {
			fmt.Printf("文档信息:\n")
			fmt.Printf("  ID: %s\n", segment.Document.ID)
			fmt.Printf("  数据源类型: %s\n", segment.Document.DataSourceType)
			fmt.Printf("  名称: %s\n", segment.Document.Name)
			if segment.Document.DocType != nil {
				fmt.Printf("  类型: %s\n", *segment.Document.DocType)
			}
		}
		fmt.Printf("相关度分数: %.10f\n", record.Score)
		if record.TsnePosition != nil {
			fmt.Printf("TSNE位置: %.10f\n", *record.TsnePosition)
		}
	}
}
