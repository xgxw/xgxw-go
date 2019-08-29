package xgxw

import (
	"context"

	"github.com/xgxw/foundation-go/storage"
)

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
	// Del is 删除文件
	Del(ctx context.Context, fid string) (err error)
	// DelFiles is 删除多个文件
	DelFiles(ctx context.Context, fids []string) (err error)
	// GetCatalog is 获取文件列表, 返回目录树的json字符串
	GetCatalog(ctx context.Context, path string, opts storage.ListOption) (catalog string, paths []string, err error)
}
