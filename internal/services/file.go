package services

import (
	"context"

	fstorage "github.com/xgxw/foundation-go/storage"
	"github.com/xgxw/xgxw-go"
)

// FileService is ...
type FileService struct {
	defaultExpired int64
	storage        fstorage.ClientInterface
}

// NewFileService create FileService
func NewFileService(storage fstorage.ClientInterface, defaultExpired int64) *FileService {
	return &FileService{
		storage:        storage,
		defaultExpired: defaultExpired,
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
	opts fstorage.ListOption) (catalog string, paths []string, err error) {
	data, paths, err := this.storage.GetCatalog(ctx, path, opts)
	return string(data), paths, err
}

// SignURL is ...
func (this *FileService) SignURL(ctx context.Context, fid string,
	method fstorage.HTTPMethod, expiredInSec int64) (url string, err error) {
	if expiredInSec == 0 {
		expiredInSec = this.defaultExpired
	}
	return this.storage.SignURL(ctx, fid, method, expiredInSec, 0)
}
