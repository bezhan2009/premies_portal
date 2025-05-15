package errs

import "errors"

// Uniqueness Errors
var (
	ErrUsernameUniquenessFailed      = errors.New("ErrUsernameUniquenessFailed")
	ErrEmailUniquenessFailed         = errors.New("ErrEmailUniquenessFailed")
	ErrPhoneUniquenessFailed         = errors.New("ErrPhoneUniquenessFailed")
	ErrKnowledgeBaseUniquenessFailed = errors.New("ErrKnowledgeBaseUniquenessFailed")
	ErrKnowledgeAlreadyExists        = errors.New("ErrKnowledgeAlreadyExists")
)
