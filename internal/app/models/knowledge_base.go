package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type KnowledgeBase struct {
	gorm.Model

	Title       string         `json:"title"`
	Description string         `json:"description"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`

	Knowledge []Knowledge `json:"knowledge" gorm:"foreignKey:KnowledgeBaseID"`
}

type Knowledge struct {
	gorm.Model

	Title       string         `json:"title"`
	Description string         `json:"description"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`

	KnowledgeBaseID uint          `json:"knowledge_base_id"`
	KnowledgeBase   KnowledgeBase `json:"-" gorm:"foreignKey:KnowledgeBaseID"`

	KnowledgeDocs []KnowledgeDocs `json:"knowledge_docs" gorm:"foreignKey:KnowledgeID"`
}

type KnowledgeDocs struct {
	gorm.Model

	Title    string `json:"title" gorm:"not null"`
	FilePath string `json:"file_path" gorm:"not null"`

	KnowledgeID uint      `json:"knowledge_id" gorm:"not null"`
	Knowledge   Knowledge `json:"knowledge" gorm:"foreignKey:KnowledgeID"`
}
