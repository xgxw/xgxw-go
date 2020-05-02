package xgxw

import "context"

type (
	ArticleBlock struct {
		ID      int64    `gorm:"column:id" json:"id"`
		FID     int64    `gorm:"column:fid" json:"fid"` // file id
		Visited int      `gorm:"column:visited" json:"visited"`
		Score   int      `gorm:"column:score" json:"score"`
		Tag     []string `gorm:"column:tag" json:"tag"`
		Desc    string   `gorm:"column:desc" json:"desc"`
	}
	DirectoryBlock struct {
		ID   int64   `gorm:"column:id" json:"id"`
		Docs []int64 `gorm:"column:docs" json:"docs"` // 路径下的文档列表
	}
	CreateArticleBlockRequest struct {
		File *struct {
			Name string `json:"name"`
			Data string `json:"data"`
		} `json:"file"`
		Score int      `json:"score"`
		Tag   []string `json:"tag"`
		Desc  string   `json:"desc"`
	}
	CreateDirBlockRequest struct{}
)

func (ArticleBlock) TableName() string {
	return "article_blocks"
}

func (DirectoryBlock) TableName() string {
	return "dir_blocks"
}

type BlockService interface {
	Get(ctx context.Context, id int64) (block interface{}, err error)
	Create(ctx context.Context, req interface{}) (block interface{}, err error)
	// Delete(ctx context.Context, id int64) (err error)
}
