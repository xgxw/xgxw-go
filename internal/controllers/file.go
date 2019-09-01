package controllers

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	flog "github.com/xgxw/foundation-go/log"
	fstorage "github.com/xgxw/foundation-go/storage"
	"github.com/xgxw/xgxw-go"
	"github.com/xgxw/xgxw-go/internal/utils"
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

func getFidFromPath(ctx echo.Context, prefix string) string {
	path := ctx.Request().URL.Path
	path = utils.CleanPath(path)
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	return path[len(prefix):]
}

// Get is 获取File.md
func (this *FileController) GetURL(ctx echo.Context) error {
	fid := getFidFromPath(ctx, "/v1/url/")
	if fid == "" {
		return ctx.NoContent(http.StatusNotFound)
	}
	if strings.HasSuffix(fid, "/") {
		return ctx.String(http.StatusRequestedRangeNotSatisfiable, "path is dir")
	}
	url, err := this.fileSvc.SignURL(context.Background(), fid, fstorage.HTTPGet, 0)
	if err != nil {
		this.logger.Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.String(http.StatusOK, url)
}

// Get is 获取File.md
func (this *FileController) Get(ctx echo.Context) error {
	fid := getFidFromPath(ctx, "/v1/file/")
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
	fid := getFidFromPath(ctx, "/v1/file/")
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
	fid := getFidFromPath(ctx, "/v1/file/")
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
	Options fstorage.ListOption `json:"options" form:"options" query:"options"`
}
type GetCatalogResopnseCarrier struct {
	Catalog string   `json:"catalog" form:"catalog" query:"catalog"`
	Paths   []string `json:"paths" form:"paths" query:"paths"`
}

// GetCatalog is ...
func (this *FileController) GetCatalog(ctx echo.Context) error {
	path := getFidFromPath(ctx, "/v1/catalog/")
	if path == "" {
		return ctx.NoContent(http.StatusNotFound)
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
