package models

type FilePath struct {
	FilePath string `json:"file_path"`
}

type ResponseWithStatusCode struct {
	Resp       interface{} `json:"response"`
	StatusCode int         `json:"status_code"`
}

type OfficeAndUsers struct {
	Office      Office       `json:"office"`
	OfficeUsers []OfficeUser `json:"office_user"`
}

type WorkerPreloadOptions struct {
	LoadCardTurnovers  bool
	LoadCardSales      bool
	LoadServiceQuality bool
	LoadMobileBank     bool
	LoadCardDetails    bool
	LoadUser           bool
}

type OptionUtil struct {
	QuestionID uint   `json:"question_id"`
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct"`
}

type NewUsersPassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ApplicationReportsRequest struct {
	ApplicationIDS []int32 `json:"application_ids"`
}
