package xgxw

import "context"

// Resume is ...
type Resume struct {
	Content  string `json:"content" gorm:"column:content"`
	UpdateAt string `json:"update_at" gorm:"column:update_at"`
}

// ResumeService is ...
type ResumeService interface {
	Get(ctx context.Context) (resume *Resume, err error)
	// 先作成常量的, 后续改为可以diff的.
	Put(ctx context.Context, content string) (err error)
}
