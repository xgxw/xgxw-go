package xgxw

import (
	"github.com/jinzhu/gorm"
)

// User is ...
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Password string `json:"-" gorm:"column:password"`
}

// UserService is ...
type UserService interface{}
