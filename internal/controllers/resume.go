package controllers

import (
	"net/http"

	"github.com/everywan/xgxw"
	"github.com/labstack/echo"
)

type (
	// ResumeController is ...
	ResumeController struct {
		fileSvc xgxw.FileService
	}
)

// NewResumeController is ...
func NewResumeController(fileSvc xgxw.FileService) *ResumeController {
	return &ResumeController{
		fileSvc: fileSvc,
	}
}

// GetResume is ...
func (e *ResumeController) GetResume(ctx echo.Context) error {
	file, err := e.fileSvc.GetFile("2")
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.String(http.StatusOK, file.URL)
}
