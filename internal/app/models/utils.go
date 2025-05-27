package models

type FilePath struct {
	FilePath string `json:"file_path"`
}

type ResponseWithStatusCode struct {
	Resp       interface{} `json:"response"`
	StatusCode int         `json:"status_code"`
}
