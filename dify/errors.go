package dify

import (
	"fmt"
)

// DifyError represents an error returned by the Dify API
type DifyError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *DifyError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Status, e.Code, e.Message)
}

// Error codes
const (
	ErrCodeInvalidParam          = "invalid_param"
	ErrCodeAppUnavailable        = "app_unavailable"
	ErrCodeProviderNotInitialize = "provider_not_initialize"
	ErrCodeProviderQuotaExceeded = "provider_quota_exceeded"
	ErrCodeModelNotSupport       = "model_currently_not_support"
	ErrCodeCompletionRequest     = "completion_request_error"
	ErrCodeTooManyFiles          = "too_many_files"
	ErrCodeUnsupportedPreview    = "unsupported_preview"
	ErrCodeFileTooLarge          = "file_too_large"
	ErrCodeUnsupportedFileType   = "unsupported_file_type"
)

var (
	ErrInvalidParam          = NewDifyError(400, ErrCodeInvalidParam, "invalid parameter")
	ErrAppUnavailable        = NewDifyError(503, ErrCodeAppUnavailable, "app is unavailable")
	ErrProviderNotInitialize = NewDifyError(500, ErrCodeProviderNotInitialize, "provider not initialized")
	ErrTooManyFiles          = NewDifyError(400, ErrCodeTooManyFiles, "too many files")
	ErrUnsupportedPreview    = NewDifyError(400, ErrCodeUnsupportedPreview, "unsupported preview")
	ErrFileTooLarge          = NewDifyError(400, ErrCodeFileTooLarge, "file too large")
	ErrUnsupportedFileType   = NewDifyError(400, ErrCodeUnsupportedFileType, "unsupported file type")
	ErrProviderQuotaExceeded = NewDifyError(429, ErrCodeProviderQuotaExceeded, "provider quota exceeded")
)

// IsInvalidParam checks if the error is an invalid parameter error
func IsInvalidParam(err error) bool {
	if difyErr, ok := err.(*DifyError); ok {
		return difyErr.Code == ErrCodeInvalidParam
	}
	return false
}

// IsQuotaExceeded checks if the error is a quota exceeded error
func IsQuotaExceeded(err error) bool {
	if difyErr, ok := err.(*DifyError); ok {
		return difyErr.Code == ErrCodeProviderQuotaExceeded
	}
	return false
}

// NewDifyError creates a new DifyError
func NewDifyError(status int, code, message string) *DifyError {
	return &DifyError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
