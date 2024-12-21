package article

import (
	"our_blog/model/dao"
	"our_blog/model/dto"
	"time"
)

type ArticleCommentFlow struct {
	CommentId  int64
	UserId     int64
	ArticleId  int64
	ParentId   *int64
	Content    string
	CreateTime time.Time
}

func ArticleComment(req *dto.ArticleCommentReq, userid int64) (*dto.ArticleCommentResp, error) {
	return NewArticleCommentFlow(req, userid).Do()
}

func NewArticleCommentFlow(req *dto.ArticleCommentReq, userid int64) *ArticleCommentFlow {
	return &ArticleCommentFlow{
		UserId:    userid,
		ArticleId: req.ArticleId,
		ParentId:  req.ParentId,
		Content:   req.Content,
	}
}

func (f *ArticleCommentFlow) Do() (*dto.ArticleCommentResp, error) {
	if err := f.CheckData(); err != nil {
		return nil, err
	}
	if err := f.Publish(); err != nil {
		return nil, err
	}
	return &dto.ArticleCommentResp{
		CommentId:  f.CommentId,
		CreateTime: f.CreateTime,
	}, nil
}

func (f *ArticleCommentFlow) CheckData() error {
	_, err := dao.NewArticleDaoInstance().GetArticleById(f.ArticleId)
	if err != nil {
		return err
	}
	return nil
}

func (f *ArticleCommentFlow) Publish() error {
	comment := &dao.Comment{
		ArticleId:  f.ArticleId,
		UserId:     f.UserId,
		Content:    f.Content,
		ParentId:   f.ParentId,
		CreateTime: time.Now(),
	}
	if err := dao.NewCommentDaoInstance().Create(comment); err != nil {
		return err
	}
	f.CommentId = comment.CommentId
	f.CreateTime = comment.CreateTime
	return nil
}
