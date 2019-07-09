package xgxw

import "context"

// File is ...
type File struct {
	FileID   string `json:"fid" gorm:"column:fid"`
	Name     string `json:"name" gorm:"column:name"`
	Content  string `json:"content" gorm:"column:content"`
	UpdateAt string `json:"update_at" gorm:"column:update_at"`
}

// FileService is ...
type FileService interface {
	Get(ctx context.Context, fid string) (todo *File, err error)
	// 先作成常量的, 后续改为可以diff的.
	Put(ctx context.Context, fid, content string) (err error)
}
