package errs

import "errors"

// General Errors
var (
	ErrRecordNotFound        = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong    = errors.New("ErrSomethingWentWrong")
	ErrUserNotFound          = errors.New("ErrUserNotFound")
	ErrDeleteFailed          = errors.New("ErrDeleteFailed")
	ErrKnowledgeBaseNotFound = errors.New("ErrKnowledgeBaseNotFound")
	ErrKnowledgeDocNotFound  = errors.New("ErrKnowledgeDocNotFound")
	ErrInvalidFilePath       = errors.New("ErrInvalidFilePath")
	ErrFileNotFound          = errors.New("ErrFileNotFound")
	ErrTempReportNotFound    = errors.New("ErrTempReportNotFound")
)
