package article

import (
	"errors"
	"our_blog/model/dao"
	"our_blog/model/dto"
	"time"
)

var (
	EmptyDataErr = errors.New("标题不能为空")
)

type ArticlePublishFlow struct {
	Title     string    // 文章标题
	ArticleId int64     // 文章 ID
	UserId    int64     // 作者 ID
	Excerpt   string    // 文章摘要
	Category  string    // 文章分类
	Content   string    // 文章内容
	Status    string    // 文章状态
	CreateAt  time.Time //创建时间
	UpdateAt  time.Time //更新时间
}

func ArticlePublish(req dto.ArticlePublishReq, userid int64) (*dto.ArticlePublishResp, error) {
	return NewArticlePublishFlow(req, userid).Do()
}

func NewArticlePublishFlow(req dto.ArticlePublishReq, userid int64) *ArticlePublishFlow {
	return &ArticlePublishFlow{
		Title:    req.Title,
		Excerpt:  req.Excerpt,
		Category: req.Category,
		Content:  req.Content,
		Status:   req.Status,
		UserId:   userid,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

func (f *ArticlePublishFlow) Do() (*dto.ArticlePublishResp, error) {
	if err := f.Publish(); err != nil {
		return nil, err
	}
	return &dto.ArticlePublishResp{
		ArticleId: f.ArticleId,
		Title:     f.Title,
		CreateAt:  f.CreateAt,
	}, nil
}

func (f *ArticlePublishFlow) CheckData() error {
	if f.Title == "" || f.Category == "" {
		return EmptyDataErr
	}
	return nil
}

func (f *ArticlePublishFlow) Publish() error {
	article := &dao.Article{
		Title:      f.Title,
		Excerpt:    f.Excerpt,
		Category:   f.Category,
		Content:    f.Content,
		Status:     f.Status,
		UserId:     f.UserId,
		CreateTime: f.CreateAt,
		UpdateTime: f.UpdateAt,
		View:       0,
	}
	if err := dao.NewArticleDaoInstance().CreateAnArticle(article); err != nil {
		return err
	}
	f.ArticleId = article.ArticleId
	return nil
}
