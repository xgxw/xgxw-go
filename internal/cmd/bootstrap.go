package cmd

import (
	"log"
	"os"

	"github.com/everywan/xgxw"
	"github.com/everywan/xgxw/internal/services"

	flog "github.com/everywan/foundation-go/log"
	fstorage "github.com/everywan/foundation-go/storage"
)

// Bootstrap 公用实例初始化
type bootstrap struct {
	Logger    *flog.Logger
	TodoSvc   xgxw.TodoService
	ResumeSvc xgxw.ResumeService
	Options   *ApplicationOps
}

func newBootstrap(opts *ApplicationOps) (boot *bootstrap, err error) {
	logger := flog.NewLogger(opts.Logger, os.Stdout)
	storage, err := fstorage.NewOssClient(&opts.Oss)
	if err != nil {
		return boot, err
	}
	todoSvc := services.NewTodoService(storage)
	resumeSvc := services.NewResumeService(storage)
	boot = &bootstrap{
		Logger:    logger,
		TodoSvc:   todoSvc,
		ResumeSvc: resumeSvc,
		Options:   opts,
	}
	return boot, nil
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
