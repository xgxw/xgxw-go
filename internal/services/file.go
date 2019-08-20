package services

import (
	"context"

	"github.com/everywan/foundation-go/storage"
	fstorage "github.com/everywan/foundation-go/storage"
	"github.com/everywan/xgxw"
)

// FileService is ...
type FileService struct {
	storage fstorage.ClientInterface
}

// NewFileService create FileService
func NewFileService(storage fstorage.ClientInterface) *FileService {
	return &FileService{
		storage: storage,
	}
}

// 判断 UserService 是否实现了外层定义的所有接口
var _ xgxw.FileService = &FileService{}

// Get is ...
func (this *FileService) Get(ctx context.Context, fid string) (file *xgxw.File, err error) {
	file = new(xgxw.File)
	buf, err := this.storage.GetObject(ctx, fid)
	if err != nil {
		return file, err
	}
	file.FileID = fid
	file.Name = fid
	file.Content = string(buf)
	return file, nil
}

// Put is ...
func (this *FileService) Put(ctx context.Context, fid, content string) (err error) {
	return this.storage.PutObject(ctx, fid, []byte(content))
}

// Del is ...
func (this *FileService) Del(ctx context.Context, fid string) (err error) {
	return this.storage.DelObject(ctx, fid)
}

// DelFiles is ...
func (this *FileService) DelFiles(ctx context.Context, fids []string) (err error) {
	return this.storage.DelObjects(ctx, fids)
}

// GetList is ...
func (this *FileService) GetCatalog(ctx context.Context, path string,
	opts storage.ListOption) (catalog string, paths []string, err error) {
	data, paths, err := this.storage.GetCatalog(ctx, path, opts)
	return string(data), paths, err
}
