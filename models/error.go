package models

type TrickError struct {
	Message string `json:"message"`
	Success bool `json:"success"`
	Data string `json:"data"`
}