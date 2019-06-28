package controllers

import (
	"context"
	"net/http"

	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw"
	"github.com/labstack/echo"
)

type (
	// ResumeController is ...
	ResumeController struct {
		logger    *flog.Logger
		resumeSvc xgxw.ResumeService
	}
)

// NewResumeController is ...
func NewResumeController(logger *flog.Logger, resumeSvc xgxw.ResumeService) *ResumeController {
	return &ResumeController{
		logger:    logger,
		resumeSvc: resumeSvc,
	}
}

// Get is ...
func (e *ResumeController) Get(ctx echo.Context) error {
	resume, err := e.resumeSvc.Get(context.Background())
	if err != nil {
		e.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, resume)
}

// Put is ...
func (e *ResumeController) Put(ctx echo.Context) error {
	type requestCarrier struct {
		Content string `json:"content" form:"content" query:"content"`
	}
	r := new(requestCarrier)
	if err := ctx.Bind(r); err != nil {
		e.logger.Error(err)
		return ctx.String(http.StatusBadRequest, "please input content")
	}
	err := e.resumeSvc.Put(context.Background(), r.Content)
	if err != nil {
		e.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusOK)
}
