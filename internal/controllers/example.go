package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	// ExampleController is ...
	ExampleController struct {
	}
)

// NewExampleController is ...
func NewExampleController() *ExampleController {
	return &ExampleController{}
}

// SayHello is ...
func (this *ExampleController) SayHello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello")
}
