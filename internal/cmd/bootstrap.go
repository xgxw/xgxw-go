package cmd

import (
	"log"
	"os"

	"github.com/everywan/xgxw"
	"github.com/everywan/xgxw/internal/services"

	"github.com/everywan/foundation-go/database"
	flog "github.com/everywan/foundation-go/log"
)

// Bootstrap 公用实例初始化
type bootstrap struct {
	Logger    *flog.Logger
	Database  *database.MysqlDB
	FileSvc   xgxw.FileService
	UserSvc   xgxw.UserService
	TodoSvc   xgxw.TodoService
	ResumeSvc xgxw.ResumeService
}

func newBootstrap(opts *ApplicationOps) (*bootstrap, error) {
	// Register
	db, err := database.NewMysqlDatabase(opts.Database.Mysql)
	handleInitError("database", err)
	logger := flog.NewLogger(opts.Logger, os.Stdout)
	fileSvc := services.NewFileService(db, logger)
	userSvc := services.NewFileService(db, logger)
	todoSvc := services.NewTodoService(db, logger, fileSvc, userSvc)
	resumeSvc := services.NewResumeService(db, logger, fileSvc, userSvc)
	return &bootstrap{
		Logger:    logger,
		Database:  db,
		FileSvc:   fileSvc,
		TodoSvc:   todoSvc,
		ResumeSvc: resumeSvc,
	}, nil
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
