package xgxw

import (
	"context"
	"time"
)

type (
	File struct {
		ID   int64  `gorm:"column:id" json:"id"`
		Name string `gorm:"column:name" json:"name"`
		URL  string `gorm:"column:url" json:"url"`

		CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
		UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
		DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	}
)

func (File) TableName() string {
	return "files"
}

type FileService interface {
	Save(ctx context.Context, name string, data []byte) (f *File, err error)
	Get(ctx context.Context, id int64) (f *File, err error)
	//Delete(ctx context.Context, id int64) (err error)
}
