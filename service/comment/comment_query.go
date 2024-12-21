package comment

import (
	"our_blog/model/dao"
	"our_blog/model/dto"
	"time"
)

type CommentGetFlow struct {
	CommentId  int64
	UserId     int64
	ArticleId  int64
	ParentId   *int64
	Content    string
	CreateTime time.Time
}

// CommentGet 根据评论 ID 获取评论的详细信息
func CommentQuery(commentId int64) (*dto.CommentGetResp, error) {
	return NewCommentGetFlow(commentId).Do()
}

// NewCommentGetFlow 创建 CommentGetFlow 实例
func NewCommentGetFlow(commentId int64) *CommentGetFlow {
	return &CommentGetFlow{
		CommentId: commentId,
	}
}

// Do 执行获取评论的操作
func (f *CommentGetFlow) Do() (*dto.CommentGetResp, error) {
	comment, err := f.GetComment()
	if err != nil {
		return nil, err
	}

	return &dto.CommentGetResp{
		CommentId:  comment.CommentId,
		UserId:     comment.UserId,
		ArticleId:  comment.ArticleId,
		ParentId:   comment.ParentId,
		Content:    comment.Content,
		CreateTime: comment.CreateTime,
	}, nil
}

// GetComment 从数据库中获取评论
func (f *CommentGetFlow) GetComment() (*dao.Comment, error) {
	comment, err := dao.NewCommentDaoInstance().QueryById(f.CommentId)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
