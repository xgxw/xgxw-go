package services

import (
	"context"

	fstorage "github.com/everywan/foundation-go/storage"
	"github.com/everywan/xgxw"
)

const (
	todoFileID = "todo.md"
)

// TodoService is ...
type TodoService struct {
	storage fstorage.StorageClientInterface
}

// NewTodoService create TodoService
func NewTodoService(storage fstorage.StorageClientInterface) *TodoService {
	return &TodoService{
		storage: storage,
	}
}

// 判断 UserService 是否实现了外层定义的所有接口
var _ xgxw.TodoService = &TodoService{}

// Get is ...
func (t *TodoService) Get(ctx context.Context) (todo *xgxw.Todo, err error) {
	todo = new(xgxw.Todo)
	buf, err := t.storage.GetObject(ctx, todoFileID)
	if err != nil {
		return todo, err
	}
	todo.Content = string(buf)
	return todo, nil
}

// Put is ...
func (t *TodoService) Put(ctx context.Context, content string) (err error) {
	return t.storage.PutObject(ctx, todoFileID, []byte(content))
}
