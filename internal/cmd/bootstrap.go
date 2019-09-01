package cmd

import (
	"log"
	"os"

	"github.com/xgxw/xgxw-go"
	"github.com/xgxw/xgxw-go/internal/services"

	flog "github.com/xgxw/foundation-go/log"
	fstorage "github.com/xgxw/foundation-go/storage"
)

// Bootstrap 公用实例初始化
type bootstrap struct {
	Logger  *flog.Logger
	FileSvc xgxw.FileService
	Options *ApplicationOps
}

func newBootstrap(opts *ApplicationOps) (boot *bootstrap, err error) {
	logger := flog.NewLogger(opts.Logger, os.Stdout)
	storage, err := fstorage.NewOssClient(&opts.Oss)
	if err != nil {
		return boot, err
	}
	defaultExpired := opts.Server.DefaultExpired
	if defaultExpired == 0 {
		defaultExpired = 60
	}
	fileSvc := services.NewFileService(storage, defaultExpired)
	boot = &bootstrap{
		Logger:  logger,
		FileSvc: fileSvc,
		Options: opts,
	}
	return boot, nil
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
