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
	// FileController is ...
	FileController struct {
		logger  *flog.Logger
		fileSvc xgxw.FileService
	}
)

// NewFileController is ...
func NewFileController(logger *flog.Logger, fileSvc xgxw.FileService) *FileController {
	return &FileController{
		logger:  logger,
		fileSvc: fileSvc,
	}
}

func (this *FileController) getFidFromPath(ctx echo.Context) string {
	path := ctx.Request().URL.Path
	if len(path) < 9 {
		return ""
	}
	return path[9:]
}

func (this *FileController) getPathFromPath(ctx echo.Context) string {
	path := ctx.Request().URL.Path
	if len(path) < 12 {
		return ""
	}
	return path[12:]
}

// Get is 获取File.md
func (this *FileController) Get(ctx echo.Context) error {
	fid := this.getFidFromPath(ctx)
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	if strings.HasSuffix(fid, "/") {
		return ctx.String(http.StatusRequestedRangeNotSatisfiable, "path is dir")
	}
	file, err := this.fileSvc.Get(context.Background(), fid)
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, file)
}

type putRequestCarrier struct {
	Content string `json:"content" form:"content" query:"content"`
}

// Put is ...
func (this *FileController) Put(ctx echo.Context) error {
	fid := this.getFidFromPath(ctx)
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	if strings.HasSuffix(fid, "/") {
		return ctx.String(http.StatusRequestedRangeNotSatisfiable, "path is dir")
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
func (this *FileController) Del(ctx echo.Context) error {
	fid := this.getFidFromPath(ctx)
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	if strings.HasSuffix(fid, "/") {
		return ctx.String(http.StatusRequestedRangeNotSatisfiable, "path is dir")
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
func (this *FileController) DelFiles(ctx echo.Context) error {
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
	Options storage.ListOption `json:"options" form:"options" query:"options"`
}
type GetCatalogResopnseCarrier struct {
	Catalog string   `json:"catalog" form:"catalog" query:"catalog"`
	Paths   []string `json:"paths" form:"paths" query:"paths"`
}

// GetCatalog is ...
func (this *FileController) GetCatalog(ctx echo.Context) error {
	path := this.getPathFromPath(ctx)
	if path == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	if !strings.HasSuffix(path, "/") {
		return ctx.String(http.StatusRequestedRangeNotSatisfiable, "path must be dir")
	}
	r := new(GetCatalogRequestCarrier)
	if err := ctx.Bind(r); err != nil {
		this.logger.Error(err)
		return ctx.String(http.StatusBadRequest, "please input params")
	}
	catalog, paths, err := this.fileSvc.GetCatalog(context.Background(), path, r.Options)
	resp := &GetCatalogResopnseCarrier{
		Catalog: catalog,
		Paths:   paths,
	}
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, resp)
}
