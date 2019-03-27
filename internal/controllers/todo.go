package controllers

import (
	"net/http"
	"strconv"

	"github.com/everywan/xgxw"
	"github.com/everywan/xgxw/internal/constants"
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
func (e *TodoController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	_id, _ := strconv.ParseUint(id, 10, 32)
	userID := ctx.Get(constants.UserID).(uint)
	todo, err := e.todoSvc.Get(uint(_id), userID)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, todo)
}

// GetTodos is ..
func (e *TodoController) GetTodos(ctx echo.Context) error {
	userID := ctx.Get(constants.UserID).(uint)
	todos, err := e.todoSvc.GetTodos(userID)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, todos)
}
