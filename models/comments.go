package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text     string `json:"text"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name" gorm:"-"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
}
