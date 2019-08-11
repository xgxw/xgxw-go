package controllers

import (
	"context"
	"net/http"
	"strings"

	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/foundation-go/storage"
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

// Del is ...
func (this *ArticleController) Del(ctx echo.Context) error {
	fid := this.getFidFromPath(ctx)
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	err := this.fileSvc.Del(context.Background(), fid)
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusOK)
}

type DelFilesRequestCarrier struct {
	Fids []string `json:"fids" form:"fids" query:"fids"`
}

// DelFiles is ...
func (this *ArticleController) DelFiles(ctx echo.Context) error {
	r := new(DelFilesRequestCarrier)
	if err := ctx.Bind(r); err != nil {
		this.logger.Error(err)
		return ctx.String(http.StatusBadRequest, "please input fids")
	}
	err := this.fileSvc.DelFiles(context.Background(), r.Fids)
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusOK)
}

type GetCatalogRequestCarrier struct {
	Path    string             `json:"path" form:"path" query:"path"`
	Options storage.ListOption `json:"options" form:"options" query:"options"`
}

// GetPublicCatalog is ...
func (this *ArticleController) GetPublicCatalog(ctx echo.Context) error {
	r := new(GetCatalogRequestCarrier)
	if err := ctx.Bind(r); err != nil {
		this.logger.Error(err)
		return ctx.String(http.StatusBadRequest, "please input fids")
	}
	if strings.Contains(r.Path, "..") {
		return ctx.String(http.StatusBadRequest, "path can't contains `..`!")
	}
	r.Path = "public/" + r.Path
	catalog, err := this.fileSvc.GetCatalog(context.Background(), r.Path, r.Options)
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.String(http.StatusOK, catalog)
}

// GetCatalog is ...
func (this *ArticleController) GetCatalog(ctx echo.Context) error {
	r := new(GetCatalogRequestCarrier)
	if err := ctx.Bind(r); err != nil {
		this.logger.Error(err)
		return ctx.String(http.StatusBadRequest, "please input fids")
	}
	catalog, err := this.fileSvc.GetCatalog(context.Background(), r.Path, r.Options)
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.String(http.StatusOK, catalog)
}
