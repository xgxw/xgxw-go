package xgxw

import (
	"context"
	"time"
)

type InodeType int

const (
	InodeTypeArticle InodeType = 1 << 0
	InodeTypeDir     InodeType = 1 << 1
	InodeTypeLink    InodeType = 1 << 2
)

type (
	Inode struct {
		ID       int64     `gorm:"column:id" json:"id"`
		Type     InodeType `gorm:"column:type" json:"type"`
		Name     string    `gorm:"column:name" json:"name"`
		BlockID  int64     `gorm:"column:block_id" json:"block_id" `
		AuthorID int64     `gorm:"column:author_id" json:"author_id"`

		CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
		UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
		DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	}
	CreateInodeRequest struct {
		Type     InodeType `json:"type"`
		Name     string    `json:"name"`
		BlockID  int64     `json:"block_id" `
		AuthorID int64     `json:"author_id"`
	}
)

func (Inode) TableName() string {
	return "inodes"
}

type InodeService interface {
	Get(ctx context.Context, id int64) (inode *Inode, err error)
	Create(ctx context.Context, req *CreateInodeRequest) (inode *Inode, err error)
	// Delete(ctx context.Context, id int64) (err error)
}
