package article

import (
	"our_blog/model/dao"
)

type ArtQueryByPageFlow struct {
	Page     int
	PageSize int
	UserId   int64
}

func QueryByPage(Page int, PageSize int, UserId int64) ([]dao.Article, error) {
	return NewArtQueryByPageFlow(Page, PageSize, UserId).Do()
}

func NewArtQueryByPageFlow(Page int, PageSize int, UserId int64) *ArtQueryByPageFlow {
	return &ArtQueryByPageFlow{
		Page:     Page,
		PageSize: PageSize,
		UserId:   UserId,
	}
}

func (f *ArtQueryByPageFlow) Do() ([]dao.Article, error) {
	if err := f.CheckData(); err != nil {
		return nil, err
	}
	return f.QueryByPage()
}

func (f *ArtQueryByPageFlow) CheckData() error {
	if f.Page < 0 {
		return BadDataErr
	}
	if f.PageSize < 0 {
		return BadDataErr
	}
	return nil
}

func (f *ArtQueryByPageFlow) QueryByPage() ([]dao.Article, error) {
	ArticleList, err := dao.NewArticleDaoInstance().GetAricleListByPages(f.Page, f.PageSize)
	if err != nil {
		return nil, err
	}
	return ArticleList, nil
}
