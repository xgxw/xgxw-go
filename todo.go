package xgxw

import "context"

// Todo is ...
type Todo struct {
	Content  string `json:"content" gorm:"column:content"`
	UpdateAt string `json:"update_at" gorm:"column:update_at"`
}

// TodoService is ...
type TodoService interface {
	Get(ctx context.Context) (todo *Todo, err error)
	// 先作成常量的, 后续改为可以diff的.
	Put(ctx context.Context, content string) (err error)
}
