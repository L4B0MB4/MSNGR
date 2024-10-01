package models

type MessageModel struct {
	Type        string `json:"Type" binding:"required"`
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description" binding:"required"`
}
