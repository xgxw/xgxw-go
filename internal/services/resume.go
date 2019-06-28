package services

import (
	"context"

	fstorage "github.com/everywan/foundation-go/storage"
	"github.com/everywan/xgxw"
)

const (
	resumeFileID = "resume.md"
)

// ResumeService is ...
type ResumeService struct {
	storage fstorage.StorageClientInterface
}

// NewResumeService create ResumeService
func NewResumeService(storage fstorage.StorageClientInterface) *ResumeService {
	return &ResumeService{
		storage: storage,
	}
}

// 判断 ResumeService 是否实现了外层定义的所有接口
var _ xgxw.ResumeService = &ResumeService{}

// Get is ...
func (r *ResumeService) Get(ctx context.Context) (resume *xgxw.Resume, err error) {
	resume = new(xgxw.Resume)
	buf, err := r.storage.GetObject(ctx, resumeFileID)
	if err != nil {
		return resume, err
	}
	resume.Content = string(buf)
	return resume, nil
}

// Put is ...
func (r *ResumeService) Put(ctx context.Context, content string) (err error) {
	return r.storage.PutObject(ctx, resumeFileID, []byte(content))
}
