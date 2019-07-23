package controllers

import (
	"context"
	"net/http"

	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw"
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

func (this *ArticleController) getFidFromPath(ctx echo.Context) string {
	path := ctx.Request().URL.Path
	if len(path) < 9 {
		return ""
	}
	return path[9:]
}

// Get is 获取Article.md
func (this *ArticleController) Get(ctx echo.Context) error {
	fid := this.getFidFromPath(ctx)
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	article, err := this.fileSvc.Get(context.Background(), fid)
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, article)
}

type putRequestCarrier struct {
	Content string `json:"content" form:"content" query:"content"`
}

// Put is ...
func (this *ArticleController) Put(ctx echo.Context) error {
	fid := this.getFidFromPath(ctx)
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	r := new(putRequestCarrier)
	if err := ctx.Bind(r); err != nil {
		this.logger.Error(err)
		return ctx.String(http.StatusBadRequest, "please input content")
	}
	err := this.fileSvc.Put(context.Background(), fid, r.Content)
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusOK)
}
