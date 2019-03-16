package cmd

import (
	"log"
	"os"

	"github.com/everywan/xgxw/internal/services"

	"github.com/everywan/foundation-go/database"
	flog "github.com/everywan/foundation-go/log"
)

// Bootstrap 公用实例初始化
type Bootstrap struct {
	Logger   *flog.Logger
	Database *database.MysqlDB
	FileSvc  *services.FileService
}

// NewBootstrap is ...
func NewBootstrap(opts ApplicationOps) (*Bootstrap, error) {
	// Register
	db, err := database.NewMysqlDatabase(opts.Database.Mysql)
	handleInitError("database", err)
	logger := flog.NewLogger(opts.Logger, os.Stdout)
	fileSvc := services.NewFileService(db, logger)
	return &Bootstrap{
		Logger:   logger,
		Database: db,
		FileSvc:  fileSvc,
	}, nil
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
