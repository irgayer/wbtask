package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique_index"`
	Password string `gorm:"not null"`
	Comments []Comment
}
