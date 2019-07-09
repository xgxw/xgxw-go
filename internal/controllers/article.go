package controllers

import (
	"context"
	"net/http"

	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw"
	"github.com/everywan/xgxw/internal/constants"
	"github.com/labstack/echo"
)

type (
	// ArticleController is ...
	ArticleController struct {
		logger  *flog.Logger
		fileSvc xgxw.FileService
	}
)

// NewArticleController is ...
func NewArticleController(logger *flog.Logger, fileSvc xgxw.FileService) *ArticleController {
	return &ArticleController{
		logger:  logger,
		fileSvc: fileSvc,
	}
}

// Get is 获取Article.md
func (e *ArticleController) Get(ctx echo.Context) error {
	fid := ctx.Param("fid")
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	// Guest 用户只能访问 public 文件夹
	if ctx.Get(constants.IsGuest).(bool) {
		fid = "/public/" + fid
	}
	article, err := e.fileSvc.Get(context.Background(), fid)
	if err != nil {
		e.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, article)
}

type putRequestCarrier struct {
	Content string `json:"content" form:"content" query:"content"`
}

// Put is ...
func (e *ArticleController) Put(ctx echo.Context) error {
	fid := ctx.QueryParam("fid")
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	r := new(putRequestCarrier)
	if err := ctx.Bind(r); err != nil {
		e.logger.Error(err)
		return ctx.String(http.StatusBadRequest, "please input content")
	}
	err := e.fileSvc.Put(context.Background(), fid, r.Content)
	if err != nil {
		e.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusOK)
}
