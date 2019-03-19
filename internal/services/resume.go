package services

import (
	"github.com/everywan/foundation-go/database"
	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw"
)

// ResumeService is ...
type ResumeService struct {
	db      *database.MysqlDB
	logger  *flog.Logger
	fileSvc xgxw.FileService
	userSvc xgxw.UserService
}

// NewResumeService create ResumeService
func NewResumeService(db *database.MysqlDB, logger *flog.Logger, fileSvc xgxw.FileService, userSvc xgxw.UserService) *ResumeService {
	return &ResumeService{
		db:      db,
		logger:  logger,
		fileSvc: fileSvc,
		userSvc: userSvc,
	}
}

// 判断 UserService 是否实现了外层定义的所有接口
var _ xgxw.ResumeService = &ResumeService{}
