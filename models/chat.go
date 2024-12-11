package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"many2many:chat_users;" json:"users"`
}
