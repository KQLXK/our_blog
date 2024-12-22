package comment

import (
	"our_blog/model/dao"
	"our_blog/model/dto"
)

// CommentListFlow 处理获取文章评论及回复的逻辑
type CommentListFlow struct {
	ArticleId int64
	//Comments  []dto.CommentWithReplies // 包含评论和回复的结构
}

// CommentList 根据文章 ID 获取该文章的所有评论及其回复
func CommentList(articleId int64) (*dto.CommentListResp, error) {
	return NewCommentListFlow(articleId).Do()
}

// NewCommentListFlow 创建 CommentListFlow 实例
func NewCommentListFlow(articleId int64) *CommentListFlow {
	return &CommentListFlow{
		ArticleId: articleId,
	}
}

// Do 执行获取评论及其回复的操作
func (f *CommentListFlow) Do() (*dto.CommentListResp, error) {
	comments, err := f.GetCommentsAndReplies()
	if err != nil {
		return nil, err
	}

	return &dto.CommentListResp{
		Comments: comments,
	}, nil
}

// GetCommentsAndReplies 从数据库中获取文章的所有评论及其回复
func (f *CommentListFlow) GetCommentsAndReplies() ([]*dao.Comment, error) {
	RootCommentList, err := dao.NewCommentDaoInstance().GetRootCommentsByArticle(f.ArticleId)
	if err != nil {
		return nil, err
	}
	var CommentsList []*dao.Comment
	for _, Comment := range RootCommentList {
		if err = dao.NewCommentDaoInstance().RecursiveGetReplies(Comment); err != nil {
			return nil, err
		}
		CommentsList = append(CommentsList, Comment)
	}
	return CommentsList, nil
}
