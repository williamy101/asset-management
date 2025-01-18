package util

type Request struct {
	Page  int `json:"page" binding:"required,min=1"`
	Limit int `json:"limit" binding:"required,min=1"`
}
