package services

import (
	"github.com/everywan/foundation-go/database"
	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw"
)

// TodoService is ...
type TodoService struct {
	db      *database.MysqlDB
	logger  *flog.Logger
	fileSvc xgxw.FileService
	userSvc xgxw.UserService
}

// NewTodoService create TodoService
func NewTodoService(db *database.MysqlDB, logger *flog.Logger, fileSvc xgxw.FileService, userSvc xgxw.UserService) *TodoService {
	return &TodoService{
		db:      db,
		logger:  logger,
		fileSvc: fileSvc,
		userSvc: userSvc,
	}
}

// 判断 UserService 是否实现了外层定义的所有接口
var _ xgxw.TodoService = &TodoService{}

// GetTodo is ...
func (t *TodoService) GetTodo(id, userID uint) (todo *xgxw.Todo, err error) {
	todo = &xgxw.Todo{}
	err = t.db.Where("id=? and user_id=?", id, userID).First(todo).Error
	return todo, err
}

// GetTodos is ...
func (t *TodoService) GetTodos(userID uint) (todos []*xgxw.Todo, err error) {
	todos = make([]*xgxw.Todo, 1)
	err = t.db.Where("user_id=?", userID).Find(todos).Error
	return todos, err
}
