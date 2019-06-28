package controllers

import (
	"context"
	"net/http"

	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw"
	"github.com/labstack/echo"
)

type (
	// TodoController is ...
	TodoController struct {
		logger  *flog.Logger
		todoSvc xgxw.TodoService
	}
)

// NewTodoController is ...
func NewTodoController(logger *flog.Logger, todoSvc xgxw.TodoService) *TodoController {
	return &TodoController{
		logger:  logger,
		todoSvc: todoSvc,
	}
}

// Get is 获取Todo.md
func (e *TodoController) Get(ctx echo.Context) error {
	todo, err := e.todoSvc.Get(context.Background())
	if err != nil {
		e.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, todo)
}

// Put is ...
func (e *TodoController) Put(ctx echo.Context) error {
	type requestCarrier struct {
		Content string `json:"content" form:"content" query:"content"`
	}
	r := new(requestCarrier)
	if err := ctx.Bind(r); err != nil {
		e.logger.Error(err)
		return ctx.String(http.StatusBadRequest, "please input content")
	}
	err := e.todoSvc.Put(context.Background(), r.Content)
	if err != nil {
		e.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusOK)
}
