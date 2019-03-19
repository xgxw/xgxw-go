package xgxw

import (
	"github.com/jinzhu/gorm"
)

// Resume is ...
type Resume struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
	FileID uint   `json:"file_id" gorm:"column:file_id"`

	User User `json:"user" gorm:"foreignkey:UserID;association_foreignkey:ID"`
	File File `json:"file" gorm:"foreignkey:FileID;association_foreignkey:ID"`
}

// ResumeService is ...
type ResumeService interface{}
