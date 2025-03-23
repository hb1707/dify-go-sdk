package knowledge

import "fmt"

// 错误类型
var (
	ErrInvalidKnowledgeID = fmt.Errorf("invalid knowledge ID")
	ErrInvalidDocumentID  = fmt.Errorf("invalid document ID")
	ErrInvalidParagraphID = fmt.Errorf("invalid paragraph ID")
	ErrInvalidRequest     = fmt.Errorf("invalid request")
	ErrInvalidResponse    = fmt.Errorf("invalid response")
	ErrNotFound           = fmt.Errorf("resource not found")
	ErrUnauthorized       = fmt.Errorf("unauthorized")
	ErrForbidden          = fmt.Errorf("forbidden")
	ErrTooManyRequests    = fmt.Errorf("too many requests")
	ErrInternalServer     = fmt.Errorf("internal server error")
)

// APIError 表示 API 错误
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("API error: %s - %s (%s)", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("API error: %s - %s", e.Code, e.Message)
}
