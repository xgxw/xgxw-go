package services

import (
	"context"

	fstorage "github.com/everywan/foundation-go/storage"
	"github.com/everywan/xgxw"
)

// FileService is ...
type FileService struct {
	storage fstorage.StorageClientInterface
}

// NewFileService create FileService
func NewFileService(storage fstorage.StorageClientInterface) *FileService {
	return &FileService{
		storage: storage,
	}
}

// 判断 UserService 是否实现了外层定义的所有接口
var _ xgxw.FileService = &FileService{}

// Get is ...
func (t *FileService) Get(ctx context.Context, fid string) (file *xgxw.File, err error) {
	// 判断 fid 格式, .. 等可能影响权限?
	file = new(xgxw.File)
	buf, err := t.storage.GetObject(ctx, fid)
	if err != nil {
		return file, err
	}
	file.Content = string(buf)
	return file, nil
}

// Put is ...
func (t *FileService) Put(ctx context.Context, fid, content string) (err error) {
	// 判断 fid 格式, .. 等可能影响权限?
	return t.storage.PutObject(ctx, fid, []byte(content))
}
