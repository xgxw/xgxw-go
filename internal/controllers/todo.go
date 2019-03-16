package controllers

import (
	"net/http"

	"github.com/everywan/xgxw/internal/services"

	"github.com/labstack/echo"
)

type (
	// TodoController is ...
	TodoController struct {
		fileSvc *services.FileService
	}
)

// NewTodoController is ...
func NewTodoController(fileSvc *services.FileService) *TodoController {
	return &TodoController{
		fileSvc: fileSvc,
	}
}

// GetTodo is 获取Todo.md
func (e *TodoController) GetTodo(ctx echo.Context) error {
	file, err := e.fileSvc.GetFile("1")
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.String(http.StatusOK, file.URL)
}
