package article

import (
	"our_blog/model/dao"
)

type ArticleCategoryFlow struct {
	CategoryStr string
}

func ArticleCategory(CategoryStr string) ([]*dao.Article, error) {
	return NewArticleSearchFlow(CategoryStr).Do()
}

func NewArticleCategoryFlow(CategoryStr string) *ArticleCategoryFlow {
	return &ArticleCategoryFlow{
		CategoryStr: CategoryStr,
	}
}

func (f *ArticleCategoryFlow) Do() ([]*dao.Article, error) {
	return f.Category()
}

func (f *ArticleCategoryFlow) Category() ([]*dao.Article, error) {
	ArticleList, err := dao.NewArticleDaoInstance().Category(f.CategoryStr)
	if err != nil {
		return nil, err
	}
	return ArticleList, nil
}
