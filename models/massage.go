package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ChatID  uint   `json:"chat_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
	User    User   `json:"user"`
}
