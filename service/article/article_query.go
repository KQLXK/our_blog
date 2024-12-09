package article

import (
	"errors"
	"gorm.io/gorm"
	"our_blog/model/dao"
	"our_blog/model/dto"
)

var (
	BadDataErr           = errors.New("参数不合法")
	QueryUnauthorizedErr = errors.New("查询未被授权")
	ArticleNotFoundErr   = errors.New("文章不存在")
)

type ArtQueryByIdFlow struct {
	ArticleId int64
}

func ArtQueryById(ArticleId int64) (*dto.ArtQueryByIdResp, error) {
	return NewArtQueryByIdFlow(ArticleId).Do()
}

func NewArtQueryByIdFlow(ArticleId int64) *ArtQueryByIdFlow {
	return &ArtQueryByIdFlow{
		ArticleId: ArticleId,
	}
}

func (f *ArtQueryByIdFlow) Do() (*dto.ArtQueryByIdResp, error) {
	if err := f.CheckData(); err != nil {
		return nil, err
	}
	article, err := f.QueryById()
	if err != nil {
		return nil, err
	}
	return &dto.ArtQueryByIdResp{
		Title:      article.Title,
		ArticleId:  article.ArticleId,
		UserId:     article.UserId,
		Excerpt:    article.Excerpt,
		Category:   article.Category,
		Content:    article.Content,
		Status:     article.Status,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
	}, nil
}

//查询文章不需要验证userid
//func (f *ArtQueryByIdFlow) CheckUserid() error {
//	article, err := dao.NewArticleDaoInstance().GetArticleById(f.ArticleId)
//	if err != nil {
//		return err
//	}
//	if article.UserId != f.UserId {
//		return QueryUnauthorizedErr
//	}
//	return nil
//}

func (f *ArtQueryByIdFlow) CheckData() error {
	if f.ArticleId < 0 {
		return BadDataErr
	}
	return nil
}

func (f *ArtQueryByIdFlow) QueryById() (*dao.Article, error) {
	article, err := dao.NewArticleDaoInstance().GetArticleById(f.ArticleId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ArticleNotFoundErr
		} else {
			return nil, err
		}
	}
	return article, nil
}
