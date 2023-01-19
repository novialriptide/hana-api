package models

type Result struct {
	IsSuccessful bool   `json:"is_successful"`
	Message      string `json:"message"`
}
