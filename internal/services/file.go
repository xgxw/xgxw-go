package services

import (
	"github.com/everywan/foundation-go/database"
	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw"
)

// FileService is ...
type FileService struct {
	db     *database.MysqlDB
	logger *flog.Logger
}

// NewFileService create FileService
func NewFileService(db *database.MysqlDB, logger *flog.Logger) *FileService {
	return &FileService{
		db:     db,
		logger: logger,
	}
}

// 判断 UserService 是否实现了外层定义的所有接口
var _ xgxw.FileService = &FileService{}

// GetFile is 获取文件
func (f *FileService) GetFile(id string) (file *xgxw.File, err error) {
	file = &xgxw.File{}
	err = f.db.Where("id=?", id).First(file).Error
	if err != nil {
		f.logger.Error(err)
	}
	return file, err
}

// GetFilesByName is 获取文件
func (f *FileService) GetFilesByName(name string) (file []*xgxw.File, err error) {
	return nil, nil
}
