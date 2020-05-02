package xgxw

import "context"

type (
	Document struct {
		Inode *Inode      `json:"inode"`
		Block interface{} `json:"block"`
	}
	CreateDocRequest struct {
		*CreateInodeRequest
		Block interface{}
	}
	CreateArticleDocRequest struct {
		*CreateInodeRequest
		*CreateArticleBlockRequest
	}
	CreateDirDocRequest struct {
		*CreateInodeRequest
		*CreateDirBlockRequest
	}
)

// DocumentService 文档服务. 其实就是 inode/block 的聚合, 为了保持 inode
// 的独立性, 从而拆分出 DocumentService 承担聚合操作的职责.
type DocumentService interface {
	Get(ctx context.Context, id int64) (doc *Document, err error)
	Create(ctx context.Context, req *CreateDocRequest) (doc *Document, err error)
	// Delete(ctx context.Context, id int64) (err error)
	// GetArticle(ctx context.Context, id int64) (doc *Document, err error)
	// GetDir(ctx context.Context, id int64) (doc *Document, err error)
	//CreateArticle(ctx context.Context, req *CreateArticleDocRequest) (doc *Document, err error)
	//CreateDir(ctx context.Context, req *CreateDirDocRequest) (doc *Document, err error)
}
