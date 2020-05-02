package xgxw

import (
	"context"
	"time"
)

type Author struct {
	ID   int64  `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

type AuthorService interface {
	Create(ctx context.Context, name string) (author *Author, err error)
	Get(ctx context.Context, id int64) (author *Author, err error)
	//Delete(ctx context.Context, id int64) (err error)
}
