package article

import (
	"errors"
	"gorm.io/gorm"
	"our_blog/model/dao"
)

type ArticleDeleteFlow struct {
	ArticleId int64
	UserId    int64
}

var (
	DeleteUnauthorizedErr = errors.New("未被授权的删除")
)

func ArticleDelete(articleid int64, userid int64) error {
	return NewArticleDeleteFlow(articleid, userid).Do()
}

func NewArticleDeleteFlow(articleid int64, userid int64) *ArticleDeleteFlow {
	return &ArticleDeleteFlow{
		ArticleId: articleid,
		UserId:    userid,
	}
}

func (f *ArticleDeleteFlow) Do() error {
	if err := f.CheckData(); err != nil {
		return err
	}
	if err := f.Delete(); err != nil {
		return err
	}
	return nil
}

func (f *ArticleDeleteFlow) CheckData() error {
	//检查文章是否存在
	article, err := dao.NewArticleDaoInstance().GetArticleById(f.ArticleId)
	if err == gorm.ErrRecordNotFound {
		return ArticleNotFoundErr
	} else if err != nil {
		return err
	}
	//如果是管理员，不用验证userid
	if isadmin, _ := dao.NewUserDaoInstance().IsAdmin(f.UserId); isadmin {
		return nil
	}
	if article.UserId != f.UserId {
		return DeleteUnauthorizedErr
	}
	return nil
}

func (f *ArticleDeleteFlow) Delete() error {
	err := dao.NewArticleDaoInstance().DeleteAnArticle(f.ArticleId)
	if err != nil {
		return err
	}
	return nil
}
