package article

import (
	"errors"
	"gorm.io/gorm"
	"our_blog/model/dao"
	"our_blog/model/dto"
	"time"
)

type ArticleUpdateFlow struct {
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

var (
	UpdateUnautherizedErr = errors.New("没有权限更新")
)

func ArticleUpdate(req dto.ArticleUpdateReq, userid int64) (*dto.ArticleUpdateResp, error) {
	return NewArticleUpdateFlow(req, userid).Do()
}

func NewArticleUpdateFlow(req dto.ArticleUpdateReq, userid int64) *ArticleUpdateFlow {
	return &ArticleUpdateFlow{
		Title:     req.Title,
		UserId:    userid,
		ArticleId: req.ArticleId,
		Excerpt:   req.Excerpt,
		Category:  req.Category,
		Content:   req.Content,
		Status:    req.Status,
		UpdateAt:  time.Now(),
	}
}

func (f *ArticleUpdateFlow) Do() (*dto.ArticleUpdateResp, error) {
	if err := f.CheckUserId(); err != nil {
		return nil, err
	}
	if err := f.Update(); err != nil {
		return nil, err
	}
	return &dto.ArticleUpdateResp{
		ArticleId: f.ArticleId,
		Title:     f.Title,
		UpdateAt:  f.UpdateAt,
	}, nil
}

func (f *ArticleUpdateFlow) CheckUserId() error {
	article, err := dao.NewArticleDaoInstance().GetArticleById(f.ArticleId)
	if err == gorm.ErrRecordNotFound {
		return ArticleNotFoundErr
	} else if err != nil {
		return err
	}
	if article.UserId != f.UserId {
		return UpdateUnautherizedErr
	}
	return nil
}

func (f *ArticleUpdateFlow) Update() error {
	article := &dao.Article{
		Title:      f.Title,
		ArticleId:  f.ArticleId,
		Excerpt:    f.Excerpt,
		Category:   f.Category,
		Content:    f.Content,
		Status:     f.Status,
		UpdateTime: f.UpdateAt,
	}
	if err := dao.NewArticleDaoInstance().UpdateAnArticle(article); err != nil {
		return err
	}
	return nil
}
