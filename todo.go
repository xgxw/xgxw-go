package xgxw

import (
	"github.com/jinzhu/gorm"
)

// Todo is ...
type Todo struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
	FileID uint   `json:"file_id" gorm:"column:file_id"`

	User User `json:"user" gorm:"foreignkey:UserID;association_foreignkey:ID"`
	File File `json:"file" gorm:"foreignkey:FileID;association_foreignkey:ID"`
}

// TodoService is ...
type TodoService interface {
	Get(id, userID uint) (todo *Todo, err error)
	GetTodos(userID uint) (todos []*Todo, err error)
}
