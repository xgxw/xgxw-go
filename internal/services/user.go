package services

import (
	"github.com/everywan/foundation-go/database"
	flog "github.com/everywan/foundation-go/log"
	"github.com/everywan/xgxw"
)

// UserService is ...
type UserService struct {
	db     *database.MysqlDB
	logger *flog.Logger
}

// NewUserService create UserService
func NewUserService(db *database.MysqlDB, logger *flog.Logger) *UserService {
	return &UserService{
		db:     db,
		logger: logger,
	}
}

// 判断 UserService 是否实现了外层定义的所有接口
var _ xgxw.UserService = &UserService{}
