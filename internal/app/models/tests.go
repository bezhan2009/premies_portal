package models

import (
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Description string `json:"description"`

	QuestionID uint     `json:"question_id" gorm:"uniqueIndex"` // связь 1:1
	Question   Question `gorm:"constraint:OnDelete:CASCADE;foreignKey:QuestionID"`

	Answers []Answer `gorm:"foreignKey:TestID"`
}

type QuestionType string

const (
	TextQuestion           QuestionType = "text"
	SingleChoiceQuestion   QuestionType = "single_choice"
	MultipleChoiceQuestion QuestionType = "multiple_choice"
)

type Question struct {
	gorm.Model

	Text string       `json:"text" gorm:"not null"`
	Type QuestionType `json:"type" gorm:"type:varchar(20);not null"`

	Options []Option `gorm:"foreignKey:QuestionID"`
}

type Option struct {
	gorm.Model

	QuestionID uint     `json:"question_id" gorm:"not null;index"`
	Question   Question `json:"-" gorm:"foreignKey:QuestionID"`
	Text       string   `json:"text" gorm:"not null"`
	IsCorrect  bool     `json:"-"`
}

type AnswerType string

const (
	TextAnswerType           AnswerType = "text"
	SingleChoiceAnswerType   AnswerType = "single_choice"
	MultipleChoiceAnswerType AnswerType = "multiple_choice"
)

type Answer struct {
	gorm.Model

	UserID     uint       `json:"user_id" gorm:"not null;index"`
	User       User       `json:"user" gorm:"foreignKey:UserID"`
	TestID     uint       `json:"test_id" gorm:"not null;index"`
	Test       Test       `json:"-" gorm:"foreignKey:TestID"`
	QuestionID uint       `json:"question_id" gorm:"not null;index"`
	Question   Question   `json:"question" gorm:"foreignKey:QuestionID"`
	Type       AnswerType `json:"type" gorm:"type:varchar(20);not null"`

	TextAnswer string `json:"text_answer"`

	SelectedOptions []SelectedOption `gorm:"foreignKey:AnswerID"`
}

type SelectedOption struct {
	gorm.Model

	AnswerID uint   `json:"answer_id" gorm:"not null"`
	Answer   Answer `json:"-" gorm:"foreignKey:AnswerID"`
	OptionID uint   `json:"option_id" gorm:"not null"`
	Option   Option `gorm:"foreignKey:OptionID"`
}
