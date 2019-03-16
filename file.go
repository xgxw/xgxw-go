package xgxw

import (
	"github.com/jinzhu/gorm"
)

// File is ...
type File struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name"`
	URL  string `json:"url" gorm:"column:url"`
}

// FileService is ...
type FileService interface {
	GetFile(id string) (file *File, err error)
	GetFilesByName(name string) (file []*File, err error)
}
