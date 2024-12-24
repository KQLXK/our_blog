package article

import (
	"errors"
	"our_blog/model/dao"
)

type ArticleSearchFlow struct {
	SearchStr string
}

var (
	ArticleNotExistErr = errors.New("article not found")
)

func ArticleSearch(SearchStr string) ([]*dao.Article, error) {
	return NewArticleSearchFlow(SearchStr).Do()
}

func NewArticleSearchFlow(SearchStr string) *ArticleSearchFlow {
	return &ArticleSearchFlow{
		SearchStr: SearchStr,
	}
}

func (f *ArticleSearchFlow) Do() ([]*dao.Article, error) {
	return f.Search()
}

func (f *ArticleSearchFlow) Search() ([]*dao.Article, error) {
	ArticleList, err := dao.NewArticleDaoInstance().Search(f.SearchStr)
	if err != nil {
		return nil, err
	}
	return ArticleList, nil
}
