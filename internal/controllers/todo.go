package controllers

import (
	"net/http"
	"strconv"

	"github.com/everywan/xgxw"
	"github.com/labstack/echo"
)

type (
	// TodoController is ...
	TodoController struct {
		fileSvc xgxw.FileService
		todoSvc xgxw.TodoService
	}
)

// NewTodoController is ...
func NewTodoController(fileSvc xgxw.FileService, todoSvc xgxw.TodoService) *TodoController {
	return &TodoController{
		fileSvc: fileSvc,
		todoSvc: todoSvc,
	}
}

// GetTodo is 获取Todo.md
func (e *TodoController) GetTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	_id, _ := strconv.ParseUint(id, 10, 32)
	// 在中间件校验 user_id 存在且通过 ctx.set 设置
	userID := ctx.Get("user_id").(uint)
	todo, err := e.todoSvc.GetTodo(uint(_id), userID)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, todo)
}

// GetTodos is ..
func (e *TodoController) GetTodos(ctx echo.Context) error {
	userID := ctx.Get("user_id").(uint)
	todos, err := e.todoSvc.GetTodos(userID)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, todos)
}
